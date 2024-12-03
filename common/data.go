package common

import (
	"crypto/md5"
	"encoding/json"
	"strings"
)

// Checksum calculates a MD5 sum for given interface
func Checksum(data interface{}) [16]byte {
	jsonBytes, _ := json.Marshal(data)
	return md5.Sum(jsonBytes)
}

func GetData(data string) []string {
	return strings.Split(strings.TrimSpace(data), "\n")
}
