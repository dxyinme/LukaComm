package vector

type Vector interface {
	PushFront(item interface{})
	PopFront() interface{}
	Size() int

	PushBack(item interface{})
	PopBack() interface{}

	Get(id int) interface{}
}

func New() Vector {
	ret := &Impl{}
	ret.New(10)
	return ret
}
