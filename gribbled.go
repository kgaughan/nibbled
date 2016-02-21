package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"strings"
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
	defer conn.Close()

	log.Printf("Connection accepted")

	reader := bufio.NewReader(conn)
	if line, err := reader.ReadString('\n'); err != nil {
		log.Print(err)
	} else {
		// Format is <selector>TAB<query>CRLF
		parts := strings.SplitN(strings.TrimRight(line, "\r\n"), "\t", 2)
		log.Printf("%q", parts)
		conn.Write([]byte("iMessage\t\t\t\r\n"))
	}
}
