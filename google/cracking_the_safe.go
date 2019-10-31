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

// finds all the disjoint paths
func (d *GraphPath) DisjointHamiltonPaths() [][]string {
    marked := make([]bool, d.graph.V())
    
    paths := make([][]string, 0)
    for i := 0; i < d.graph.V(); i += 1 {
        if marked[i] {
            continue
        }
        
        path := d.dfs(marked, i)
        if len(path) == 0 {
            continue
        }
        
        // marking already traversed path to exclude it from the next pass
        for _, w := range path {
            marked[w] = true
        }
        
        namedpath := make([]string, 0)
        for j := 0; j < len(path); j += 1 {
            namedpath = append(namedpath, d.graph.Name(path[j]))
        }
        paths = append(paths, namedpath)
    }
    
    return paths
}

func (d *GraphPath) dfs(marked []bool, s int) []int {
    marked[s] = true
    
    maxp := make([]int, 0)
    for _, w := range d.graph.Adj(s) {
        if marked[w] == true {
            // visited vertex
            continue
        }
        
        p := d.dfs(marked, w)
        
        if len(p) > len(maxp) {
            maxp = p
        }
    }
    
    marked[s] = false
    
    maxp = append([]int{s}, maxp...)
    return maxp
}

func crackSafe(n int, k int) string {
    // generate all the passwords, k^n
    //N := int(math.Pow(float64(k), float64(n)))
    passwords := generatePasswords(n, k)
    
    passpref := make(map[string][]int)
    for i, p := range passwords {
        pref := p[:n-1]
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
    
    paths := ep.DisjointHamiltonPaths()
    
    res := ""
    for _, path := range paths {
        res += path[0]
        for i := 1; i < len(path); i += 1 {
            res += path[i][n-1:]
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