package v1

import (
	"context"
	"log"
	"time"
)

// server is used to implement ImagesServer.
type Server struct {
	UnimplementedImagesServer
}

func (s *Server) CreateImage(ctx context.Context, req *CreateImageRequest) (*CreateImageResponse, error) {

	log.Printf("Received: %v", req.GetName())

	time.Sleep(6 * time.Second)

	// TODO: this just example, need implemented soon
	return &CreateImageResponse{
		Message: "return ok",
		Image: &Image{
			Id:   "de4acdad-7187-41e6-966c-089d43ad3e6d",
			Name: req.Name,
		},
	}, nil
}
