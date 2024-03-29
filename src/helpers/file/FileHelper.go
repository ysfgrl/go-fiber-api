package file

import (
	"github.com/gofiber/fiber/v2"
	response2 "github.com/ysfgrl/go-fiber-api/src/models"
	"github.com/ysfgrl/go-fiber-api/src/pkg/response"
)

const (
	TempUploadPath   = "./TEMP/"
	TempUploadBucket = "./TEMP/"
)

func UploadTmpFile(c *fiber.Ctx, key string) (*response2.UploadedFile, *response2.Error) {

	file, err := c.FormFile(key)
	if err != nil {
		return nil, response.GetError(err)
	}

	//if err := c.SaveFile(file, TempUploadPath+file.Filename); err != nil {
	//	return nil, response.GetError("error", err)
	//}
	//

	return &response2.UploadedFile{
		Name: file.Filename,
		Size: file.Size,
		Type: file.Header["Content-Type"][0],
	}, nil
}
