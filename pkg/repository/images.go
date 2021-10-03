package repository

import (
	"context"
	"fmt"

	"go.etcd.io/bbolt"

	imagesv1 "github.com/potacloud/pota/api/images/v1"
	"github.com/potacloud/pota/pkg/converter"
)

type imageRepository struct {
	DB *bbolt.DB
}

func NewImageRepository(db *bbolt.DB) ImageRepository {
	return &imageRepository{DB: db}
}

func (r *imageRepository) Create(ctx context.Context, image *imagesv1.Image) error {

	// Create bucket if not exist
	err := r.DB.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("images"))
		if err != nil {
			return fmt.Errorf("error creating bucket: %s", err)
		}

		// return if object already exists
		bytes := b.Get([]byte(image.Id))
		if bytes != nil {
			return fmt.Errorf("error object already exists : %s", image.Id)
		}

		// convert image into []byte format
		byteImage, err := converter.EncodeToBytes(image)
		if err != nil {
			return fmt.Errorf("error converting object : %s", err)
		}

		// insert object into db
		err = b.Put([]byte(image.Id), byteImage)
		if err != nil {
			return fmt.Errorf("error inserting object key %s : %s", image.Id, err)
		}
		return nil
	})

	return err
}

func (r *imageRepository) List(ctx context.Context) ([]*imagesv1.Image, error) {

	var images []*imagesv1.Image

	err := r.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("images"))

		b.ForEach(func(k, v []byte) error {
			image, err := converter.DecodeToImage(v)
			if err != nil {
				return fmt.Errorf("error decoding object  %s : %s", image.Id, err)
			}
			images = append(images, image)
			return nil
		})

		return nil
	})

	return images, err
}

func (r *imageRepository) Detail(ctx context.Context, id string) (*imagesv1.Image, error) {

	var image *imagesv1.Image = new(imagesv1.Image)
	var err error

	err = r.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("images"))

		bytes := b.Get([]byte(id))

		if bytes == nil {
			return fmt.Errorf("image id %s not found", id)
		}

		image, err = converter.DecodeToImage(bytes)
		if err != nil {
			return fmt.Errorf("error decoding object  %s : %s", id, err)
		}

		return nil
	})

	return image, err

}

func (r *imageRepository) Update(ctx context.Context, image *imagesv1.Image) error {

	err := r.DB.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("images"))
		if err != nil {
			return fmt.Errorf("error creating bucket: %s", err)
		}

		bytes := b.Get([]byte(image.Id))
		if bytes == nil {
			return fmt.Errorf("error object not found : %s", image.Id)
		}

		byteImage, err := converter.EncodeToBytes(image)
		if err != nil {
			return fmt.Errorf("error converting object %s : %s", image.Id, err)
		}

		err = b.Put([]byte(image.Id), byteImage)
		if err != nil {
			return fmt.Errorf("error updating object %s : %s", image.Id, err)
		}

		return nil
	})

	return err
}

func (r *imageRepository) Delete(ctx context.Context, id string) error {

	err := r.DB.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("images"))
		if err != nil {
			return fmt.Errorf("error creating bucket: %s", err)
		}

		bytes := b.Get([]byte(id))
		if bytes == nil {
			return fmt.Errorf("error object not found : %s", id)
		}

		err = b.Delete([]byte(id))
		if err != nil {
			return fmt.Errorf("error deleting object  %s : %s", id, err)
		}

		return nil
	})

	return err
}
