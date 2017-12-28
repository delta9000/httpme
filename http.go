package main

import (
	"fmt"
	"github.com/phayes/freeport"
	"net"
	"net/http"
	"strconv"
)

func main() {
	port, _ := freeport.GetFreePort()
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)
	fmt.Println("Now serving current directory on http://" + GetOutboundIP().String() + ":" + strconv.Itoa(port) + "/\nCTRL+C to exit")
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func GetOutboundIP() net.IP {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}
