package clients

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/ysfgrl/go-fiber-api/src/config"
)

func initMinio(ctx context.Context) (*minio.Client, error) {
	client, errInit := minio.New(config.AppConf.Minio.Url, &minio.Options{
		Creds: credentials.NewStaticV4(
			config.AppConf.Minio.AccessKey,
			config.AppConf.Minio.SecretKey, "",
		),
		Secure: false,
	})
	if errInit != nil {
		return nil, errInit
	}

	return client, nil
}

//func (m *minioHelper) IsExistBucket(name string) bool {
//	exists, err := m.client.BucketExists(m.ctx, name)
//	if err == nil && exists {
//		return true
//	}
//	return false
//}
//
//func (m *minioHelper) NewBucket(name string) *response2.Error {
//
//	if m.IsExistBucket(name) {
//		return nil
//	}
//	err := m.client.MakeBucket(m.ctx, name, minio.MakeBucketOptions{Region: "us-east-1"})
//	if err != nil {
//		return response.GetError(err)
//	}
//	return nil
//}
//
//func (m *minioHelper) PutHeaderObject(bucketName string, file *multipart.FileHeader) (minio.UploadInfo, *response2.Error) {
//
//	buffer, err := file.Open()
//
//	if err != nil {
//		return minio.UploadInfo{}, response.GetError(err)
//	}
//	defer buffer.Close()
//
//	objectName := file.Filename
//	fileBuffer := buffer
//	contentType := file.Header["Content-Type"][0]
//	fileSize := file.Size
//
//	info, err1 := m.client.PutObject(m.ctx, bucketName, "test/"+objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})
//
//	if err1 != nil {
//		return minio.UploadInfo{}, response.GetError(err1)
//	}
//	return info, nil
//}
