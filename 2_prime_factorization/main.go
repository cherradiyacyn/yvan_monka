package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// the largest number in "prime_numbers.txt" is 989999; 989999² = 980098020001
// max <= 980098020001
const max = math.MaxInt32

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "No arguments...")
		os.Exit(1)
	}

	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if input > max || input < 2 {
		fmt.Printf("Retry with : 1 < n < %d\n", max)
		os.Exit(1)
	}

	primeDecomposition := func(number int) map[int]int {
		m := make(map[int]int)
		var isPrime bool
		for number > 1 {
			if isPrime {
				m[number]++
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
				if number%pn == 0 {
					m[pn]++
					number /= pn
					break
				}
				if pn > int(math.Sqrt(max)) {
					isPrime = true
				}
			}
			primesFile.Close()
		}
		return m
	}(input)

	for _, v := range primeDecomposition {
		if len(primeDecomposition) == 1 {
			if v == 1 {
				fmt.Printf("%d is prime.\n", input)
				os.Exit(1)
			}
		}
	}

	primeFactorization := func(m map[int]int) string {
		s := make([]string, 0)
		for k, v := range m {
			for i := v; i > 0; i-- {
				s = append(s, strconv.Itoa(k))
			}
		}
		str := strings.Join(s, "*")
		return str
	}(primeDecomposition)

	fmt.Printf("%d = %s\n", input, primeFactorization)
}

// https://youtu.be/k0rhj8fwdjs
