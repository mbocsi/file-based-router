package router

import (
	"fmt"
	"net/http"
	"os"
	"path"
)

type FBHandler struct {
	root string
}

func NewFBHandler(root string) http.Handler {
	return &FBHandler{root}
}

func (h *FBHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ext := path.Ext(r.URL.Path)
		var p string
		if ext == "" {
			p = path.Join(h.root, r.URL.Path, "index.html")
		} else {
			p = path.Join(h.root, r.URL.Path)
		}
		content, err := os.ReadFile(p)
		if err != nil {
			http.Error(w, "Content not found", http.StatusNotFound)
		}
		switch ext {
		case ".css":
			w.Header().Set("Content-Type", "text/css")
		case ".js":
			w.Header().Set("Content-Type", "text/js")
		case "":
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
		default:
			w.Header().Set("Content-Type", "text")
		}
		fmt.Fprint(w, string(content))
	default:
		http.Error(w, "Path not found", http.StatusNotFound)
	}
}
