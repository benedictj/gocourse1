package basics

import "fmt"

func main() {
	sum := 0
	for {
		sum += 10
		fmt.Println("Sum:", sum)
		if sum >= 50 {
			break
		}
	}
	i := 1
	for i <= 5 {
		fmt.Println("Iteration:", i)
		i++
	}

	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("OddNumber:", num)
		num++
	}
}
