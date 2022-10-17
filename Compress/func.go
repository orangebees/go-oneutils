package Compress

import (
	"archive/tar"
	"bytes"
	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasthttp"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const Separator = string(filepath.Separator)

// TarDirToByteBuffer tar归档目录，并写入到ByteBuffer中
func TarDirToByteBuffer(rpath string, compress string) (*bytebufferpool.ByteBuffer, error) {
	b := bytebufferpool.Get()
	defer bytebufferpool.Put(b)
	tw := tar.NewWriter(b)
	defer tw.Close()
	//tar压缩
	err := filepath.Walk(rpath,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				//跳过文件夹
				return nil
			}
			header, err := tar.FileInfoHeader(info, "")
			if err != nil {
				return err
			}
			header.Name = strings.TrimPrefix(path, rpath+Separator)
			err = tw.WriteHeader(header)
			if err != nil {
				return err
			}
			filebytes, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			_, err = io.Copy(tw, bytes.NewReader(filebytes))
			if err != nil {
				return err
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	b2 := bytebufferpool.Get()
	switch compress {
	case "br":
		_, err = fasthttp.WriteBrotliLevel(b2, b.B, fasthttp.CompressBrotliBestCompression)
		if err != nil {
			return nil, err
		}
		return b2, nil
	case "gz":
		_, err = fasthttp.WriteGzipLevel(b2, b.B, fasthttp.CompressBestCompression)
		if err != nil {
			return nil, err
		}
		return b2, nil
	default:
		b2.Write(b.B)
	}
	return b2, nil
}
