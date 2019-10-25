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