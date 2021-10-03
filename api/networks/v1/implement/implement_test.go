package implement

import (
	"context"
	"log"
	"testing"

	"github.com/potacloud/pota/pkg/dbcon"
	"github.com/potacloud/pota/pkg/repository"
)

func TestServiceCreateNetwork(t *testing.T) {
	db, err := dbcon.BboltConnect("../../../../pota/pota.db")
	if err != nil {
		log.Fatal("error connecting to database")
	}
	defer db.Close()

	networkRepository := repository.NewNetworkRepository(db)

	networkServer := NetworkServer{
		NetworkRepository: networkRepository,
	}

	networkServer.CreateNetwork(context.Background(), nil)

}
