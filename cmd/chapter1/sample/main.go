package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"os"
	"path/filepath"
	"strings"
)

type Info struct {
	mapset.Set "学生名字" // 结构体标签
}

type Image struct {
	infos mapset.Set
}

var images = new(Image)

func (image Image) getInfos() {
	filepath.Walk("F:/jason_local/go_project/study_project/go_web_demo", image.explorer)
}

func (image Image) explorer(path string, info os.FileInfo, err error) error {

	if err != nil {
		fmt.Println("ERROR: %v", err)
		return err
	}
	if !info.IsDir() {
		const log = ".log"
		const img = ".png"
		if strings.Contains(info.Name(), log) {
			fmt.Println("找到日志: %sn", path)
		} else if strings.Contains(info.Name(), img) {
			image.infos.Add(path)
		}
	}
	return nil
}

func main() {
	images.getInfos()
	for img := range images.infos.Iter() {
		fmt.Println(img)
	}
}
