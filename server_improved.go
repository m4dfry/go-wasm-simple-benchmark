package main

import (
	"fmt"
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// directory to serve.
var dir = "."
var d = http.Dir(dir)
var fileserver = http.FileServer(d)
var wasmFile = regexp.MustCompile("\\.wasm$")

// server port
var port = "3000"

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

// Use the Writer part of gzipResponseWriter to write the output.
func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func GzipHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the client can accept the gzip encoding.
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// The client cannot accept it, so return the output
			// uncompressed.
			next.ServeHTTP(w, r)
			return
		}
		// Set the HTTP header indicating encoding.
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		next.ServeHTTP(gzipResponseWriter{Writer: gz, ResponseWriter: w}, r)
	})
}

func CustomFs(w http.ResponseWriter, req *http.Request) {
	ruri := req.RequestURI
	if wasmFile.MatchString(ruri) {
		w.Header().Set("Content-Type", "application/wasm")
	}
	fileserver.ServeHTTP(w, req)
}

func main() {
	fmt.Println("Server start on port " + port + " ..")
	fshandler := http.HandlerFunc(CustomFs)
	http.Handle("/", GzipHandler(fshandler))
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

// ~ Thanks to https://www.lemoda.net/go/gzip-handler/index.html for the proper gzip handler
