package base

import "go-jvm/ch05/rtda"

type Instruction interface {
	FetchOperands(reader *ByteCodeReader)
	Execute(frame *rtda.Frame)
}
