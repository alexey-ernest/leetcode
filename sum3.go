package leetcode

import "sort"

type IntSort []int
func (a IntSort) Len() int { return len(a) }
func (a IntSort) Less(i, j int) bool { return a[i] < a[j] }
func (a IntSort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func threeSum(nums []int) [][]int {
    // O(N*lgN)
    sort.Sort(IntSort(nums))
    
    res := make([][]int, 0)
    // O(N^2)
    for i := 0; i < len(nums)-2; i += 1 {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        
        j := i + 1
        k := len(nums)-1
        
        for j < k {
            for j > i + 1 && j < k && nums[j] == nums[j-1] {
                j++
            }
            if j >= k {
                break
            }
            if nums[i] + nums[j] + nums[k] == 0 {
                res = append(res, []int{nums[i], nums[j], nums[k]})
                j++
            } else if nums[i] + nums[j] + nums[k] < 0 {
                j++
            } else {
                k--
            }
        }
    }
    
    return res
}