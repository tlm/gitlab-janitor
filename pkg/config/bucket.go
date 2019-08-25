package config

import (
	"context"
	"errors"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/azureblob"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/s3blob"
)

type Bucket struct {
	URL string `json:"url" yaml:"url"`
}

func NewBucket() *Bucket {
	return &Bucket{}
}

func ToBucket(conf *Bucket) (*blob.Bucket, error) {
	return ToBucketWithContext(context.Background(), conf)
}

func ToBucketWithContext(ctx context.Context, conf *Bucket) (*blob.Bucket, error) {
	if conf.URL == "" {
		return nil, errors.New("bucket url cannot be null")
	}

	return blob.OpenBucket(ctx, conf.URL)
}
