//import "math"
import "fmt"

type Graph struct {
    e int
    v int
    indexes map[int]string
    keys map[string]int
    adj [][]int
}

func NewGraph(v []string) Graph {
    n := len(v)
    indexes := make(map[int]string)
    keys := make(map[string]int)
    for i, name := range v {
        indexes[i] = name
        keys[name] = i
    }
    
    return Graph{
        v: n,
        indexes: indexes,
        keys: keys,
        adj: make([][]int, n),
    }
}

func (g *Graph) AddEdge(v, w string) {
    vi := g.keys[v]
    wi := g.keys[w]
    g.adj[vi] = append(g.adj[vi], wi)
    g.e++
}

func (g *Graph) Name(v int) string {
    return g.indexes[v]
}

func (g *Graph) Index(name string) int {
    return g.keys[name]
}

func (g *Graph) Adj(v int) []int {
    return g.adj[v]
}

func (g *Graph) E() int {
    return g.e
}

func (g *Graph) V() int {
    return g.v
}

type GraphPath struct {
    graph *Graph
}

func NewGraphPath(g *Graph) GraphPath {
    
    return GraphPath{
        graph: g,
    }
}

func (d *GraphPath) Path(source int) []string {
    path := make([]int, d.graph.V())
    marked := make([]bool, d.graph.V())
    
    // find a directed path containing all the vertices (Hamilton path)
    d.dfs(marked, path, source, 1, d.graph.V())
    
    res := make([]string, len(path))
    for i := range path {
        res[i] = d.graph.Name(path[i])
    }
    return res
}

func (d *GraphPath) dfs(marked []bool, path []int, s int, counter int, total int) bool {
    marked[s] = true
    
    if counter == total {
        path[counter-1] = s
        return true
    }
    
    for _, w := range d.graph.Adj(s) {
        if marked[w] == true {
            // visited vertex
            continue
        }
        
        if d.dfs(marked, path, w, counter+1, total) {
            path[counter-1] = s
            return true
        } else {
            marked[w] = false
        }
    }
    
    return false
}

func crackSafe(n int, k int) string {
    // generate all the passwords, k^n
    //N := int(math.Pow(float64(k), float64(n)))
    passwords := generatePasswords(n, k)
    
    passpref := make(map[string][]int)
    for i, p := range passwords {
        pref := p[:k-1]
        passpref[pref] = append(passpref[pref], i)
    }
    
    // building graph
    g := NewGraph(passwords)
    for _, p := range passwords {
        pref := p[1:]
        for _, w := range passpref[pref] {
            if p == passwords[w] {
                continue
            }
            g.AddEdge(p, passwords[w])
        }
    }
    
    ep := NewGraphPath(&g)
    source := passwords[0]
    
    path := ep.Path(g.Index(source))
    
    res := ""
    for _, p := range path {
        if res == "" {
            res += p
        } else {
            res += p[k-1:]
        }
    }
    
    return res
}

func generatePasswords(n,k int) []string {
    if n == 0 {
        return nil
    }
    
    gen := generatePasswords(n-1, k)
    res := make([]string, 0)
    for i := 0; i < k; i += 1 {
        res = append(res, fmt.Sprintf("%d", i))
    }
    
    if gen == nil {
        return res
    }

    newres := make([]string, 0, len(res)*len(gen))
    for _, r := range res {
        for _, g := range gen {
            newres = append(newres, r + g)
        }
    }
    res = newres
    
    return res
}