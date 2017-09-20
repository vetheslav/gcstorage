# Simple Google Cloud Storage Client Library for Go


```go
import "github.com/vetheslav/gcstorage"
```

To install the package on your system,

```
$ go get -u github.com/vetheslav/gcstorage
```

For read one file from your bucket:

```go
storageClient := gcstorage.NewBucket(<BUCKET-NAME>)
if err != nil {
		// TODO: Handle error.
}
defer storageClient.Close()
file, err := storageClient.ReadFile(<FILE-NAME>)
if err != nil {
		// TODO: Handle error.
}
```

For upload one file to your bucket:

```go
storageClient := gcstorage.NewBucket(<BUCKET-NAME>)
if err != nil {
		// TODO: Handle error.
}
defer storageClient.Close()
err := storageClient.UploadFile(<FILE-NAME>, <CONTENT-TYPE>, <FILE-CONTENT>)
if err != nil {
		// TODO: Handle error.
}
```
