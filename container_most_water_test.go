// Leetcode: https://leetcode.com/problems/container-with-most-water/#/description
// Given a histogram, represented as an integer array, find the pair of lines which together
// with the x-axis traps the most water.

// Brute Force: Examine all pairs of lines, calculating the trapped water. Track max. O(n^2)
// Better: Iterate the array from both ends. At each iteration, discard the shorter line in
// the pair, advancing that endpoint inward. Track the max. O(n) time and O(1) space.

package leetcode

import (
	"testing"
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	max_area, i, j := 0, 0, len(height)-1

	for i < j {
		area := min(height[i], height[j]) * (j - i)
		max_area = max(area, max_area)

		if height[i] < height[j] {
			i++
		} else if height[i] > height[j] {
			j--
		} else { // height[i] == height[j]
			i++
			j--
		}
	}
	return max_area
}

func test_max_area(t *testing.T, height []int, expected_area int) {
	area := maxArea(height)
	if area != expected_area {
		t.Errorf("Expected area to be %d, but got %d\n", expected_area, area)
	}
}

func TestMaxArea(t *testing.T) {
	test_max_area(t, []int{1, 2, 5, 4, 0, 5}, 15)
	test_max_area(t, []int{20, 20, 5, 4, 0, 5}, 25)
	test_max_area(t, []int{50, 50, 5, 4, 0, 5}, 50)
	test_max_area(t, []int{5, 5, 5, 4, 0, 5}, 25)
}
