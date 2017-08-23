package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	list := Eratosthenes(num)

	for i, v := range list {
		if v {
			fmt.Println(i)
		}
	}
}

func Eratosthenes(n int) []bool {
	numberes_bool := make([]bool, n)
	numberes_bool[0] = false
	numberes_bool[1] = false

	for i := 2; i < n; i++ {
		numberes_bool[i] = true
	}

	for i := 2; i < n; i++ {
		if numberes_bool[i] {
			for j := i * 2; j < n; j += i {
				numberes_bool[j] = false
			}
		}
	}

	return numberes_bool
}
