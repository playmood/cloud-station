package aliyun

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/playmood/cloud-station/store"
)

var (
	// 对象是否实现接口的强制约束
	_ store.Uploader = &AliOssStore{}
)

// 构造函数
func NewAliOssStore(endpoint, accessKey, accessSecret string) (*AliOssStore, error) {
	c, err := oss.New(endpoint, accessKey, accessSecret)
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
