package main

import (
	"encoding/binary"
	"encoding/json"
	"net"
)

func QueryAdminSocket(path string) (*PerfCounter, error) {
	Debug.Println("Dialing.....", path)
	conn, err := net.Dial("unix", path)
	defer conn.Close()
	if err != nil {
		Error.Println("Connecting to socket returned ", err)
		return nil, err
	} else {
		Info.Println("Connection successful!")
	}

	n, err := conn.Write([]byte("{\"prefix\":\"perf dump\"}\x00"))

	if err != nil {
		Error.Println("Writing to socket returned ", err)
		return nil, err
	} else {
		Debug.Println("Writing successful!", n)
	}
	len_buff := make([]byte, 4)
	_, err = conn.Read(len_buff)
	if err != nil {
		Error.Println("Reading length from socket failed", err)
		return nil, err
	}
	buff := make([]byte, binary.BigEndian.Uint32(len_buff))
	_, err = conn.Read(buff)

	if err != nil {
		Error.Println("Reading to socket returned ", err)
	} else {
		perf := &PerfCounter{}
		err = json.Unmarshal(buff, &perf)
		if err != nil {
			Error.Println("Unmarshalling json errored with ", err)
		}
		Debug.Println(perf)
		return perf, nil
	}

	return nil, err
}
