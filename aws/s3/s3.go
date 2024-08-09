package s3

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Config struct {
	Bucket string
}

type Client struct {
	bucket        string
	instance      *s3.Client
	presignClient *s3.PresignClient
}

func New(c Config) *Client {
	// Read the aws configuration file, the default path is ~/.aws.
	cfg, err := awscfg.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	return &Client{
		bucket:        c.Bucket,
		instance:      client,
		presignClient: s3.NewPresignClient(client),
	}
}

func (c *Client) PresignPutObject(key string) (*PresignRequest, error) {
	req, err := c.presignClient.PresignPutObject(context.Background(), &s3.PutObjectInput{
		Bucket: &c.bucket,
		Key:    &key,
	}, func(po *s3.PresignOptions) {
		// Signature valid for 5 minutes.
		po.Expires = 5 * time.Minute
	})
	if err != nil {
		return nil, fmt.Errorf("failed to presign request: %w", err)
	}

	return &PresignRequest{
		URL: req.URL,
	}, nil
}

func (c *Client) PutObject(key string, content []byte) (*Object, error) {
	output, err := c.instance.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: &c.bucket,
		Key:    &key,
		Body:   bytes.NewReader(content),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to put object: %w", err)
	}

	return &Object{
		Key:  key,
		ETag: *output.ETag,
	}, nil
}
