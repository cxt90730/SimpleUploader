package main

import (
	"html/template"
	"io"
	"net/http"
	"os"
)

func h(w http.ResponseWriter, r *http.Request) {
	file, handle, err := r.FormFile("upload")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	f, err := os.OpenFile("MyFile/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	defer file.Close()
	defer f.Close()
	w.Write([]byte("ok"))
	return
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./pages/index.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	t.Execute(w, nil)
}

func main() {

    if  _ , err := os.Stat("MyFile"); err != nil{
        os.Mkdir("MyFile", 0777)
    }

	http.HandleFunc("/", index)
	http.HandleFunc("/upload", h)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}
