package encryption

import "crypto/sha256"

func CheckPassword(p string, e [32]byte) bool {
	s := sha256.Sum256([]byte(p))
	if s == e {
		return true
	}
	return false
}

func Encryption(p string) [32]byte {
	return sha256.Sum256([]byte(p))
}
