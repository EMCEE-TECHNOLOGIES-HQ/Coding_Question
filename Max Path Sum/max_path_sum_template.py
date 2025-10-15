# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def maxPathSum(self, root: TreeNode) -> int:
        """
        Given the root of a binary tree, return the maximum path sum of any non-empty path.
        
        A path in a binary tree is a sequence of nodes where each pair of adjacent nodes 
        in the sequence has an edge connecting them. A node can only appear in the sequence 
        at most once. The path does not need to pass through the root.
        
        The path sum of a path is the sum of the node's values in the path.
        
        Args:
            root: TreeNode - The root of the binary tree
            
        Returns:
            int - The maximum path sum
        """
        # TODO: Implement your solution here
        pass


# ---------- Example Usage ----------
def main():
    # Example 1
    # Tree: [1, 2, 3]
    #       1
    #      / \
    #     2   3
    root1 = TreeNode(1)
    root1.left = TreeNode(2)
    root1.right = TreeNode(3)
    
    print("Example 1 Output:", Solution().maxPathSum(root1))  # Expected: 6 (2 + 1 + 3)

    # Example 2
    # Tree: [-10, 9, 20, None, None, 15, 7]
    #        -10
    #        /  \
    #       9    20
    #           /  \
    #          15   7
    root2 = TreeNode(-10)
    root2.left = TreeNode(9)
    root2.right = TreeNode(20)
    root2.right.left = TreeNode(15)
    root2.right.right = TreeNode(7)

    print("Example 2 Output:", Solution().maxPathSum(root2))  # Expected: 42 (15 + 20 + 7)

if __name__ == "__main__":
    main()

