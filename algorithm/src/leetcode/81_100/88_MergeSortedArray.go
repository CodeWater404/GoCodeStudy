package _1_100

import "sort"

/**
  @author: CodeWater
  @since: 2023/5/30
  @desc: 88. 合并两个有序数组
**/

//==========================first method========================
func merge(nums1 []int, m int, nums2 []int, n int) {
	temp, q := make([]int, m+n), append(nums1[:m], nums2[:n]...)
	mergeSort(q, temp, 0, m+n-1)
	copy(nums1, q)
}

func mergeSort(q, temp []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	mergeSort(q, temp, l, mid)
	mergeSort(q, temp, mid+1, r)
	k, i, j := 0, l, mid+1
	for i <= mid && j <= r {
		if q[i] < q[j] {
			temp[k] = q[i]
			i++
		} else {
			temp[k] = q[j]
			j++
		}
		k++
	}
	for i <= mid {
		temp[k] = q[i]
		i++
		k++
	}
	for j <= r {
		temp[k] = q[j]
		j++
		k++
	}
	for i, j = l, 0; i <= r; j++ {
		q[i] = temp[j]
		i++
	}
}

//==========================second method========================
func merge2(nums1 []int, m int, nums2 []int, n int) {
	// nums1 = append(nums1[:m] , nums2...)
	// 另外一种合并写法
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}
