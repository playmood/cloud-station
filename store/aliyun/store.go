package aliyun

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/playmood/cloud-station/store"
	"os"
)

var (
	// 对象是否实现接口的强制约束
	_ store.Uploader = &AliOssStore{}
)

type Options struct {
	Endpoint     string
	AccessKey    string
	AccessSecret string
}

func (o *Options) Validate() error {
	if o.Endpoint == "" || o.AccessKey == "" || o.AccessSecret == "" {
		return fmt.Errorf("endpoint, access_key secret_key is empty")
	}
	return nil
}

func NewDefaultAliOssStore() (*AliOssStore, error) {
	endPoint := os.Getenv("ALI_OSS_ENDPOINT")
	accessKey := os.Getenv("ALI_AK")
	secretKey := os.Getenv("ALI_SK")
	return NewAliOssStore(&Options{
		Endpoint:     endPoint,
		AccessKey:    accessKey,
		AccessSecret: secretKey,
	})
}

// 构造函数
func NewAliOssStore(opts *Options) (*AliOssStore, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}
	c, err := oss.New(opts.Endpoint, opts.AccessKey, opts.AccessSecret)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{
		client: c,
	}, nil
}

type AliOssStore struct {
	client *oss.Client
}

func (s *AliOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	// 获取bucket对象
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return err
	}
	// 上传文件文件到bucket
	if err := bucket.PutObjectFromFile(objectKey, fileName); err != nil {
		return err
	}
	// 打印下载链接
	downloadURL, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("download URL for one day: %s\n", downloadURL)
	return nil
}
