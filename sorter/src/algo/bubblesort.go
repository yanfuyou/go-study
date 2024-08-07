package algo

// BubbleSort 冒泡排序
func BubbleSort(vals []int) {
	flag := true
	for i := 0; i < len(vals)-1; i++ {
		flag = true
		for j := 0; j < len(vals)-i-1; j++ {
			if vals[j] > vals[j+1] {
				vals[j], vals[j+1] = vals[j+1], vals[j]
				flag = false
			}
		}
		if flag == true {
			break
		}
	}
}
