// Leetcode: https://leetcode.com/problems/h-index/#/description
// Given an array A, where A[i] is the # of citations for a researcher's i'th paper,  write a
// function to compute the researcher's h-index. A scientist has index h if h of her N papers
// have at least h citations, and the other N − h papers h or fewer citations."

// Brute Force: For h = n-1..0, iterate the array to calculate if h is a valid h-index. O(n^2) time.
// Better: Sort the array, then iterate. At each iteration, partition the array into two subarrays.
// Check if left subarray has <= h citations, and right subarray has h items with >= h citations.
// O(n*logn) time and O(n) space.

package leetcode

import (
  "sort"
  "testing"
)

func hIndex(citations []int) int {
  sort.Ints(citations)
  size := len(citations)
  for i := 0; i < size; i++ {
    h := size - i
    if citations[i] >= h { return h }
  }
  return 0
}

func testHIndex(t *testing.T, citations []int, expected_h int) bool {
  h := hIndex(citations)
  if h != expected_h {
    t.Errorf("Expected h-index to be %d, but got %d\n", expected_h, h)
  }
  return true
}

func TestHIndex(t *testing.T) {
  testHIndex(t, []int {}, 0)
  testHIndex(t, []int {0}, 0)
  testHIndex(t, []int {0,0,0,0}, 0)
  testHIndex(t, []int {1}, 1)
  testHIndex(t, []int {11}, 1)
  testHIndex(t, []int {0,1,1}, 1)
  testHIndex(t, []int {0,1,3,3,3}, 3)
  testHIndex(t, []int {3,2,3,3,0}, 3)
  testHIndex(t, []int {0,3,3,3,3}, 3)
  testHIndex(t, []int {3,3,3,3,3}, 3)
  testHIndex(t, []int {5,0,5,5,5}, 4)
  testHIndex(t, []int {5,5,5,5,5}, 5)
  testHIndex(t, []int {6,5,5,5,5,5}, 5)
  testHIndex(t, []int {100,100}, 2)
  testHIndex(t, []int {3,0,6,1,5}, 3)
  testHIndex(t, []int {0,0,0,0,2,2}, 2)
  testHIndex(t, []int {0,0,0,0,2,2,3}, 2)
  testHIndex(t, []int {0,0,0,0,3,4,5}, 3)
}
