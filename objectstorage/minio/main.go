package main

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	cfg := aws.NewConfig().WithCredentials(credentials.NewStaticCredentials("access key", "secret key", ""))
	s, err := session.NewSession(cfg)
	if err != nil {
		panic(err)
	}

	client := awss3.New(s, cfg.WithRegion("any"), cfg.WithEndpoint("minio.example.com"), cfg.WithS3ForcePathStyle(true), cfg.WithDisableSSL(true))
	resp, err := client.GetObjectWithContext(context.Background(), &awss3.GetObjectInput{
		Bucket: aws.String("d7y"),
		Key:    aws.String("config.toml"),
	})
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
