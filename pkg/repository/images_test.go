package repository

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	imagesv1 "github.com/potacloud/pota/api/images/v1"
	"github.com/potacloud/pota/pkg/dbcon"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCreateImage(t *testing.T) {
	db, err := dbcon.BboltConnect("../../pota/pota.db")
	if err != nil {
		log.Fatal("error connecting to database")
	}
	defer db.Close()

	img := NewImageRepository(db)
	err = img.Create(context.Background(), &imagesv1.Image{
		Id:        "830659d9-bdfb-4e74-b85a-75557759cbef",
		Name:      "test",
		Size:      10,
		Path:      "/path/to/image",
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
	})

	assert.NoError(t, err)
}

func TestListImage(t *testing.T) {
	db, err := dbcon.BboltConnect("../../pota/pota.db")
	if err != nil {
		log.Fatal("error connecting to database")
	}
	defer db.Close()

	img := NewImageRepository(db)
	images, err := img.List(context.Background())
	for _, v := range images {
		fmt.Println(v.Id, v.Name)
	}
	assert.NoError(t, err)

}

func TestDetailImage(t *testing.T) {
	db, err := dbcon.BboltConnect("../../pota/pota.db")
	if err != nil {
		log.Fatal("error connecting to database")
	}
	defer db.Close()

	img := NewImageRepository(db)
	image, err := img.Detail(context.Background(), "830659d9-bdfb-4e74-b85a-75557759cbef")
	fmt.Println(image.Id, image.Name)

	assert.NoError(t, err)
	assert.NotEmpty(t, image)
}

func TestUpdateImage(t *testing.T) {
	db, err := dbcon.BboltConnect("../../pota/pota.db")
	if err != nil {
		log.Fatal("error connecting to database")
	}
	defer db.Close()

	img := NewImageRepository(db)
	err = img.Update(context.Background(), &imagesv1.Image{
		Id:        "830659d9-bdfb-4e74-b85a-75557759cbef",
		Name:      "test-updated",
		Size:      10,
		Path:      "/path/to/image",
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
	})

	assert.NoError(t, err)
}

func TestDeleteImage(t *testing.T) {
	db, err := dbcon.BboltConnect("../../pota/pota.db")
	if err != nil {
		log.Fatal("error connecting to database")
	}
	defer db.Close()

	img := NewImageRepository(db)
	err = img.Delete(context.Background(), "830659d9-bdfb-4e74-b85a-75557759cbef")

	assert.NoError(t, err)
}
