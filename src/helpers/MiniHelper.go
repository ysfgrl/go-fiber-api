package helpers

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go-fiber-api/src/config"
	response2 "go-fiber-api/src/models"
	"go-fiber-api/src/utils/response"
	"mime/multipart"
)

type minioHelper struct {
	BaseHelper[minio.Client]
}

func NewMinioHelper(ctx context.Context) (*minioHelper, *response2.MyError) {
	client, errInit := minio.New(config.Minio.Url, &minio.Options{
		Creds: credentials.NewStaticV4(
			config.Minio.AccessKey,
			config.Minio.SecretKey, "",
		),
		Secure: false,
	})
	if errInit != nil {
		return nil, response.GetError(errInit)
	}

	return &minioHelper{
		BaseHelper[minio.Client]{
			client: client,
			ctx:    context.TODO(),
		},
	}, nil
}

func (m *minioHelper) IsExistBucket(name string) bool {
	exists, err := m.client.BucketExists(m.ctx, name)
	if err == nil && exists {
		return true
	}
	return false
}

func (m *minioHelper) NewBucket(name string) *response2.MyError {

	if m.IsExistBucket(name) {
		return nil
	}
	err := m.client.MakeBucket(m.ctx, name, minio.MakeBucketOptions{Region: "us-east-1"})
	if err != nil {
		return response.GetError(err)
	}
	return nil
}

func (m *minioHelper) PutHeaderObject(bucketName string, file *multipart.FileHeader) (minio.UploadInfo, *response2.MyError) {

	buffer, err := file.Open()

	if err != nil {
		return minio.UploadInfo{}, response.GetError(err)
	}
	defer buffer.Close()

	objectName := file.Filename
	fileBuffer := buffer
	contentType := file.Header["Content-Type"][0]
	fileSize := file.Size

	info, err1 := m.client.PutObject(m.ctx, bucketName, "test/"+objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})

	if err1 != nil {
		return minio.UploadInfo{}, response.GetError(err1)
	}
	return info, nil
}
