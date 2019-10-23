package leetcode

func merge(a, b []int) []int {
    res := make([]int, len(a)+len(b))
    i, j, k := 0, 0, 0
    for i < len(a) && j < len(b) {
        if a[i] <= b[j] {
            res[k] = a[i]
            k+=1
            i+=1
        } else {
            res[k] = b[j]
            k+=1
            j+=1
        }
    }
    for ; i < len(a); i+=1 {
        res[k] = a[i]
        k+=1
    }
    for ; j < len(b); j+=1 {
        res[k] = b[j]
        k+=1
    }
    
    return res
}

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

func nextIndex(a []int, i int) int {
    if i < len(a)-1 {
        return i+1
    }
    return -1
}

func condition(nums1 []int, nums2 []int, i, j int) int {
    // if min(nums1[i+1], nums2[j+1]) >= max(nums1[i], nums2[j]) then cond == 0
    nexti := nextIndex(nums1, i)
    nextj := nextIndex(nums2, j)
    var minnext int 
    if nexti == -1 && nextj == -1 {
        // corner case when two arrays has len == 1
        return 0
    }
    
    if nexti == -1 {
        minnext = nums2[nextj]
    } else if nextj == -1 {
        minnext = nums1[nexti]
    } else {
        minnext = _min(nums1[nexti], nums2[nextj])
    }
    
    var maxcurrent int
    if i == -1 {
        maxcurrent = nums2[j]
    } else if j == -1 {
        maxcurrent = nums1[i]
    } else {
        maxcurrent = _max(nums1[i], nums2[j])
    }
    
    if minnext >= maxcurrent {
        return 0
    }
    
    // if nums1[i] > nums2[j+1] then cond > 0
    pj := j
    if pj < 0 {
        pj = nextj
    }
    if i >= 0 && nums1[i] > nums2[pj] {
        return 1
    }
    
    // if nums2[j] > nums1[i+1] then cond < 0
    pi := i
    if pi < 0 {
        pi = nexti
    }
    if j >= 0 && nums2[j] > nums1[pi] {
        return -1
    }
    
    return 0
}

func getSortedCandidates(nums1 []int, nums2 []int, i, j int) []int {
    cands1, cands2 := []int{}, []int{}
    if (len(nums1)+len(nums2))%2 == 0 {
        // examining last 2 elements for each array
        for k := i; k >= 0 && k >= i-1; k -= 1 {
            cands1 = append([]int{nums1[k]}, cands1...)
        }
        for k := j; k >= 0 && k >= j-1; k -= 1 {
            cands2 = append([]int{nums2[k]}, cands2...)
        }
    } else {
        // examining last 1 element for each array
        if i >= 0 {
            cands1 = []int{nums1[i]}    
        }
        if j >= 0 {
            cands2 = []int{nums2[j]}    
        }
    }
    cands := merge(cands1, cands2)
    return cands
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m := len(nums1)
    n := len(nums2)
    
    maxelements := (n+m)/2 + 1
    var cands []int
    if n == 0 {
        if maxelements >= 2 {
            cands = nums1[maxelements-2:maxelements]
        } else {
            cands = nums1[maxelements-1:maxelements]
        }
    } else if m == 0 {
        if maxelements >= 2 {
            cands = nums2[maxelements-2:maxelements]
        } else {
            cands = nums2[maxelements-1:maxelements]
        }
    } else {
        // both parts non-empty
        
        i := _min(maxelements, m)
        // i and j are numbers of elements to contribute from each array, not an indexes
        j := maxelements-i
        step := i/2
        for i >= 0 && j >= 0 {
            j = maxelements-i
            if j > len(nums2) {
                // check if we out of arr2
                i += j-len(nums2)
                j = len(nums2)    
            }

            cond := condition(nums1, nums2, i-1, j-1)
            if cond == 0 {
                break
            }

            if step == 0 {
                step = 1
            }
            if cond < 0 {
                i += step
            } else {
                i -= step
            }
            step = step/2
        }

        cands = getSortedCandidates(nums1, nums2, i-1, j-1)
    }
    
    var median float64
    if (n+m)%2 == 0 {
        median = (float64(cands[len(cands)-1]) + float64(cands[len(cands)-2]))/2
    } else {
        median = float64(cands[len(cands)-1])   
    }
    return median
}