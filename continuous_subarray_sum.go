import "fmt"

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func checkSubarraySum(nums []int, k int) bool {
    if k == 0 {
        for i := range nums {
            if i > 0 && nums[i] == 0 && nums[i-1] == 0 {
                return true
            }
        }
        return false
    }
    
    cache := make(map[int]int, len(nums)+1)
    cache[0] = -1
    sum := 0
    for i := range nums {
        sum += nums[i]
        rem := abs(sum)%abs(k)
        if c, ok := cache[rem]; ok {
            if i-c > 1 {
                return true    
            }
        }
        if _, ok := cache[rem]; !ok {
            cache[rem] = i    
        }
    }
    
    return false
}