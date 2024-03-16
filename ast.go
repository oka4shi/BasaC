package main

type ConstantVal struct {
	val int
}

func newConstantVal(val int) *ConstantVal{
	cv := new(ConstantVal)
	cv.val = val
	return cv
}

func (cv *ConstantVal)get() int {
	return cv.val
}


