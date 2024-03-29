# httpme - Quick HTTP Server
![Build Status](https://github.com/delta9000/httpme/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/delta9000/httpme)](https://goreportcard.com/report/github.com/delta9000/httpme)
[![License: BSD 3-Clause](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)]( https://opensource.org/licenses/BSD-3-Clause)

httpme is a simple command-line HTTP server that serves files from a specified directory. 

It can automatically select an available port or use a port defined by the user.

## Usage
* `-port`: Specify the port to listen on (default: random available port).
* `-address`: Specify the address to bind to (default: localhost).
* `-path`: Specify the directory path to serve (default: current directory).
* `-tls`: Enable TLS (default: false).
* `-cert`: Specify the TLS certificate file (default: none).
* `-key`: Specify the TLS key file (default: none).
* `-help`: Display help message.
* `-version`: Display version information.


### Examples
- Serve the current directory on a random free port on localhost:
```bash
httpme 
```

- Serve files from /path/to/directory on http://0.0.0.0:8080/
```bash
httpme -port=8080 -address=0.0.0.0 -path=/path/to/directory
```

- Serve files from /path/to/directory on https://
```bash
httpme -tls -cert /path/to/cert.pem -key /path/to/key.pem -path /path/to/directory
```

## Installing
To install httpme, use the following command:
```bash
go install github.com/delta9000/httpme/cmd/httpme@latest
```

## License
BSD 3-Clause License