import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(pre []int, post []int) *TreeNode {
    //fmt.Printf("pre %+v post %+v\n", pre, post)
    if len(pre) == 0 {
        return nil
    }
    x := &TreeNode{
        Val: pre[0],
    }
    if len(pre) == 1 {
        return x
    }
    
    L := 0
    for i := 0; i < len(post); i += 1 {
        if post[i] == pre[1] {
            L = i+1
            break
        }
    }
    x.Left = constructFromPrePost(pre[1:L+1], post[:L])
    x.Right = constructFromPrePost(pre[L+1:], post[L:len(post)-1])
    return x
}