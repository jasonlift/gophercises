package dynamics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximumSubarray(t *testing.T) {
	input := []int{-2,1,-3,4,-1,2,1,-5,4}
	desired := 6

	assert.Equal(t, desired, maxSubArray(input))
}
