package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func SetSignature(key []byte, s []byte) string {
	m := hmac.New(sha256.New, key)
	m.Write(s)
	return hex.EncodeToString(m.Sum(nil))
}