
func minWindow(S string, T string) string {
    if S == T {
        return T
    }
    
    dp := make([]int, len(T))
    idx := make(map[byte][]int)
    for i := 0; i < len(T); i += 1 {
        idx[T[i]] = append(idx[T[i]], i)
        dp[i] = -1
    }
    
    min := S+T
    // [1,1,1]
    for i := 0; i < len(S); i += 1 {
        idxs, ok := idx[S[i]]
        if !ok {
            continue
        }
        for j := len(idxs)-1; j >= 0; j -= 1 {
            id := idxs[j]
            if id == 0 {
                // starting new sub-sequence, trying to increase index to shorten the subsequence
                dp[id] = i
            } else {
                // using index of the previous char of T (max index in S which contains T[:j])
                dp[id] = dp[id-1]
            }
            if id == len(T)-1 && dp[id] >= 0 {
                // just found subsequence containing T
                if i-dp[id]+1 < len(min) {
                    min = S[dp[id]:i+1]
                }
            }
        }
    }
    
    if min == S+T {
        return ""
    }
    
    return min
}