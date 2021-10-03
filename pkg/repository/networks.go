package repository

import (
	"context"
	"fmt"

	"go.etcd.io/bbolt"

	networksv1 "github.com/potacloud/pota/api/networks/v1"
	"github.com/potacloud/pota/pkg/converter"
)

type networkRepository struct {
	DB *bbolt.DB
}

func NewNetworkRepository(db *bbolt.DB) NetworkRepository {
	return &networkRepository{DB: db}
}

func (r *networkRepository) Create(ctx context.Context, network *networksv1.Network) error {

	err := r.DB.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("networks"))
		if err != nil {
			return fmt.Errorf("error creating bucket: %s", err)
		}

		// return if object already exists
		bytes := b.Get([]byte(network.Id))
		if bytes != nil {
			return fmt.Errorf("error object already exists : %s", network.Id)
		}

		// convert image into []byte format
		byteNetwork, err := converter.EncodeToBytes(network)
		if err != nil {
			return fmt.Errorf("error converting object : %s", err)
		}

		// insert object into db
		err = b.Put([]byte(network.Id), byteNetwork)
		if err != nil {
			return fmt.Errorf("error inserting object key %s : %s", network.Id, err)
		}
		return nil
	})

	return err
}

func (r *networkRepository) List(ctx context.Context) ([]*networksv1.Network, error) {
	var networks []*networksv1.Network

	err := r.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("networks"))

		b.ForEach(func(k, v []byte) error {
			network, err := converter.DecodeToNetwork(v)
			if err != nil {
				return fmt.Errorf("error decoding object  %s : %s", network.Id, err)
			}
			networks = append(networks, network)
			return nil
		})

		return nil
	})

	return networks, err
}

func (r *networkRepository) Detail(ctx context.Context, id string) (*networksv1.Network, error) {

	var network *networksv1.Network = new(networksv1.Network)
	var err error

	err = r.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("networks"))

		bytes := b.Get([]byte(id))

		if bytes == nil {
			return fmt.Errorf("network id %s not found", id)
		}

		network, err = converter.DecodeToNetwork(bytes)
		if err != nil {
			return fmt.Errorf("error decoding object  %s : %s", id, err)
		}

		return nil
	})

	return network, err
}

func (r *networkRepository) Update(ctx context.Context, network *networksv1.Network) error {

	err := r.DB.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("networks"))
		if err != nil {
			return fmt.Errorf("error creating bucket: %s", err)
		}

		bytes := b.Get([]byte(network.Id))
		if bytes == nil {
			return fmt.Errorf("error object not found : %s", network.Id)
		}

		byteNetwork, err := converter.EncodeToBytes(network)
		if err != nil {
			return fmt.Errorf("error converting object %s : %s", network.Id, err)
		}

		err = b.Put([]byte(network.Id), byteNetwork)
		if err != nil {
			return fmt.Errorf("error updating object %s : %s", network.Id, err)
		}

		return nil
	})

	return err
}

func (r *networkRepository) Delete(ctx context.Context, id string) error {

	err := r.DB.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("networks"))
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
