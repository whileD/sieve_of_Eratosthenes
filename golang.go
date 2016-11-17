package main
import "fmt"

const START_NUM int = 2

func main(){
	var num int
	fmt.Scan(&num)
	Eratosthenes(num)
}

func Eratosthenes(n int) {
	var numberes_bool = make([]bool, n)
	numberes_bool[0] = false
	numberes_bool[1] = false
	for i := START_NUM; i < len(numberes_bool); i++{
		numberes_bool[i] = true;
	}
	for i := START_NUM; i < len(numberes_bool); i++{
		if numberes_bool[i] {
			for j := i*2; j < len(numberes_bool); j += i{
				numberes_bool[j] = false
			}
		}
	}

	for i := 0; i < len(numberes_bool); i++{
		if numberes_bool[i]{
			fmt.Println(i)
		}
	}
}
