package store

import (
	"encoding/json"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
)

// PrepareDB prepares a DB, loading a directory in.
func PrepareDB(rootDir string, kittyCount int, db DB) error {
	for i := 0; i < kittyCount; i++ {
		var (
			iStr     = strconv.Itoa(i)
			jsonName = path.Join(rootDir, iStr+".json")
			pngName  = path.Join(rootDir, iStr+".png")
			jsonObj  KittyMeta
			pngRaw   []byte
		)
		{
			f, e := os.Open(pngName)
			if e != nil {
				return e
			}
			pngRaw, e = ioutil.ReadAll(f)
			if e != nil {
				return e
			}
			f.Close()
		}
		{
			f, e := os.Open(jsonName)
			if e != nil {
				return e
			}
			if e := json.NewDecoder(f).Decode(&jsonObj); e != nil {
				return e
			}
			f.Close()
		}
		hash, e := db.AddKitty(pngRaw, &jsonObj)
		if e != nil {
			return e
		}
		log.Printf(
			"[%d] kitty added : id(%d) image(%s)",
			i, i, hash.Hex())
	}
	return nil
}

// InitDir initializes a directory.
func InitDir(rootDir string, count int) error {
	if e := os.RemoveAll(rootDir); e != nil {
		return e
	}
	if e := os.MkdirAll(rootDir, os.FileMode(0755)); e != nil {
		return e
	}
	var rect = image.NewRGBA(image.Rect(0, 0, 1200, 1200))
	for i := 0; i < count; i++ {

		var (
			iStr     = strconv.Itoa(i)
			jsonName = path.Join(rootDir, iStr+".json")
			pngName  = path.Join(rootDir, iStr+".png")
			raw, _   = json.MarshalIndent(
				KittyMeta{ID: uint64(i), Attributes: []string{}},
				"", "    ")
		)
		{
			f, e := os.Create(jsonName)
			if e != nil {
				return e
			}
			if _, e := f.Write(raw); e != nil {
				return e
			}
			f.Close()
		}
		{
			f, e := os.Create(pngName)
			if e != nil {
				return e
			}
			if e := png.Encode(f, rect); e != nil {
				return e
			}
			f.Close()
		}
	}
	return nil
}
