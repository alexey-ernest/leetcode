package leetcode

import "testing"

func TestTwoSortedMedian(t *testing.T) {
	res := findMedianSortedArrays([]int{2,3,4,5,6,9,10}, []int{1,7,8})
	exp := 5.5
	if res != exp {
		t.Errorf("expected %0.2f != %0.2f", exp, res)
	}
}