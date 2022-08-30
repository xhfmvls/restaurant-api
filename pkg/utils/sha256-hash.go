package utils

import(
	"crypto/sha256"
	"fmt"
)

func Sha256(val string) string {
	h := sha256.New()
	h.Write([]byte(val))
	hashedString := fmt.Sprintf("%x", h.Sum(nil))
	return hashedString
}