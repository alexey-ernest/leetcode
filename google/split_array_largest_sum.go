
func splitArray(nums []int, m int) int {
    if len(nums) == 0 {
        return 0
    }
    
    // scan to identify total sum and max element: O(N)
    sum := 0
    largest := nums[0]
    for i := range nums {
        sum += nums[i]
        if nums[i] > largest {
            largest = nums[i]
        }
    }
    
    // binary search to find exact value x: largest <= x0 <= sum
    // satisfying condition, that tryDivideNotLargerThan(x0) == true
    // and tryDivideNotLargerThan(x) == false for each x < x0
    l, r := largest, sum
    x0 := sum
    for l <= r {
        x := (l+r)/2
        if tryDivideNotLargerThan(nums, m, x) == false {
            // number of parts >= m, so we can't divide into m parts
            // with largest sum <= x
            l = x+1
        } else {
            // trying to find solution with less x
            r = x-1
            if x < x0 {
                x0 = x
            }
        }
    }
    
    return x0
}

func tryDivideNotLargerThan(nums []int, m int, max int) bool {
    parts := 0
    currsum := 0
    for i := range nums {
        currsum += nums[i]
        if i < len(nums)-1 && currsum + nums[i+1] > max {
            // start new part
            parts++
            currsum = 0
        }
    }
    if currsum > 0 {
        parts++
    }
    return parts <= m
}
