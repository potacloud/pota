package main

import (
	"context"
	"fmt"
	"log"
	"time"

	imagesv1 "github.com/potacloud/pota/api/images/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {
	viper.SetConfigName("pota")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("pota/config") // TODO: ubah ke direktory yang sebenarnya
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func main() {
	conn, err := grpc.Dial("unix:///"+viper.GetString("socket"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial error: %v", err)
	}
	defer conn.Close()

	// define new grpc client
	c := imagesv1.NewImagesClient(conn)

	// client context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(viper.GetInt("potac.timeout"))*time.Second)
	defer cancel()

	img := &imagesv1.CreateImageRequest{
		Name: "test-image",
		Url:  "https://github.com/potacloud/pota/test-image/qcow2",
	}

	res, err := c.CreateImage(ctx, img)
	if err != nil {
		s, ok := status.FromError(err)
		if !ok {
			fmt.Println("unknown error: ", s.Code(), s.Message())
			return
		}
		switch s.Code() {
		case codes.DeadlineExceeded:
			fmt.Println("Deadline Exceeded: ", s.Code(), s.Message())
			return
		case codes.Aborted:
			fmt.Println("Request Aborted: ", s.Code(), s.Message())
			return
		}
	}

	fmt.Println(res.Message)
	fmt.Println("Image Name: ", res.Image.Name)
	fmt.Println("Image ID: ", res.Image.Id)
}
