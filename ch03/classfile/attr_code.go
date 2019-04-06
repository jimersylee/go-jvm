package classfile

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

type CodeAttribute struct {
	cp                  ConstantPool
	maxStacks           uint16
	maxLocals           uint16
	code                []byte
	exceptionTableEntry []*ExceptionTableEntry
	attributes          []AttributeInfo
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStacks = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTableEntry = readeExceptionTable(reader)
	self.readAttributes(reader, self.cp)
}

func (self *CodeAttribute) readAttributes(reader *ClassReader, pools ConstantPool) {

}

func readeExceptionTable(reader *ClassReader) []*ExceptionTableEntry {

}
