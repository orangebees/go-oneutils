package Convert

import (
	"math"
	"strings"
	"unsafe"
)

const hextable = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// S2B string转[]byte
func S2B(s string) (b []byte) {
	*(*string)(unsafe.Pointer(&b)) = s
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + 2*unsafe.Sizeof(&b))) = len(s)
	return
}

// B2S []byte转string
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Uint64ToBase62Bytes uint转62进制字节数组
func Uint64ToBase62Bytes(num uint64) []byte {
	var bytes = make([]byte, 11)
	bytes = bytes[:0]
	for num > 0 {
		bytes = append(bytes, hextable[num%62])
		num = num / 62
	}
	for left, right := 0, len(bytes)-1; left < right; left, right = left+1, right-1 {
		bytes[left], bytes[right] = bytes[right], bytes[left]
	}
	return bytes
}

// Base62BytesToUint64 62进制字节数组转uint64
func Base62BytesToUint64(str []byte) uint64 {
	var num uint64
	n := len(str)
	for i := 0; i < n; i++ {
		pos := strings.IndexByte(hextable, str[i])
		num += uint64(math.Pow(62, float64(n-i-1)) * float64(pos))
	}
	return num
}
