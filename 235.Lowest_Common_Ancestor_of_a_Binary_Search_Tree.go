package leetcode

import "container/list"

// the tree is BST
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if nil == root {
		return root
	}
	for nil != root {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}

// the tree is common tree
func lowestCommonAncestor1(node1, node2, root *TreeNode2) *TreeNode2 {
	path1 := list.New()
	path2 := list.New()
	if nil == root || nil == node1 || nil == node2 {
		return nil
	}
	getPath(node1, root, path1)
	getPath(node2, root, path2)
	var lastParentNode *TreeNode2
	for path1.Front() != nil && path2.Front() != nil {
		if path1.Front().Value == path2.Front().Value {
			lastParentNode = path1.Front().Value.(*TreeNode2)
		}
		path1.Remove(path1.Front())
		path2.Remove(path2.Front())
	}
	return lastParentNode
}

func getPath(node, root *TreeNode2, path *list.List) bool {
	if node == root {
		return true
	}
	path.PushBack(root)
	found := false
	for i:= 0; i < len(root.Child) && !found; i++ {
		found = getPath(node, root.Child[i], path)
	}
	if !found {
		path.Remove(path.Back())
	}
	return found
}

