package offer

/**
  @author: CodeWater
  @since: 2023/6/26
  @desc: 最小的k个数
**/
func getLeastNumbers(arr []int, k int) []int {
	quickSort(arr, 0, len(arr)-1)

	return arr[:k]
}

func quickSort(q []int, l, r int) {
	if l >= r {
		return
	}
	i, j := l-1, r+1
	x := q[(i+j)>>1]
	for i < j {
		for {
			i++
			if q[i] >= x {
				break
			}
		}
		for {
			j--
			if q[j] <= x {
				break
			}
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}
	quickSort(q, l, j)
	quickSort(q, j+1, r)

}
