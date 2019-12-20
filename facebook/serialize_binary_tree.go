package facebook

import (
	"strings"
)

type TreeNode struct {
	Value int32
	Left *TreeNode
	Right *TreeNode
}

func TreesAreEqual(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil  {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Value != t2.Value {
		return false
	}

	l := TreesAreEqual(t1.Left, t2.Left)
	r := TreesAreEqual(t1.Right, t2.Right)
	return l && r
}

func SerializeBinaryTree(t *TreeNode) []byte {
	var sb strings.Builder
	var helper func(*TreeNode)
	helper = func(t *TreeNode) {
		if t == nil {
			sb.Write(make([]byte, 4))
			return
		}
		sb.Write(itobytes(t.Value))
		helper(t.Left)
		helper(t.Right)
	}
	helper(t)

	return []byte(sb.String())
}

func DeserializeBinaryTree(b []byte) *TreeNode {
	offset := 0
	var helper func() *TreeNode
	helper = func() *TreeNode {
		if offset >= len(b) {
			return nil
		}

		buf := b[offset:offset+4]
		offset += 4
		val := bytestoi(buf)
		if val == 0 {
			// terminal node
			return nil
		}

		node := &TreeNode{
			Value: val,
		}
		node.Left = helper()
		node.Right = helper()

		return node
	}
	root := helper()
	return root
}

func itobytes(n int32) []byte {
	if n <= 0 {
		panic("only n > 0 is supported")
	}
	res := make([]byte, 4)
	mask := int32(1) << 8 - 1
	for i := len(res) - 1; i >= 0; i -= 1 {
		res[i] = byte(mask & n)
		n = n >> 8
	}
	return res
}

func bytestoi(b []byte) int32 {
	res := int32(0)
	for i := 0; i < 4; i += 1 {
		res = res * 256 + int32(b[i])
	}
	return res
}
