package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"
)

var version = "invalid"

var (
	printVersion bool
	root         string
	hostname     string
	port         string
)

var errCantResolve = errors.New("can't resolve selector")

func init() {
	flag.BoolVarP(&printVersion, "version", "V", false, "Print version and exit")
	flag.StringVarP(&root, "root", "r", "/srv/gopher", "Root directory of server")
	flag.StringVarP(&hostname, "hostname", "h", "localhost", "Hostname to present")
	flag.StringVarP(&port, "port", "p", "70", "Port to bind to")

	flag.Usage = func() {
		name := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "%s - a server for the Gopher protocol.\n\n", name)
		fmt.Fprintf(os.Stderr, ":\n\n")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if printVersion {
		fmt.Println(version)
		return
	}

	if _, err := os.Stat(root); os.IsNotExist(err) {
		log.Fatalf("Root directory '%v' not found", root)
	}

	ln, err := net.Listen("tcp", net.JoinHostPort(hostname, port))
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		if conn, err := ln.Accept(); err != nil {
			log.Print(err)
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
		if strings.HasPrefix(parts[0], "..") {
			if err := writeError(conn, "Bad selector"); err != nil {
				log.Print(err)
			}
			return
		}

		localPath := filepath.Join(root, parts[0])
		if err := resolve(conn, localPath, parts[0]); err != nil {
			log.Print(err)
			if err := writeError(conn, err.Error()); err != nil {
				log.Print(err)
			}
		}
	} else if err := scanner.Err(); err != nil {
		log.Print(err)
	}
}

func resolve(out io.Writer, localPath, selector string) error {
	if fi, err := os.Stat(localPath); err != nil {
		return fmt.Errorf("%q: %w", selector, errCantResolve)
	} else if fi.IsDir() {
		gophermap := filepath.Join(localPath, "gophermap")
		if _, err := os.Stat(gophermap); err == nil {
			if err := loadGopherMap(out, localPath, selector); err != nil {
				return err //nolint:wrapcheck
			}
			if _, err := out.Write([]byte(".\r\n")); err != nil {
				return err //nolint:wrapcheck
			}
		}
		catalogue, err := listDirectory(localPath, selector)
		if err != nil {
			return err
		}
		if _, err := write(out, catalogue); err != nil {
			return err
		}
		return nil
	}

	return sendFile(out, localPath)
}

func sendFile(out io.Writer, localPath string) error {
	f, err := os.Open(localPath)
	if err != nil {
		return err //nolint:wrapcheck
	}
	defer f.Close()
	if _, err := io.Copy(out, f); err != nil {
		return err //nolint:wrapcheck
	}
	return nil
}
