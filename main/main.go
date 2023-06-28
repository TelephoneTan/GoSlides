package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

type htmlHandler struct {
	fileServer http.Handler
}

const html = "html"
const worker = "worker.js"
const urlsPlaceHolder = "/*INSERT URLS*/"

var images = []string{
	"1218645.jpg",
	"1371163.jpg",
	"2082121.jpg",
	"2188578.jpg",
	"2188588.jpg",
	"2188651.jpg",
	"3b5d9eaa6b80cf4983b709a28662975c_720w.webp",
	"42277.jpg",
	"433c53c6212c9650583456df9833fcb1_r.jpg",
	"4e6anenk0-dark.png",
	"4e6anenk0-light.png",
	"645620.png",
	"660691.jpg",
	"666179.jpg",
	"685219.jpg",
	"890306.png",
	"967668.jpg",
	"967671.jpg",
	"967673.jpg",
	"967674.jpg",
	"967678.jpg",
	"967681.jpg",
	"967686.jpg",
	"967708.png",
	"967709.jpg",
	"967712.jpg",
	"earth.jpg",
	"impellent-dark.png",
	"star.png",
	"v2-927e127abdfaed47d8c172d7ff5dde44_r.jpg",
	"v2-a21eaf8e757c55b1b34214af2568cf11_r.jpg",
	"v2-a97ecfddfc8bb7334ec5ac9ab9a67b95_r.jpg",
	"v2-d858191577356128b31c88e186eea0db_720w.webp",
	"v2-e3ff581c4f19196977248b7297a044f5_r.jpg",
	"vanilla-default.webp",
}

var urls string

func init() {
	urlsBuilder := strings.Builder{}
	for _, image := range images {
		urlsBuilder.WriteString("\"/images/")
		urlsBuilder.WriteString(image)
		urlsBuilder.WriteString("\",")
	}
	urls = urlsBuilder.String()
}

func (h *htmlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch paths := strings.Split(r.URL.Path, "/"); {
	case strings.HasPrefix(paths[1], worker):
		content, _ := os.ReadFile(html + "/" + worker)
		content = []byte(strings.ReplaceAll(string(content), urlsPlaceHolder, urls))
		w.Header().Set("Content-Type", "application/javascript")
		_, _ = w.Write([]byte("// " + strings.TrimPrefix(paths[1], worker) + "\n"))
		_, _ = w.Write(content)
	default:
		h.fileServer.ServeHTTP(w, r)
	}
}

func main() {
	const addr = "localhost:8082"
	log.Printf("在浏览器中查看: http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, &htmlHandler{
		fileServer: http.FileServer(http.Dir(html)),
	}))
}
