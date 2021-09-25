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

	divisors := func(m map[int]int) []int {
		ss := make([][]int, 0)
		for k, v := range m {
			temp := make([]int, 0)
			for i := v; i >= 0; i-- {
				temp = append(temp, int(math.Pow(float64(k), float64(i))))
			}
			ss = append(ss, temp)
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

	fmt.Printf("Divisors : %v\n", divisors)
}

// https://youtu.be/k0rhj8fwdjs
