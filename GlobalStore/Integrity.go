package GlobalStore

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
)

type Integrity string

func (it Integrity) GetType() string {
	for i := 0; i < len(it); i++ {
		if it[i] == '-' {
			return string(it[:i])
		}
	}
	return ""
}
func (it Integrity) GetValue() string {
	for i := 0; i < len(it); i++ {
		if it[i] == '-' {
			return string(it[i+1:])
		}
	}
	return ""

}
func NewIntegrity(integrityType string, data []byte) Integrity {
	var integrityBytes = make([]byte, 64)
	integrityBytes = integrityBytes[:0]
	switch integrityType {
	case "md5":
		t := md5.Sum(data)
		integrityBytes = append(integrityBytes, t[:]...)
	case "sha1":
		t := sha1.Sum(data)
		integrityBytes = append(integrityBytes, t[:]...)
	case "sha256":
		t := sha256.Sum256(data)
		integrityBytes = append(integrityBytes, t[:]...)
	case "sha512":
		t := sha512.Sum512(data)
		integrityBytes = append(integrityBytes, t[:]...)
	}

	return Integrity(integrityType + "-" + base64.StdEncoding.EncodeToString(integrityBytes))
}
