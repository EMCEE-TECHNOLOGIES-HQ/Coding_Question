class Solution {
    /**
     * Given n non-negative integers height[i], where each represents a vertical line at coordinate (i, height[i]),
     * find two lines that together with the x-axis form a container such that the container contains the most water.
     * 
     * Return the maximum amount of water a container can store.
     *
     * Constraints:
     * - 2 <= n <= 10^5
     * - 0 <= height[i] <= 10^4
     * 
     * @param height Array of non-negative integers representing heights
     * @return The maximum amount of water a container can store
     */
    public int maxArea(int[] height) {
        // TODO: Implement your solution here
        return 0;
    }
}

// ---------- Example Usage ----------
public class ContainerWithMostWaterTemplate {
    public static void main(String[] args) {
        Solution solution = new Solution();
        
        // Example 1
        int[] height1 = {1, 8, 6, 2, 5, 4, 8, 3, 7};
        System.out.println("Example 1 Output: " + solution.maxArea(height1)); // Expected: 49
        
        // Example 2
        int[] height2 = {1, 1};
        System.out.println("Example 2 Output: " + solution.maxArea(height2)); // Expected: 1
    }
}

