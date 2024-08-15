package src

// Rect 矩形
type Rect struct {
	x, y          float64
	width, height float64
}

// Area 成员方法用于计算面积
func (r *Rect) Area() float64 {
	return r.width * r.height
}

// NewRect 使用NewXxx开头标识构造函数
func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}
