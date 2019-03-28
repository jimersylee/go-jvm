package classfile

import "fmt"

//常量池
type ConstantPool struct {
}

func (pool *ConstantPool) getClassName(u uint16) string {
	return ""
}

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

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

//获取类名
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

//获取超类名字
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)

}

func readAttributes(reader *ClassReader, constantPool interface{}) []interface{} {

}

func readMembers(reader *ClassReader, constantPool interface{}) []*interface{} {

}

func readConstantPool(reader *ClassReader) interface{} {
	return ConstantPool{}
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
