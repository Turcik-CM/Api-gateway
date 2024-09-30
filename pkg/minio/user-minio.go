package minio

import (
	pb "api-gateway/genproto/user"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"net/http"

	"log"
)

var MinioClient *minio.Client
var BucketName = "profile-image"
var Endpoint = "3.120.111.217:9001"

func InitUser() error {
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"

	minioClient, err := minio.New(Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Println(err)
		return err
	}

	MinioClient = minioClient

	return nil
}

func AAAA() {
	file, err := req.Open()
	if err != nil {
		h.logger.Error("Error occurred while opening file", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	_, err = MinioClient.PutObject(context.Background(), BucketName, req.Filename, file, req.Size, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		h.logger.Error("Error occurred while uploading file", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imageUrl := fmt.Sprintf("http://%s/%s/%s", Endpoint, BucketName, req.Filename)

	change := pb.URL{
		UserId: c.MustGet("user_id").(string),
		Url:    imageUrl,
	}

	log.Println(imageUrl)
	log.Println(imageUrl)
	log.Println(imageUrl)

}
