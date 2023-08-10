package main

import "net/http"

func index_route(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Trigger", "evProcListRefresh")

	params := map[string]interface{}{
		"Name": "Davide",
	}
	ExecuteTemplate("index.xhtml", w, params)
}
