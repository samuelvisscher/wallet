package wallet

import (
	"errors"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/encoder"
	"github.com/skycoin/skycoin/src/util/file"
	"io/ioutil"
	"os"
	"time"
	"io"
)

type (
	// AssetType determines the asset type that the wallet holds.
	AssetType string

	// Extension determines a file's extension.
	Extension string
)

const (
	// Version determines the wallet file's version.
	Version uint64 = 0

	// KittyAsset represents the "kittycash" asset type.
	KittyAsset AssetType = "kittycash"

	// FileExt is the kittycash file extension.
	FileExt Extension = ".kcw"
)

/*
	<<< TYPES >>>
*/

type FloatingMeta struct {
	Version   uint64 `json:"version"`
	Label     string `json:"label"`
	Encrypted bool   `json:"encrypted"`
	Password  string `json:"-"`
	Meta
}

type Meta struct {
	AssetType AssetType `json:"type"`
	Seed      string    `json:"seed"`
	TS        int64     `json:"timestamp"`
}

type FloatingWallet struct {
	Meta    FloatingMeta
	Entries []Entry
}

type Wallet struct {
	Meta    Meta
	Entries []Entry
}

func (w Wallet) Serialize() []byte {
	return encoder.Serialize(w)
}

/*
	<<< CREATION >>>
*/

type FloatingWalletOptions struct {
	Label     string `json:"string"`
	Seed      string `json:"seed"`
	Encrypted bool   `json:"encrypted"`
	Password  string `json:"password,omitempty"`
}

func (o *FloatingWalletOptions) Verify() error {
	if o.Label == "" {
		return errors.New("invalid label")
	}
	if o.Seed == "" {
		return errors.New("invalid seed")
	}
	if o.Encrypted && o.Password == "" {
		return errors.New("invalid password")
	}
	return nil
}

func NewFloatingWallet(options *FloatingWalletOptions) (*FloatingWallet, error) {
	if e := options.Verify(); e != nil {
		return nil, e
	}

	return &FloatingWallet{
		Meta: FloatingMeta{
			Version:   Version,
			Label:     options.Label,
			Encrypted: options.Encrypted,
			Password:  options.Password,
			Meta: Meta{
				AssetType: KittyAsset,
				Seed:      options.Seed,
				TS:        time.Now().UnixNano(),
			},
		},
		Entries: []Entry{},
	}, nil
}

func LoadFloatingWallet(f io.Reader, label, password string) (*FloatingWallet, error) {
	raw, e := ioutil.ReadAll(f)
	if e != nil {
		return nil, e
	}
	prefix, data, e := ExtractPrefix(raw)
	if e != nil {
		return nil, e
	}
	encrypted := prefix.Encrypted()
	if encrypted {
		pHash := cipher.SumSHA256([]byte(password))
		data, e = cipher.Chacha20Decrypt(data, pHash[:], prefix.Nonce())
		if e != nil {
			return nil, e
		}
	} else {
		password = ""
	}

	var wallet Wallet
	if e := encoder.DeserializeRaw(data, &wallet); e != nil {
		return nil, e
	}
	return &FloatingWallet{
		Meta: FloatingMeta{
			Version:   prefix.Version(),
			Label:     label,
			Encrypted: encrypted,
			Password:  password,
			Meta:      wallet.Meta,
		},
		Entries: wallet.Entries,
	}, nil
}

func (fw *FloatingWallet) Save() error {
	version := fw.Meta.Version

	nonce := EmptyNonce()
	if fw.Meta.Encrypted {
		nonce = RandNonce()
	}

	prefix := NewPrefix(version, nonce)

	data := fw.ToWallet().Serialize()
	if fw.Meta.Encrypted {
		var e error
		pHash := cipher.SumSHA256([]byte(fw.Meta.Password))
		data, e = cipher.Chacha20Encrypt(data, pHash[:], nonce)
		if e != nil {
			return e
		}
	}

	return SaveBinary(
		LabelPath(fw.Meta.Label),
		append(prefix[:], data...),
	)
}

func (fw *FloatingWallet) GenerateEntries(n int) {
	sks := cipher.GenerateDeterministicKeyPairs([]byte(fw.Meta.Seed), n)
	fw.Entries = make([]Entry, n)
	for i := 0; i < n; i++ {
		entry, _ := NewEntry(sks[i])
		fw.Entries[i] = *entry
	}
}

func (fw *FloatingWallet) Count() int {
	return len(fw.Entries)
}

func (fw *FloatingWallet) ToWallet() *Wallet {
	return &Wallet{
		Meta:    fw.Meta.Meta,
		Entries: fw.Entries,
	}
}

/*
	<<< HELPERS >>>
*/

func SaveBinary(fn string, data []byte) error {
	return file.SaveBinary(fn, data, os.FileMode(0600))
}
