package problems

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"
)

type arrayProblems struct {
}

func NewarrayProblems() *arrayProblems {
	return &arrayProblems{}
}

// ---------------------------------  easy  ---------------------------------- //

func (a *arrayProblems) CountSubarrays(nums []int) int {
	result := 0

	for i := 0; i < len(nums); i++ {
		if i+2 >= len(nums) {
			break
		}
		if nums[i+1]%2 != 0 {
			continue
		}

		if nums[i]+nums[i+2] == 0 {
			if nums[i+1] == 0 {
				result++
			}
		} else if nums[i]+nums[i+2] == nums[i+1]/2 {
			result++
		}

	}

	return result
}

func (a *arrayProblems) GetFinalState(nums []int, k int, multiplier int) []int {
	minimum := func(nums []int) (min int, ind int, ok bool) {
		if len(nums) == 0 {
			return 0, 0, false
		}
		m := nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] < m {
				m = nums[i]
				ind = i
			}
		}
		return m, ind, true
	}

	for i := 0; i < k; i++ {
		min, ind, ok := minimum(nums)
		if !ok {
			break
		}
		nums[ind] = min * multiplier
	}

	return nums
}

func (a *arrayProblems) FinalPrices(prices []int) []int {
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				prices[i] = prices[i] - prices[j]
				break
			}
		}
	}

	return prices
}

func (a *arrayProblems) StringMatching(words []string) []string {
	result := make([]string, 0)
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i != j && strings.Contains(words[j], words[i]) && !slices.Contains(result, words[i]) {
				result = append(result, words[i])
			}
		}
	}
	return result
}

func (a *arrayProblems) CountPrefixSuffixPairs(words []string) int {
	isPrefixAndSuffix := func(str1, str2 string) bool {
		return strings.HasPrefix(str2, str1) && strings.HasSuffix(str2, str1) && len(str2) >= len(str1)
	}

	result := 0
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if i != j && isPrefixAndSuffix(words[i], words[j]) {
				result++
			}
		}
	}

	return result
}

func (a *arrayProblems) RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	result := 1
	curr := nums[0]
	for i := 1; i < len(nums); i++ {
		if curr != nums[i] {
			curr = nums[i]
			result++
		} else {
			nums[i] = 100
		}
	}

	slices.Sort(nums)

	return result
}

func (a *arrayProblems) RemoveElement(nums []int, val int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = 51
		} else {
			result++
		}
	}

	slices.Sort(nums)
	return result
}

func (a *arrayProblems) SearchInsert(nums []int, target int) int {
	for ii, i := range nums {
		if i == target {
			return ii
		} else if i > target {
			return ii
		}
	}
	return len(nums)
}

func (a *arrayProblems) PlusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func (a *arrayProblems) Merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	slices.Sort(nums1)
}

func (a *arrayProblems) SingleNumber(nums []int) int {
	result := nums[0]
	for _, num := range nums {
		result ^= num
	}
	return result
}

func (a *arrayProblems) SummaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	sumRanges := map[int]int{}
	start, curr := nums[0], nums[0]
	sumRanges[start] = curr

	for _, num := range nums[1:] {
		if curr+1 == num {
			sumRanges[start] = num
			curr = num
		} else {
			sumRanges[num] = num
			start, curr = num, num
		}
	}

	// Create a slice of the start points in their original order
	var starts []int
	seen := make(map[int]bool)
	for _, num := range nums {
		if _, exists := sumRanges[num]; exists && !seen[num] {
			starts = append(starts, num)
			seen[num] = true
		}
	}

	// Build the result in the correct order
	var result []string
	for _, start := range starts {
		end := sumRanges[start]
		if start == end {
			result = append(result, fmt.Sprintf("%d", start))
		} else {
			result = append(result, fmt.Sprintf("%d->%d", start, end))
		}
	}

	return result
}

func (a *arrayProblems) MaxRepeating(sequence string, word string) int {
	counter := 0
	sub := word
	for range len(sequence) / len(word) {
		if strings.Contains(sequence, sub) {
			counter++
		}
		sub = sub + word
	}
	return counter
}

func (a *arrayProblems) GetLongestSubsequence(words []string, groups []int) []string {
	result := make([]string, 1)
	result[0] = words[0]

	for i := 1; i < len(words); i++ {
		if groups[i] != groups[i-1] {
			result = append(result, words[i])
		}
	}

	return result
}

// --------------------------------- medium ---------------------------------- //

func (a *arrayProblems) EliminateMaximum(dist []int, speed []int) int {
	n := len(dist)
	time := make([]int, n)
	for i := range n {
		time[i] = (dist[i] + speed[i] - 1) / speed[i]
	}

	sort.Ints(time)
	fmt.Printf("time: %v\n", time)

	for i := range n {
		if time[i] <= i {
			return i
		}
	}

	return n
}

func (a *arrayProblems) PivotArray(nums []int, pivot int) []int {
	less := make([]int, 0)
	equal := make([]int, 0)
	greater := make([]int, 0)

	for _, num := range nums {
		switch {
		case num < pivot:
			less = append(less, num)
		case num == pivot:
			equal = append(equal, num)
		default:
			greater = append(greater, num)
		}
	}

	return append(append(less, equal...), greater...)
}

func (a *arrayProblems) MinOperations(boxes string) []int {
	input, result := strings.Split(boxes, ""), make([]int, len(boxes))
	fmt.Printf("input: %v\n", input)

	for i := range input {
		for j := range input {
			if input[j] == "1" {
				result[i] += int(math.Abs(float64(i - j)))
			}
		}
	}

	return result
}

func (a *arrayProblems) FindArray(pref []int) []int {
	for i := len(pref) - 1; i > 0; i-- {
		pref[i] ^= pref[i-1]
	}
	return pref
}

func (a *arrayProblems) SingleNumbers(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	bit := xor & -xor
	x, y := 0, 0

	for _, num := range nums {
		if num&bit != 0 {
			x ^= num
		} else {
			y ^= num
		}
	}

	return []int{x, y}
}

func (a *arrayProblems) CountMaxOrSubsetsBitmask(nums []int) int {
	or := 0
	for _, num := range nums {
		or |= num
	}

	result := 0
	for mask := range 1 << len(nums) {
		tempOr := 0
		for i := 0; i < len(nums) && tempOr != or; i++ {
			if mask&(1<<i) != 0 {
				tempOr |= nums[i]
			}
		}
		if tempOr == or {
			result++
		}
	}

	return result
}

func (a *arrayProblems) CountMaxOrSubsetsDFS(nums []int) int {
	maxOr := 0
	for _, num := range nums {
		maxOr |= num
	}

	count := 0
	stack := []struct{ idx, currentOr int }{{0, 0}}

	for len(stack) > 0 {
		fmt.Printf("> stack: %v, count: %d\n", stack, count)
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.currentOr == maxOr {
			// All remaining subsets are valid: iterate count by 2^(n - idx)
			count += 1 << (len(nums) - node.idx)
			continue
		}

		if node.idx == len(nums) {
			continue
		}

		// Option 1: Include nums[node.idx]
		newOr := node.currentOr | nums[node.idx]
		stack = append(stack, struct{ idx, currentOr int }{node.idx + 1, newOr})

		// Option 2: Exclude nums[node.idx]
		stack = append(stack, struct{ idx, currentOr int }{node.idx + 1, node.currentOr})
	}

	return count
}

func (a *arrayProblems) MaxIncreaseKeepingSkyline(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	count := 0
	n := len(grid)
	maxRow := make([]int, n)
	maxCol := make([]int, n)

	for i := range n {
		maxRow[i] = slices.Max(grid[i])
		for j := range n {
			maxCol[j] = max(maxCol[j], grid[i][j])
		}
	}

	for i := range n {
		for j := range n {
			count += min(maxRow[i], maxCol[j]) - grid[i][j]
		}
	}
	return count
}

func (a *arrayProblems) DeleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	freq := make(map[int]int)
	maxNum := 0
	for _, num := range nums {
		freq[num]++
		if num > maxNum {
			maxNum = num
		}
	}

	dp := make([]int, maxNum+1)
	dp[1] = freq[1] * 1

	for num := 2; num <= maxNum; num++ {
		points := freq[num] * num
		dp[num] = max(dp[num-1], dp[num-2]+points)
	}

	return dp[maxNum]
}

func (a *arrayProblems) IsZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)
	diff := make([]int, n+1)

	for _, q := range queries {
		l, r := q[0], q[1]
		for i := l; i <= r; i++ {
			diff[i]++
		}
	}

	for i := range n {
		nums[i] -= diff[i]
		if nums[i] < 0 {
			nums[i] = 0
		}
	}

	return slices.Max(nums) == 0
}

func (a *arrayProblems) LongestSubsequence(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	// dp is a slice of maps where dp[i] maps differences to maximum lengths
	dp := make([]map[int]int, n)
	for i := range dp {
		dp[i] = make(map[int]int)
	}

	maxLen := 2

	for i := 1; i < n; i++ {
		for j := range i {
			diff := int(math.Abs(float64(nums[i] - nums[j])))
			maxPrev := 0

			// Find the maximum length in dp[j] where prev_diff >= current diff
			for prevDiff, length := range dp[j] {
				if prevDiff >= diff && length > maxPrev {
					maxPrev = length
				}
			}

			// Update dp[i][diff] to be the maximum of:
			// - its current value
			// - maxPrev + 1 (extending a previous subsequence)
			// - 2 (starting a new subsequence with nums[j] and nums[i])
			dp[i][diff] = max(2, max(maxPrev+1, dp[i][diff]))
			maxLen = max(maxLen, dp[i][diff])
		}
	}

	return maxLen
}

// ---------------------------------  hard  ---------------------------------- //

// func (a *arrayProblems) MaximumScore(nums []int, multipliers []int) int {

// }
