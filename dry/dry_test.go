package dry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortedUniqueUnion(t *testing.T) {
	expected := []int{1, 3, 4, 5, 6, 7, 8, 9}
	result := sortedUniqueUnion()
	assert.Equal(t, expected, result)
}

func Test_sortedDescendingUniqueIntersect(t *testing.T) {
	expected := []int{7, 5, 3}
	result := sortedDescendingUniqueIntersect()
	assert.Equal(t, expected, result)
}
