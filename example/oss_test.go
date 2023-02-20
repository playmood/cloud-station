package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"testing"
)

var (
	// 全局client实例
	client *oss.Client
)

var (
	AccessKey    = os.Getenv("ALI_AK")
	AccessSecret = os.Getenv("ALI_SK")
	OssEndpoint  = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName   = os.Getenv("ALI_BUCKET_NAME")
)

// 测试阿里云OssSDK BucketList接口
func TestBucketList(t *testing.T) {

	lsRes, err := client.ListBuckets()
	if err != nil {
		t.Log(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
}

// 测试上传
func TestUploadFile(t *testing.T) {
	bucket, err := client.Bucket("my-bucket")
	if err != nil {
		t.Log(err)
	}

	err = bucket.PutObjectFromFile("my-object", "LocalFile")
	if err != nil {
		t.Log(err)
	}
}

// 初始化一个OSS Client，等下给所有测试用例使用
func init() {

	c, err := oss.New(OssEndpoint, AccessKey, AccessSecret)
	if err != nil {
		panic(err)
	}
	client = c
}
