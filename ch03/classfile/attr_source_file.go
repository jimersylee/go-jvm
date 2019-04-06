package classfile

type SourceFileAttrbute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (self *SourceFileAttrbute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}

func (self *SourceFileAttrbute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}
