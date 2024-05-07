package main

import "fmt"

type TreeNode struct {
	left, right, sum int
	LeftNode         *TreeNode
	RightNode        *TreeNode
}

func main() {
	nums := []int{3, 5, 8, 1, 2, 6, 7}
	tree := buildTree(nums, 0, len(nums)-1)
	ans := lookup(tree, 2, 6)
	fmt.Println(ans)

	modify(tree, 2, 10)
	ans = lookup(tree, 2, 6)
	fmt.Println(ans)
}

func buildTree(nums []int, left, right int) *TreeNode {
	if left == right {
		return &TreeNode{
			left:  left,
			right: right,
			sum:   nums[left],
		}
	}
	root := &TreeNode{
		left:  left,
		right: right,
	}
	mid := (left + right) / 2
	root.LeftNode = buildTree(nums, left, mid)
	root.RightNode = buildTree(nums, mid+1, right)
	root.sum = root.LeftNode.sum + root.RightNode.sum
	return root
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func lookup(root *TreeNode, left, right int) int {
	if overlap(root.left, root.right, left, right) {
		i, j := max(root.left, left), min(right, root.right)
		if root.left == i && root.right == j {
			return root.sum
		}
		leftSum := lookup(root.LeftNode, i, j)
		rightSum := lookup(root.RightNode, i, j)
		return leftSum + rightSum
	}
	return 0
}

func modify(root *TreeNode, idx, val int) int {
	if root.left == root.right && root.right == idx {
		root.sum = val
		return val
	}

	if overlap(root.left, root.right, idx, idx) {
		root.sum = modify(root.LeftNode, idx, val) + modify(root.RightNode, idx, val)
	}
	return root.sum
}

func overlap(left, right, i, j int) bool {
	if i >= left && i <= right {
		return true
	}
	if j >= left && j <= right {
		return true
	}
	return false
}
