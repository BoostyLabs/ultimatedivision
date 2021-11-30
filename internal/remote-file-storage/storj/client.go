package storj

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zeebo/errs"
)

var (
	// MinioError is class for minio errors.
	MinioError = errs.Class("minio")
	// Error is class for remove file storage error.
	Error = errs.Class("remote file storage error")
)

// Config is the setup for a particular client.
type Config struct {
	S3Gateway string
	AccessKey string
	SecretKey string
}

// Client implements basic S3 Client with minio.
type Client struct {
	API *minio.Client
}

// NewClient creates new Client.
func NewClient(conf Config) (*Client, error) {
	opts := &minio.Options{
		Creds: credentials.New(
			&credentials.Static{
				Value: credentials.Value{
					AccessKeyID:     conf.AccessKey,
					SecretAccessKey: conf.SecretKey,
				},
			},
		),
		Secure: true,
		Region: "us-east-1",
	}

	c, err := minio.New(conf.S3Gateway, opts)
	if err != nil {
		return &Client{}, MinioError.Wrap(err)
	}

	return &Client{API: c}, nil
}
