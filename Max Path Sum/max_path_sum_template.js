// Definition for a binary tree node.
class TreeNode {
    constructor(val = 0, left = null, right = null) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
}

class Solution {
    /**
     * Given the root of a binary tree, return the maximum path sum of any non-empty path.
     * 
     * A path in a binary tree is a sequence of nodes where each pair of adjacent nodes 
     * in the sequence has an edge connecting them. A node can only appear in the sequence 
     * at most once. The path does not need to pass through the root.
     * 
     * The path sum of a path is the sum of the node's values in the path.
     * 
     * @param {TreeNode} root - The root of the binary tree
     * @return {number} - The maximum path sum
     */
    maxPathSum(root) {
        // TODO: Implement your solution here
        return 0;
    }
}

// ---------- Example Usage ----------
function main() {
    const solution = new Solution();
    
    // Example 1
    // Tree: [1, 2, 3]
    //       1
    //      / \
    //     2   3
    const root1 = new TreeNode(1);
    root1.left = new TreeNode(2);
    root1.right = new TreeNode(3);
    
    console.log("Example 1 Output:", solution.maxPathSum(root1)); // Expected: 6 (2 + 1 + 3)
    
    // Example 2
    // Tree: [-10, 9, 20, null, null, 15, 7]
    //        -10
    //        /  \
    //       9    20
    //           /  \
    //          15   7
    const root2 = new TreeNode(-10);
    root2.left = new TreeNode(9);
    root2.right = new TreeNode(20);
    root2.right.left = new TreeNode(15);
    root2.right.right = new TreeNode(7);
    
    console.log("Example 2 Output:", solution.maxPathSum(root2)); // Expected: 42 (15 + 20 + 7)
}

main();

