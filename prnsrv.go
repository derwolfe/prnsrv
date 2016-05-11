package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"strconv"
	"strings"
)

var port int

func main() {
	flag.IntVar(&port, "port", 8000, "the port on which to listen")
	flag.Parse()

	log.Printf("starting on %d\n", port)
	server, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if server == nil {
		log.Fatal(err)
	}
	conns := clientConns(server)
	for {
		go handleConn(<-conns)
	}
}

func clientConns(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	i := 0
	go func() {
		for {
			client, err := listener.Accept()
			if client == nil {
				log.Fatal(err)
				continue
			}
			i++
			log.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}

func handleConn(client net.Conn) {
	b := bufio.NewReader(client)
	for {
		line, err := b.ReadString('\n')
		if err != nil { // EOF, or worse
			break
		}
		log.Println(strings.TrimSpace(line))
	}
}
