package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const max = 42

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

	primeDecompositions := func(s []int) []map[int]int {
		sm := make([]map[int]int, 0)
		for i := 0; i < len(s); i++ {
			onePrimeDecomposition := func(n int) map[int]int {
				m := make(map[int]int)
				var isPrime bool
				for n > 1 {
					if isPrime {
						m[n]++
						break
					}
					primesFile, err := os.Open("prime_numbers.txt")
					if err != nil {
						fmt.Fprintf(os.Stderr, "%v\n", err)
						os.Exit(1)
					}
					scanr := bufio.NewScanner(primesFile)
					for scanr.Scan() {
						pn, _ := strconv.Atoi(scanr.Text())
						if n%pn == 0 {
							m[pn]++
							n = n / pn
							break
						}
						if pn > int(math.Sqrt(max)) {
							isPrime = true
						}
					}
					primesFile.Close()
				}
				return m
			}(s[i])
			sm = append(sm, onePrimeDecomposition)
		}
		return sm
	}(numbers)

	limit := numbers[0]
	if numbers[1] > numbers[0] {
		limit = numbers[1]
	}

	lcm := func(l int, sm []map[int]int) int {
		m := make(map[int]int, 0)
		for i := 2; i <= l; i++ {
			c := 0
			for _, m := range sm {
				if m[i] > c {
					c = m[i]
				}
			}
			if c != 0 {
				m[i] = c
			}
		}
		lcm := 1
		for k, v := range m {
			lcm *= int(math.Pow(float64(k), float64(v)))
		}
		return lcm
	}(limit, primeDecompositions)

	gcd := func(l int, sm []map[int]int) int {
		m := make(map[int]int, 0)
		for i := 2; i <= l; i++ {
			c := l
			for _, m := range sm {
				if m[i] < c {
					c = m[i]
				}
			}
			if c != 0 {
				m[i] = c
			}
		}
		gcd := 1
		for k, v := range m {
			gcd *= int(math.Pow(float64(k), float64(v)))
		}
		return gcd
	}(limit, primeDecompositions)

	fmt.Printf("lcm : %d\n", lcm)
	fmt.Printf("gcd : %d\n", gcd)
}

// https://youtu.be/2bIK1KkQ1k0
