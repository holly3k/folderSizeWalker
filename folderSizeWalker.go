package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type File struct {
	FullPath     string
	Size         float64
	SizeReadable string
	Childs       []File
}

func main() {
	var dirPath string = "E:\\"
	t := time.Now()
	folderDetail := scanDir(dirPath)
	fmt.Println(folderDetail.FullPath)
	fmt.Println(folderDetail.Size)
	fmt.Println(len(folderDetail.Childs))
	jsonFolderDetails, err := json.Marshal(folderDetail)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	} else {
		err = ioutil.WriteFile("out.json", jsonFolderDetails, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("time spend:" + time.Since(t).String())
}

func scanDir(path string) File {

	dirAry, err := ioutil.ReadDir(path)
	folderDetail := File{}
	folderDetail.FullPath = path
	folderDetail.Size = 0
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
			folderDetail.Childs = append(folderDetail.Childs, File{FullPath: filepath.Join(path, e.Name()), Size: float64(e.Size()), SizeReadable: formatSize(float64(e.Size()))})
		}
	}
	folderDetail.SizeReadable = formatSize(folderDetail.Size)
	return folderDetail
}

func formatSize(size float64) string {
	if size < 1024 {
		return strconv.FormatFloat(size, 'f', 5, 32) + "bit"
	} else {
		size = size / 1024
		if size < 1024 {
			return strconv.FormatFloat(size, 'f', 5, 32) + "k"
		} else {
			size = size / 1024
			if size < 1024 {
				return strconv.FormatFloat(size, 'f', 5, 32) + "M"
			} else {
				size = size / 1024
				if size < 1024 {
					return strconv.FormatFloat(size, 'f', 5, 32) + "G"
				} else {
					size = size / 1024
					if size < 1024 {
						return strconv.FormatFloat(size, 'f', 5, 32) + "T"
					} else {
						size = size / 1024
						return strconv.FormatFloat(size, 'f', 5, 32) + "P"
					}
				}
			}
		}
	}
}
