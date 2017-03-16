package utils

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"time"
)

func IsLetter(c uint8) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func CountOneBits(x int64) int {
	cnt := 0
	for ; x > 0; cnt++ {
		x = x & (x - 1)
	}
	return cnt
}

func EncodeString(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func GenerateUserSalt(name string) string {
	return GetSha1Hash(fmt.Sprintf("%s%d", name, time.Now().Unix()))
}

func GenerateUserPassword(salt string, password string) string {
	return GetSha1Hash(salt, password)
}

func GenerateLoginTokenKey(salt string, uid int64, dname string) string {
	return fmt.Sprintf("%s%d%s", salt, uid, dname)
}

func GenerateGameTokenKey(salt string, uid int64, gid int64, dname string) string {
	return fmt.Sprintf("%s%d%d%s", salt, uid, gid, dname)
}

func GenerateTokenValue(salt string) string {
	return fmt.Sprintf("%s%d", salt, time.Now().Unix())
}

func GetSha1Hash(strings ...string) string {
	h := sha1.New()
	for _, s := range strings {
		h.Write([]byte(s))
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
