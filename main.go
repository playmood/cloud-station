package main

import (
	"flag"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

var (
	endPoint   = os.Getenv("ALI_OSS_ENDPOINT")
	accessKey  = os.Getenv("ALI_AK")
	secretKey  = os.Getenv("ALI_SK")
	bucketName = os.Getenv("ALI_BUCKET_NAME")
	uploadFile = ""
)

// 实现文件上传
func upload(filePath string) error {
	// 实例化客户端
	client, err := oss.New(endPoint, accessKey, secretKey)
	if err != nil {
		return err
	}

	// 获取bucket对象
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}
	// 上传文件文件到bucket
	return bucket.PutObjectFromFile(filePath, filePath)
}

func validate() error {
	if endPoint == "" || accessKey == "" || secretKey == "" {
		return fmt.Errorf("endpoint, access_key secret_key is empty")
	}
	if uploadFile == "" {
		return fmt.Errorf("upload file required")
	}
	return nil
}

func loadParams() {
	flag.StringVar(&uploadFile, "f", "", "give the name of upload file")
	flag.Parse()
}

func main() {
	// 参数加载
	loadParams()

	if err := validate(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := upload(uploadFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("upload %s successful\n", uploadFile)

}
