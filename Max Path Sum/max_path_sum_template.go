package main

import (
	"fmt"
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Solution struct
type Solution struct{}

/**
 * Given the root of a binary tree, return the maximum path sum of any non-empty path.
 *
 * A path in a binary tree is a sequence of nodes where each pair of adjacent nodes
 * in the sequence has an edge connecting them. A node can only appear in the sequence
 * at most once. The path does not need to pass through the root.
 *
 * The path sum of a path is the sum of the node's values in the path.
 *
 * @param root The root of the binary tree
 * @return The maximum path sum
 */
func (s *Solution) maxPathSum(root *TreeNode) int {
	// TODO: Implement your solution here
	return 0
}

// ---------- Example Usage ----------
func main() {
	solution := Solution{}

	// Example 1
	// Tree: [1, 2, 3]
	//       1
	//      / \
	//     2   3
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}

	fmt.Println("Example 1 Output:", solution.maxPathSum(root1)) // Expected: 6 (2 + 1 + 3)

	// Example 2
	// Tree: [-10, 9, 20, null, null, 15, 7]
	//        -10
	//        /  \
	//       9    20
	//           /  \
	//          15   7
	root2 := &TreeNode{Val: -10}
	root2.Left = &TreeNode{Val: 9}
	root2.Right = &TreeNode{Val: 20}
	root2.Right.Left = &TreeNode{Val: 15}
	root2.Right.Right = &TreeNode{Val: 7}

	fmt.Println("Example 2 Output:", solution.maxPathSum(root2)) // Expected: 42 (15 + 20 + 7)
}

