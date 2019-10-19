package google

import "fmt"

// Given an array A that is a permutation of n numbers [1-n]. 
// Find the number of subarrarys S that meets the following condition: 
// max(S) - min(S) = length(S) - 1.

func _min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func _max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func SubarraysSatisfyCondition(arr []int) int {	
	return subarraysSatisfyCondition(arr, 0, len(arr) - 1)
}

func subarraysSatisfyCondition(arr []int, l, r int) int {
	fmt.Printf("l=%d, r=%d\n", l, r)
	if l == r {
		return 1
	}

	// building min and max arr prefixes: O(r-l)
	lminpref := make([]int, r-l+1)
	lmaxpref := make([]int, r-l+1)
	lmin := arr[l]
	lmax := arr[l]
	for i := l; i <= r; i += 1 {
		lmin = _min(lmin, arr[i])
		lminpref[i-l] = lmin

		lmax = _max(lmax, arr[i])
		lmaxpref[i-l] = lmax
	}

	rminpref := make([]int, r-l+1)
	rmaxpref := make([]int, r-l+1)
	rmin := arr[r]
	rmax := arr[r]
	for i := r; i >= l; i -= 1 {
		rmin = _min(rmin, arr[i])
		rminpref[i-l] = rmin

		rmax = _max(rmax, arr[i])
		rmaxpref[i-l] = rmax
	}

	split := l
	splity := arr[l]
	leftlower := false
	for k := l; k < r; k += 1 {
		if lmaxpref[k-l] < rminpref[k+1-l] {
			split = k
			splity = lmaxpref[k-l]
			leftlower = true
			break
		}
		if lminpref[k-l] > rmaxpref[k+1-l] {
			split = k
			splity = lminpref[k-l]
			break
		}
	}
	fmt.Printf("split=%d, splity_y=%d\n", split, splity)

	lcount := 0
	// go from split to the left and count all squares relatively to splitting point
	smin :=  arr[split]
	smax := arr[split]
	for i := split; i >= l; i -= 1 {
		smin = _min(smin, arr[i])
		smax = _max(smax, arr[i])
		if (leftlower && (splity - smin) == (split - i)) ||
		   (!leftlower && (smax - splity) == (split - i)) {
			lcount++
		}
	}

	rcount := 0
	// go from split to the right and count all squares relatively to splitting point
	smin = arr[split]
	smax = arr[split]
	for i := split; i <= r; i += 1 {
		smin = _min(smin, arr[i])
		smax = _max(smax, arr[i])
		if (leftlower && (smax - splity) == (i - split)) ||
		   (!leftlower && (splity - smin) == (i - split)) {
			if i == split && arr[split] == splity {
				// skipping split point as it was used on the left side to avoid counting it twice
				continue
			}
			rcount++
		}
	}
	fmt.Printf("lc=%d, rc=%d\n", lcount, rcount)
	// multiply squares from both sides as they are asymmetric to the splitting point,
	// meaning we count all combinations from both quadrants
	count := lcount*rcount

	// divide and conquer: O(N*logN) in average, O(N^2) worst case
	return count + subarraysSatisfyCondition(arr, l, split) + subarraysSatisfyCondition(arr, split + 1, r)
}
