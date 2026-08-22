package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dustin/go-humanize"
)

type FileMeta struct {
	Name    string `json:"filename"`
	Path    string `json:"filepath"`
	Type    string `json:"filetype"`
	Size    string `json:"size"`
	Perm    string `json:"permissions"`
	Mod     string `json:"last_modified"`
	RawSize int64  `json:"raw_size"`
}

func getMeta(path string) (*FileMeta, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("'%s' not found", path)
		} else {
			return nil, err
		}
	}

	abs, _ := filepath.Abs(path)
	var size int64
	var ftype string
	if info.IsDir() {
		size, _ = dirSize(path)
		count := dirFileCount(path)
		ftype = fmt.Sprintf("directory, %d files", count)
	} else {
		size = info.Size()
		ftype, _ = fileCmd(path)
	}

	return &FileMeta{
		Name:    info.Name(),
		Path:    shortenHome(abs),
		Type:    ftype,
		Size:    humanize.Bytes(uint64(size)),
		Perm:    fmt.Sprintf("%o", info.Mode().Perm()),
		Mod:     info.ModTime().Format("02 Jan 06 15:04"),
		RawSize: size,
	}, nil
}
