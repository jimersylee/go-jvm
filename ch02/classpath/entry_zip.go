package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

//压缩文件入口实现,实现Entry接口
type ZipEntry struct {
	absDir string
}

func (self *ZipEntry) string() string {
	return self.absDir
}

//构造函数
func newZipEntry(path string) *ZipEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{absDir}
}

//读取类的字节集
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absDir)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil

		}
	}

	return nil, nil, errors.New("class not found:" + className)

}
