package src

import (
	"testing"
)

func TestInit(t *testing.T) {
	rect1 := new(Rect)
	rect1.width = 3
	rect1.height = 4
	if rect1.Area() != 12 {
		t.Error("rect1 Area error")
	}

	rect2 := &Rect{}
	rect2.width = 3
	rect2.height = 4
	if rect2.Area() != 12 {
		t.Error("rect2 Area error")
	}
	rect3 := &Rect{0, 0, 3, 4}
	if rect3.Area() != 12 {
		t.Error("rect3 Area error")
	}
	rect4 := &Rect{width: 3, height: 4}
	if rect4.Area() != 12 {
		t.Error("rect4 Area error")
	}

}
