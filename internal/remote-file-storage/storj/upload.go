package storj

import (
	"bytes"
	"context"

	"github.com/minio/minio-go/v7"
)

// Upload uploads provided data into object with specific name into provided bucket.
func (client *Client) Upload(ctx context.Context, bucket, objectName string, data []byte) error {
	_, err := client.API.PutObject(ctx, bucket, objectName, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return MinioError.Wrap(err)
	}

	return nil
}
