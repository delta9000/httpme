# httpme - Quick HTTP Server
httpme is a simple command-line HTTP server that serves files from a specified directory. 

It can automatically select an available port or use a port defined by the user.

## Usage
* `-port`: Specify the port to listen on (default: random available port).
* `-address`: Specify the address to bind to (default: localhost).
* `-path`: Specify the directory path to serve (default: current directory).

### Examples
- Serve the current directory on a random free port on localhost:
```bash
httpme 
```

- Serve files from /path/to/directory on http://0.0.0.0:8080/
```bash
httpme -port=8080 -address=0.0.0.0 -path=/path/to/directory
```

## Installing
To install httpme, use the following command:
```bash
go install github.com/delta9000/httpme@latest
```

## License
BSD 3-Clause License