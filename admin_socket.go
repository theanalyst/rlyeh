package main

import (
	"fmt"
	"net"
	"os"
)

func QueryAdminSocket(path string) error {
	fmt.Println("Dialing.....", path)
	conn, err := net.Dial("unix", path)
	defer conn.Close()
	if err != nil {
		fmt.Println("Connecting to socket returned ", err)
	} else {
		fmt.Println("Connection successful!")
	}

	n, err := conn.Write([]byte("{\"prefix\":\"perf dump\"}\x00"))

	if err != nil {
		fmt.Println("Writing to socket returned ", err)
	} else {
		fmt.Println("Writing successful!", n)
	}

	buff := make([]byte, 102400)
	var result string

	n, err = conn.Read(buff)
	if err != nil {
		fmt.Println("Reading to socket returned ", err)
	} else {
		result = string(buff[:n])
		fmt.Println("Reading successful!", result)
	}

	return err
}

func main() {
	err := QueryAdminSocket(os.Args[1])
	if err != nil {
		fmt.Println("We errored out", err)
	}
}
