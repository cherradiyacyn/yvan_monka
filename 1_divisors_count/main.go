package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// the largest number in "prime_numbers.txt" is 989999; 989999² = 980098020001
// max <= 980098020001
const max = math.MaxInt32

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "No arguments...")
		os.Exit(1)
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if number > max || number < 2 {
		fmt.Printf("Retry with : 1 < n < %d\n", max)
		os.Exit(1)
	}

	primeDecomposition := func(n int) map[int]int {
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
					n /= pn
					break
				}
				if pn > int(math.Sqrt(max)) {
					isPrime = true
				}
			}
			primesFile.Close()
		}
		return m
	}(number)

	divisorsCount := func(m map[int]int) int {
		n := 1
		for _, v := range m {
			n *= (v + 1)
		}
		return n
	}(primeDecomposition)

	if divisorsCount == 2 {
		fmt.Printf("%d is prime.\n", divisorsCount)
		os.Exit(1)

	}

	fmt.Printf("Divisors' count : %d\n", divisorsCount)
}

// https://youtu.be/k0rhj8fwdjs