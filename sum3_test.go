package leetcode

import "testing"

func Test3Sum(t *testing.T) {
	res := threeSum([]int{-4,-2,1,-5,-4,-4,4,-2,0,4,0,-2,3,1,-5,0})
	exp := [][]int{
		[]int{-5,1,4},
		[]int{-4,0,4},
		[]int{-4,1,3},
		[]int{-2,-2,4},
		[]int{-2,1,1},
		[]int{0,0,0},
	}
	for i := range res {
		for j := range res[i] {
			if exp[i][j] != res[i][j] {
				t.Errorf("expected %+v, got %+v", exp, res)
				return
			}
		}
	}
}