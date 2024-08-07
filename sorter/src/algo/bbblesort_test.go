package algo

import "testing"

func TestBubbleSort(t *testing.T) {
	vals := []int{5, 4, 2, 3, 1}
	BubbleSort(vals)
	if vals[0] != 1 || vals[4] != 5 {
		t.Error("BubbleSort failed,got ", vals)
	}
}
