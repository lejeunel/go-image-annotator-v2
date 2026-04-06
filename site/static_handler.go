package site

import (
	"net/http"
)

func RegisterStaticFiles(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("site/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
}
