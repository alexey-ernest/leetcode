import "fmt"

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func new21Game(N int, K int, W int) float64 {
    dp := make(map[int]float64)
    for i := K; i <= N; i += 1 {
        dp[i] = 1.0
    }

    s := float64(min(N-K+1, W))
    for i := K-1; i >= 0; i -= 1 {
        dp[i] = float64(s)/float64(W)
        s = s + dp[i] - dp[i+W]
    }
    
    return dp[0]
}