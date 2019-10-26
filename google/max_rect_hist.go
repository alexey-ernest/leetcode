import "fmt"

func maximalRectangle(matrix [][]byte) int {
    n := len(matrix)
    if n == 0 {
        return 0
    }
    
    m := len(matrix[0])
    
    // O(n*m) building histograms for each column
    leftadj := make([][]int, n)
    for i := 0; i < n; i += 1 {
        left := make([]int, m)
        leftadj[i] = left
        for j := 0; j < m; j += 1 {
            val := 0
            if matrix[i][j] == '1' {
                val = 1
            }
            
            // counting left adjacent pixels
            if j == 0 || val == 0 {
                left[j] = val
            } else {
                left[j] = left[j-1]+1
            }
        }
    }
    
    maxsq := 0
    for i := 0; i < m; i += 1 {
        h := make([]int, n)
        for j := 0; j < n; j += 1 {
            h[j] = leftadj[j][i]
        }
        sq := largestRectangleArea(h)
        if sq > maxsq {
            maxsq = sq
        }
    }
    
    return maxsq
}

func largestRectangleArea(heights []int) int {
    if len(heights) == 0 {
        return 0
    }
    if len(heights) == 1 {
        return heights[0]
    }
    
    stack := make([][2]int, 0, len(heights))
    res := 0
    for _, h := range heights {
        if len(stack) == 0 || h > stack[len(stack)-1][0] {
            stack = append(stack, [2]int{h, 1})
            continue
        }
        
        sq := h
        w := 1
        j := len(stack)-1
        minh := stack[j][0]
        for ; j >= 0; j -= 1 {
            top := stack[j]
            if h > top[0] {
                break
            }
            
            if top[0] < minh {
                minh = top[0]
            }
            
            sq += top[1]*h
            w += top[1]
            
            if minh*(w-1) > res {
                res = minh*(w-1)
            }
        }
        
        if sq > res {
            res = sq
        }
        
        stack = stack[:j+1]
        stack = append(stack, [2]int{h, w})
    }
    
    if len(stack) > 1 {
        minh := stack[len(stack)-1][0]
        w := 0
        for j := len(stack)-1; j >= 0; j -= 1 {
            top := stack[j]
            if top[0] < minh {
                minh = top[0]
            }
            w += top[1]
            if minh*w > res {
                res = minh*w
            }
        }
    }
    
    return res
}