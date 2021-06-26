package count

func Sort(arr []int) {
	copy(arr, sorting(arr))
}

// 为不浪费空间 数组每个元素是某一个范围内 有点大材小用
// 下标计算就需要偏移量了min
func sorting(arr []int) []int {
	// 计算出数列的最大长度
	max, min := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] > min {
			min = arr[i]
		}
	}
	// 统计出现的次数 max-min+1
	countArrays := make([]int, max-min+1)
	for i := 0; i < len(arr); i++ {
		countArrays[arr[i]]++
	}
	sortedArrays := make([]int, len(arr))
	// 排序
	index := 0
	for i := 0; i < len(countArrays); i++ {
		for j := 0; j < countArrays[i]; j++ {
			sortedArrays[index] = i
			index += 1
		}
	}
	return sortedArrays
}
