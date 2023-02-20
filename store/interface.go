package store

// 定义如何上传文件到bucket
// 抽象后，并不关心需要上传到哪个OSS的bucket
type Uploader interface {
	Upload(bucketName string, objectKey string, fileName string) error
}
