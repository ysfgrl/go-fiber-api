package helpers

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go-fiber-api/config"
	"go-fiber-api/models"
	"go-fiber-api/myUtils/response"
	"mime/multipart"
)

type MinioHelper struct {
	Config *config.Config
	Ctx    context.Context
	client *minio.Client
}

func (m *MinioHelper) Connect() *models.MyError {
	minioClient, errInit := minio.New(m.Config.MinioUrl, &minio.Options{
		Creds: credentials.NewStaticV4(
			m.Config.MinioAccessKey,
			m.Config.MinioSecretKey, "",
		),
		Secure: false,
	})
	if errInit != nil {
		return response.GetError(errInit)
	}
	m.client = minioClient
	return nil
}

func (m *MinioHelper) IsExistBucket(name string) bool {
	exists, err := m.client.BucketExists(m.Ctx, name)
	if err == nil && exists {
		return true
	}
	return false
}

func (m *MinioHelper) NewBucket(name string) *models.MyError {

	if m.IsExistBucket(name) {
		return nil
	}
	err := m.client.MakeBucket(m.Ctx, name, minio.MakeBucketOptions{Region: "us-east-1"})
	if err != nil {
		return response.GetError(err)
	}
	return nil
}

func (m *MinioHelper) PutHeaderObject(bucketName string, file *multipart.FileHeader) (minio.UploadInfo, *models.MyError) {

	buffer, err := file.Open()

	if err != nil {
		return minio.UploadInfo{}, response.GetError(err)
	}
	defer buffer.Close()

	objectName := file.Filename
	fileBuffer := buffer
	contentType := file.Header["Content-Type"][0]
	fileSize := file.Size

	info, err1 := m.client.PutObject(m.Ctx, bucketName, objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})

	if err1 != nil {
		return minio.UploadInfo{}, response.GetError(err1)
	}
	return info, nil
}
