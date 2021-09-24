package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const max = math.MaxInt32

func nextInt(s string) int {
	var number int
	number, err := strconv.Atoi(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if number > max || number < 2 {
		fmt.Printf("Retry with : 1 < n < %d\n", max)
		os.Exit(1)
	}
	return number
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "No arguments...")
		os.Exit(1)
	}

	number1 := nextInt(os.Args[1])
	number2 := nextInt(os.Args[2])

	gcd := func(n1, n2 int) int {
		dividend, divisor := n1, n2
		if n2 > n1 {
			dividend = n2
			divisor = n1
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
	}(number1, number2)

	if gcd == 1 {
		fmt.Printf("%d and %d are co-prime to each other.\n", number1, number2)
		os.Exit(1)
	}

	fmt.Printf("gcd : %d\n", gcd)
}

// https://youtu.be/GRspofhTrfQ
