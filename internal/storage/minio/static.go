package minio

import (
	"context"

	"github.com/minio/minio-go/v7"
)

const (
	UserAvatarsBucket = "user-avatars"
	AudioFilesBucket  = "audio-files"
)

type StaticStorage struct {
	bucketName  string
	MinioClient *minio.Client
}

func NewStaticStorage(ctx context.Context, minioClient *minio.Client, bucketName string) (storage *StaticStorage, err error) {
	bucketExists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return
	}

	if !bucketExists {
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{
			Region: "ru-central1",
		})
		if err != nil {
			return
		}

		policy := `{
        "Version": "2012-10-17",
        "Statement": [
            {
                "Sid": "PublicRead",
                "Effect": "Allow",
                "Principal": "*",
                "Action": ["s3:GetObject"],
                "Resource": ["arn:aws:s3:::` + bucketName + `/*"]
            }
        ]
    }`

		err = minioClient.SetBucketPolicy(ctx, bucketName, policy)
		if err != nil {
			return
		}
	}

	storage = &StaticStorage{
		bucketName:  bucketName,
		MinioClient: minioClient,
	}
	return
}

func (a *StaticStorage) Store(ctx context.Context, fileName string, filePath string, contentType string) (err error) {
	_, err = a.MinioClient.FPutObject(ctx, a.bucketName, fileName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return
	}

	return
}

func (a *StaticStorage) Delete(ctx context.Context, fileName string) (err error) {
	err = a.MinioClient.RemoveObject(ctx, a.bucketName, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return
	}

	return
}
