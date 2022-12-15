package main

import (
	"fmt"
	"os"
	"sort"
)

type ByModTime []os.FileInfo

func (fis ByModTime) Len() int {
	return len(fis)
}

func (fis ByModTime) Swap(i, j int) {
	fis[i], fis[j] = fis[j], fis[i]
}

func (fis ByModTime) Less(i, j int) bool {
	return fis[i].ModTime().Before(fis[j].ModTime())
}

func SortFile(path string) (files ByModTime) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fis, err := f.Readdir(-1)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	files = make(ByModTime, len(fis)+10)
	j := 0
	for _, v := range fis {
		files[j] = v
		j++
	}
	files = files[:j]
	sort.Sort(ByModTime(files))
	return
}

func main() {
	files := SortFile("F:\\jason_local\\11_3")
	for _, v := range files {
		fmt.Println(v.Name())
	}
}
