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
}

func (km KittyMeta) Encode() []byte {
	raw, _ := json.Marshal(km)
	return raw
}

func KittyMetaFromRaw(raw []byte) (KittyMeta, error) {
	var km KittyMeta
	if e := json.Unmarshal(raw, &km); e != nil {
		return km, e
	}
	return km, nil
}
