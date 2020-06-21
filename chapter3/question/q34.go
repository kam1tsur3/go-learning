package main

import (
	"archive/zip"
	"io"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("content-Dispostion", "attachment; filename=ascii_sample.zip")

	z_wtr := zip.NewWriter(w)
	defer z_wtr.Close()

	f, _ := z_wtr.Create("test")
	reader := strings.NewReader("test zip on web")
	io.Copy(f, reader)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
