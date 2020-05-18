package reserved

import "jvm-engine/engine/instructions/base"
import "jvm-engine/engine/rtda"
import "jvm-engine/engine/native"
import _ "jvm-engine/engine/native/java/io"
import _ "jvm-engine/engine/native/java/lang"
import _ "jvm-engine/engine/native/java/security"
import _ "jvm-engine/engine/native/java/util/concurrent/atomic"
import _ "jvm-engine/engine/native/sun/io"
import _ "jvm-engine/engine/native/sun/misc"
import _ "jvm-engine/engine/native/sun/reflect"

// Invoke native method
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
