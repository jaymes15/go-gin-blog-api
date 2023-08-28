// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package storjservice

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"storj.io/uplink"
	"storj.io/uplink/edge"
)

const defaultExpiration = 7 * 24 * time.Hour

// UploadAndDownloadData uploads the specified data to the specified key in the
// specified bucket, using the specified Satellite, API key, and passphrase.
func UploadAndDownloadData(ctx context.Context,
	accessGrant, bucketName, uploadKey string, dataToUpload []byte) error {

	// Parse access grant, which contains necessary credentials and permissions.
	access, err := uplink.ParseAccess(accessGrant)
	if err != nil {
		return fmt.Errorf("could not request access grant: %v", err)
	}

	// Open up the Project we will be working with.
	project, err := uplink.OpenProject(ctx, access)
	if err != nil {
		return fmt.Errorf("could not open project: %v", err)
	}
	defer project.Close()

	// Ensure the desired Bucket within the Project is created.
	_, err = project.EnsureBucket(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("could not ensure bucket: %v", err)
	}

	// Intitiate the upload of our Object to the specified bucket and key.
	upload, err := project.UploadObject(ctx, bucketName, uploadKey, &uplink.UploadOptions{
		// It's possible to set an expiration date for data.
		Expires: time.Now().Add(defaultExpiration),
	})
	if err != nil {
		return fmt.Errorf("could not initiate upload: %v", err)
	}

	// Copy the data to the upload.
	buf := bytes.NewBuffer(dataToUpload)
	_, err = io.Copy(upload, buf)
	if err != nil {
		_ = upload.Abort()
		return fmt.Errorf("could not upload data: %v", err)
	}

	// Commit the uploaded object.
	err = upload.Commit()
	if err != nil {
		return fmt.Errorf("could not commit uploaded object: %v", err)
	}

	// Initiate a download of the same object again
	download, err := project.DownloadObject(ctx, bucketName, uploadKey, nil)
	if err != nil {
		return fmt.Errorf("could not open object: %v", err)
	}
	defer download.Close()

	// Read everything from the download stream
	receivedContents, err := io.ReadAll(download)
	if err != nil {
		return fmt.Errorf("could not read data: %v", err)
	}

	// Check that the downloaded data is the same as the uploaded data.
	if !bytes.Equal(receivedContents, dataToUpload) {
		return fmt.Errorf("got different object back: %q != %q", dataToUpload, receivedContents)
	}

	return nil
}

func CreatePublicSharedLink(ctx context.Context, accessGrant, bucketName, objectKey string) (string, error) {
	// Define configuration for the storj sharing site.
	config := edge.Config{
		AuthServiceAddress: "auth.storjshare.io:7777",
	}

	// Parse access grant, which contains necessary credentials and permissions.
	access, err := uplink.ParseAccess(accessGrant)
	if err != nil {
		return "", fmt.Errorf("could not parse access grant: %w", err)
	}

	// Restrict access to the specified paths.
	restrictedAccess, err := access.Share(
		uplink.Permission{
			// only allow downloads
			AllowDownload: true,
			// this allows to automatically cleanup the access grants
			NotAfter: time.Now().Add(defaultExpiration),
		}, uplink.SharePrefix{
			Bucket: bucketName,
			Prefix: objectKey,
		})
	if err != nil {
		return "", fmt.Errorf("could not restrict access grant: %w", err)
	}

	// RegisterAccess registers the credentials to the linksharing and s3 sites.
	// This makes the data publicly accessible, see the security implications in https://docs.storj.io/dcs/concepts/access/access-management-at-the-edge.
	credentials, err := config.RegisterAccess(ctx, restrictedAccess, &edge.RegisterAccessOptions{Public: true})
	if err != nil {
		return "", fmt.Errorf("could not register access: %w", err)
	}

	// Create a public link that is served by linksharing service.
	url, err := edge.JoinShareURL("https://link.storjshare.io", credentials.AccessKeyID, bucketName, objectKey, nil)
	if err != nil {
		return "", fmt.Errorf("could not create a shared link: %w", err)
	}

	return url, nil
}

func SaveImageToStorj(imagePath string, objectKeyName string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}

	ctx := context.Background()
	accessGrant := flag.String("access", os.Getenv("STORJ_ACCESS_GRANT"), "access grant from satellite")
	flag.Parse()

	bucketName := os.Getenv("STORJ_BUCKET_NAME")
	objectKey := fmt.Sprintf("images/%s", objectKeyName)

	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return "", fmt.Errorf("could not read image file: %w", err)
	}

	err = UploadAndDownloadData(ctx, *accessGrant, bucketName, objectKey, imageData)
	if err != nil {
		return "", fmt.Errorf("upload failed: %w", err)
	}

	url, err := CreatePublicSharedLink(ctx, *accessGrant, bucketName, objectKey)
	if err != nil {
		return "", fmt.Errorf("creating public link failed: %w", err)
	}

	fmt.Println("Image upload and sharing successful!")
	fmt.Println("Public link:", url)

	return url, nil
}
