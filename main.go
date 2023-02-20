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
	help       = false
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
	if err := bucket.PutObjectFromFile(filePath, filePath); err != nil {
		return err
	}
	// 打印下载链接
	downloadURL, err := bucket.SignURL(filePath, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("download URL for one day: %s\n", downloadURL)
	return nil
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
	flag.BoolVar(&help, "h", false, "help")
	flag.StringVar(&uploadFile, "f", "", "give the name of upload file")
	flag.Parse()

	if help {
		usage()
	}
}

// 打印使用说明
func usage() {
	fmt.Fprintf(os.Stderr, `cloud-station version: 0.0.1
Usage: cloud-station [-h] -f <uplaod_file_path>
Options:
`)
	flag.PrintDefaults()
}

func main() {
	// 参数加载
	loadParams()

	if err := validate(); err != nil {
		fmt.Println(err)
		usage()
		os.Exit(1)
	}
	if err := upload(uploadFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("upload %s successful\n", uploadFile)

}
