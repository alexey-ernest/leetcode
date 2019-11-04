import "fmt"

const MaxInt = int(^uint(0) >> 1)

// Dijkstra algorithm: MlgN
func networkDelayTime(times [][]int, N int, K int) int {
    
    // build weighted digraph
    adj := make([][]*edge, N+1)
    for _, t := range times {
        adj[t[0]] = append(adj[t[0]], &edge{ from: t[0], to: t[1], weight: t[2] })
    }
    
    // for storing distances from the source
    distTo := make([]int, N+1)
    for i := range distTo {
        distTo[i] = MaxInt
    }
    distTo[0] = 0
    distTo[K] = 0
    
    pq := NewIndexMinPQ()
    pq.insert(K, 0)
    
    for !pq.isEmpty() {
        relax(pq.delMin(), adj, distTo, &pq)
    }
    //fmt.Printf("%+v\n", distTo)
    
    max := 0
    for _, r := range distTo {
        if r == MaxInt {
            return -1
        }
        if r > max {
            max = r
        }
    }
    
    return max
}

func relax(v int, adj [][]*edge, distTo []int, pq *indexMinPQ) {
    for _, e := range adj[v] {
        w := e.to
        if distTo[w] > distTo[v] + e.weight {
            distTo[w] = distTo[v] + e.weight
            pq.insert(w, distTo[w]) // insert or update
        }
    }
}

type edge struct {
    from int
    to int
    weight int
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