package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
)

func QueryAdminSocket(path string) error {
	fmt.Println("Dialing.....", path)
	log.Println("Dialing.....", path)
	conn, err := net.Dial("unix", path)
	defer conn.Close()
	if err != nil {
		log.Println("Connecting to socket returned ", err)
	} else {
		log.Println("Connection successful!")
	}

	n, err := conn.Write([]byte("{\"prefix\":\"perf dump\"}\x00"))

	if err != nil {
		log.Println("Writing to socket returned ", err)
	} else {
		log.Println("Writing successful!", n)
	}

	buff := make([]byte, 102400)
	var result string

	n, err = conn.Read(buff)
	if err != nil {
		fmt.Println("Reading to socket returned ", err)
		log.Println("Reading to socket returned ", err)
	} else {
		result = string(buff[:n])
		fmt.Println("Reading successful!", result)
	}

	return err
}

func main() {
	osd_sockets := make([]string, 100)
	osd_sockets, err := filepath.Glob(os.Args[1])
	if err != nil {
		log.Println("No sockets fount exiting")
	}
	log.Println(osd_sockets)
	for _, osd_socket := range osd_sockets {
		QueryAdminSocket(osd_socket)
	}
}
