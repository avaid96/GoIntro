package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	nowfib := 0
	prevfib := 0
	return func() int {
		if nowfib == 0 {
			nowfib = 1
			return nowfib
		} else {
			tnowfib := nowfib + prevfib
			prevfib = nowfib
			nowfib = tnowfib
			return nowfib
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 15; i++ {
		fmt.Println(f())
	}
}
