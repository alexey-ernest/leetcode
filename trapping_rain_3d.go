import "fmt"

func trapRainWater(heightMap [][]int) int {
    if len(heightMap) == 0 {
        return 0
    }
    
    n := len(heightMap)
    m := len(heightMap[0])
    
    visits := make([][]bool, n)
    for i := range visits {
        visits[i] = make([]bool, m)
    }
    
    // building min pq for block heights because min height blocks define min level
    // minimum height border defines all adjacent heights, so add them first
    // pop min height and and adjust current level, explore adjacent blocks to
    // calculate water above them
    pq := NewIndexMinPQ()
    for i := range heightMap {
        visits[i][0] = true
        visits[i][m-1] = true
        pq.insert(i*m, heightMap[i][0])
        pq.insert(i*m+m-1, heightMap[i][m-1])
    }
    for j := 1; j < m-1; j += 1 {
        visits[0][j] = true
        visits[n-1][j] = true
        pq.insert(j, heightMap[0][j])
        pq.insert((n-1)*m+j, heightMap[n-1][j])
    }
        
    adj := [][]int{
        []int{0,1},
        []int{0,-1},
        []int{1,0},
        []int{-1,0},
    }
    current := 0
    water := 0
    for !pq.isEmpty() {
        height := pq.min()
        ind := pq.delMin()
        i := ind/m
        j := ind%m
        
        current = _max(current, height)
        
        // explore adjacent blocks
        for _, d := range adj {
            ia := i+d[0]
            ja := j+d[1]
            if ia == n || ia < 0 || ja == m || ja < 0 || visits[ia][ja] {
                continue
            }

            if current > heightMap[ia][ja] {
                water += current-heightMap[ia][ja]
            }
            pq.insert(ia*m+ja, heightMap[ia][ja])
            visits[ia][ja] = true
        }
    }
    
    return water
}

func _max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

type indexMinPQ struct {
    keys []int
    ind []int
    invind []int
    n int
}

func NewIndexMinPQ() indexMinPQ {
    return indexMinPQ{
        keys: make([]int, 1),
        ind: make([]int, 0),
        invind: make([]int, 1),
    }
}

func (pq *indexMinPQ) isEmpty() bool {
    return pq.n == 0
}

func (pq *indexMinPQ) less(i, j int) bool {
    k1 := pq.keys[i]
    k2 := pq.keys[j]
    if k1 < k2 {
        return true
    }
    if k1 > k2 {
        return false
    }
    
    // k1 == k2
    ind1 := pq.invind[i]
    ind2 := pq.invind[j]
    if ind1 < ind2 {
        return true
    }
    return false
}

func (pq *indexMinPQ) sink(i int) {
    k := pq.ind[i]
    for 2*k <= pq.n {
        ch := 2*k
        if ch < pq.n && pq.less(ch+1, ch) {
            ch += 1
        }
        
        if pq.less(ch, k) {
            // swap
            ki := pq.invind[k]
            chi := pq.invind[ch]
            pq.keys[k], pq.keys[ch] = pq.keys[ch], pq.keys[k]
            pq.invind[k], pq.invind[ch] = pq.invind[ch], pq.invind[k]
            pq.ind[ki], pq.ind[chi] = ch, k
        } else {
            break
        }
        k = ch
    }
}

func (pq *indexMinPQ) swim(i int) {
    k := pq.ind[i]
    for k > 1 {
        p := k/2
        if pq.less(k, p) {
            // swap
            ki := pq.invind[k]
            pi := pq.invind[p]
            pq.keys[k], pq.keys[p] = pq.keys[p], pq.keys[k]
            pq.invind[k], pq.invind[p] = pq.invind[p], pq.invind[k]
            pq.ind[ki], pq.ind[pi] = p, k
        } else {
            break
        }
        k = p
    }
}

func (pq *indexMinPQ) insert(i int, val int) {
    if i < pq.n && pq.ind[i] > 0 {
        // update key
        k := pq.ind[i]
        oldval := pq.keys[k]
        pq.keys[k] = val
        if val < oldval {
            pq.swim(i)
        } else if val > oldval {
            pq.sink(i)
        }
        return
    } 
    
    // insert new key
    
    if i >= len(pq.ind) {
        // resize index array
        newind := make([]int, i+1)
        for k := range newind {
            if k < pq.n {
                newind[k] = pq.ind[k]
            } else {
                newind[k] = -1
            }
        }
        pq.ind = newind
    }
    pq.ind[i] = pq.n+1
    
    pq.keys = append(pq.keys, val)
    pq.invind = append(pq.invind, i)
    pq.n += 1
    
    pq.swim(i)
}

// returns min element
func (pq *indexMinPQ) min() int {
    return pq.keys[1]
}

// return index of min element
func (pq *indexMinPQ) minIndex() int {
    return pq.invind[1]
}

// returns index of min element after deletion
func (pq *indexMinPQ) delMin() int {
    minind := pq.invind[1]
    maxind := pq.invind[pq.n]
    
    pq.keys[1] = pq.keys[pq.n]
    pq.keys = pq.keys[:pq.n]
    pq.invind[1] = pq.invind[pq.n]
    pq.invind = pq.invind[:pq.n]
    pq.ind[minind] = -1
    pq.ind[maxind] = 1
    
    pq.n -= 1
    if pq.n > 0 {
        pq.sink(pq.invind[1])    
    }
    
    return minind
}