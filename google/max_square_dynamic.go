func _min(a,b int) int {
    if a < b {
        return a
    }
    return b
}

func maximalSquare(matrix [][]byte) int {
    n := len(matrix)
    if n == 0 {
        return 0
    }
    
    m := len(matrix[0])
    maxsq := 0
    var prevrow []int
    for i := 0; i < n; i += 1 {
        row := make([]int, m)
        for j := 0; j < m; j += 1 {
            if matrix[i][j] == '0' {
                row[j] = 0
                continue
            }
            if i == 0 {
                row[j] = 1
                if maxsq == 0 {
                    maxsq = 1
                }
                continue
            }
            side := prevrow[j]
            if j > 0 {
                side = _min(side, row[j-1])
                side = _min(side, prevrow[j-1])
            } else {
                side = 0
            }
            side += 1
            row[j] = side
            
            if side*side > maxsq {
                maxsq = side*side
            }
        }
        prevrow = row
    }
    
    return maxsq
}