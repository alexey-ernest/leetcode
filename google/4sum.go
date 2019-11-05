import "fmt"
import "sort"

func fourSum(nums []int, target int) [][]int {
    
    if len(nums) < 4 {
        return nil
    }
    
    sort.Ints(nums)
    
    res := make([][]int, 0)
    
    for i := 0; i < len(nums)-2; i++ {
        for ; i > 0 && i < len(nums)-2 && nums[i] == nums[i-1]; i++ {}
        for j := i + 1; j < len(nums)-1; j += 1 {
            for ; j > i+1 && j < len(nums)-1 && nums[j] == nums[j-1]; j++ {}
            left := j+1
            right := len(nums)-1
            for left < right {
                sum := nums[i] + nums[j] + nums[left] + nums[right]
                if sum < target {
                    left++
                    for ; left < right && nums[left] == nums[left-1]; left++ {}
                } else if sum > target {
                    right--
                    for ; right > left && nums[right] == nums[right+1]; right-- {}
                } else {
                    res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
                    left++
                    for ; left < right && nums[left] == nums[left-1]; left++ {}
                    right--
                    for ; right > left && nums[right] == nums[right+1]; right-- {}
                }
            }
        }
    }
    
    return res
}