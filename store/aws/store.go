package aws

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/playmood/cloud-station/store"
)

var (
	// 对象是否实现接口的强制约束
	_ store.Uploader = &AwsOssStore{}
)

// 构造函数
func NewAwsOssStore() *AwsOssStore {
	return &AwsOssStore{}
}

type AwsOssStore struct {
	client *oss.Client
}

func (s *AwsOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	return nil
}
