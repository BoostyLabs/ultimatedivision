package storj

import (
	"context"
	"errors"
	"io"
	"io/ioutil"

	"github.com/minio/minio-go/v7"
)

// Download downloads object from specific bucket and returns it as byte slice.
func (client *Client) Download(ctx context.Context, bucket, objectName string, buffer []byte) ([]byte, error) {
	reader, err := client.API.GetObject(ctx, bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, MinioError.Wrap(err)
	}
	defer func() { _ = reader.Close() }()

	n, err := reader.Read(buffer[:cap(buffer)])
	if !errors.Is(err, io.EOF) {
		rest, err := ioutil.ReadAll(reader)
		if errors.Is(err, io.EOF) {
			err = nil
		}
		if err != nil {
			return nil, Error.Wrap(err)
		}
		buffer = append(buffer, rest...)
		n = len(buffer)
	}

	buffer = buffer[:n]
	return buffer, nil
}
