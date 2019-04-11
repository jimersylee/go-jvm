package classfile

//异常表入口
type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

//代码属性
type CodeAttribute struct {
	cp                  ConstantPool
	maxStacks           uint16
	maxLocals           uint16
	code                []byte
	exceptionTableEntry []*ExceptionTableEntry
	attributes          []AttributeInfo
}

//读取内容
func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStacks = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTableEntry = readeExceptionTable(reader)
	self.readAttributes(reader, self.cp)
}

//读取属性
func (self *CodeAttribute) readAttributes(reader *ClassReader, pools ConstantPool) {

}

//读取异常表
func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {

}
