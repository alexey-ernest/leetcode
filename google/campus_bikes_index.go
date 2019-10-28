import (
    "math"
)

func assignBikes(workers [][]int, bikes [][]int) []int {
    n := len(workers)
    
    // n and m <= 1000 so we simply index them by distance
    dist := make([][][2]int, 2000)
    for i := range workers {
        for j := range bikes {
            d := manhdist(workers[i], bikes[j])
            dist[d] = append(dist[d], [2]int{i, j})
        }
    }
    
    res := make([]int, n)
    workersassigned := make(map[int]bool)
    bikestaken := make(map[int]bool)
    for _, pairs := range dist {
        if pairs == nil {
            continue
        }
        for j := range pairs {
            w, b := pairs[j][0], pairs[j][1]
            if workersassigned[w] || bikestaken[b] {
                continue
            }
            res[w] = b
            workersassigned[w] = true
            bikestaken[b] = true
        }
    }
    
    return res
}

func manhdist(a, b []int) int {
    return int(math.Abs(float64(a[0]-b[0])) + math.Abs(float64(a[1]-b[1])))
}