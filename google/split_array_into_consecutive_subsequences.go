import "fmt"

// [1,2,3,3,4,4,5,5]
func isPossible(nums []int) bool {
    counts := make(map[int]int, len(nums))
    for _, n := range nums {
        counts[n]++
    }
    tails := make(map[int]int, len(counts))
    
    for _, x := range nums {
        if counts[x] == 0 {
            continue
        }
        if tails[x] > 0 {
            // a chain just ended before x, adding current number to that chain by moving end of the chain to next distinct number
            tails[x] -= 1
            tails[x+1] += 1
        } else if counts[x+1] > 0 && counts[x+2] > 0 {
            // new sequence started and we have 3 consecutive numbers after that, unite them to the chain and put end of the chain to the next distinct number greater than the 3d integer
            counts[x+1] -= 1
            counts[x+2] -= 1
            tails[x+3] += 1
        } else {
            // we had previous sequence with length < 3
            return false
        }
        counts[x] -= 1
    }
    
    return true
}
