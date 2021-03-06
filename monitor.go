package main

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/ceph/go-ceph/rados"
)

// Return a rados connection handle, given a conf object
func GetRadosHandle(conf string) (conn *rados.Conn, err error) {
	conn, err = rados.NewConn()
	if err != nil {
		return nil, err
	}

	conn.ReadConfigFile(conf)

	err = conn.Connect()
	return conn, err
}

func GetClusterStatus(conf string) {
	conn, err := GetRadosHandle(conf)

	command, err := json.Marshal(map[string]string{"prefix": "report"})
	buf, _, err := conn.MonCommand(command)
	if err == nil {
		var message map[string]interface{}

		err = json.Unmarshal(buf, &message)
		//log.Debug(message)
	}
}

func GetClusterIOPs(conf string) {
	conn, err := GetRadosHandle(conf)
	if err == nil {
		cmd, err := json.Marshal(map[string]string{"prefix": "osd pool stats", "format": "json"})
		buf, _, err := conn.MonCommand(cmd)
		if err == nil {
			var message []map[string]interface{}
			json.Unmarshal(buf, &message)
			log.Debug(message)
		}
	}
}
