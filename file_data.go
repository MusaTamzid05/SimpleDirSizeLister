package main

import "fmt"

type FileData struct {
	Name     string
	SizeByte int64
	Size     int64
	SizeUnit string
}

func (f FileData) Show() {
	fmt.Println("Name : ", f.Name, " , Size : ", f.Size, "  ", f.SizeUnit)
}

type FileSorter []FileData

func (s FileSorter) Len() int {
	return len(s)
}

func (s FileSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s FileSorter) Less(i, j int) bool {
	return s[i].SizeByte > s[j].SizeByte
}
