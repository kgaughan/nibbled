package main

import (
	"flag"
	"log"
	"net"
	"time"
)

var bind, root, userDir string
var cgi bool

func main() {
	flag.StringVar(&bind, "bind", ":70", "Interface/port to bind to.")
	flag.StringVar(&root, "root", "/srv/gopher", "Directory to serve from.")
	flag.BoolVar(&cgi, "cgi", false, "Allow CGI scripts.")
	flag.StringVar(&userDir, "userdir", "", "Expose user directories over gopher.")
	flag.Parse()

	ln, err := net.Listen("tcp", bind)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	var delay time.Duration
	for {
		conn, err := ln.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if delay == 0 {
					delay = 5 * time.Millisecond
				} else {
					delay *= 2
				}
				if max := 1 * time.Second; delay > max {
					delay = max
				}
				log.Printf("Accept error: %v; retrying in %v", err, delay)
				time.Sleep(delay)
			} else {
				panic(err)
			}
		} else {
			delay = 0
			go handleConnection(conn)
		}
	}
}

func handleConnection(conn net.Conn) {
	log.Printf("Connection accepted")
}
