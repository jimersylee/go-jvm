package classpath

import (
	"path/filepath"
)

//目录型入口实现,实现Entry接口
type WildcardEntry struct {
	absDir string
}

func (self *DirEntry) string() string {
	return self.absDir
}

//构造函数
func newWildcardEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &DirEntry{absDir}
}

func (self *WildcardEntry) readClass(className string) ([]byte, Entry, error) {

	return nil, nil, nil
}
