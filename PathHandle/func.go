package PathHandle

import (
	"crypto/sha512"
	"errors"
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
//		mod取自hash转换为可见字符串之后的值
//	 已优化 可以被内联
func Bucket256AllocatedUseSha512(fileBytes []byte) (hash string, mod string) {
	sha512hash, dst, j := sha512.Sum512(fileBytes), make([]byte, 128), 0
	//转为可见字符
	for _, v := range sha512hash {
		dst[j], dst[j+1] = hextable[v>>4], hextable[v&0x0f]
		j += 2
	}
	//
	t, hash := xxhash.Sum64(dst)%256, string(dst)
	mod = string([]byte{hextable[t>>3], hextable[t%16]})
	return
}
func FilePathToDirPath(str string, s string) string {

	return str
}

// KeepDirsExist  确保某些目录一定存在
func KeepDirsExist(paths ...string) error {
	for i := 0; i < len(paths); i++ {
		err := KeepDirExist(paths[i])
		if err != nil {
			return err
		}
	}
	return nil
}

// KeepDirExist 确保某目录一定存在
func KeepDirExist(path string) error {
	f, err := os.Stat(path)
	if err == nil {
		//存在
		if f.IsDir() {
			//是已经存在的目录 不处理
			return nil
		}
		//是文件,返回错误
		return errors.New("file exists instead of directory")
	}
	if os.IsNotExist(err) {
		//不存在
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
		return nil
	}
	//其他错误
	return err
}
func EncodeToString(src [64]byte) string {
	dst := make([]byte, 128)
	j := 0
	for _, v := range src {
		dst[j] = hextable[v>>4]
		dst[j+1] = hextable[v&0x0f]
		j += 2
	}
	return string(dst)
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
