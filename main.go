package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/alecthomas/kingpin"
	"os"
	"path/filepath"
)

var (
	osd_sockets = kingpin.Flag("osd_sock", "pattern for osd sockets").Default("/var/run/ceph/osd*asok").String()
	mon_sockets = kingpin.Flag("mon_sock", "pattern for mon sockets").Default("/var/run/ceph/mon*asok").String()
	conf        = kingpin.Flag("conf", "path to configuration file").Default("/etc/ceph/ceph.conf").String()
	verbose     = kingpin.Flag("verbose", "verbose logging").Short('v').Bool()
	logfile     = kingpin.Flag("log", "path to logfile").Default("/var/log/ceph/rlyeh.log").String()
)

type perf struct {
	PerfCounter
	error
}

func main() {

	kingpin.Version("0.1")
	kingpin.Parse()

	f, err := os.OpenFile(*logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	LogInit(*verbose, f)
	osd_sockets, err := filepath.Glob(*osd_sockets)
	if err != nil {
		log.Error("No sockets found exiting")
	}
	log.Info("Found following sockets", osd_sockets)
	c := make(chan perf)
	perfcounters := make([]PerfCounter, len(osd_sockets))

	for _, osd_socket := range osd_sockets {
		go func() {
			result, err := QueryAdminSocket(osd_socket)
			if err != nil {
				log.Fatal("Querying Socket failed with ", err)
			}
			c <- perf{*result, err}
		}()
		query := <-c
		if query.error == nil {
			perfcounters = append(perfcounters, query.PerfCounter)
		}

	}
	GetClusterStatus(*conf)
	GetClusterIOPs(*conf)
}
