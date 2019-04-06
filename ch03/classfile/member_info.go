package classfile


//表示字段和方法的结构体
type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	membersCount := reader.readUint16()
	members := make([]*MemberInfo, membersCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

//读取字段或方法数据
func readMember(reader *ClassReader, constantPool ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              constantPool,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, constantPool)
	}
}
