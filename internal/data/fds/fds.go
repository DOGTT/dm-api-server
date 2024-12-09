package fds

import (
	"bytes"
	"context"
	"net/http"
	"time"

	"github.com/DOGTT/dm-api-server/internal/conf"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	BucketList = []string{}
)

const (
	BucketNameAvatar = "avatar"
)

func init() {
	BucketList = append(BucketList,
		BucketNameAvatar,
	)
}

// 文件型数据
type FDSClient struct {
	c   *conf.FDSConfig
	svc *s3.S3
}

func New(c *conf.FDSConfig) (fc *FDSClient, err error) {
	// 读取配置文件

	// 创建新的 session
	sess, err := session.NewSession(&aws.Config{
		Endpoint:         aws.String(c.Endpoint),
		Region:           aws.String("us-east-1"), // MinIO 不需要真实的区域，但需要设置
		Credentials:      credentials.NewStaticCredentials(c.AccessKey, c.SecretKey, ""),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return nil, err
	}
	// 创建 S3 客户端
	svc := s3.New(sess)
	// 创建 bucket
	fdsClient := &FDSClient{
		c:   c,
		svc: svc,
	}
	if err := fdsClient.initCreateBucket(); err != nil {
		return nil, err
	}
	return fdsClient, nil
}

// createBucket 创建 bucket
func (c *FDSClient) initCreateBucket() error {
	for i := range BucketList {
		_, err := c.svc.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(BucketList[i]),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *FDSClient) PutObject(ctx context.Context, bucket string, fileID string, fileBytes []byte) error {
	contentType := detectContentType(fileBytes)
	_, err := c.svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(fileID),
		Body:        bytes.NewReader(fileBytes),
		ContentType: aws.String(contentType),
	})
	return err
}

// GeneratePresignedURL 生成预签名 URL
func (c *FDSClient) GeneratePresignedURL(ctx context.Context, bucket string, fileID string, expires time.Duration) (string, error) {
	req, _ := c.svc.GetObjectRequest(&s3.GetObjectInput{
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
