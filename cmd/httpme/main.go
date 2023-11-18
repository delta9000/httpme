package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
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
	tls := flag.Bool("tls", false, "Enable HTTPS")
	cert := flag.String("cert", "", "Path to TLS certificate")
	key := flag.String("key", "", "Path to TLS key")

	flag.Parse()

	if _, err := os.Stat(*rootPath); os.IsNotExist(err) {
		log.Fatalf("Error: Path does not exist: %s", *rootPath)
	}

	if _, err := os.Open(*rootPath); err != nil {
		log.Fatalf("Error: Path is not readable: %s", *rootPath)
	}

	absRootPath, err := filepath.Abs(*rootPath)
	if err != nil {
		log.Fatalf("Error resolving absolute path of root directory: %v", err)
	}

	if tls != nil && *tls {
		if cert == nil || *cert == "" {
			log.Fatalf("Error: HTTPS enabled but no certificate provided")
		}
		if key == nil || *key == "" {
			log.Fatalf("Error: HTTPS enabled but no key provided")
		}

		if _, err := os.Stat(*cert); os.IsNotExist(err) {
			log.Fatalf("Error: TLS certificate file does not exist: %s", *cert)
		}
		if _, err := os.Stat(*key); os.IsNotExist(err) {
			log.Fatalf("Error: TLS key file does not exist: %s", *key)
		}
	}

	if !*tls && (*cert != "" || *key != "") {
		log.Printf("Warning: TLS certificate or key provided but -tls not provided")
	}

	if *version {
		fmt.Println("httpme v0.1.2")
		return
	}

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
	protocol := "http"
	if *tls {
		protocol = "https"
	}

	if *address == "0.0.0.0" {
		fmt.Printf("Serving %s on all interfaces (0.0.0.0:%s) and accessible at %s://localhost:%s\n", absRootPath, actualPort, protocol, actualPort)
	} else {
		fmt.Printf("Serving %s on %s://%s:%s\n", absRootPath, protocol, *address, actualPort)
	}

	fs := http.FileServer(http.Dir(*rootPath))
	http.Handle("/", logRequest(fs))

	if *tls {
		err = http.ServeTLS(listener, nil, *cert, *key)
	} else {
		err = http.Serve(listener, nil)
	}

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}

func logRequest(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		log.Printf("[%s] %s %s %s", r.RemoteAddr, r.Method, r.URL, time.Since(start))
	}
}
