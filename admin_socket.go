package main

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"net"
)

func QueryAdminSocket(path string) error {
	log.Println("Dialing.....", path)
	conn, err := net.Dial("unix", path)
	defer conn.Close()
	if err != nil {
		log.Println("Connecting to socket returned ", err)
		return err
	} else {
		log.Println("Connection successful!")
	}

	n, err := conn.Write([]byte("{\"prefix\":\"perf dump\"}\x00"))

	if err != nil {
		log.Println("Writing to socket returned ", err)
		return err
	} else {
		log.Println("Writing successful!", n)
	}
	len_buff := make([]byte, 4)
	_, err = conn.Read(len_buff)
	if err != nil {
		log.Println("Reading length from socket failed", err)
		return err
	}
	buff := make([]byte, binary.BigEndian.Uint32(len_buff))
	var perf PerfCounter
	_, err = conn.Read(buff)
	if err != nil {
		log.Println("Reading to socket returned ", err)
	} else {
		err = json.Unmarshal(buff, &perf)
		if err != nil {
			log.Println("Unmarshalling json errored with ", err)
		}
		log.Println(perf)
	}

	return err
}
