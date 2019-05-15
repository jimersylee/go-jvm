package classfile

/**
异常表结构体
*/
type ExceptionAttribute struct {
	exceptionIndexTable []uint16
}

func (self *ExceptionAttribute) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}
func (self *ExceptionAttribute) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}
