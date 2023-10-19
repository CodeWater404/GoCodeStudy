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

//=======================Third method=========================
//有序，直接从尾部扫描往nums1比较添加。
//三个变量标记两个数组以及合并数组的尾部，最后只需要判断nums2是否还有剩余，因为nums1还有剩余也是在结果数组中。
func merge3(nums1 []int, m int, nums2 []int, n int) {
	k, i, j := m+n-1, m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			k--
			i--
		} else {
			nums1[k] = nums2[j]
			k--
			j--
		}
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}

}
