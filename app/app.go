package main

import (
	"flag"

	"github.com/czertbytes/ga"
)

func main() {
	var (
		listenUDPServer  = flag.String("listen-udp", ":4263", "UDP server address")
		listenHTTPServer = flag.String("listen-http", ":8080", "HTTP server address")
	)
	flag.Parse()

	ga.Run(*listenUDPServer, *listenHTTPServer)
}
