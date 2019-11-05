import "sort"
import "fmt"

func fourSumCount(A []int, B []int, C []int, D []int) int {
    
    sort.Ints(C)
    sort.Ints(D)
    
    // O(N^2)
    cdsums := make(map[int]int, len(C)*len(D))
    for i := 0; i < len(C); i += 1 {
        i1 := i+1
        for ; i1 < len(C) && C[i1] == C[i]; i1 += 1 {}
        for j := 0; j < len(D); j += 1 {
            j1 := j+1
            for ; j1 < len(D) && D[j1] == D[j]; j1 += 1 {}
            
            sum := C[i]+D[j]
            cdsums[sum] += (i1-i)*(j1-j)
            j = j1-1
        }
        
        i = i1-1
    }
    
    // O(N^2)
    count := 0
    for i := range A {
        for j := range B {
            count += cdsums[0-A[i]-B[j]]
        }
    }
    
    return count
}