package main

import "fmt"

type Entry struct {
	Type     byte
	Name     string
	Selector string
	Host     string
	Port     string
}

func (e Entry) String() string {
	return fmt.Sprintf("%c%s\t%s\t%s\t%s\r\n", e.Type, e.Name, e.Selector, e.Host, e.Port)
}
