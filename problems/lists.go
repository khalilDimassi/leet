package problems

import "sort"

// ---------------------------------  easy  ---------------------------------- //

// --------------------------------- medium ---------------------------------- //
type RangeFreqQuery struct {
	mapping map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	valueIdx := make(map[int][]int)
	for idx, value := range arr {
		valueIdx[value] = append(valueIdx[value], idx)
	}

	return RangeFreqQuery{mapping: valueIdx}
}

func (RFQ *RangeFreqQuery) Query(left int, right int, value int) int {
	idxs, exists := RFQ.mapping[value]
	if !exists {
		return 0
	}

	l := sort.Search(len(idxs), func(i int) bool {
		return idxs[i] >= left
	})

	r := sort.Search(len(idxs), func(i int) bool {
		return idxs[i] > right
	})

	return r - l
}

//                     ---------------    ----------------                     //

// ---------------------------------  hard  ---------------------------------- //
