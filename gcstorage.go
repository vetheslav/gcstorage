package gcstorage

import (
	"context"
	"io/ioutil"
	"sync"

	"cloud.google.com/go/storage"
)

var (
	storageClient *storage.Client
	storageError  error
	once          sync.Once
)

type StorageBucket struct {
	*storage.BucketHandle
}

func NewBucket(name string) (*StorageBucket, error) {
	once.Do(func() {
		storageClient, storageError = storage.NewClient(context.Background())
	})
	if storageError != nil {
		return nil, storageError
	}

	return &StorageBucket{storageClient.Bucket(name)}, nil
}

func (b *StorageBucket) UploadFile(fileName, contentType string, content []byte, rule ...storage.ACLRule) error {
	ctx := context.Background()
	w := b.Object(fileName).NewWriter(ctx)
	defer w.Close()

	w.ContentType = contentType
	w.ACL = rule
	_, err := w.Write(content)

	return err
}

func (b *StorageBucket) ReadFile(fileName string) ([]byte, error) {
	ctx := context.Background()
	r, err := b.Object(fileName).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return ioutil.ReadAll(r)
}

func Close() error {
	return storageClient.Close()
}
