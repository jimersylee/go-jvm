package classfile

import "fmt"

//类文件结构体
type ClassFile struct {
	//magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func (self ClassFile) read(reader *ClassReader) {

}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}

		}
	}()
	classReader := &ClassReader{classData}
	classFile := ClassFile{}
	classFile.read(classReader)
	return
}
