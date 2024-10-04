package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"mime/multipart"

	"log"
)

var MinioClient *minio.Client

var Endpoint = "3.120.111.217:9000"

var UserBucketName = "profile-image"
var PostBucketName = "post-image"
var Nationality = "nationality-image"
var FlagBucketName = "flags"

func InitUserMinio() error {
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

func UploadUser(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		log.Println("1-", err)
		return "", err
	}
	defer file.Close()

	_, err = MinioClient.PutObject(context.Background(), UserBucketName, fileHeader.Filename, file, fileHeader.Size, minio.PutObjectOptions{
		ContentType: "image/png",
	})

	if err != nil {
		log.Println("2-", err)
		return "", err
	}

	imageUrl := fmt.Sprintf("http://%s/%s/%s", Endpoint, UserBucketName, fileHeader.Filename)

	return imageUrl, nil
}

func UploadPost(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		log.Println("1-", err)
		return "", err
	}
	defer file.Close()

	_, err = MinioClient.PutObject(context.Background(), PostBucketName, fileHeader.Filename, file, fileHeader.Size, minio.PutObjectOptions{
		ContentType: "image/png",
	})

	if err != nil {
		log.Println("2-", err)
		return "", err
	}

	imageUrl := fmt.Sprintf("http://%s/%s/%s", Endpoint, PostBucketName, fileHeader.Filename)

	return imageUrl, nil
}

func UploadNationality(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		log.Println("1-", err)
		return "", err
	}
	defer file.Close()

	_, err = MinioClient.PutObject(context.Background(), Nationality, fileHeader.Filename, file, fileHeader.Size, minio.PutObjectOptions{
		ContentType: "image/png",
	})

	if err != nil {
		log.Println("2-", err)
		return "", err
	}

	imageUrl := fmt.Sprintf("http://%s/%s/%s", Endpoint, Nationality, fileHeader.Filename)

	return imageUrl, nil
}

func UploadFlag(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		log.Println("1-", err)
		return "", err
	}

	defer file.Close()

	_, err = MinioClient.PutObject(context.Background(), FlagBucketName, fileHeader.Filename, file, fileHeader.Size, minio.PutObjectOptions{
		ContentType: "image/png",
	})

	if err != nil {
		log.Println("2-", err)
		return "", err
	}

	imageUrl := fmt.Sprintf("http://%s/%s/%s", Endpoint, FlagBucketName, fileHeader.Filename)

	return imageUrl, nil
}
