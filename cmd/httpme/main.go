package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	// Define command-line flags
	port := flag.Int("port", 0, "Port to listen on, default is a random port")
	address := flag.String("address", "localhost", "Address to bind to")
	rootPath := flag.String("path", "./", "Path to serve")
	version := flag.Bool("version", false, "Print version and exit")

	absRootPath, err := filepath.Abs(*rootPath)
	if err != nil {
		log.Fatalf("Error resolving absolute path of root directory: %v", err)
	}

	flag.Parse()

	if *version {
		fmt.Println("httpme v0.1.1")
		return
	}
	// Create a listener on the specified port or a free port
	listener, err := net.Listen("tcp", net.JoinHostPort(*address, strconv.Itoa(*port)))
	if err != nil {
		log.Fatalf("Error creating TCP listener: %v", err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatalf("Error closing TCP listener: %v", err)
		}
	}(listener)

	actualPort := strconv.Itoa(listener.Addr().(*net.TCPAddr).Port)

	if *address == "0.0.0.0" {
		fmt.Printf("Serving %s on all interfaces (0.0.0.0:%s) and accessible at http://localhost:%s\n", absRootPath, actualPort, actualPort)
	} else {
		fmt.Printf("Serving %s on http://%s:%s\n", absRootPath, *address, actualPort)
	}

	fs := http.FileServer(http.Dir(*rootPath))
	http.Handle("/", logRequest(fs))

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)
	}
}

func logRequest(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		log.Printf("[%s] %s %s %s", r.RemoteAddr, r.Method, r.URL, time.Since(start))
	}
}
