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
	}(gcd)

	divisors := func(m map[int]int) []int {
		ss := make([][]int, 0)
		for k, v := range m {
			t := make([]int, 0)
			for i := v; i >= 0; i-- {
				t = append(t, int(math.Pow(float64(k), float64(i))))
			}
			ss = append(ss, t)
		}
		s := ss[0]
		for i := 1; i < len(ss); i++ {
			s = func(s1, s2 []int) []int {
				s := make([]int, 0)
				for i := 0; i < len(s1); i++ {
					for j := 0; j < len(s2); j++ {
						s = append(s, s1[i]*s2[j])
					}
				}
				return s
			}(s, ss[i])
		}
		return s
	}(primeDecomposition)

	fmt.Printf("Common divisors : %v\n", divisors)
}

// https://youtu.be/GRspofhTrfQ
// https://youtu.be/Rno1eANN7aY
// https://youtu.be/k0rhj8fwdjs
