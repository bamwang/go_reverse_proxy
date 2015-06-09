package main

func main() {
	power := genPower(2)
	for i := 0; i < 8; i++ {
		println(i, power())
	}
}

func genPower(n int) func() int {
	x := 1
	return func() int {
		x *= n
		return x
	}
}
