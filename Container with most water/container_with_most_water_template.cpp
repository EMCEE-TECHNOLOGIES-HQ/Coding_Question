#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

class Solution {
public:
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
     * @param height Vector of non-negative integers representing heights
     * @return The maximum amount of water a container can store
     */
    int maxArea(vector<int>& height) {
        // TODO: Implement your solution here
        return 0;
    }
};

// ---------- Example Usage ----------
int main() {
    Solution solution;
    
    // Example 1
    vector<int> height1 = {1, 8, 6, 2, 5, 4, 8, 3, 7};
    cout << "Example 1 Output: " << solution.maxArea(height1) << endl; // Expected: 49
    
    // Example 2
    vector<int> height2 = {1, 1};
    cout << "Example 2 Output: " << solution.maxArea(height2) << endl; // Expected: 1
    
    return 0;
}

