package main

import (
	"fmt"
	"log"
	"net"
	"os"

	imagesv1 "github.com/potacloud/pota/api/images/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	viper.SetConfigName("pota")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("pota/config") // TODO: ubah ke directory yang sebenarnya
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

// server is used to implement ImagesServer.
type Server struct {
	imagesv1.UnimplementedImagesServer
}

func main() {
	// remove prev unix socket (if exists) before start the new one
	if err := os.RemoveAll(viper.GetString("socket")); err != nil {
		log.Fatalf("failed to remove potad.sock: %v", err)
	}

	// define new unix socket
	l, err := net.Listen("unix", viper.GetString("socket"))
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()

	// define new grpc server and register ImageServer to this
	s := grpc.NewServer()
	imagesv1.RegisterImagesServer(s, &imagesv1.Server{})

	log.Printf("server listening at %v", l.Addr())

	// listen the connection
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
