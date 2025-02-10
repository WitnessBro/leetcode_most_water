package move_zero_to_the_end

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	var prices = []int{1, 8, 6, 2, 5, 4, 8, 3, 7} // Input array
	var expectedResult = 49                       // The expected answer with correct length.
	res := maxArea(prices)                        // Calls your implementation

	assert.Equal(t, expectedResult, res)
}
func maxArea(height []int) int {
	result := 0
	leftP := 0
	rightP := len(height) - 1
	for i := 0; i < len(height)-1; i++ {
		if height[leftP] <= height[rightP] {
			if (rightP-leftP)*height[leftP] >= result {
				result = (rightP - leftP) * height[leftP]
			}
			leftP++
		} else {
			if (rightP-leftP)*height[rightP] >= result {
				result = (rightP - leftP) * height[rightP]
			}
			rightP--
		}
	}
	return result
}
