package aliyun_test

import (
	"github.com/playmood/cloud-station/store"
	"github.com/playmood/cloud-station/store/aliyun"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	uploader store.Uploader
)

var (
	bucketName = os.Getenv("ALI_BUCKET_NAME")
)

// Aliyun OSS Store upload 测试用例
func TestUpload(t *testing.T) {
	// test断言
	// 通过new获取一个断言实例
	should := assert.New(t)
	err := uploader.Upload(bucketName, "test.txt", "store_test.go")
	if should.NoError(err) {
		// 没有error，开始下一步骤
		t.Log("upload ok")
	}
}

func TestUploadError(t *testing.T) {
	should := assert.New(t)
	err := uploader.Upload(bucketName, "test.txt", "store_testxxxxxxxx.go")
	if should.Error(err, "ssssss") {
		t.Log("upload not ok")
	}
}

// uploader实例化逻辑
func init() {
	ali, err := aliyun.NewDefaultAliOssStore()
	if err != nil {
		panic(err)
	}
	uploader = ali
}
