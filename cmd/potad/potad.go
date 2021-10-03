package main

import (
	"fmt"
	"log"
	"net"
	"os"

	imagesv1 "github.com/potacloud/pota/api/images/v1"
	imagesv1Impl "github.com/potacloud/pota/api/images/v1/implement"
	networksv1 "github.com/potacloud/pota/api/networks/v1"
	networksv1Impl "github.com/potacloud/pota/api/networks/v1/implement"
	"github.com/potacloud/pota/pkg/dbcon"
	"github.com/potacloud/pota/pkg/repository"
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

	db, err := dbcon.BboltConnect(viper.GetString("db"))
	if err != nil {
		log.Fatal("database connection error: ", err)
	}

	// repository layer
	imageRepository := repository.NewImageRepository(db)
	networkRepository := repository.NewNetworkRepository(db)

	// service layer
	// define new grpc server and register ImageServer to this
	s := grpc.NewServer()
	imagesv1.RegisterImagesServer(s, &imagesv1Impl.ImageServer{ImageRepository: imageRepository})
	networksv1.RegisterNetworksServer(s, &networksv1Impl.NetworkServer{NetworkRepository: networkRepository})

	log.Printf("server listening at %v", l.Addr())

	// listen the connection
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
