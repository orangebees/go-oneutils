package GlobalStore

// FileInfoMap 文件相对路径 文件信息
type FileInfoMap map[string]FileInfo

// Equal 检测两个FileInfoMap是否相等
func (m FileInfoMap) Equal(fim FileInfoMap) bool {
	if len(m) != len(fim) {
		return false
	}
	for k, info := range m {
		fileInfo, ok := fim[k]
		if !ok {
			return false
		}
		if info.Integrity != fileInfo.Integrity {
			return false
		}
		if info.Size != fileInfo.Size {
			return false
		}
		if info.Mode != fileInfo.Mode {
			return false
		}
		if info.CheckedAt != fileInfo.CheckedAt {
			return false
		}
	}
	return true
}

// FileEqual 仅检测两个FileInfoMap的文件是否相等
func (m FileInfoMap) FileEqual(fim FileInfoMap) bool {
	if len(m) != len(fim) {
		return false
	}
	for k, info := range m {
		fileInfo, ok := fim[k]
		if !ok {
			return false
		}
		if info.Integrity != fileInfo.Integrity {
			return false
		}
		if info.Size != fileInfo.Size {
			return false
		}
		if info.Mode != fileInfo.Mode {
			return false
		}
		if info.CheckedAt != fileInfo.CheckedAt {
			return false
		}
	}
	return true
}
