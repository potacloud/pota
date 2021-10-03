package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	imagesv1 "github.com/potacloud/pota/api/images/v1"
	networksv1 "github.com/potacloud/pota/api/networks/v1"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
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

func errorHandler(err error) {
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
}

func main() {
	conn, err := grpc.Dial("unix:///"+viper.GetString("socket"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial error: %v", err)
	}
	defer conn.Close()

	// define new grpc client
	imageClient := imagesv1.NewImagesClient(conn)
	networkClient := networksv1.NewNetworksClient(conn)

	// client context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(viper.GetInt("potac.timeout"))*time.Second)
	defer cancel()

	app := &cli.App{
		Name:  "pota cli",
		Usage: "client for pota",
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "create object",
				Subcommands: []*cli.Command{
					{
						Name:    "image",
						Aliases: []string{"images"},
						Usage:   "create image",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "url", Aliases: []string{"u"}, Required: true},
						},
						Action: func(c *cli.Context) error {
							img := &imagesv1.CreateImageRequest{
								Name: c.Args().Get(0),
								Url:  c.String("url"),
							}

							res, err := imageClient.CreateImage(ctx, img)
							errorHandler(err)

							fmt.Println("Message   : ", res.Message)
							fmt.Println("ID        : ", res.Image.Id)
							fmt.Println("Name      : ", res.Image.Name)
							fmt.Println("Size      : ", res.Image.Size)
							fmt.Println("Path      : ", res.Image.Path)
							fmt.Println("CraetedAt : ", res.Image.CreatedAt.AsTime())

							return nil
						},
					},
					{
						Name:    "network",
						Aliases: []string{"networks"},
						Usage:   "create network",
						Action: func(c *cli.Context) error {
							net := &networksv1.CreateNetworkRequest{
								Name: c.Args().Get(0),
							}

							res, err := networkClient.CreateNetwork(ctx, net)
							errorHandler(err)

							fmt.Println("Message   : ", res.Message)
							fmt.Println("ID        : ", res.Network.Id)
							fmt.Println("Name      : ", res.Network.Name)
							fmt.Println("Bridge    : ", res.Network.Bridge)
							fmt.Println("CIDR      : ", res.Network.Cidr)
							fmt.Println("MTU       : ", res.Network.Mtu)
							fmt.Println("Gateway   : ", res.Network.Gateway)
							fmt.Println("CraetedAt : ", res.Network.CreatedAt.AsTime())

							return nil
						},
					},
				},
			},
			{
				Name:  "get",
				Usage: "get object",
				Subcommands: []*cli.Command{
					{
						Name:    "image",
						Aliases: []string{"images"},
						Usage:   "list image",
						Action: func(c *cli.Context) error {

							images, err := imageClient.ListImage(ctx, &imagesv1.ListImageRequest{})
							errorHandler(err)

							for i, v := range images.Image {
								fmt.Printf("======== %d ========\n", i)

								fmt.Println("ID        : ", v.Id)
								fmt.Println("Name      : ", v.Name)
								fmt.Println()
							}

							return nil
						},
					},
				},
			},
			{
				Name:  "describe",
				Usage: "describe object",
				Subcommands: []*cli.Command{
					{
						Name:    "image",
						Aliases: []string{"images"},
						Usage:   "describe image",
						Action: func(c *cli.Context) error {

							res, err := imageClient.DetailImage(ctx, &imagesv1.DetailImageRequest{Id: c.Args().Get(0)})
							errorHandler(err)

							if res == nil {
								fmt.Printf("image id %s not found\n", c.Args().Get(0))
								return nil
							}

							fmt.Println("ID        : ", res.Image.Id)
							fmt.Println("Name      : ", res.Image.Name)
							fmt.Println("Size      : ", res.Image.Size)
							fmt.Println("Path      : ", res.Image.Path)
							fmt.Println("CraetedAt : ", res.Image.CreatedAt.AsTime())
							fmt.Println("UpdatedAt : ", res.Image.UpdatedAt.AsTime())

							return nil
						},
					},
				},
			},
			{
				Name:  "edit",
				Usage: "edit object",
				Subcommands: []*cli.Command{
					{
						Name:    "image",
						Aliases: []string{"images"},
						Usage:   "edit image",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Required: true},
						},
						Action: func(c *cli.Context) error {

							res, err := imageClient.UpdateImage(ctx, &imagesv1.UpdateImageRequest{Id: c.Args().Get(0), Name: c.String("name")})
							errorHandler(err)

							if res == nil {
								fmt.Printf("image id %s not found\n", c.Args().Get(0))
								return nil
							}

							fmt.Println(res.Message)

							return nil
						},
					},
				},
			},
			{
				Name:  "delete",
				Usage: "delete object",
				Subcommands: []*cli.Command{
					{
						Name:    "image",
						Aliases: []string{"images"},
						Usage:   "delete image",
						Action: func(c *cli.Context) error {

							res, err := imageClient.DeleteImage(ctx, &imagesv1.DeleteImageRequest{Id: c.Args().Get(0)})
							errorHandler(err)

							if res == nil {
								fmt.Printf("image id %s not found\n", c.Args().Get(0))
								return nil
							}

							fmt.Print(res.Message)

							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
