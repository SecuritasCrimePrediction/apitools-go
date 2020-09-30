package blobx

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"io"
	"path"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/storage"
	"gopkg.in/go-playground/validator.v9"
)

type ContainerConfig struct {
	AccountName   string `validate:"required,min=2"`
	AccountKey    string `validate:"required,min=2"`
	ContainerName string `validate:"required,min=2"`
}

type SASOptions struct {
	ValidFrom    time.Time
	ValidTo      time.Time
	ReadAccess   bool
	AddAccess    bool
	CreateAccess bool
	WriteAccess  bool
	DeleteAccess bool
	ListAccess   bool
}

type ContainerConn struct {
	container *storage.Container
}

// Get a connection to a container in an azure storage account.
// Use this func if you already have a connection to an Azure storage account.
func NewContainerConn(storage storage.BlobStorageClient, containerName string) (*ContainerConn, error) {
	container := storage.GetContainerReference(containerName)
	exists, err := container.Exists()
	if err != nil {
		return nil, fmt.Errorf("could not connect to the storage account, err: %v", err)
	}

	if !exists {
		return nil, fmt.Errorf("no container with the name %q exist in this storage account", containerName)
	}

	return &ContainerConn{
		container: container,
	}, nil
}

// Get a connection to a container in an azure storage account.
// Use this func if you don't have a connection to an Azure account already.
func NewAccountContainerConn(conf ContainerConfig) (*ContainerConn, error) {
	v := validator.New()
	if err := v.Struct(conf); err != nil {
		return nil, fmt.Errorf("config error: %v", err)
	}

	client, err := storage.NewBasicClient(conf.AccountName, conf.AccountKey)
	if err != nil {
		return nil, fmt.Errorf("could not connect to azure, err: %v", err)
	}

	return NewContainerConn(client.GetBlobService(), conf.ContainerName)
}

func (c *ContainerConn) GetContainerSASURI(ctx context.Context, opts SASOptions) (string, error) {
	return c.container.GetSASURI(storage.ContainerSASOptions{
		ContainerSASPermissions: storage.ContainerSASPermissions{
			BlobServiceSASPermissions: storage.BlobServiceSASPermissions{
				Read:   opts.ReadAccess,
				Add:    opts.AddAccess,
				Create: opts.CreateAccess,
				Write:  opts.WriteAccess,
				Delete: opts.DeleteAccess,
			},
			List: opts.ListAccess,
		},
		OverrideHeaders: storage.OverrideHeaders{},
		SASOptions: storage.SASOptions{
			Start:  opts.ValidFrom,
			Expiry: opts.ValidTo,
		},
	})
}

func (c *ContainerConn) GetBlobSASURI(ctx context.Context, blobName string, opts SASOptions) (string, error) {
	blob := c.container.GetBlobReference(blobName)

	exist, err := blob.Exists()
	if err != nil {
		return "", err
	}

	if !exist {
		return "", fmt.Errorf("blob with the name %q does not exist in this container", blobName)
	}

	return blob.GetSASURI(storage.BlobSASOptions{
		BlobServiceSASPermissions: storage.BlobServiceSASPermissions{
			Read:   opts.ReadAccess,
			Add:    opts.AddAccess,
			Create: opts.CreateAccess,
			Write:  opts.WriteAccess,
			Delete: opts.DeleteAccess,
		},
		OverrideHeaders: storage.OverrideHeaders{},
		SASOptions: storage.SASOptions{
			Start:  opts.ValidFrom,
			Expiry: opts.ValidTo,
		},
	})
}

func (c *ContainerConn) BlobReader(ctx context.Context, blobName string) (io.ReadCloser, error) {
	blob := c.container.GetBlobReference(blobName)

	exist, err := blob.Exists()
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, fmt.Errorf("blob with the name %q does not exist in this container", blobName)
	}

	return blob.Get(&storage.GetBlobOptions{})
}

func (c *ContainerConn) BlobBytes(ctx context.Context, blob string) ([]byte, error) {
	var err error
	reader, err := c.BlobReader(ctx, blob)
	if err != nil {
		return nil, err
	}

	buf := bytes.Buffer{}
	_, err = buf.ReadFrom(reader)
	if err != nil {
		return nil, err
	}

	if err := reader.Close(); err != nil {
		return nil, fmt.Errorf("failed to close blob bytes reader, %v", err)
	}

	return buf.Bytes(), nil
}

func (c *ContainerConn) ListBlobs(ctx context.Context, prefix string) ([]string, error) {
	var blobNames []string

	resp, err := c.container.ListBlobs(storage.ListBlobsParameters{
		Prefix:  prefix,
		Timeout: 15, // seconds
	})
	if err != nil {
		return nil, err
	}

	for _, blob := range resp.Blobs {
		blobNames = append(blobNames, blob.Name)
	}

	return blobNames, nil
}

func (c *ContainerConn) ListBlobsByPattern(ctx context.Context, pattern string) ([]string, error) {
	wildcardParts := strings.Split(pattern, "*")
	dirPrefix := path.Dir(wildcardParts[0])

	// Fix to work with files not in a directory
	if dirPrefix == "." {
		dirPrefix = ""
	}

	blobNames, err := c.ListBlobs(ctx, dirPrefix)
	if err != nil {
		return nil, err
	}

	var matchingBlobNames []string
	for _, blobName := range blobNames {
		matched, err := path.Match(pattern, blobName)
		if err != nil {
			return nil, fmt.Errorf("unexpected error when matching patterns: %w", err)
		}
		if matched {
			matchingBlobNames = append(matchingBlobNames, blobName)
		}
	}

	return matchingBlobNames, nil
}

// this method will handle acquire and release of the lease of the file
// if you already have the lease - then send in the leaseID
func (c *ContainerConn) TruncateBlob(ctx context.Context, reader io.Reader, blobName string, leaseId string) error {
	blob := c.container.GetBlobReference(blobName)
	if _, err := blob.DeleteIfExists(&storage.DeleteBlobOptions{
		LeaseID: leaseId,
		Timeout: 15,
	}); err != nil {
		return err
	}

	return c.AppendBlob(ctx, reader, blobName, leaseId)
}

// this method will handle acquire and release of the lease of the file
// if you already have the lease - send in the current leaseID
func (c *ContainerConn) AppendBlob(ctx context.Context, reader io.Reader, blobName string, leaseId string) error {
	releaseLease := false
	leaseStr := leaseId
	blob := c.container.GetBlobReference(blobName)

	exist, err := blob.Exists()
	if err != nil {
		return err
	}

	if !exist {
		if err := blob.PutAppendBlob(&storage.PutBlobOptions{}); err != nil {
			return err
		}
	}

	if leaseStr == "" {
		releaseLease = true
		leaseStr, err = c.AcquireLease(ctx, blobName)
		if err != nil {
			return err
		}
	}

	buf := make([]byte, 1*1024*1024)
	for {
		_, err := reader.Read(buf)
		if err == io.EOF {
			break
		}

		if err := blob.AppendBlock(buf, &storage.AppendBlockOptions{
			LeaseID: leaseStr}); err != nil {
			return err
		}
	}

	if releaseLease {
		return c.ReleaseLease(ctx, blobName, leaseStr)
	}

	return nil
}

// returns the leaseId for the file, and error if one occurred
func (c *ContainerConn) AcquireLease(ctx context.Context, blobName string) (string, error) {
	blob := c.container.GetBlobReference(blobName)
	randomLeaseId := uuid.New()

	leaseId, err := blob.AcquireLease(-1, randomLeaseId.String(), &storage.LeaseOptions{
		Timeout: 15,
	})

	if err != nil {
		return "", ParseAzureError(err)
	}

	return leaseId, nil
}

func (c *ContainerConn) ReleaseLease(ctx context.Context, blobName string, leaseId string) error {
	blob := c.container.GetBlobReference(blobName)

	return blob.ReleaseLease(leaseId, &storage.LeaseOptions{
		Timeout: 15,
	})
}
