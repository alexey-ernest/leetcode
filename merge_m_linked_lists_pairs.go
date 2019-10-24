// merges two sorted linked lists
func mergeLists(l1, l2 *ListNode) *ListNode {
    var res *ListNode
    var prev *ListNode
    for l1 != nil && l2 != nil {
        var next *ListNode
        if l2.Val < l1.Val {
            next = l2
            l2 = l2.Next
        } else {
            next = l1
            l1 = l1.Next
        }
        if res == nil {
            res = next
        }
        if prev != nil {
            prev.Next = next
        }
        prev = next
    }
    
    var next *ListNode
    if l1 != nil {
        next = l1
    }
    if l2 != nil {
        next = l2
    }
    if res == nil {
        res = next
    }
    if prev != nil {
        prev.Next = next
    }
    return res
}

func mergeKListsPairs(lists []*ListNode) *ListNode {
    // O(N*lg(m)) with merging lists by pairs
    if len(lists) == 0 {
        return nil
    }
    
    for len(lists) > 1 {
        // pairing lists
        pairs := len(lists)/2
        if len(lists)%2 > 0 {
            pairs += 1
        }
        newlists := make([]*ListNode, pairs)
        for i := 0; i < pairs; i+=1 {
            if 2*i == len(lists)-1 {
                newlists[i] = lists[2*i]
                continue
            }
            newlists[i] = mergeLists(lists[2*i], lists[2*i+1])
        }
        lists = newlists
    }
    
    return lists[0]
}