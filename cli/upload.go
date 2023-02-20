package cli

import (
	"fmt"
	"github.com/playmood/cloud-station/store"
	"github.com/playmood/cloud-station/store/aliyun"
	"github.com/playmood/cloud-station/store/aws"
	"github.com/playmood/cloud-station/store/tx"
	"github.com/spf13/cobra"
)

var (
	ossProvider     string
	ossEndpoint     string
	ossAccessKey    string
	ossAccessSecret string
	ossBucketName   string
	uploadFile      string
)

var UploadCmd = &cobra.Command{
	Use:     "upload",
	Long:    "upload 文件上传",
	Short:   "upload 文件上传",
	Example: "upload -f filename",
	RunE: func(cmd *cobra.Command, args []string) error {
		var uploader store.Uploader
		var err error
		switch ossProvider {
		case "aliyun":
			uploader, err = aliyun.NewAliOssStore(&aliyun.Options{
				Endpoint:     ossEndpoint,
				AccessSecret: ossAccessSecret,
				AccessKey:    ossAccessKey,
			})
		case "tx":
			uploader = tx.NewAwsOssStore()
		case "aws":
			uploader = aws.NewAwsOssStore()
		default:
			return fmt.Errorf("not support oss storage")
		}
		if err != nil {
			return err
		}
		// 使用uploader上传文件
		return uploader.Upload(ossBucketName, uploadFile, uploadFile)
	},
}

func init() {
	f := UploadCmd.PersistentFlags()
	f.StringVarP(&ossProvider, "provider", "p", "aliyun", "oss storage provider [aliyun/tx/aws]")
	f.StringVarP(&ossEndpoint, "endpoint", "e", "oss-cn-beijing.aliyuncs.com", "oss storage provider endpoint")
	f.StringVarP(&ossAccessKey, "access_key", "k", "", "oss storage provider ak")
	f.StringVarP(&ossAccessSecret, "access_secret", "s", "", "oss storage provider sk")
	f.StringVarP(&ossBucketName, "bucket_name", "b", "devcloud-station-1", "oss storage provider bucket name")
	f.StringVarP(&uploadFile, "upload_file", "f", "", "upload file name")
	RootCmd.AddCommand(UploadCmd)
}
