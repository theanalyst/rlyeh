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
	c := make(chan struct {
		PerfCounter
		error
	})
	perfcounters := make([]PerfCounter, len(osd_sockets))

	for _, osd_socket := range osd_sockets {
		go func() {
			result, err := QueryAdminSocket(osd_socket)
			c <- struct {
				PerfCounter
				error
			}{*result, err}
		}()
		query := <-c
		if query.error == nil {
			perfcounters = append(perfcounters, query.PerfCounter)
		}

	}

}
