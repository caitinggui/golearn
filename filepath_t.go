package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func getFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}

func main() {
	files, _ := filepath.Glob("*")
	log.Println(files)
	log.Println(getFileSize(files[0]))

	ifiles, _ := ioutil.ReadDir("./logs")
	log.Println(ifiles)
	for _, f := range ifiles {
		if !f.IsDir() {
			log.Println(f.Name(), f.Size(), f.ModTime())
		}
	}
}
