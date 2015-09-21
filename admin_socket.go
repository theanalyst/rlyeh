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
	len := make([]byte, 4)

	buff := make([]byte, 102400)
	n, err = conn.Read(buff)
	if err != nil {
		fmt.Println("Reading to socket returned ", err)
	} else {
		fmt.Println("Reading successful!", string(buff[0:n]))
	}

	return err
}

func main() {
	err := QueryAdminSocket(os.Args[1])
	if err != nil {
		fmt.Println("We errored out", err)
	}
}
