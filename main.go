package main

import (
	"flag"
	"log"
	"net/http"
)

var port = flag.String("port", ":80", "Port to listen on")
var dir = flag.String("dir", "./", "Directory to share")

func init() {
	flag.Parse()
}

func main() {
	fs := http.FileServer(http.Dir(*dir))
	log.Println("Serving", *dir, "on HTTP port", *port)
	err := http.ListenAndServe(*port, loggingMiddleware(fs))
	if err != nil {
		log.Fatal(err)
	}
}

// 日志中间件
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
