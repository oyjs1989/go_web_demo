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
	"sort"
	"strings"
)

const (
	TEMPLATE_DIR = "./template"
	HOST_ADDR    = ":5000"
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

type ByModTime []os.FileInfo

func (fis ByModTime) Len() int {
	return len(fis)
}

func (fis ByModTime) Swap(i, j int) {
	fis[i], fis[j] = fis[j], fis[i]
}

func (fis ByModTime) Less(i, j int) bool {
	return fis[j].ModTime().Before(fis[i].ModTime())
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

func readJsonFile(filePaht string) string {
	file, err := os.Open(filePaht)
	check(err)
	defer file.Close() // 关闭文本流
	var info Result
	// 创建json解码器
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&info)
	ret_json, _ := json.Marshal(info)
	return string(ret_json)
}

func getDir() {
	const project_name = "aipocket_new"
	const image_dir = "image"
	const model_type = "request_mark_cards_model_recognition"
	const logName = "log"
	const ipynb = "ipynb_checkpoints"

	fileInfoArr := SortFile(PROJECT_DIR)
	var temName, temPath, logPath, timestampDirName string
	var files []string
	for _, fileInfo := range fileInfoArr {
		temName = fileInfo.Name()
		if !strings.Contains(temName, project_name) {
			continue
		}
		// 项目下面
		temPath = PROJECT_DIR + "/" + temName + "/" + image_dir + "/" + model_type
		imageInfoArr := SortFile(temPath)
		imageDirs[temName] = map[string][]string{}
		for _, timestampDir := range imageInfoArr {
			// 时间戳文件夹
			timestampDirName = timestampDir.Name()
			logPath = temPath + "/" + timestampDirName
			logsInfoArr := SortFile(logPath)
			files = make([]string, len(logsInfoArr))
			for _, eachFile := range logsInfoArr {
				var fileName = eachFile.Name()
				if strings.Contains(fileName, logName) {
					fileName = readJsonFile(logPath + "/" + fileName)
				} else if strings.Contains(fileName, ipynb) {
					continue
				} else {
					fileName = logPath + "/" + fileName
				}
				files = append(files, fileName)
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
		funcs := map[string]any{
			"contains":  strings.Contains,
			"hasPrefix": strings.HasPrefix,
			"hasSuffix": strings.HasSuffix}
		t := template.Must(template.New(temName).Funcs(funcs).ParseFiles(temPath))
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
	getDir()
	renderHtml(w, "list.html", imageDirs)
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

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]map[string][]string) error {
	//t, err := template.ParseFiles("template/" + tmpl + ".html")
	//if err != nil {
	//	return err
	//}
	//err = t.Execute(w, locals)
	//return err
	err := templates[tmpl].Execute(w, locals)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
