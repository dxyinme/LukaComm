package vector

type Impl struct {
	l int
	r int

	tmpSlice []interface{}
}

func (i *Impl) New(size int) {
	if size < 10 {
		size = 10
	}
	i.tmpSlice = make([]interface{}, size)
	i.l = size / 2 + 1
	i.r = size / 2
}

func (i *Impl) PushFront(item interface{}) {
	if i.l == 0 {
		Len := cap(i.tmpSlice)
		tmp := make([]interface{}, Len)
		i.l += Len
		i.r += Len
		tmp = append(tmp, i.tmpSlice...)
		i.tmpSlice = tmp
	}
	i.l--
	i.tmpSlice[i.l] = item
}

func (i *Impl) PopFront() interface{} {
	if i.r-i.l+1 <= 0 {
		panic("vector length is 0")
	}
	ret := i.tmpSlice[i.l]
	i.l++
	i.shrink()
	return ret
}

func (i *Impl) Size() int {
	return i.r - i.l + 1
}

func (i *Impl) PushBack(item interface{}) {
	i.r++
	if i.r >= len(i.tmpSlice) {
		i.tmpSlice = append(i.tmpSlice, item)
	} else {
		i.tmpSlice[i.r] = item
	}
}

func (i *Impl) PopBack() interface{} {
	if i.r-i.l+1 <= 0 {
		panic("vector length is 0")
	}
	ret := i.tmpSlice[i.r]
	i.r--
	i.shrink()
	return ret
}

func (i *Impl) Get(id int) interface{} {
	if id >= i.r-i.l+1 {
		return nil
	}
	return i.tmpSlice[i.l+id]
}

func (i *Impl) shrink() {
	Len := cap(i.tmpSlice)
	if Len < 5 {
		return
	}
	if i.l > 3*Len/4 {
		pre := Len / 2
		i.l -= pre
		i.r -= pre
		i.tmpSlice = i.tmpSlice[pre:]
		return
	}
	if i.r < Len/4 {
		back := Len / 2
		i.tmpSlice = i.tmpSlice[:back]
	}
}
