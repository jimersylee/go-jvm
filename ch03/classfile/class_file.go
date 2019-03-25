package classfile

import "fmt"

//类文件类
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

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)

}

func readConstantPool(reader *ClassReader) interface{} {

}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {

}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {

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
