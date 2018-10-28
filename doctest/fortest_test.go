package fortest

import (
	"testing"
)

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 8}
	expected := 18
	actual := Sum(nums)
	if actual != expected {
		t.Errorf("Expected the sum of %v to be %d but instead got %d!", nums, expected, actual)
	}
}
