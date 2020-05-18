package io

import "jvm-engine/engine/native"
import "jvm-engine/engine/rtda"

func init() {
	native.Register("sun/io/Win32ErrorMode", "setErrorMode", "(J)J", setErrorMode)
}

func setErrorMode(frame *rtda.Frame) {
	// todo
	frame.OperandStack().PushLong(0)
}
