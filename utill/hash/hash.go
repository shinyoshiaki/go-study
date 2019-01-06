package hash

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1(str string) string {
	encoder := sha1.New()
	encoder.Write([]byte(str))
	hash := encoder.Sum(nil)
	hexHash := hex.EncodeToString(hash)
	return hexHash
}
