package tx

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/playmood/cloud-station/store"
)

var (
	// 对象是否实现接口的强制约束
	_ store.Uploader = &TxOssStore{}
)

// 构造函数
func NewAwsOssStore(endpoint, accessKey, accessSecret string) (*TxOssStore, error) {
	c, err := oss.New(endpoint, accessKey, accessSecret)
	if err != nil {
		return nil, err
	}
	return &TxOssStore{
		client: c,
	}, nil
}

type TxOssStore struct {
	client *oss.Client
}

func (s *TxOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	return nil
}
