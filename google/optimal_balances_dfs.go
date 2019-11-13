import "fmt"

const maxInt = 1 << 31

func minTransfers(transactions [][]int) int {
    balancesmap := make(map[int]int, len(transactions))
    for _, t := range transactions {
        balancesmap[t[0]] -= t[2]
        balancesmap[t[1]] += t[2]
    }
    balances := make([]int, len(balancesmap))
    i := 0
    for _, b := range balancesmap {
        balances[i] = b
        i++
    }
    
    return dfs(0, balances)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func dfs(i int, balances []int) int {
    // ignore 0 balances
    for i < len(balances) && balances[i] == 0 {
        i++
    }
    
    if i == len(balances) {
        return 0
    }
    
    mintrans := maxInt
    for j := i + 1; j < len(balances); j += 1 {
        if balances[j] * balances[i] > 0 {
            // same sign
            continue
        }
        
        // making a transaction from current balance to some other balance
        balances[j] += balances[i]
        
        // exploring other balances with current transaction
        // keeping minimum number of transactions
        mintrans = min(mintrans, 1 + dfs(i+1, balances))
        
        // restoring balances
        balances[j] -= balances[i]
    }

    return mintrans
}

