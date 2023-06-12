package file

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/models"
	"go-fiber-api/myUtils/response"
)

const (
	TempUploadPath   = "./TEMP/"
	TempUploadBucket = "./TEMP/"
)

func UploadTmpFile(c *fiber.Ctx, key string) (*models.UploadedFile, *models.MyError) {

	file, err := c.FormFile(key)
	if err != nil {
		return nil, response.GetError(err)
	}

	//if err := c.SaveFile(file, TempUploadPath+file.Filename); err != nil {
	//	return nil, response.GetError("error", err)
	//}
	//

	return &models.UploadedFile{
		Name: file.Filename,
		Size: file.Size,
		Type: file.Header["Content-Type"][0],
	}, nil
}
