package service

import (
	"car-rental-backend/pkg/minio"
	"context"
	"fmt"
	"log/slog"
	"mime/multipart"
	"time"

	miniogo "github.com/minio/minio-go/v7"
)

type MediaService struct {
	minioClient *minio.MinioClient
}

func NewMediaService(minioClient *minio.MinioClient) *MediaService {
	return &MediaService{minioClient: minioClient}
}

func (s *MediaService) Upload(collection string, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
	path := collection + "/" + filename
	ctx := context.Background()
	_, err = s.minioClient.Client.PutObject(
		ctx,
		s.minioClient.Bucket,
		path,
		src,
		file.Size,
		miniogo.PutObjectOptions{
			ContentType: file.Header.Get("Content-Type"),
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to minio: %w", err)
	}

	url := fmt.Sprintf("%s/%s/%s", s.minioClient.PublicURL, s.minioClient.Bucket, path)
	return url, nil
}

func (s *MediaService) Delete(collection, filename string) error {
	path := collection + "/" + filename
	slog.Info("DELETING IMAGE", "filename", path)
	ctx := context.Background()
	err := s.minioClient.Client.RemoveObject(ctx, s.minioClient.Bucket, path, miniogo.RemoveObjectOptions{})
	if err != nil {
		slog.Error("FAILED TO DELETE IMAGE", "error", err.Error())
		return fmt.Errorf("failed to delete file from minio: %w", err)
	}
	slog.Info("DELETED IMAGE", "filename", path)
	return nil
}
