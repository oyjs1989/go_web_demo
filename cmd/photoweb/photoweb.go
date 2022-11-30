package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"strings"
)

const (
	TEMPLATE_DIR = "./template"
	HOST_ADDR    = ":8899"
	PROJECT_DIR  = "/project"
)

type INFO map[string]interface{}

var templates = make(map[string]*template.Template)
var imageDirs = make(map[string]map[string][]string)

type Result struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
	Data      []struct {
		Prob  float64 `json:"prob"`
		Label string  `json:"label"`
	} `json:"data"`
	ExecTime float64 `json:"exec_time"`
}

func getDir() {
	const project_name = "aipocket"
	const image_dir = "image"
	const model_type = "request_mark_cards_model_recognition"

	fileInfoArr, err := ioutil.ReadDir(PROJECT_DIR)
	check(err)
	var temName, temPath, logPath, timestampDirName string
	var files []string
	for _, fileInfo := range fileInfoArr {
		temName = fileInfo.Name()
		if !strings.Contains(temName, project_name) {
			continue
		}
		// 项目下面
		temPath = PROJECT_DIR + "/" + temName + "/" + image_dir + "/" + model_type
		imageInfoArr, _ := ioutil.ReadDir(PROJECT_DIR)
		imageDirs[temName] = map[string][]string{}
		for _, timestampDir := range imageInfoArr {
			// 时间戳文件夹
			timestampDirName = timestampDir.Name()
			logPath = temPath + "/" + timestampDirName
			logsInfoArr, _ := ioutil.ReadDir(logPath)
			files = make([]string, len(logsInfoArr))
			for _, eachFile := range logsInfoArr {
				files = append(files, logPath+"/"+eachFile.Name())
			}
			imageDirs[temName][timestampDirName] = files
		}
	}
}

func getTemplate() {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	check(err)
	var temName, temPath string
	for _, fileInfo := range fileInfoArr {
		temName = fileInfo.Name()
		if ext := path.Ext(temName); ext != ".html" {
			continue
		}
		temPath = TEMPLATE_DIR + "/" + temName
		log.Println("Loading Template:", temPath)
		t := template.Must(template.ParseFiles(temPath))
		templates[temName] = t
	}
}

func init() {
	getTemplate()
	getDir()
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Println("WARN: panic in %v - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	locals := make(INFO)
	for dir, path := range imageDirs {
		log.Println("Loading imageDirs:", path)
		var eachList []INFO
		fileInfoArr, err := ioutil.ReadDir(path)
		//checkError(err, w)
		//var listHtml string
		check(err)
		for _, eachDir := range fileInfoArr {
			log.Println("Loading eachDir:", eachDir)
			if !IsDir(eachDir.Name()) {
				continue
			}
			dir_infos := make(INFO)
			images := []string{}
			var logs []Result
			for _, fileInfo := range fileInfoArr {
				log.Println("Loading fileInfo:", fileInfo)
				if strings.Contains(fileInfo.Name(), "jpg") {
					images = append(images, fileInfo.Name())
				}
				if strings.Contains(fileInfo.Name(), "log") {
					file, err := os.Open(UPLOAD_DIR + "/" + fileInfo.Name())
					if err != nil {
						fmt.Println("文件打开失败 = ", err)
						continue
					}
					defer file.Close() // 关闭文本流
					var info Result
					// 创建json解码器
					decoder := json.NewDecoder(file)
					err = decoder.Decode(&info)
					if err != nil {
						fmt.Println("解码失败", err.Error())
					} else {
						fmt.Println("解码成功")
						fmt.Println(info)
						logs = append(logs, info)
					}
				}
			}
			dir_infos["images"] = images
			dir_infos["logs"] = logs
			eachList = append(eachList, dir_infos)
		}
		locals[dir] = eachList
	}

	renderHtml(w, "list.html", locals)
	//checkError(err, w)
	//io.WriteString(w, "<html><body><ol>"+listHtml+"</ol></body></html>")
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := imageId

	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	logId := r.FormValue("id")
	file, err := os.Open(logId)
	check(err)
	defer file.Close() // 关闭文本流
	var info Result
	// 创建json解码器
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&info)
	ret_json, _ := json.Marshal(info)
	w.Header().Set("Content-Type", "application/json")
	http.ServeFile(w, r, string(ret_json))
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsExist(err) {
		return false, nil
	}
	return false, err
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()

}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) error {
	//t, err := template.ParseFiles("template/" + tmpl + ".html")
	//if err != nil {
	//	return err
	//}
	//err = t.Execute(w, locals)
	//return err

	err := templates[tmpl].Execute(w, locals)
	return err

}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func checkError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("run go web server")
	http.HandleFunc("/", safeHandler(listHandler))
	http.HandleFunc("/log", safeHandler(logHandler))
	http.HandleFunc("/image", safeHandler(imageHandler))
	// http.HandleFunc("/upload", safeHandler(uploadHandler))
	err := http.ListenAndServe(HOST_ADDR, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
