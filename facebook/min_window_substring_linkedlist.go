func minWindow(s string, t string) string {
    if len(t) == 0 {
        return ""
    }
    
    chars := make(map[byte]int)
    for i := 0; i < len(t); i += 1 {
        chars[t[i]]++
    }
    
    res := ""
    listindexes := list{}
    mapidx := make(map[byte]*queue)
    charscount := 0
    
    for i := 0; i < len(s); i += 1 {
        if _, ok := chars[s[i]]; !ok {
            continue
        }
        
        idx := &dnode{
            val: i,
        }
        listindexes.append(idx)
        
        if mapidx[s[i]] == nil {
            mapidx[s[i]] = &queue{}
        }
        mapidx[s[i]].enqueue(&node{ val: idx })
        charscount++
        
        if mapidx[s[i]].size > chars[s[i]] {
            // dequeue old index
            oldidx := mapidx[s[i]].dequeue()
            listindexes.remove(oldidx.val)
            charscount--
        }
        
        if charscount < len(t) {
            continue
        }
        
        l := listindexes.head.val
        r := listindexes.tail.val
        if len(res) == 0 || r - l + 1 < len(res) {
            res = s[l:r+1]
        }
    }
    
    return res
}

type node struct {
    next *node
    val *dnode
}

type queue struct {
    head *node
    tail *node
    size int
}

func (q *queue) enqueue(x *node) {
    if q.size == 0 {
        q.head = x
        q.tail = x
        q.size++
        return
    }
    
    q.tail.next = x
    q.tail = x
    q.size++
}

func (q *queue) dequeue() *node {
    if q.size == 0 {
        return nil
    }
    
    if q.size == 1 {
        x := q.head
        q.head = nil
        q.tail = nil
        q.size--
        return x
    }
    
    x := q.head
    q.head = q.head.next
    x.next = nil
    q.size--
    return x
}

type dnode struct {
    next *dnode
    prev *dnode
    val int
}

type list struct {
    head *dnode
    tail *dnode
    size int
}

func (l *list) append(x *dnode) {
    if l.size == 0 {
        l.head = x
        l.tail = x
        l.size++
        return
    }
    
    l.tail.next = x
    x.prev = l.tail
    l.tail = x
    l.size++
}

func (l *list) remove(x *dnode) {
    pr := x.prev
    ne := x.next
    
    if pr != nil {
        pr.next = ne    
    }
    if ne != nil {
        ne.prev = pr   
    }
    
    x.next = nil
    x.prev = nil
    l.size--
    
    if l.head == x {
        l.head = ne
    }
    if l.tail == x {
        l.tail = pr
    }
}