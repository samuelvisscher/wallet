package store

import (
	"sync"
)

type ImmutableStorer interface {
	Append(meta KittyMeta, img []byte) (int, bool)
	ListKitties(pageIndex, pageSize int) []KittyMeta
	GetKitty(kID int) (KittyMeta, bool)
	GetImage(kID int, resolution string) ([]byte, bool)
}

type ImmutableBasic struct {
	sync.RWMutex
	metas  []KittyMeta
	images [][]byte
}

func (ib *ImmutableBasic) Append(meta KittyMeta, img []byte) (int, bool) {
	ib.Lock()
	defer ib.Unlock()

	// TODO: Check.

	ib.metas, ib.images = append(ib.metas, meta), append(ib.images, img)
	return len(ib.metas)-1, true
}

func (ib *ImmutableBasic) ListKitties(pageIndex, pageSize int) []KittyMeta {
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

func (ib *ImmutableBasic) GetKitty(kID int) (KittyMeta, bool) {
	ib.RLock()
	defer ib.RUnlock()

	if kID >= len(ib.metas) {
		return ib.metas[kID], false
	}
	return ib.metas[kID], true
}

func (ib *ImmutableBasic) GetImage(kID int, resolution string) ([]byte, bool) {
	ib.RLock()
	defer ib.RUnlock()

	if kID >= len(ib.images) {
		return nil, false
	}
	return ib.images[kID], true
}
