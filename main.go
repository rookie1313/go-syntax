package main

import (
	"container/list"
)

func main() {
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		Val:   value,
		Left:  nil,
		Right: nil,
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* 层序遍历 */
func levelOrder(root *TreeNode) []int {
	// 初始化队列，加入根节点
	queue := list.New()
	queue.PushBack(root)
	// 初始化一个切片，用于保存遍历序列
	nums := make([]int, 0)
	for queue.Len() > 0 {
		// 队列出队
		node := queue.Remove(queue.Front()).(*TreeNode)
		// 保存节点值
		nums = append(nums, node.Val)
		if node.Left != nil {
			// 左子节点入队
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			// 右子节点入队
			queue.PushBack(node.Right)
		}
	}
	return nums
}
