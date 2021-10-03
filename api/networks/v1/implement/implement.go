package implement

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	networksv1 "github.com/potacloud/pota/api/networks/v1"
	"github.com/potacloud/pota/pkg/repository"
	"github.com/vishvananda/netlink"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type NetworkServer struct {
	networksv1.UnimplementedNetworksServer
	repository.NetworkRepository
}

func (s *NetworkServer) CreateNetwork(ctx context.Context, req *networksv1.CreateNetworkRequest) (*networksv1.CreateNetworkResponse, error) {

	uuid := uuid.New().String()
	uuidSplit := strings.Split(uuid, "-")

	net := &networksv1.Network{
		Id:        uuid,
		Name:      req.Name,
		Bridge:    "br-" + uuidSplit[0],
		Cidr:      "10.101.100.0/24",
		Gateway:   "10.101.100.1/24",
		Mtu:       1500,
		Snat:      true,
		Generated: true,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}

	// network logic
	bridge := &netlink.Bridge{LinkAttrs: netlink.LinkAttrs{
		TxQLen: -1,
		Name:   "br-" + uuidSplit[0],
		MTU:    1500,
	}}

	err := netlink.LinkAdd(bridge)
	if err != nil {
		fmt.Printf("could not add %s: %v\n", bridge.Name, err)
		return &networksv1.CreateNetworkResponse{
			Message: fmt.Sprintf("could not add %s: %v\n", bridge.Name, err),
			Network: nil,
		}, err
	}

	err = netlink.LinkSetUp(bridge)
	if err != nil {
		fmt.Printf("could not set dev %s up: %v\n", bridge.Name, err)
		return &networksv1.CreateNetworkResponse{
			Message: fmt.Sprintf("could not set dev %s up: %v\n", bridge.Name, err),
			Network: nil,
		}, err
	}

	addr, err := netlink.ParseAddr(net.Gateway)
	if err != nil {
		fmt.Printf("could not parse addres %s : %v\n", net.Cidr, err)
		return &networksv1.CreateNetworkResponse{
			Message: fmt.Sprintf("could not parse addres %s : %v\n", net.Cidr, err),
			Network: nil,
		}, err
	}

	err = netlink.AddrAdd(bridge, addr)
	if err != nil {
		fmt.Printf("could not add addres %s : %v\n", bridge.Name, err)
		return &networksv1.CreateNetworkResponse{
			Message: fmt.Sprintf("could not add addres %s : %v\n", bridge.Name, err),
			Network: nil,
		}, err
	}

	return &networksv1.CreateNetworkResponse{
		Message: "create network success",
		Network: net,
	}, nil
}

func (s *NetworkServer) ListNetwork(ctx context.Context, req *networksv1.ListNetworkRequest) (*networksv1.ListNetworkResponse, error) {
	return nil, nil
}

func (s *NetworkServer) DetailNetwork(ctx context.Context, req *networksv1.DetailNetworkRequest) (*networksv1.DetailNetworkResponse, error) {
	return nil, nil
}

func (s *NetworkServer) UpdateNetwork(ctx context.Context, req *networksv1.UpdateNetworkRequest) (*networksv1.UpdateNetworkResponse, error) {
	return nil, nil
}

func (s *NetworkServer) DeleteNetwork(ctx context.Context, req *networksv1.DeleteNetworkRequest) (*networksv1.DeleteNetworkResponse, error) {
	return nil, nil
}
