package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed public/*
var static embed.FS

func main() {
	http.Handle("/", http.FileServer(getFileSystem()))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// getFileSystem strips out prefixes
func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(static, "public")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(fsys)
}
