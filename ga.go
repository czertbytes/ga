package ga

import (
	"os"
	"os/signal"
)

var storage *Storage

func init() {
	storage = NewStorage()
}

func Run(listenUDPServer, listenHTTPServer string) {
	go startUDPServer(listenUDPServer)
	go startHTTPServer(listenHTTPServer)

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, os.Interrupt, os.Kill)
	<-quitChan
}
