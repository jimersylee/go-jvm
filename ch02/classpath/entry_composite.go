package classpath

import (
	"errors"
	"strings"
)

//目录型入口实现,实现Entry接口
type CompositeEntry []Entry

func (self *CompositeEntry) String() string {
	return self.String()
}

//构造函数
func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry

}

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		{
			if err == nil {
				return data, from, nil
			}
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}
