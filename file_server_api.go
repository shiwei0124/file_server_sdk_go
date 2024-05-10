package file_server_sdk_go

import (
	"context"
	"io"
)

type FileServerApi interface {
	//Login
	Login(accessKeyId string, accessKeySecret string) error
	// PutObject
	PutObject(ctx context.Context, objectKey string, reader io.Reader) error
	// PutLargeObject
	PutLargeObject(ctx context.Context, objectKey string, reader io.Reader) error
	// GetObject
	GetObject(ctx context.Context, objectKey string) (io.ReadCloser, int64, error)
}

type FileServerApiImpl struct {
	token string
}

func NewFileServerApi() FileServerApi {
	return &FileServerApiImpl{}
}

func (f *FileServerApiImpl) Login(accessKeyId string, accessKeySecret string) error {
	f.token = "token"
	return nil
}

func (f *FileServerApiImpl) PutObject(ctx context.Context, objectKey string, reader io.Reader) error {

	return nil
}

func (f *FileServerApiImpl) PutLargeObject(ctx context.Context, objectKey string, reader io.Reader) error {
	return nil
}

func (f *FileServerApiImpl) GetObject(ctx context.Context, objectKey string) (io.ReadCloser, int64, error) {
	return nil, 0, nil
}
