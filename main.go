package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	LogInit(os.Stderr, os.Stdout, os.Stdout, os.Stderr)
	osd_sockets, err := filepath.Glob(os.Args[1])
	if err != nil {
		Error.Println("No sockets found exiting")
	}
	Debug.Println("Found following sockets", osd_sockets)
	c := make(chan struct {
		PerfCounter
		error
	})
	perfcounters := make([]PerfCounter, len(osd_sockets))

	for _, osd_socket := range osd_sockets {
		go func() {
			result, err := QueryAdminSocket(osd_socket)
			if err != nil {
				log.Fatal("Querying Socket failed with ", err)
			}
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
	GetClusterStatus(os.Args[2])
}
