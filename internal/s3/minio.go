package s3

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type StorageInterface interface {
	Put(ctx context.Context, key string, reader io.Reader, size int64) error
	Get(ctx context.Context, key string) (io.ReadCloser, error)
	Remove(ctx context.Context, key string) error
}

type Client struct {
	*minio.Client
	Bucket string
}

func New(endpoint, accessKey, secretKey, bucket string) (*Client, error) {
	cli, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	return &Client{Client: cli, Bucket: bucket}, nil
}

func (c *Client) Put(ctx context.Context, key string, reader io.Reader, size int64) error {
	_, err := c.PutObject(ctx, c.Bucket, key, reader, size, minio.PutObjectOptions{})
	return err
}

func (c *Client) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	return c.GetObject(ctx, c.Bucket, key, minio.GetObjectOptions{})
}

func (c *Client) Remove(ctx context.Context, key string) error {
	return c.RemoveObject(ctx, c.Bucket, key, minio.RemoveObjectOptions{})
}