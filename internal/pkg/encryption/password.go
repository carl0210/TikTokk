package encryption

import (
	"crypto/sha256"
	"encoding/hex"
)

func CheckPassword(p, e string) bool {
	c := Encryption(p)
	if hex.EncodeToString(c[:]) == e {
		return true
	}
	return false
}

func Encryption(p string) [32]byte {
	return sha256.Sum256([]byte(p))
}
