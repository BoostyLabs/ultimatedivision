package remote_file_storage

import (
	"context"
)

// RemoteFileStorage interface to call s3 with minio.
type RemoteFileStorage interface {
	Upload(ctx context.Context, bucket, objectName string, data []byte) error
	Download(ctx context.Context, bucket, objectName string, buffer []byte) ([]byte, error)
}
