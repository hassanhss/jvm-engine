package rtda

import "jvm-engine/engine/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
