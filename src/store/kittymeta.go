package store

import (
	"encoding/json"
)

type KittyMeta struct {
	ID           uint64   `json:"ID"`
	ImageHash    string   `json:"image_hash"`
	Name         string   `json:"name"`
	Bio          string   `json:"bio"`
	Breed        string   `json:"breed"`
	Attributes   []string `json:"attributes"`
	PriceBitcoin int64    `json:"price_bitcoin"`
	PriceSkycoin int64    `json:"price_skycoin"`
}

func (km KittyMeta) Encode() []byte {
	raw, _ := json.Marshal(km)
	return raw
}

func KittyMetaFromRaw(raw []byte) (*KittyMeta, error) {
	var km = new(KittyMeta)
	if e := json.Unmarshal(raw, km); e != nil {
		return nil, e
	}
	return km, nil
}
