package cloudinaryservice

import (
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
)

func UploadToCloudinary(c *gin.Context, image *multipart.FileHeader, fileIdentifier string) (*uploader.UploadResult, error) {
	CLOUDINARYCLOUDNAME := os.Getenv("CLOUDINARY_CLOUD_NAME")
	CLOUDINARYAPIKEY := os.Getenv("CLOUDINARY_API_KEY")
	CLOUDINARYAPISECRET := os.Getenv("CLOUDINARY_API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUDINARYCLOUDNAME, CLOUDINARYAPIKEY, CLOUDINARYAPISECRET)
	return cld.Upload.Upload(c, image.Filename, uploader.UploadParams{PublicID: fileIdentifier})
}
