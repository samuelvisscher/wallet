package store

import (
	"encoding/json"
	"image"
	"image/png"
	"os"
	"path"
	"strconv"
)

// PrepareDB prepares a DB, loading a directory in.
func PrepareDB(rootDir string, db DB) error {

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
	var (
		raw, _ = json.MarshalIndent(KittyMeta{Attributes: []string{}}, "", "    ")
		rect   = image.NewRGBA(image.Rect(0, 0, 1200, 1200))
	)
	for i := 0; i < count; i++ {
		iStr := strconv.Itoa(i)
		{
			f, e := os.Create(path.Join(rootDir, iStr+".json"))
			if e != nil {
				return e
			}
			if _, e := f.Write(raw); e != nil {
				return e
			}
			f.Close()
		}
		{
			f, e := os.Create(path.Join(rootDir, iStr+".png"))
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