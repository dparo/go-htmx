package main

import (
	"net/http"
	"path/filepath"

	"embed"
	"html/template"
	"io/fs"
	"log"

	"github.com/gorilla/mux"
)

type ServerContext struct {
	staticFs  fs.FS
	templates map[string]*template.Template
}

const (
	TEMPLATES_DIR = "templates"
	STATIC_DIR    = "static"
)

//go:embed templates/* static/*
var embedFs embed.FS
var ctx = ServerContext{
	nil,
	map[string]*template.Template{},
}

func InitServerCtx() error {
	staticContentSub, err := fs.Sub(embedFs, STATIC_DIR)
	if err != nil {
		return err
	}
	ctx.staticFs = staticContentSub

	return fs.WalkDir(embedFs, TEMPLATES_DIR, func(path string, tmpl fs.DirEntry, err error) error {
		if tmpl.IsDir() {
			return nil
		}
		log.Printf("Parsing template at filepath %s", path)
		pt, err := template.ParseFS(embedFs, path)
		if err != nil {
			return err
		}

		ctx.templates[path] = pt
		return nil
	})
}

// NOTE(d.paro): Thread safety
//   - ctx.templates[key] is thread safe access.
//   - template.Execute() is thread safe
func ExecuteTemplate(key string, w http.ResponseWriter, params map[string]interface{}) error {
	t := ctx.templates[filepath.Join(TEMPLATES_DIR, key)]
	return t.Execute(w, params)
}

func main() {
	err := InitServerCtx()
	if err != nil {
		log.Fatal(err)
	}

	m := mux.NewRouter()
	m.HandleFunc("/", index_route)
	m.HandleFunc("/index.html", index_route)
	m.HandleFunc("/index", index_route)
	m.HandleFunc("/api/v1/hello", hello)

	m.PathPrefix("/").Handler(
		http.FileServer(http.FS(ctx.staticFs)),
	)


	http.ListenAndServe(":8080", m)
}
