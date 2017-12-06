package main

import "fmt"

func fibonacci() func() int {
	v1, v2 := 0, 1

	return func() int {
		v := v1
		v1 = v2
		v2 = v + v1

		return v
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
