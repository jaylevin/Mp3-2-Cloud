package storage

import (
	"log"
	"io"

	/* Vendors */
	"github.com/minio/minio-go"

	/* Internal */
	"github.com/jaylevin/Mp32Cloud/config"
)

type MinioClient struct {
	*minio.Client
}

/*
	See https://github.com/minio/minio-go
	Don't forget to set your access key, secret key and endpoint in config.json!
*/

func NewClient(conf *config.Config) *MinioClient {
	ssl := true

	client, err := minio.New(conf.Endpoint, conf.AccessKey, conf.SecretKey, ssl)
	if err != nil {
		log.Fatal(err)
	}

	return &MinioClient{client}
}

func (client *MinioClient) UploadMp3(bucketName string, fileName string, reader io.Reader, fileSize int64) (int64, error) {
	n, err := client.PutObject(bucketName, fileName, reader, fileSize, minio.PutObjectOptions{ContentType: "audio/mp3"})
	if err != nil {
		return -1, err
	}

	return n, nil
}
