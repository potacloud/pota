package repository

import (
	context "context"

	imagesv1 "github.com/potacloud/pota/api/images/v1"
	networksv1 "github.com/potacloud/pota/api/networks/v1"
)

type ImageRepository interface {
	Create(ctx context.Context, image *imagesv1.Image) error
	List(ctx context.Context) ([]*imagesv1.Image, error)
	Detail(ctx context.Context, id string) (*imagesv1.Image, error)
	Update(ctx context.Context, image *imagesv1.Image) error
	Delete(ctx context.Context, id string) error
}

type NetworkRepository interface {
	Create(ctx context.Context, network *networksv1.Network) error
	List(ctx context.Context) ([]*networksv1.Network, error)
	Detail(ctx context.Context, id string) (*networksv1.Network, error)
	Update(ctx context.Context, network *networksv1.Network) error
	Delete(ctx context.Context, id string) error
}
