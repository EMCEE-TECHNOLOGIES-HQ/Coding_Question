from typing import List

class Solution:
    """
    Given n non-negative integers height[i], where each represents a vertical line at coordinate (i, height[i]),
    find two lines that together with the x-axis form a container such that the container contains the most water.
    
    Return the maximum amount of water a container can store.

    Constraints:
    - 2 <= n <= 10^5
    - 0 <= height[i] <= 10^4
    """
    def maxArea(self, height: List[int]) -> int:
        # TODO: Implement your solution here
        return 0


# ---------- Example Usage ----------
def main():
    solution = Solution()
    
    # Example 1
    height1 = [1, 8, 6, 2, 5, 4, 8, 3, 7]
    print("Example 1 Output:", solution.maxArea(height1))  # Expected: 49
    
    # Example 2
    height2 = [1, 1]
    print("Example 2 Output:", solution.maxArea(height2))  # Expected: 1

if __name__ == "__main__":
    main()

