package control

import "jvm-engine/engine/instructions/base"
import "jvm-engine/engine/rtda"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
