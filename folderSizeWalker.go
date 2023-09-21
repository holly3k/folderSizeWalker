// go build -o folderSizeWalker.dll -buildmode=c-shared folderSizeWalker.go
package main

import "C"
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type File struct {
	FullPath     string  `json:"path"`
	Size         float64 `json:"size"`
	SizeReadable string  `json:"readableSize"`
	Childs       []File  `json:"childs"`
	Type         string  `json:"type"`
}

//export startScanning
func startScanning(s *C.char) {
	var dirPath string = C.GoString(s)
	startScanningWrapper(dirPath)
}

func startScanningWrapper(dirPath string) {
	t := time.Now()
	folderDetail := scanDir(dirPath)
	fmt.Println(folderDetail.FullPath)
	fmt.Println(folderDetail.SizeReadable)
	fmt.Println(len(folderDetail.Childs))
	jsonFolderDetails, err := json.Marshal(folderDetail)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	} else {
		if err := os.Remove("out.json"); err != nil {
			log.Println(err)
		}
		err = ioutil.WriteFile("out.json", jsonFolderDetails, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("time spend:" + time.Since(t).String())
}

func main() {
	var dirPath string = os.Args[1]
	startScanningWrapper(dirPath)
}

func scanDir(path string) File {

	dirAry, err := ioutil.ReadDir(path)
	folderDetail := File{}
	folderDetail.FullPath = path
	folderDetail.Size = 0
	folderDetail.Type = "FOLDER"
	folderDetail.SizeReadable = "0kb"
	folderDetail.Childs = []File{}
	if err != nil {
		return folderDetail
	}
	for _, e := range dirAry {
		if e.IsDir() {
			subFolder := scanDir(filepath.Join(path, e.Name()))
			folderDetail.Childs = append(folderDetail.Childs, subFolder)
			folderDetail.Size += subFolder.Size
		} else {
			folderDetail.Size += float64(e.Size())
			folderDetail.Childs = append(folderDetail.Childs, File{Type: "FILE", Childs: []File{}, FullPath: filepath.Join(path, e.Name()), Size: float64(e.Size()), SizeReadable: formatSize(float64(e.Size()))})
		}
	}
	folderDetail.SizeReadable = formatSize(folderDetail.Size)
	return folderDetail
}

func formatSize(size float64) string {
	if size < 1024 {
		return strconv.FormatFloat(size, 'f', 2, 32) + "bit"
	} else {
		size = size / 1024
		if size < 1024 {
			return strconv.FormatFloat(size, 'f', 2, 32) + "k"
		} else {
			size = size / 1024
			if size < 1024 {
				return strconv.FormatFloat(size, 'f', 2, 32) + "M"
			} else {
				size = size / 1024
				if size < 1024 {
					return strconv.FormatFloat(size, 'f', 2, 32) + "G"
				} else {
					size = size / 1024
					if size < 1024 {
						return strconv.FormatFloat(size, 'f', 2, 32) + "T"
					} else {
						size = size / 1024
						return strconv.FormatFloat(size, 'f', 2, 32) + "P"
					}
				}
			}
		}
	}
}
