import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    h := height(root)
    if h == 0 {
        return 1
    }
    
    // binary search to find the last leaf of the last level from left to right
    // O(log(2^h))=O(h)
    l, r := 0, int(math.Pow(2.0, float64(h)))-1
    lastidx := 1
    for l <= r {
        mid := (l+r)/2
        // O(log(2^h))=O(h)
        if !existsIndex(root, mid, 0, h) {
            r = mid - 1
        } else {
            lastidx = mid
            l = mid + 1
        }
    }
    count := int(math.Pow(2.0, float64(h))) + lastidx
    return count
}

func height(n *TreeNode) int {
    if n == nil {
        return 0
    }
    
    h := 0
    for n.Left != nil {
        n = n.Left
        h += 1
    }
    return h
}

// checks if the node exists with specified index
func existsIndex(n *TreeNode, idx int, height int, maxheight int) bool {
    if n == nil {
        return false
    }
    if height == maxheight {
        return true
    }
    
    maxidx := int(math.Pow(2.0, float64(maxheight-height)))-1
    mididx := maxidx/2
    if idx <= mididx {
        return existsIndex(n.Left, idx, height+1, maxheight)
    } else {
        return existsIndex(n.Right, idx-(mididx+1), height+1, maxheight)
    }
}
