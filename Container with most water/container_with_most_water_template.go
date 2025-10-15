package main

import "fmt"

// Solution struct
type Solution struct{}

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
 * @param height Slice of non-negative integers representing heights
 * @return The maximum amount of water a container can store
 */
func (s *Solution) maxArea(height []int) int {
	// TODO: Implement your solution here
	return 0
}

// ---------- Example Usage ----------
func main() {
	solution := Solution{}

	// Example 1
	height1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("Example 1 Output:", solution.maxArea(height1)) // Expected: 49

	// Example 2
	height2 := []int{1, 1}
	fmt.Println("Example 2 Output:", solution.maxArea(height2)) // Expected: 1
}

