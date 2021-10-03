package repository

import (
	"context"
	"fmt"
	"log"
	"testing"

	networksv1 "github.com/potacloud/pota/api/networks/v1"
	"github.com/potacloud/pota/pkg/dbcon"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCreateNetwork(t *testing.T) {
	db, err := dbcon.BboltConnect("../../pota/pota.db")
	if err != nil {
		log.Fatal("error connecting to database")
	}
	defer db.Close()

	network := NewNetworkRepository(db)
	err = network.Create(context.Background(), &networksv1.Network{
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
	})

	assert.NoError(t, err)
}

func TestListNetwork(t *testing.T) {
	db, err := dbcon.BboltConnect("../../pota/pota.db")
	if err != nil {
		log.Fatal("error connecting to database")
	}
	defer db.Close()

	network := NewNetworkRepository(db)
	networks, err := network.List(context.Background())
	for _, v := range networks {
		fmt.Println(v.Id, v.Name)
	}
	assert.NoError(t, err)
}

func TestDetailNetwork(t *testing.T) {
	db, err := dbcon.BboltConnect("../../pota/pota.db")
	if err != nil {
		log.Fatal("error connecting to database")
	}
	defer db.Close()

	net := NewNetworkRepository(db)
	network, err := net.Detail(context.Background(), "cba1ac03-e4ed-47f9-8665-0551cc5a5af0")
	fmt.Println(network.Id, network.Name)

	assert.NoError(t, err)
	assert.NotEmpty(t, network)
}

func TestUpdateNetwork(t *testing.T) {
	db, err := dbcon.BboltConnect("../../pota/pota.db")
	if err != nil {
		log.Fatal("error connecting to database")
	}
	defer db.Close()

	net := NewNetworkRepository(db)
	err = net.Update(context.Background(), &networksv1.Network{
		Id:        "cba1ac03-e4ed-47f9-8665-0551cc5a5af0",
		Name:      "test-network-update",
		Cidr:      "10.100.101.0/24",
		Gateway:   "10.100.100.1",
		Bridge:    "br-cba1ac03",
		Mtu:       9000,
		Snat:      false,
		Generated: true,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	})
	assert.NoError(t, err)
}

func TestDeleteNetwork(t *testing.T) {
	db, err := dbcon.BboltConnect("../../pota/pota.db")
	if err != nil {
		log.Fatal("error connecting to database")
	}
	defer db.Close()

	net := NewNetworkRepository(db)
	err = net.Delete(context.Background(), "cba1ac03-e4ed-47f9-8665-0551cc5a5af0")

	assert.NoError(t, err)
}
