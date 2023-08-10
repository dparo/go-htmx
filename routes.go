package main

import "net/http"

func index_route(w http.ResponseWriter, r *http.Request) {
	t := templates["templates/index.xhtml"]

	m := map[string]string{"Name": "foobar"}

	w.Header().Add("HX-Trigger", "evProcListRefresh")
	t.Execute(w, m)
}
