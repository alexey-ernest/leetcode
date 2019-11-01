import "math"

var ops = map[byte]func(i,j float64) float64 {
    '+': func(i,j float64) float64 { return i+j },
    '-': func(i,j float64) float64 { return i-j },
    '*': func(i,j float64) float64 { return i*j },
    '/': func(i,j float64) float64 { return i/j },
}

func judgePoint24(nums []int) bool {
    fnums := make([]float64, len(nums))
    for i := range nums {
        fnums[i] = float64(nums[i])
    }
    return solve(fnums)
}

func solve(nums []float64) bool {
    if len(nums) == 1 {
        return math.Abs(nums[0]-24.0) < 0.000001
    }
    
    // picking two numbers and applying different operations to them
    for i := range nums {
        for j := range nums {
            if i == j {
                continue
            }
            rem := make([]float64, 0)
            for k := range nums {
                if k != i && k != j {
                    rem = append(rem, nums[k])
                }
            }
            
            // applying one of four operations
            for op, f := range ops {
                // filter out symmetric + and * pairs
                if (op == '*' || op == '+') && j < i {
                    continue
                }
                
                opval := f(nums[i], nums[j])
                rem = append(rem, opval)
                if solve(rem) == true {
                    return true
                }
                rem = rem[:len(rem)-1]
            }
        }
    }
    
    return false
}