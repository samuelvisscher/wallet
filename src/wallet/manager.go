package wallet

import (
	"io"
	"sync"
)

type Manager struct {
	mux     sync.Mutex
	labels  []string
	wallets map[string]*FloatingWallet
}

func NewManager() (*Manager, error) {
	m := &Manager{
		wallets: make(map[string]*FloatingWallet),
	}
	e := RangeLabels(func(f io.Reader, label, fPath string, prefix Prefix) {
		var (
			ver = prefix.Version()
			enc = prefix.Encrypted()
		)
		switch {
		case ver != Version:
			log.Warningf(
				"wallet file `%s` is of version %v, while only version %v is supported",
				label, ver, Version)
			return
		case enc:
			return
		}

	})
	if e != nil {
		return nil, e
	}
}
