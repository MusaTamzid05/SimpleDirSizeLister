package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

var (
	MB_SIZE = 1000000
	GB_SIZE = 1000000000
)

func ReadDir(path string) ([]os.FileInfo, error) {

	filePtr, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer filePtr.Close()

	fileInfo, err := filePtr.Readdir(-1)
	return fileInfo, nil
}

func DirSize(path string) (int64, error) {

	// https://stackoverflow.com/questions/32482673/how-to-get-directory-total-size
	var size int64
	size = 0

	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})

	return size, err
}

func GetInfo(file os.FileInfo, dirPath string) (FileData, error) {

	name := file.Name()
	var size int64
	var sizeByte int64
	var err error
	var unit string

	if file.IsDir() {
		size, err = DirSize(dirPath + "/" + file.Name())

		if err != nil {
			return FileData{}, err
		}

	} else {
		size = file.Size()
	}

	sizeByte = size

	if size >= int64(MB_SIZE) && size < int64(GB_SIZE) {
		size /= int64(MB_SIZE)
		unit = "MB"
	} else if size > int64(GB_SIZE) {
		size /= int64(GB_SIZE)
		unit = "GB"
	} else {
		unit = "bytes"
	}
	return FileData{Name: name, SizeByte: sizeByte, Size: size, SizeUnit: unit}, nil
}

func main() {

	pathPtr := flag.String("path", "", "Path of the dir to count size")
	flag.Parse()

	if *pathPtr == "" {
		fmt.Println("Usage : -path path_of_dir")
		os.Exit(1)
	}

	path := *pathPtr

	fileInfo, err := ReadDir(path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileData := []FileData{}

	for _, info := range fileInfo {
		data, err := GetInfo(info, path)

		if err != nil {
			fmt.Println("Error opening ", info.Name())
			fmt.Println(err)
		}

		fileData = append(fileData, data)
	}

	sort.Sort(FileSorter(fileData))

	for _, data := range fileData {
		data.Show()
	}
}
