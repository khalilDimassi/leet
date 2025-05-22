package problems

import (
	"fmt"
	"math"
	"slices"
	"strconv"
)

type casualProblems struct {
}

func NewcasualProblems() *casualProblems {
	return &casualProblems{}
}

// ---------------------------------  easy  ---------------------------------- //

func (c *casualProblems) FindEvenNumbers(digits []int) []int {
	isValidNumber := func(n int) bool {
		return (n%2 == 0 && n > 99 && n < 999)
	}

	results := make([]int, 0)
	var item int
	for ii, i := range digits {
		for ij, j := range digits {
			for ik, k := range digits {
				if ii == ij || ii == ik || ij == ik {
					continue
				}

				item = (i * 100) + (j * 10) + k
				if isValidNumber(item) {
					results = append(results, item)
				}
			}
		}
	}

	slices.Sort(results)
	results = slices.Compact(results)

	return results
}

func (c *casualProblems) CountLargestGroup(n int) int {
	result := 0
	strN := fmt.Sprint(n)
	for _, i := range strN {
		ii, _ := strconv.Atoi(string(i))
		result += ii
	}
	return result
}

func (c *casualProblems) CountPairs(nums []int, k int) int {
	result := 0
	for ii, i := range nums {
		for ij, j := range nums {
			if !(ii > ij) {
				continue
			}
			if i == j && (ii*ij)%k == 0 {
				fmt.Printf("nums[%d] == nums[%d], and %d * %d == 0, which is divisible by %d\n", ii, ij, ii, ij, k)
				result++
			}
		}
	}

	return result
}

// --------------------------------- medium ---------------------------------- //

// ---------------------------------  hard  ---------------------------------- //

func (c *casualProblems) ColorTheGrid(m int, n int) int {
	const mod = 1e9 + 7
	// Step 1: Generate all valid column patterns
	var patterns []int
	for mask := range int(math.Pow(3.0, float64(m))) {
		valid := true
		prevColor := -1
		temp := mask
		for range m {
			color := temp % 3
			if color == prevColor {
				valid = false
				break
			}
			prevColor = color
			temp /= 3
		}
		if valid {
			patterns = append(patterns, mask)
		}
	}

	// Step 2: Precompute compatibility between patterns
	compatible := make(map[int][]int)
	for _, p1 := range patterns {
		for _, p2 := range patterns {
			compatibleFlag := true
			temp1, temp2 := p1, p2
			for range m {
				c1 := temp1 % 3
				c2 := temp2 % 3
				if c1 == c2 {
					compatibleFlag = false
					break
				}
				temp1 /= 3
				temp2 /= 3
			}
			if compatibleFlag {
				compatible[p1] = append(compatible[p1], p2)
			}
		}
	}

	// Step 3: Initialize DP table for the first column
	prevDP := make(map[int]int)
	for _, p := range patterns {
		prevDP[p] = 1
	}

	// Step 4: DP transitions for subsequent columns
	for j := 1; j < n; j++ {
		// fmt.Printf(">> prevDP: %v\n", prevDP)
		currDP := make(map[int]int)
		for p1, count := range prevDP {
			for _, p2 := range compatible[p1] {
				currDP[p2] = (currDP[p2] + count) % mod
			}
		}
		prevDP = currDP
	}

	// Step 5: Calculate the total number of valid colorings
	total := 0
	for _, count := range prevDP {
		total = (total + count) % mod
	}

	return total
}
