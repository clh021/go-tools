package minio

import (
	"context"
	"log"
	"time"

	"gitee.com/linakesi/source-analysis-tools-ui/cli/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioService struct {
	conf   *config.MinioConfig
	client *minio.Client
}

func New(c config.MinioConfig) *MinioService {
	m := &MinioService{
		conf: &c,
	}
	return m
}

func (m MinioService) InitClient() *minio.Client {
	var err error
	// Initialize minio client object.
	m.client, err = minio.New(m.conf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.conf.AccessKeyID, m.conf.SecretAccessKey, ""),
		Secure: m.conf.UseSSL,
	})
	if err != nil {
		log.Fatalf("conetct minio server fail %s url %s ", err.Error(), m.conf.Endpoint)
	}
	m.client.HealthCheck(time.Microsecond*5)
	log.Printf("isonline: %#v\n", m.client.IsOnline())
	ctx := context.Background()
	err = m.client.MakeBucket(ctx, m.conf.BucketName, minio.MakeBucketOptions{Region: m.conf.Location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := m.client.BucketExists(ctx, m.conf.BucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", m.conf.BucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", m.conf.BucketName)
	}

	log.Printf("client  : %#v\n", m.client) // minioClient is now set up
	log.Printf("client EndpointURL  : %#v\n", m.client.EndpointURL()) // minioClient is now set up
	return m.client
}

func (m MinioService) PrintEndpointURL() {
	log.Printf("client EndpointURL  : %#v\n", m.client.EndpointURL()) // minioClient is now set up
}

func (m MinioService) Put(filePath string, objectName string, contentType string) (minio.UploadInfo, error) {
	ctx := context.Background()
	c := m.InitClient()
	// log.Printf("client EndpointURL  : %#v\n", m.client.EndpointURL()) // minioClient is now set up
	// Upload the zip file with FPutObject
	// info, err := m.client.FPutObject(ctx, m.conf.BucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	return c.FPutObject(ctx, m.conf.BucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	// FPutObject 入参是文件路径
	// PutObject 入参是 file.Open()
}
