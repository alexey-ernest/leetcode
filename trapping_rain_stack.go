func trap(height []int) int {
    stack := make([]int,0,len(height))
    left := 0
    water := 0
    for i := 0; i < len(height); i+=1 {
        if height[i] < left {
            stack = append(stack, height[i])
            continue
        }
        // new border, calculating water
        waterlevel := left
        for j := len(stack)-1; j >= 0; j -= 1 {
            water += waterlevel-stack[j]
        }
        
        if len(stack) > 0 {
            stack[0] = height[i]    
            stack = stack[:1]
        } else {
            stack = append(stack, height[i])
        }
        
        left = height[i]
    }
    
    // handling tail
    if len(stack) > 1 {
        // reversing tail and calculating water for it
        tail := make([]int, len(stack))
        for i := range stack {
            tail[len(stack)-i-1] = stack[i]
        }
        water += trap(tail)
    }
    
    return water
}