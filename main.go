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
	staticFs fs.FS
	template *template.Template
}

const (
	TEMPLATES_DIR = "templates"
	STATIC_DIR    = "static"
)

//go:embed templates/* static/*
var embedFs embed.FS
var ctx = ServerContext{nil, nil}

func InitServerCtx() error {
	staticContentSub, err := fs.Sub(embedFs, STATIC_DIR)
	if err != nil {
		return err
	}
	ctx.staticFs = staticContentSub

	tmpl := template.New("")

	fs.WalkDir(embedFs, TEMPLATES_DIR, func(path string, root fs.DirEntry, err error) error {
		if root.IsDir() {
			return nil
		}
		log.Printf("Parsing template at filepath %s", path)
		b, err := fs.ReadFile(embedFs, path)
		if err != nil {
			return nil
		}

		relpath, err := filepath.Rel(TEMPLATES_DIR, path)
		if err != nil {
			return nil
		}

		// _, err = tmpl.ParseFS(embedFs, path)
		t := tmpl.New(relpath)
		_, err = t.Parse(string(b))

		if err != nil {
			return nil
		}

		return nil
	})

	ctx.template = tmpl

	log.Printf("%s\n", tmpl.DefinedTemplates())

	return nil
}

// NOTE(d.paro): Thread safety
//   - template.{Execute, ExcuteTemplate}() is thread safe
func ExecuteTemplate(key string, w http.ResponseWriter, params map[string]interface{}) error {
	return ctx.template.ExecuteTemplate(w, key, params)
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
