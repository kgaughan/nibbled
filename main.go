package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"path"
	"path/filepath"
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

		// Make sure the selector is safe
		parts[0] = path.Clean(parts[0])
		if parts[0] == "." {
			parts[0] = ""
		}
		log.Printf("%q", parts)
		if strings.HasPrefix(parts[0], "..") {
			writeError(conn, "Bad selector")
			return
		}

		localPath := filepath.Join(root, parts[0])
		if fi, err := os.Stat(localPath); err == nil {
			if fi.IsDir() {
				if err := listDirectory(conn, localPath, parts[0]); err != nil {
					log.Print(err)
					return
				}
			} else if f, err := os.Open(localPath); err != nil {
				log.Print(err)
				writeError(conn, "Could not open")
				return
			} else {
				defer f.Close()
				if _, err := io.Copy(conn, f); err != nil {
					log.Print(err)
					return
				}
			}
		} else {
			log.Print(err)
			writeError(conn, "Bad selector")
			return
		}
	} else if err := scanner.Err(); err != nil {
		log.Print(err)
	}
}
