package handlemedia

import (
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context, image *multipart.FileHeader, fileIdentifier string) (*uploader.UploadResult, error) {
	var res *uploader.UploadResult

	err := c.SaveUploadedFile(image, image.Filename)
	if err != nil {
		return res, err
	}

	resp, err := UploadToCloudinary(c, image, fileIdentifier)

	os.Remove(image.Filename)

	if err != nil {
		return res, err
	}

	return resp, nil

}
