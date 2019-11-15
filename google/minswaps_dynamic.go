const maxInt int = 1 << 31 - 1

func minSwap(A []int, B []int) int {
    n1 := 0
    s1 := 1
    
    for i := 1; i < len(A); i += 1 {
        n2 := maxInt
        s2 := maxInt
        
        if A[i-1] < A[i] && B[i-1] < B[i] {
            // 0 or 2 swaps
            n2 = min(n2, n1)
            s2 = min(s2, s1+1)
        }
        if A[i-1] < B[i] && B[i-1] < A[i] {
            // 1 swap of i-1 or i column
            n2 = min(n2, s1) // i-1 swapped
            s2 = min(s2, n1+1) // i swapped
        }
        n1 = n2
        s1 = s2
    }
    
    return min(n1, s1)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    
    return b
}
