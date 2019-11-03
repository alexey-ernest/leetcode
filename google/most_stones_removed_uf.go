
type uf struct {
    data []int
    size int
}

func NewUf(size int) uf {
    d := make([]int, size)
    for i := range d {
        d[i] = i
    }
    return uf{
        data: d,
        size: size,
    }
}

func (u *uf) find(x int) int {
    for u.data[x] != x {
        x = u.data[x]
    }
    return x
}

func (u *uf) union(x,y int) {
    xc := u.find(x)
    yc := u.find(y)
    if xc == yc {
        return
    }
    
    u.data[xc] = yc
    u.size--
}

func (u *uf) connected(x,y int) bool {
    xc := u.find(x)
    yc := u.find(y)
    return xc == yc
}

func removeStones(stones [][]int) int {
    if len(stones) == 0 {
        return 0
    }
    
    // x, y < 10000, so we can make x and y independent variables in 1x20000 space 
    uf := NewUf(20000)
    
    // O(N)
    for _, s := range stones {
        // O(lgN)
        uf.union(s[0], s[1] + 10000)
    }
    
    // in each connected component we can remove k-1 stones maximum
    comps := make(map[int]struct{})
    for _, s := range stones {
        c := uf.find(s[0])
        comps[c] = struct{}{}
    }
    
    // we can remove all stones except 1 in each component
    return len(stones)-len(comps)
}