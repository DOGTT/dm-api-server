package fds

import (
	"bytes"
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (c *FDSClient) PutObject(ctx context.Context, bucket string, fileID string, fileBytes []byte) error {
	contentType := detectContentType(fileBytes)
	_, err := c.s3api.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(fileID),
		Body:        bytes.NewReader(fileBytes),
		ContentType: aws.String(contentType),
	})
	return err
}

// GenerateGetPresignedURL 生成预签名 URL
func (c *FDSClient) GenerateGetPresignedURL(ctx context.Context, bucket string, fileID string, expires time.Duration) (string, error) {
	req, _ := c.s3api.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileID),
	})

	url, err := req.Presign(expires)
	if err != nil {
		return "", err
	}
	return url, nil
}

// GeneratePutPresignedURL 生成预签名 Put URL
func (c *FDSClient) GeneratePutPresignedURL(ctx context.Context, bucket string, fileID string, expires time.Duration) (string, error) {
	req, _ := c.s3api.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileID),
	})

	url, err := req.Presign(expires)
	if err != nil {
		return "", err
	}
	return url, nil
}

// detectContentType 自动检测文件类型
func detectContentType(fileBytes []byte) string {
	return http.DetectContentType(fileBytes)
}
