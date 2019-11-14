func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
func checkRecord(n int) int {
    mod := 1000000007
    
    dp := make([]int, max(n+1,4))
    dp[0] = 1
    dp[1] = 2
    dp[2] = 4
    dp[3] = 7
    for i := 4; i <= n; i += 1 {
        dp[i] = (2*dp[i-1] % mod + (mod - dp[i-4])) % mod
    }
    res := dp[n]
    
    // exluding variants with A in earch position (multiply variants to the left and to the right), as the set is symmetric we can calculate only for the half of the n
    for i := 0; i < n/2; i += 1 {
        res = (res + 2*dp[i]*dp[n-i-1]) % mod
    }
    // corner case for the middle point
    if n%2 != 0 {
        i := (n-1)/2
        res = (res + dp[i]*dp[n-i-1]) % mod
    }
    
    return res
}