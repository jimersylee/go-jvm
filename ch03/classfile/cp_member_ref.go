package classfile

//常量成员引用
type ConstantMemberRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

//常量字段引用
type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo
}

//常量方法引用
type ConstantMethodRefInfo struct {
	ConstantMemberRefInfo
}

//常量接口方法引用
type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberRefInfo
}

func (self *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}
func (self *ConstantMemberRefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}
