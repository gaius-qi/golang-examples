package main

import (
	"bytes"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func main() {
	endpoint := ""
	accessKey := ""
	secretKey := ""
	bucketName := ""
	objectKey := "test"

	client, err := oss.New(endpoint, accessKey, secretKey)
	if err != nil {
		panic(err)
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		panic(err)
	}

	data := []byte("test")
	reader := bytes.NewReader(data)

	_, err = reader.Seek(1, 0)
	if err != nil {
		panic(err)
	}

	if err := bucket.PutObject(objectKey, reader); err != nil {
		panic(err)
	}
}
