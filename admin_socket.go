package main

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"net"
	"os"
	"path/filepath"
)

// Need to confirm this types with json structure definition from the socket
type journal_latency struct {
	sum   float32 `json:"sum"`
	count uint64  `json:"avgcount"`
}

type filestore struct {
	journal_latency journal_latency `json:"journal_latency"`
}

type leveldb struct {
	leveldb_get int `json:"leveldb_get"`
}

type perf_counter struct {
	filestore filestore `json:"filestore"`
	leveldb   leveldb   `json:"leveldb"`
}

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
	//	var perf perf_counter
	var perf interface{}
	_, err = conn.Read(buff)
	if err != nil {
		log.Println("Reading to socket returned ", err)
	} else {
		err = json.Unmarshal(buff, &perf)
		if err != nil {
			log.Println("Unmarshalling json errored with ", err)
		}
		log.Println(perf.(map[string]interface{}))
		// log.Println("journal latency", perf.filestore.journal_latency.sum)
		// log.Println("journal count", perf.filestore.journal_latency.count)
		// log.Println("journal count", perf.leveldb.leveldb_get)
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
