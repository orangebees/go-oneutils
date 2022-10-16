package Random

const hextable = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandBytes 快速生成n位随机字节
func RandBytes(n int) []byte {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = hextable[RandIntn(62)]
	}
	return b
}

// RandBytes32 快速生成32位随机字节
func RandBytes32() []byte {
	b := make([]byte, 32)
	for i := 0; i < 32; i++ {
		b[i] = hextable[RandIntn(62)]
	}
	return b
}

// RandIntn  快速生成num范围内的随机数
func RandIntn(num uint32) int {
	return int(FastRand() % num)
}
