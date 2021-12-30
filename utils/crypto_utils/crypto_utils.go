package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))                // has.write takes byte as the input value corresponding to string value // so we can write that byte in hash
	return hex.EncodeToString(hash.Sum(nil)) // here we not want to return the same string because that is password. so try another approach  EncodeToString ( it gives Encodestring value corresponding to byte)
}
