package constants

import "jvm-engine/engine/instructions/base"
import "jvm-engine/engine/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
