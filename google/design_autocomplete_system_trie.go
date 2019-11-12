import "strings"

const (
    R int = 27 // ASCII
    hotMax int = 3
)

type hotnode struct {
    val int
    key string
}

type node struct {
    child []*node
    hottest []*hotnode // cache of the most hottest keys in current branch
    val int
}

func newNode() *node {
    return &node {
        child: make([]*node, R),
    }
}

func charidx(c byte) int {
    if c == ' ' {
        return 0
    }
    return int(c-'a')
}

type AutocompleteSystem struct {
    root *node
    query string
    querynode *node
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
    tr := AutocompleteSystem{}
    
    // building a trie
    for i, s := range sentences {
        tr.root = tr.insert(tr.root, s, times[i], 0)
    }
    
    tr.querynode = tr.root
    return tr
}

func (this *AutocompleteSystem) insert(x *node, key string, val int, d int) *node {
    if x == nil {
        x = newNode()
    }
    
    if d == len(key) {
        x.val = val
        
        currhottest := []*hotnode{
            &hotnode{
                val: val,
                key: key,
            },
        }   
        if x.hottest == nil {
            // new node
            x.hottest = currhottest
        } else {
            // merge exiting hottest list with the current node
            newhottest := make([]*hotnode, 0)
            for _, r := range x.hottest {
                if r.key == key {
                    continue
                }
                newhottest = append(newhottest, r)
            }
            x.hottest = this.mergetwo(newhottest, currhottest)
        }
        
        return x
    }
    
    c := charidx(key[d])
    x.child[c] = this.insert(x.child[c], key, val, d+1)
    
    // merge R hottest lists to get top 3 hottest nodes at each level
    // O(3*R)
    x.hottest = this.mergeHottest(x)
    if x.val > 0 {
        currkey := key[:d]
        currhottest := []*hotnode{
            &hotnode{
                val: x.val,
                key: currkey,
            },
        }
        x.hottest = this.mergetwo(x.hottest, currhottest)
    }
    return x
}

func (this *AutocompleteSystem) mergeHottest(x *node) []*hotnode {
    lists := make([][]*hotnode, 0, R)
    for i := range x.child {
        if x.child[i] == nil {
            continue
        }
        lists = append(lists, x.child[i].hottest)
    }
    
    for len(lists) > 1 {
        newlists := make([][]*hotnode, 0)
        for i := 0; 2*i < len(lists); i += 1 {
            var m []*hotnode
            if 2*i + 1 < len(lists) {
                m = this.mergetwo(lists[2*i], lists[2*i+1])
            } else {
                m = lists[2*i]
            }
            newlists = append(newlists, m)
        }
        lists = newlists
    }
    
    hottest := lists[0]    
    return hottest
}

func (this *AutocompleteSystem) mergetwo(l1 []*hotnode, l2 []*hotnode) []*hotnode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    
    res := make([]*hotnode, 0, len(l1) + len(l2))
    i, j := 0, 0
    for i < len(l1) && j < len(l2) && len(res) < hotMax {
        if l1[i].val > l2[j].val {
            res = append(res, l1[i])
            i++
        } else if l1[i].val < l2[j].val {
            res = append(res, l2[j])
            j++
        } else if strings.Compare(l1[i].key, l2[j].key) <= 0 {
            res = append(res, l1[i])
            i++
        } else {
            res = append(res, l2[j])
            j++
        }
    }
    
    if len(res) == hotMax {
        return res
    }
    
    for i < len(l1) {
        res = append(res, l1[i])
        i++
    }
    for j < len(l2) {
        res = append(res, l2[j])
        j++
    }
    
    if len(res) > hotMax {
        res = res[:hotMax]
    }
    
    return res
}

func (this *AutocompleteSystem) search(x *node, key string, d int) *node {
    if x == nil {
        return nil
    }
    
    if d == len(key) {
        return x
    }
    
    c := charidx(key[d])
    return this.search(x.child[c], key, d+1)
}

func (this *AutocompleteSystem) Input(c byte) []string {
    if c == '#' {
        // adding search to the history
        x := this.search(this.root, this.query, 0)
        if x == nil {
            this.insert(this.root, this.query, 1, 0)
        } else {
            this.insert(this.root, this.query, x.val+1, 0)
        }
        
        this.query = ""
        this.querynode = this.root
        return nil
    }
    
    // continue search symbol by symbol
    this.query += string(c)
    if this.querynode == nil {
        return nil
    }
    
    this.querynode = this.search(this.querynode, string(c), 0)
    if this.querynode == nil {
        return nil
    }
    
    hottest := make([]string, len(this.querynode.hottest))
    for i, hn := range this.querynode.hottest {
        hottest[i] = hn.key
    }
    return hottest
}


/**
 * Your AutocompleteSystem object will be instantiated and called as such:
 * obj := Constructor(sentences, times);
 * param_1 := obj.Input(c);
 */