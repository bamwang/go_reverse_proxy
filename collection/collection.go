package collection

type IntArray []int

func (this IntArray) Map(callback func(int) int) IntArray {
	var array = IntArray{}
	for i, n := range this {
		(array)[i] = callback(n)
	}
	return array
}
