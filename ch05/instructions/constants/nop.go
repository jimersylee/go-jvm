package constants

import "go-jvm/ch05/instructions/base"
import "go-jvm/ch05/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
