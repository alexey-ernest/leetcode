import "strings"
import "strconv"
import "fmt"
import "math"

var operators []string = []string{"*", "/", "+", "-"}
var patterns []string = []string{
    "$#$#$#$",
    "($#$)#$#$",
    "($#$#$)#$",
    "$#($#$)#$",
    "$#($#$#$)",
    "$#$#($#$)",
    "(($#$)#$)#$",
    "($#($#$))#$",
    "$#(($#$)#$)",
    "$#($#($#$))",
    "($#$)#($#$)",
}

func judgePoint24(nums []int) bool {
    if len(nums) != 4 {
        panic("expected 4 nums")
    }
    
    numstr := make([]string, 0, len(nums))
    for _, n := range nums {
        numstr = append(numstr, strconv.Itoa(n))
    }
    numsperm := permutations(numstr)
    //fmt.Printf("nums perm %d\n", len(numsperm))
    opcomb := combinations(3, operators)
    //fmt.Printf("op combinations %d\n", len(opcomb))
    
    for _, num := range numsperm {
        for _, op := range opcomb {
            for _, ptrn := range patterns {
                numstr := strings.Join(num, "")
                opstr := strings.Join(op, "")
                
                expr := buildexpr(numstr, opstr, ptrn)
                
                postfix := infixtopostfix(expr)
                res := solvepostfixexpr(postfix)
                //fmt.Printf("%s %s %.2f\n", expr, postfix, res)
                
                if math.Abs(res - 24.0) < 0.000001 {
                    return true
                }
            }
        }
    }
    return false
}

func infixtopostfix(expr string) string {
    res := ""
    stack := make([]byte, 0)
    precedence := make(map[byte]int)
    precedence['/'] = 3
    precedence['*'] = 3
    precedence['+'] = 2
    precedence['-'] = 2
    precedence['('] = 1
    for i := 0; i < len(expr); i += 1 {
        ch := expr[i]
        if ch == ')' || precedence[ch] > 0 {
            // op
            if ch == '(' {
                stack = append(stack, ch)
            } else if ch == ')' {
                for j := len(stack)-1; j >= 0; j -= 1 {
                    if stack[j] == '(' {
                        stack = stack[:j]
                        break
                    }
                    res += string(stack[j])
                }
            } else {
                // pop all operators with higher precedence to the output
                // or until we reach ( since it has the lowest precedence
                for j := len(stack)-1; j >= 0; j -= 1 {
                    if precedence[stack[j]] < precedence[ch] {
                        stack = stack[:j+1]
                        break
                    }
                    res += string(stack[j])
                    
                    if j == 0 {
                        stack = stack[:0]
                    }
                }
                stack = append(stack, ch)
            }
        } else {
            res += string(ch)
        }
    }
    for j := len(stack)-1; j >= 0; j -= 1 {
        res += string(stack[j])
    }
    
    return res
}

func solvepostfixexpr(expr string) float64 {
    vals := make([]float64, 0)
    for i := 0; i < len(expr); i += 1 {
        s := expr[i]
        if s >= '0' && s <= '9' {
            vals = append(vals, float64(s-'0'))
        } else {
            // ops
            val2 := vals[len(vals)-1]
            val1 := vals[len(vals)-2]
            
            
            res := 0.0
            switch s {
                case '+':
                    res = val1+val2
                case '-':
                    res = val1-val2
                case '*':
                    res = val1*val2
                case '/':
                    res = val1/val2
            default:
                panic(fmt.Sprintf("unknown operator in expr %s", expr))
            }
            vals[len(vals)-2] = res
            vals = vals[:len(vals)-1]
        }
    }
    
    return vals[0]
}

func buildexpr(nums string, ops string, pattern string) string {
    for i := 0; i < len(nums); i += 1 {
        pattern = strings.Replace(pattern, "$", string(nums[i]), 1)
    }
    for i := 0; i < len(ops); i += 1 {
        pattern = strings.Replace(pattern, "#", string(ops[i]), 1)
    }
    return pattern
}

func combinations(n int, list []string) [][]string {
    if n == 0 {
        return [][]string{}
    }
    
    if n == 1 {
        return [][]string{list[:]}
    }
    
    combs := combinations(n-1, list)
    res := make([][]string, 0)
    for _, l := range list {
        for _, comb := range combs {
            for i := 0; i < len(comb); i += 1 {
                res = append(res, []string{l+comb[i]})
            }
            
        }
    }
    
    return res
}

func permutations(list []string) [][]string {
    if len(list) == 0 {
        return [][]string{}
    }
    
    if len(list) == 1 {
        return [][]string{list}
    }
    
    res := make([][]string, 0)
    for i := 0; i < len(list); i += 1 {        
        rem := make([]string, len(list)-1)
        copy(rem[:i], list[:i])
        copy(rem[i:], list[i+1:])
        remp := permutations(rem)
        for _, p := range remp {
            perm := append([]string{list[i]}, p...)
            res = append(res, perm)
        }
    }
    
    return res
}