package main

import "net/http"

type Cmd struct {
	Pid  int
	Argv string
}

func index_route(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Trigger", "evProcListRefresh")

	params := map[string]interface{}{
		"Cmds":    []Cmd{{0, "systemd"}, {1, "foo"}},
		"Signals": []string{"SIGTERM", "SIGINT", "SIGKILL"},
	}

	ExecuteTemplate("index.gohtml", w, params)
}
