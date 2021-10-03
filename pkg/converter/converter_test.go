package converter

import (
	"testing"

	imagesv1 "github.com/potacloud/pota/api/images/v1"
	networksv1 "github.com/potacloud/pota/api/networks/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type test struct {
	Name string
}

func TestEncodeToBytes(t *testing.T) {

	object := test{
		Name: "lorem ipsum dolor sit amet",
	}

	bytes, err := EncodeToBytes(object)
	assert.NoError(t, err)
	assert.NotEmpty(t, bytes)
}

func TestDecodeToImage(t *testing.T) {
	img := imagesv1.Image{
		Id:        "1e35a8b3-5a79-4e93-bc29-919e2e3e2284",
		Name:      "cirros",
		Size:      uint32(324556),
		Path:      "/var/lib/pota/images/cirros.img",
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}

	bytes, _ := EncodeToBytes(img)

	imgDecode, err := DecodeToImage(bytes)
	assert.NoError(t, err)
	assert.Equal(t, img, *imgDecode)
}

func TestDecodeToNetwork(t *testing.T) {
	net := networksv1.Network{
		Id:        "cba1ac03-e4ed-47f9-8665-0551cc5a5af0",
		Name:      "test-network",
		Cidr:      "10.100.101.0/24",
		Gateway:   "10.100.100.1",
		Bridge:    "br-cba1ac03",
		Mtu:       1500,
		Snat:      true,
		Generated: true,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}

	bytes, _ := EncodeToBytes(net)

	netDecode, err := DecodeToNetwork(bytes)
	assert.NoError(t, err)
	assert.Equal(t, net, *netDecode)
}
