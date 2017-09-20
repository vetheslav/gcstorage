package gcstorage

import (
	"testing"

	"bytes"

	"golang.org/x/net/context"
)

const projectID = "dailyfrenzy-product"

func TestGCStorageGetBucket(t *testing.T) {
	bucket, err := NewBucket("sashalala")
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	err = bucket.Create(ctx, projectID, nil)
	if err != nil {
		t.Fatal(err)
	}
	err = bucket.Delete(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGCStorageUploadFile(t *testing.T) {
	bucket, err := NewBucket("sashalala")
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	err = bucket.Create(ctx, projectID, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = bucket.UploadFile("test.txt", "text/plain", []byte("hello, world"))
	if err != nil {
		t.Error(err)
	}

	err = bucket.Object("test.txt").Delete(ctx)
	if err != nil {
		t.Error(err)
	}

	err = bucket.Delete(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGCStorageReadFile(t *testing.T) {
	bucket, err := NewBucket("sashalala")
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	err = bucket.Create(ctx, projectID, nil)
	if err != nil {
		t.Fatal(err)
	}

	fileContent := []byte("hello, world")
	err = bucket.UploadFile("test.txt", "text/plain", fileContent)
	if err != nil {
		t.Error(err)
	}

	data, err := bucket.ReadFile("test.txt")
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(data, fileContent) {
		t.Errorf("read wrong data. assuming %s, got %s", fileContent, data)
	}

	err = bucket.Object("test.txt").Delete(ctx)
	if err != nil {
		t.Error(err)
	}

	err = bucket.Delete(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
