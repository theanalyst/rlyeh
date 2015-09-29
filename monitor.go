package main

import (
	"encoding/json"
	"github.com/ceph/go-ceph/rados"
)

func GetClusterStatus(conf string) {
	conn, _ := rados.NewConn()
	conn.ReadConfigFile(conf)
	conn.Connect()

	command, err := json.Marshal(map[string]string{"prefix": "report"})
	buf, _, err := conn.MonCommand(command)
	if err == nil {
		var message map[string]interface{}

		err = json.Unmarshal(buf, &message)
		Debug.Println(message)
	}
}
