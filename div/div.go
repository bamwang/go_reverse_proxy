package main

func main() {
	q, r := divAndMod(20, 3)
	print(q, r)
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	print(sum(arr[0:len(arr)]...))
}

func divAndMod(a, b int) (c, d int) {
	c = a / b
	d = a % b
	return
}

func sum(arg ...int) (sum int) {
	for _, v := range arg {
		sum += v
	}
	return
}
