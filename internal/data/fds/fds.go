package fds

import (
	"fmt"
	"time"

	grpc_api "github.com/DOGTT/dm-api-server/api/grpc"
	"github.com/DOGTT/dm-api-server/internal/conf"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	log "github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

const (
	BucketNameDefault   = "default"
	BucketNameAvatar    = "avatar"
	BucketNamePofpImage = "pofpImage"

	PreSignDurationDefault = time.Minute * 10
)

var (
	bucketInitList = []string{
		BucketNameDefault,
		BucketNameAvatar,
		BucketNamePofpImage}

	bucketNameMapForObjectType = map[grpc_api.ObjectType]string{
		grpc_api.ObjectType_OT_DEFAULT:    BucketNameDefault,
		grpc_api.ObjectType_OT_POFP_IMAGE: BucketNamePofpImage,
	}
)

func init() {
}

func GetBucketName(objectType grpc_api.ObjectType) string {
	name := bucketNameMapForObjectType[objectType]
	if name == "" {
		log.L().Warn("objectType not found in bucketNameMapForObjectType",
			zap.String("objectType", objectType.String()))
		return BucketNameDefault
	}
	return name
}

// 文件型数据
type FDSClient struct {
	c     *conf.FDSConfig
	s3api *s3.S3
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
	s3api := s3.New(sess)
	// 创建 bucket
	fdsClient := &FDSClient{
		c:     c,
		s3api: s3api,
	}
	if err := fdsClient.initCreateBucket(); err != nil {
		return nil, err
	}
	return fdsClient, nil
}

// createBucket 创建 bucket
func (c *FDSClient) initCreateBucket() error {
	for i := range bucketInitList {
		_, err := c.s3api.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(bucketInitList[i]),
		})
		if err != nil {
			aerr, ok := err.(awserr.Error)
			if !ok {
				return err
			}
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				fmt.Printf("Bucket %s already exists\n", bucketInitList[i])
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				fmt.Printf("Bucket %s already owned by you\n", bucketInitList[i])
			default:
				return err
			}
		}
	}
	return nil
}
