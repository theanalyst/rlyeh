package main

import (
	"github.com/alecthomas/kingpin"
	"log"
	"os"
	"path/filepath"
)

var (
	osd_sockets = kingpin.Flag("osd_sock", "pattern for osd sockets").Default("/var/run/ceph/osd*asok").String()
	mon_sockets = kingpin.Flag("mon_sock", "pattern for mon sockets").Default("/var/run/ceph/mon*asok").String()
	conf        = kingpin.Flag("conf", "path to configuration file").Default("/etc/ceph/ceph.conf").String()
)

func main() {

	kingpin.Version("0.1")
	kingpin.Parse()

	LogInit(os.Stderr, os.Stdout, os.Stdout, os.Stderr)
	osd_sockets, err := filepath.Glob(*osd_sockets)
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
	GetClusterStatus(*conf)
	GetClusterIOPs(*conf)
}
