package main

import (
	"encoding/binary"
	"encoding/json"
	"net"
)

func QueryAdminSocket(path string) (*PerfCounter, error) {
	conn, err := net.Dial("unix", path)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	Debug.Println("Connected to socket ", path)

	_, err = conn.Write([]byte("{\"prefix\":\"perf dump\"}\x00"))

	if err != nil {
		return nil, err
	}

	len_buff := make([]byte, 4)
	_, err = conn.Read(len_buff)
	if err != nil {
		return nil, err
	}
	buff := make([]byte, binary.BigEndian.Uint32(len_buff))
	_, err = conn.Read(buff)

	if err != nil {
		return nil, err
	} else {
		perf := &PerfCounter{}
		err = json.Unmarshal(buff, &perf)
		if err != nil {
			Error.Println("Unmarshalling json errored with ", err)
		}
		Debug.Println(perf)
		return perf, nil
	}
}
