package utils

import (
	"crypto/md5"
	"fmt"
)

var (
	salt = "service-computing-agenda-cli-go"
)

// MD5 hashes salted password and returns result in hex
func MD5(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password+salt)))
}
