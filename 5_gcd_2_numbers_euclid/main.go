package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const max = math.MaxInt32

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "No arguments...")
		os.Exit(1)
	}

	numbers := make([]int, 0)
	for i := 1; i < 3; i++ {
		number, err := strconv.Atoi(os.Args[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		if number > max || number < 2 {
			fmt.Printf("Retry with : 1 < n < %d\n", max)
			os.Exit(1)
		}
		numbers = append(numbers, number)
	}

	gcd := func(s []int) int {
		dividend, divisor := s[0], s[1]
		if s[1] > s[0] {
			dividend = s[1]
			divisor = s[0]
		}
		remainder := 1
		for remainder != 0 {
			remainder = dividend % divisor
			if remainder != 0 {
				dividend = divisor
				divisor = remainder
			}
		}
		return divisor
	}(numbers)

	if gcd == 1 {
		fmt.Printf("%d and %d are co-prime to each other.\n", numbers[0], numbers[1])
		os.Exit(1)
	}

	fmt.Printf("gcd : %d\n", gcd)
}

// https://youtu.be/GRspofhTrfQ
