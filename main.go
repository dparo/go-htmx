package main

import (
	"net/http"

	"embed"
	"html/template"
	"io/fs"
	"log"

	"github.com/gorilla/mux"
)

const (
	templatesDir = "templates"
)

var (
	//go:embed templates/*
	files     embed.FS
	templates map[string]*template.Template
)

func LoadVFS() error {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	return fs.WalkDir(files, templatesDir, func(path string, tmpl fs.DirEntry, err error) error {
		if tmpl.IsDir() {
			return nil
		}
		log.Printf("Parsing template at filepath %s", path)
		pt, err := template.ParseFS(files, path)
		if err != nil {
			return err
		}

		templates[path] = pt
		return nil
	})
}

func main() {
	err := LoadVFS()
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", index_route)
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
