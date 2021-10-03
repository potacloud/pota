package converter

import (
	"bytes"
	"encoding/gob"

	imagesv1 "github.com/potacloud/pota/api/images/v1"
	networksv1 "github.com/potacloud/pota/api/networks/v1"
)

func EncodeToBytes(p interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func DecodeToImage(s []byte) (*imagesv1.Image, error) {
	img := imagesv1.Image{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&img)
	if err != nil {
		return &img, err
	}
	return &img, nil
}

func DecodeToNetwork(s []byte) (*networksv1.Network, error) {
	net := networksv1.Network{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&net)
	if err != nil {
		return &net, err
	}
	return &net, nil
}
