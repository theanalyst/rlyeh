package main

import (
	"path"
	"strings"
)

// Returns the OSD id given the admin socket path
// assumes osd socket paths are of the form
// /path/to/osd.asok format
func GetOSDId(sock_path string) string {
	s := strings.Split(path.Base(sock_path), ".")
	return s[1]
}
