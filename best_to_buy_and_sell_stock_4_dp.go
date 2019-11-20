import "fmt"

func maxProfit(k int, prices []int) int {
    if len(prices) <= 1 {
        return 0
    }
    
    trans := 0
    if prices[0] < prices[1] {
        trans = 1
    }
    
    res := 0
    for i := range prices[:len(prices)-1] {
        if prices[i+1] > prices[i] {
            res += prices[i+1] - prices[i]
        } else {
            trans++
        }
    }
    
    if trans <= k {
        return res
    }
    
    // S O(n*k)
    dp := make([][]int, len(prices))
    for i := range dp {
        dp[i] = make([]int, k+1)
    }
    
    for j := 1; j <= k; j += 1 {
        state := -int(^uint(0) >> 1)-1
        for i := 1; i < len(prices); i += 1 {    
            state = max(state, dp[i-1][j-1]-prices[i-1])
            dp[i][j] = max(dp[i-1][j], state + prices[i])
        }
    }
    
    return dp[len(prices)-1][k]
}

func max (a, b int) int {
    if a > b {
        return a
    }
    return b
}