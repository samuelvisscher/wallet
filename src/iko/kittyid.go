package iko

import "sort"

type KittyID uint64

type KittyIDs []KittyID

func (ids KittyIDs) Sort() {
	sort.Slice(ids, func(i, j int) bool {
		return (ids)[i] < (ids)[j]
	})
}

func (ids *KittyIDs) Add(id KittyID) {
	*ids = append(*ids, id)
	ids.Sort()
}

func (ids *KittyIDs) Remove(id KittyID) {
	for i, v := range *ids {
		if v == id {
			*ids = append((*ids)[:i], (*ids)[i+1:]...)
			return
		}
	}
}