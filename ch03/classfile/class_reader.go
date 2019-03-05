package classfile

import (
	"encoding/binary"
)

//类读取器
type ClassReader struct {
	data []byte
}

//读取u1
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

//读取u2
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

//读取u3
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

//读取u8
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

//读取u16们
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s

}

//读取字节集
func (self *ClassReader) readBytes(length uint32) []byte {
	bytes := self.data[:length]
	self.data = self.data[:length]
	return bytes
}
