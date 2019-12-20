package facebook

import "testing"
//import "fmt"

func TestSerializeBinaryTree(t *testing.T) {
	root := &TreeNode {
		Value: 1,
		Left: &TreeNode {
			Value: 2,
		},
		Right: &TreeNode {
			Value: 3,
			Left: &TreeNode {
				Value: 4,
			},
			Right: &TreeNode {
				Value: 5,
			},
		},
	}
	res := SerializeBinaryTree(root)
	exp := []byte("\x00\x00\x00\x01\x00\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x05\x00\x00\x00\x00\x00\x00\x00\x00")
	if len(res) != len(exp) {
		t.Fatalf("expected byte slice of len %d, but got %d", len(exp), len(res))
	}

	for i := 0; i < len(res); i += 1 {
		if res[i] != exp[i] {
			t.Fatalf("%v != %v", res, exp)
		}
	}

	droot := DeserializeBinaryTree(res)
	if !TreesAreEqual(droot, root) {
		t.Fatalf("deserialized tree is not equal to the original one")
	}
}