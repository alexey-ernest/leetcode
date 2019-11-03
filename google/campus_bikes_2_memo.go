type state struct {
    dist int
    cache map[int]int
    workers [][]int
    bikes [][]int
}

func assignBikes(workers [][]int, bikes [][]int) int {
    s := &state{
        cache: make(map[int]int),
        workers: workers,
        bikes: bikes,
    }
    res := mindist(s, 0, 0)
    return res
}

func mindist(s *state, i int, usedbikes int) int {
    if i == len(s.workers) {
        return 0
    }
    if s.cache[usedbikes] > 0 {
        return s.cache[usedbikes]
    }
    res := int((^uint(0)) >> 1)
    for j := range s.bikes {
        if (usedbikes & (1 << uint(j))) > 0 {
            continue
        }
        newusedbikes := usedbikes | (1 << uint(j))
        res = min(res, dist(s.workers[i], s.bikes[j]) + mindist(s, i+1, newusedbikes))
    }
    
    s.cache[usedbikes] = res
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}

func dist(p1, p2 []int) int {
    return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
}