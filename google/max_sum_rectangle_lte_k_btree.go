func maxSumSubmatrix(matrix [][]int, k int) int {
    n := len(matrix)
    if n == 0 {
        return 0
    }
    m := len(matrix[0])
    if m == 0 {
        return 0
    }
    
    maxrect := -int(^uint(0) >> 1)-1
    // O(m)
    for i := range matrix[0] {
        sums := make([]int, n)
        // O(m)
        for j := i; j < m; j += 1 {
            // O(n)
            for p := range matrix {
                sums[p] += matrix[p][j]
            }
            
            sum := 0
            t := bst{}
            t.root = t.insert(t.root, 0)
            // O(n*logn)
            for p := range sums {
                sum += sums[p]
                ceiling := t.ceiling(t.root, sum - k)
                if ceiling != nil {
                    maxrect = max(maxrect, sum - ceiling.key)
                }
                t.root = t.insert(t.root, sum)
            }
        }
    }
    
    return maxrect
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

type node struct {
    left *node
    right *node
    key int
}

type bst struct {
    root *node
}

func (t *bst) insert(x *node, key int) *node {
    if x == nil {
        return &node{
            key: key,
        }
    }
    
    if x.key == key {
        return x
    }
    
    if x.key > key {
        x.left = t.insert(x.left, key)
    } else {
        x.right = t.insert(x.right, key)
    }
    
    return x
}

func (t *bst) ceiling(x *node, key int) *node {
    if x == nil {
        return nil
    }
    
    if x.key == key {
        return x
    }
    
    if x.key < key {
        return t.ceiling(x.right, key)
    }

    l := t.ceiling(x.left, key)
    if l != nil {
        return l
    }
    return x    
}