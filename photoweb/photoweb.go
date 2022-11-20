package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
)

const (
	UPLOAD_DIR = "./uploads"
	TEMPLATE_DIR = "./template"
	HOST_ADDR = ":8090"
)


var templates = make(map[string]*template.Template)

func init() {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	check(err)
	var temName, temPath string
	for _, fileInfo := range fileInfoArr {
		temName = fileInfo.Name()
		if ext := path.Ext(temName); ext != ".html" {
			continue
		}
		temPath = TEMPLATE_DIR + "/" + temName
		log.Println("Loading template:", temPath)
		t := template.Must(template.ParseFiles(temPath))
		templates[temName] = t
	}
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
	fileInfoArr, err := ioutil.ReadDir("./uploads")
	//checkError(err, w)
	//var listHtml string
	check(err)
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
		//imgId := fileInfo.Name()
		//listHtml += "<li><a href=\"/view?id="+imgId+"\">"+imgId+"</a></li>"
	}
	locals["images"] = images
	renderHtml(w, "list.html", locals)
	//checkError(err, w)
	//io.WriteString(w, "<html><body><ol>"+listHtml+"</ol></body></html>")


}
func  uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderHtml(w, "upload.html", nil)
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		check(err)

		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		check(err)
		defer t.Close()

		_, err = io.Copy(t, f)
		check(err)

		http.Redirect(w, r, "/view?id=" + filename, http.StatusFound)
	}
}


func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId

	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}


func deleteHandler(w http.ResponseWriter, r *http.Request) {

		name := r.FormValue("name")
		fmt.Println("Will delete: ", UPLOAD_DIR + "/" + name)

		pathName := UPLOAD_DIR + "/" + name
		_, err := PathExists(pathName)
		check(err)

		if err := os.Remove(pathName); err != nil {
			check(err)
		}
		fmt.Println("delete success!")
		http.Redirect(w, r, "/", http.StatusFound)

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
		panic(err)
	}

}

func main() {
	fmt.Println("test main")
	http.HandleFunc("/", safeHandler(listHandler))
	http.HandleFunc("/del", safeHandler(deleteHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	err := http.ListenAndServe(HOST_ADDR, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
