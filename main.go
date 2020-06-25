package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"os"
	"path"
	"strings"
)

var Version string

var root string
var hostname string
var port string

func init() {
	defaultHostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&root, "root", "/srv/gopher", "Root directory of server")
	flag.StringVar(&hostname, "hostname", defaultHostname, "Hostname to present")
	flag.StringVar(&port, "port", "70", "Port to bind to")
}

func main() {
	flag.Parse()

	if _, err := os.Stat(root); os.IsNotExist(err) {
		log.Fatalf("Root directory '%v' not found", root)
	}

	ln, err := net.Listen("tcp", hostname+":"+port)
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
	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		// Format is <selector>TAB<query>CRLF
		parts := strings.SplitN(strings.TrimRight(scanner.Text(), "\r\n"), "\t", 2)
		log.Printf("%q", parts)
		conn.Write([]byte("iMessage\t\t\t\r\n"))
	} else if err := scanner.Err(); err != nil {
		log.Print(err)
	}
}
