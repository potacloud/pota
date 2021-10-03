package v1

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	imagesv1 "github.com/potacloud/pota/api/images/v1"
	"github.com/potacloud/pota/pkg/downloader"
	"github.com/potacloud/pota/pkg/repository"
	"github.com/spf13/viper"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// server is used to implement ImagesServer.
type ImageServer struct {
	imagesv1.UnimplementedImagesServer
	repository.ImageRepository
}

func (s *ImageServer) CreateImage(ctx context.Context, req *imagesv1.CreateImageRequest) (*imagesv1.CreateImageResponse, error) {

	log.Printf("Received CreateImage Request\n")

	// download image from request url
	log.Printf("downloading image from %s", req.Url)
	size, path, err := downloader.DownloadImage(req.Url, viper.GetString("image.path"))
	if err != nil {
		return &imagesv1.CreateImageResponse{
			Message: fmt.Sprintf("error downloading image: %s", err.Error()),
			Image:   nil,
		}, err
	}

	// scaffold new image object
	img := &imagesv1.Image{
		Id:        uuid.New().String(),
		Name:      req.Name,
		Size:      uint32(size),
		Path:      path,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}

	if err := s.Create(ctx, img); err != nil {
		return &imagesv1.CreateImageResponse{
			Message: fmt.Sprintf("error creating image: %s", err.Error()),
			Image:   nil,
		}, err
	}

	return &imagesv1.CreateImageResponse{
		Message: "create image success",
		Image:   img,
	}, nil
}

func (s *ImageServer) ListImage(ctx context.Context, req *imagesv1.ListImageRequest) (*imagesv1.ListImageResponse, error) {

	log.Printf("Received ListImage Request\n")

	images, err := s.List(ctx)
	if err != nil {
		return &imagesv1.ListImageResponse{
			Message: fmt.Sprintf("error listing image: %s", err.Error()),
			Image:   nil,
		}, err
	}

	return &imagesv1.ListImageResponse{
		Message: "list image success",
		Image:   images,
	}, nil
}

func (s *ImageServer) DetailImage(ctx context.Context, req *imagesv1.DetailImageRequest) (*imagesv1.DetailImageResponse, error) {

	log.Printf("Received DetailImage Request\n")

	image, err := s.Detail(ctx, req.Id)
	if err != nil {
		return &imagesv1.DetailImageResponse{
			Message: fmt.Sprintf("error detailing image: %s", err.Error()),
			Image:   nil,
		}, err
	}

	return &imagesv1.DetailImageResponse{
		Message: "detail image success",
		Image:   image,
	}, nil
}

func (s *ImageServer) UpdateImage(ctx context.Context, req *imagesv1.UpdateImageRequest) (*imagesv1.UpdateImageResponse, error) {

	log.Printf("Received UpdateImage Request\n")

	image, err := s.Detail(ctx, req.Id)
	if err != nil {
		return &imagesv1.UpdateImageResponse{
			Message: fmt.Sprintf("error detailing image: %s", err.Error()),
		}, err
	}

	imageNew := &imagesv1.Image{
		Id:        req.Id,
		Name:      req.Name,
		Size:      image.Size,
		Path:      image.Path,
		CreatedAt: image.CreatedAt,
		UpdatedAt: timestamppb.Now(),
	}

	err = s.Update(ctx, imageNew)
	if err != nil {
		return &imagesv1.UpdateImageResponse{
			Message: fmt.Sprintf("error updating image: %s", err.Error()),
		}, err
	}

	return &imagesv1.UpdateImageResponse{
		Message: "update image success",
	}, nil
}

func (s *ImageServer) DeleteImage(ctx context.Context, req *imagesv1.DeleteImageRequest) (*imagesv1.DeleteImageResponse, error) {

	log.Printf("Received DeleteImage Request\n")

	err := s.Delete(ctx, req.Id)
	if err != nil {
		return &imagesv1.DeleteImageResponse{
			Message: fmt.Sprintf("error deleting image: %s", err.Error()),
		}, err
	}

	return &imagesv1.DeleteImageResponse{
		Message: "delete image success",
	}, nil
}
