package ga

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func startUDPServer(listenAddress string) {
	addr, err := net.ResolveUDPAddr("udp", listenAddress)
	if err != nil {
		panic(err)
	}

	listen, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	log.Printf("http: UDP server listening on %s\n", listenAddress)

	buf := make([]byte, 120)
	for {
		n, err := listen.Read(buf)
		if err != nil {
			log.Println("http: Reading from UDP failed. Error %s", err.Error())
			continue
		}

		buffer := bytes.NewBuffer(buf[0:n])
		content := strings.Trim(buffer.String(), " \r\n")
		if len(content) > 0 {
			log.Printf("http: UDP content '%s'\n", content)

			go func(value string) {
				timestamp, userId, err := NewParser(value).ParseTimestampUserId()
				if err != nil {
					return
				}

				storage.Insert(NewRecord(timestamp, userId))
			}(content)
		}
	}
}

func startHTTPServer(listenAddress string) {
	log.Printf("http: HTTP server listening on %s\n", listenAddress)

	http.HandleFunc("/users", usersHandler)
	err := http.ListenAndServe(listenAddress, nil)
	if err != nil {
		panic(err)
	}
}

func usersHandler(responseWriter http.ResponseWriter, request *http.Request) {
	start, end, err := parseQuery(request.URL.Query())
	if err != nil {
		fmt.Fprintf(responseWriter, err.Error())
		return
	}

	records := storage.Find(start, end)
	uniqUserIds := records.UniqueUserIds()

	log.Printf("http: Timestamp range <%d,%d) has %d unique users", start, end, uniqUserIds)

	fmt.Fprintf(responseWriter, "%d", uniqUserIds)
}

func parseQuery(params map[string][]string) (int64, int64, error) {
	start, err := parseInt64Param("start", params)
	if err != nil {
		return 0, 0, err
	}

	end, err := parseInt64Param("end", params)
	if err != nil {
		return 0, 0, err
	}

	return start, end, nil
}

func parseInt64Param(name string, params map[string][]string) (int64, error) {
	valueStr, found := params[name]
	if !found {
		return 0, fmt.Errorf("http: Bad request! Required parameter '%s' is missing!", name)
	}

	if len(valueStr) != 1 {
		return 0, fmt.Errorf("http: Bad request! Parameter '%s' is not valid!", name)
	}

	value, err := strconv.ParseInt(valueStr[0], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("http: Bad request! Parameter '%s' is not valid!", name)
	}

	return value, nil
}
