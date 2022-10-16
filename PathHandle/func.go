package PathHandle

import (
	"crypto/sha512"
	"github.com/cespare/xxhash/v2"
	"github.com/orangebees/go-oneutils/Convert"
	"github.com/orangebees/go-oneutils/random"
	"os"
	"path/filepath"
	"strings"
)

const Separator = string(filepath.Separator)
const hextable = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RunInTempDir 在缓存目录中工作
func RunInTempDir(f func(tmppath string) error) error {
	p := os.TempDir() + Separator + Convert.B2S(Random.RandBytes32())
	err := f(p)
	if err != nil {
		return err
	}
	err = os.RemoveAll(p)
	if err != nil {
		return err
	}
	return nil
}

// UnifyPathSlashSeparator 统一路径分隔符为斜杠
func UnifyPathSlashSeparator(b []byte) {
	for i := 0; i < len(b); i++ {
		if b[i] == '\\' {
			b[i] = '/'
		}
	}
}

// UnifyPathBackSlashSeparator 统一路径分隔符为反斜杠
func UnifyPathBackSlashSeparator(b []byte) {
	for i := 0; i < len(b); i++ {
		if b[i] == '/' {
			b[i] = '\\'
		}
	}
}

// Bucket256AllocatedUseSha512  使用sha512作为hash算法为文件分配桶（256个桶下）
//
//	mod取自hash转换为可见字符串之前的值
func Bucket256AllocatedUseSha512(fileBytes []byte) (hash string, mod string) {
	sha512hash, dst, c := sha512.Sum512(fileBytes), make([]byte, 128), make([]byte, 2)
	for i, j := 0, 0; i < 64; i++ {
		dst[j], dst[j+1] = hextable[sha512hash[i]>>4], hextable[sha512hash[i]&0x0f]
		j += 2
	}
	t := xxhash.Sum64(dst) % 256
	t1, t2 := t%16, t>>3
	c[0], c[1] = hextable[t2], hextable[t1]
	hash, mod = string(dst), string(c)
	return
}
func FilePathToDirPath(str string, s string) string {

	return str
}

// URLToLocalDirPath uri转本地相对路径
func URLToLocalDirPath(url string) string {
	tmp := strings.Split(url, "://")
	tmplen := len(tmp)
	var tmpbytes []byte
	if tmplen == 2 {
		tmpbytes = append(tmpbytes, tmp[1]...)

	} else {
		tmpbytes = append(tmpbytes, tmp[0]...)
	}
	for i := 0; i < len(tmpbytes); i++ {
		if tmpbytes[i] == '/' {
			tmpbytes[i] = filepath.Separator
		}
	}
	return string(tmpbytes)
}

// PathExists 路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
