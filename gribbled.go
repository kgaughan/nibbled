package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"os"
	"strings"
)

var bind = flag.String("bind", ":70", "Interface/port to bind to")
var root = flag.String("root", "/srv/gopher", "Root directory of server")

func main() {
	flag.Parse()

	if _, err := os.Stat(*root); os.IsNotExist(err) {
		log.Fatalf("Root directory '%v' not found", *root)
	}

	ln, err := net.Listen("tcp", *bind)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		if conn, err := ln.Accept(); err != nil {
			log.Println(err)
		} else {
			go handleConnection(conn)
		}
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
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
