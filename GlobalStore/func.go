package GlobalStore

import (
	"crypto/sha512"
	"github.com/orangebees/go-oneutils/Fetch"
	"github.com/orangebees/go-oneutils/PathHandle"
	"os"
	"path/filepath"
	"sort"
)

const Separator = string(filepath.Separator)

type FileStore struct {
	//工作根目录
	root string
	//元数据目录（在根目录下的相对目录）
	metadata string
	//构建目录
	build string
	//存储目录
	store string
}

type FileInfo struct {
	//日期
	CheckedAt int64 `json:"checked_at"`
	//校验和
	Integrity string `json:"integrity"`
	//权限
	Mode int `json:"mode"`
	//文件大小
	Size int64 `json:"size"`
}

type Metadata interface {
	GetFileInfoMap() FileInfoMap
	SetFileInfoMap(fim FileInfoMap)
	Fetch.EasyJsonSerialization
}

// NewFileStore 初始化文件存储
func NewFileStore(root string, metadata string, build string, store string) (*FileStore, error) {
	err := PathHandle.KeepDirsExist(root,
		root+Separator+metadata,
		root+Separator+build,
		root+Separator+store,
	)
	if err != nil {
		return nil, nil
	}
	return &FileStore{
		root:     root,
		metadata: metadata,
		build:    build,
		store:    store,
	}, nil
}

// AddDir 添加目录到全局文件存储
func (s FileStore) AddDir(absPath string) (FileInfoMap, error) {
	//路径不应该为root目录
	return nil, nil
}

// NewFileInfoMapFromDir 仅从文件夹生成FileInfoMap
func (s FileStore) NewFileInfoMapFromDir(absPath string) (FileInfoMap, error) {
	//路径不应该为root目录
	err := filepath.Walk(absPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//获取相对路径到结构体切片
			if info.IsDir() {
				//跳过文件夹
				return nil
			}
			file, err2 := os.ReadFile(path)
			if err2 != nil {
				return err2
			}
			println(file)
			rel, err := filepath.Rel(absPath, path)
			if err != nil {
				return err
			}
			rp := []byte(rel)
			//统一为Linux下的分隔符
			PathHandle.UnifyPathSlashSeparator(rp)
			//

			return nil
		})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// BuildDir 通过FileInfoMap构建文件目录
func (s FileStore) BuildDir(fim FileInfoMap) error {
	return nil
}

// VerifyDir 验证文件目录
func (s FileStore) VerifyDir(fim FileInfoMap, absPath string) error {

	return nil
}

// PathIntegrityVerificationUseSha512  验证计算目录的sha512
func PathIntegrityVerificationUseSha512(rpath string) (string, error) {
	var sumlists []string
	err := filepath.Walk(rpath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//fmt.Println(path)
			//获取相对路径到结构体切片
			if info.IsDir() {
				//跳过文件夹
				return nil
			}
			file, err2 := os.ReadFile(path)
			if err2 != nil {
				return err2
			}

			rel, err := filepath.Rel(rpath, path)
			if err != nil {
				return err
			}
			rp := []byte(rel)
			//统一为Linux下的分隔符
			PathHandle.UnifyPathSlashSeparator(rp)
			//
			rph := PathHandle.EncodeToString(sha512.Sum512(rp))
			fh := PathHandle.EncodeToString(sha512.Sum512(file))
			sum := PathHandle.EncodeToString(sha512.Sum512([]byte(rph + fh)))
			sumlists = append(sumlists, sum)
			return nil
		})
	if err != nil {
		return "", err
	}
	//排序
	sort.Strings(sumlists)
	var sumstr string
	for i := 0; i < len(sumlists); i++ {
		sumstr += sumlists[i]
	}
	s := PathHandle.EncodeToString(sha512.Sum512([]byte(sumstr)))
	return s, nil
}

// GetMetadataPath 获取metadata的理论路径
func (s FileStore) GetMetadataPath(path string) (string, error) {
	p := s.root + Separator + s.metadata + Separator + path + ".json"
	dir, _ := filepath.Split(p)
	err := PathHandle.KeepDirExist(dir)
	if err != nil {
		return "", err
	}
	return p, nil
}
