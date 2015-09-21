package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	osd_sockets, err := filepath.Glob(os.Args[1])
	if err != nil {
		log.Println("No sockets fount exiting")
	}
	log.Println(osd_sockets)
	for _, osd_socket := range osd_sockets {
		QueryAdminSocket(osd_socket)
	}
}
