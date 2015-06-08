package main

func main() {

}

func genPower() (a func() int) {
	b := func() int {
		return 1
	}
	return b
}
