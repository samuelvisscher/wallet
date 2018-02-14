package store

import (
	"sync"
)

type MetaStorer interface {
	Append(meta KittyMeta, img []byte) (uint64, bool)
	ListKitties(pageIndex, pageSize int) []KittyMeta
	GetKitty(kID uint64) (KittyMeta, bool)
	GetImage(kID uint64, resolution string) ([]byte, bool)
}

type MetaMemoryDB struct {
	sync.RWMutex
	metas  []KittyMeta
	images [][]byte
}

func (ib *MetaMemoryDB) Append(meta KittyMeta, img []byte) (uint64, bool) {
	ib.Lock()
	defer ib.Unlock()

	// TODO: Check.
	meta.ID = uint64(len(ib.metas))

	ib.metas, ib.images = append(ib.metas, meta), append(ib.images, img)
	return meta.ID, true
}

func (ib *MetaMemoryDB) ListKitties(pageIndex, pageSize int) []KittyMeta {
	ib.RLock()
	defer ib.RUnlock()

	if pageIndex < 0 || pageSize < 0 {
		return []KittyMeta{}
	}
	var (
		start = pageSize * pageIndex
		end   = start + pageSize
	)
	if end > len(ib.metas) {
		end = len(ib.metas)
	}
	if end <= start || start >= len(ib.metas) {
		return []KittyMeta{}
	}
	return ib.metas[start:end]
}

func (ib *MetaMemoryDB) GetKitty(kID uint64) (KittyMeta, bool) {
	ib.RLock()
	defer ib.RUnlock()

	if kID >= uint64(len(ib.metas)) {
		return ib.metas[kID], false
	}
	return ib.metas[kID], true
}

func (ib *MetaMemoryDB) GetImage(kID uint64, resolution string) ([]byte, bool) {
	ib.RLock()
	defer ib.RUnlock()

	if kID >= uint64(len(ib.metas)) {
		return nil, false
	}
	return ib.images[kID], true
}
