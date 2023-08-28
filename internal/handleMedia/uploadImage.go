package handlemedia

import (
	storjservice "blog/pkg/cloudStorageServices/storjService"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context, image *multipart.FileHeader, fileIdentifier string) (string, error) {

	err := c.SaveUploadedFile(image, image.Filename)
	if err != nil {
		return "", err
	}

	// resp, err := cloudinaryservice.UploadToCloudinary(c, image, fileIdentifier)
	url, err := storjservice.SaveImageToStorj(image.Filename, fileIdentifier)
	os.Remove(image.Filename)

	if err != nil {
		return "", err
	}

	return url, nil

}
