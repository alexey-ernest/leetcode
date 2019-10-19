package google

import "testing"

func TestSubarraysSatisfyCondition(t *testing.T) {
	res := SubarraysSatisfyCondition([]int{4, 3, 1, 2, 5})
	if res != 10 {
		t.Errorf("%d != 10", res)
	}
}