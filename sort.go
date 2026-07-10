package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

func sortFiles(sortBy string, topN int, dir string) {
	var files []FileMeta
	var mu sync.Mutex
	var wg sync.WaitGroup
	paths := make(chan string, 100)

	workerCount := getWorkerCount()

	for range workerCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range paths {
				meta, err := getMeta(path)
				if err == nil {
					mu.Lock()
					files = append(files, *meta)
					mu.Unlock()
				}
			}
		}()
	}

	filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err == nil && !d.IsDir() {
			paths <- path
		}
		return nil
	})
	close(paths)
	wg.Wait()

	if len(files) == 0 {
		fmt.Println("no files found")
		return
	}

	sort.Slice(files, func(i, j int) bool {
		switch sortBy {

		case "biggest":
			return files[i].RawSize > files[j].RawSize
		case "smallest":
			return files[i].RawSize < files[j].RawSize
		}
		return false
	})

	if topN > len(files) {
		topN = len(files)
	}
	if dir == "." {
		dir = "current directory"
	}

	fmt.Printf("%d %s files in %s\n\n", topN, bold(sortBy), shortenHome(dir))
	for i := 0; i < topN; i++ {
		f := files[i]
		fmt.Printf("%-10s %s %s\n", green(f.Size), pink(shortenHome(f.Path)), blue(f.Mod))
		fmt.Printf("%s\n", yellow(f.Type))
	}
}
