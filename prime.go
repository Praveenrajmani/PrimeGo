package main

import "fmt"
import "time"

func main() {
	t0 := time.Now()
	lprime := 0
	resulting := 0

	for i := 2; resulting < 60; i++ {
		if validateprime(i) == 1 {
			fmt.Println("\n", i)
			lprime = i
			t1 := time.Now()
			resulting = int(t1.Sub(t0) / (time.Millisecond * 1000))
		} else {
			t1 := time.Now()
			resulting = int(t1.Sub(t0) / (time.Millisecond * 1000))
		}
	}
	fmt.Println("\n\n The largest prime is ", lprime)
}

func validateprime(number int) int {
	var count int = 0
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			count++
		}
	}
	if count == 2 {
		return 1
	} else {
		return 0
	}
}
