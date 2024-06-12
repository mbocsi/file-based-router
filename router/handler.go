package router

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cgi"
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
	log.Println(r.Method, r.URL.Path, r.RemoteAddr)
	switch r.Method {
	case "GET":
		ext := path.Ext(r.URL.Path)
		p := path.Join(h.root, r.URL.Path)
		var content []byte
		if ext != "" {
			var err error
			content, err = os.ReadFile(p)
			if err != nil {
				http.Error(w, "Content not found", http.StatusNotFound)
				return
			}

		}
		switch ext {
		case ".css":
			w.Header().Set("Content-Type", "text/css")
			fmt.Fprint(w, string(content))
		case ".js":
			w.Header().Set("Content-Type", "text/js")
			fmt.Fprint(w, string(content))
		case ".html":
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprint(w, string(content))
		case ".cgi":
			cgiHandler := cgi.Handler{Path: p}
			cgiHandler.ServeHTTP(w, r)
		case "":
			files, err := os.ReadDir(p)
			if err != nil {
				http.Error(w, "Path not found", http.StatusNotFound)
				return
			}
			f := ""
			prio := 5
			for _, file := range files {
				if prio == 0 {
					break
				}
				switch file.Name() {
				case "index.html":
					f = "index.html"
					prio = 0
				case "index.cgi":
					if 1 < prio {
						prio = 1
						f = "index.cgi"
					}
				case "index.js":
					if 2 < prio {
						prio = 2
						f = "index.js"
					}
				case "index.css":
					if 3 < prio {
						prio = 3
						f = "index.css"
					}
				}
			}
			if prio == 5 {
				http.Error(w, "Not found", http.StatusNotFound)
				return
			}
			r.URL.Path = path.Join(r.URL.Path, f)
			h.ServeHTTP(w, r)
		default:
			http.Error(w, "Unsupported file type", http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "Path not found", http.StatusNotFound)
	}
}
