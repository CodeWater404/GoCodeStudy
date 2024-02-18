package __20

/**
  @author: CodeWater
  @since: 2024/2/18
  @desc: 4. 寻找两个正序数组的中位数
**/
// 1. 首先对于一个升序数组找到第k个元素是谁，当k=（n + m）/2时，也就是成为中间点
// 2. 然后对于第一个A和第二个B数组找到第k/2个元素。（暂时不考虑数组的奇偶个数）
// 3. 对于这两种位置的k/2，有三种情况：
//      a. A[k/2] < B[k/2]:那么A数组小于A[k/2]的前面元素一共有k/2个；因为B[k/2]>A[k/2]，所以B数组小于
//         A[k/2]的元素个数少于k/2个。所以A数组和B数组的前半段0-k/2个组合起来的k个数中小于A[k/2]的个数，
//         少于k个，那么也就是说A数组的前半段0-k/2个元素一定不是中位数，可以把A前k/2部分去掉。
//      b. A[k/2] > B[k/2]:那么B数组小于A[k/2]的前面元素一共有k/2个；因为B[k/2]<A[k/2]，所以A数组小于
//         B[k/2]的元素个数少于k/2个。所以A数组和B数组的前半段0-k/2个组合起来的k个数中小于A[k/2]的个数，
//         少于k个，那么也就是说B数组的前半段0-k/2个元素一定不是中位数，可以把B前k/2部分去掉。
//      c. A[k/2] = B[k/2]:那么A数组小于A[k/2]的前面元素一共有k/2个；因为B[k/2]=A[k/2]，所以B数组小于
//         A[k/2]的元素个数也有k/2个。所以A[k/2]和B[k/2]恰好就是第k个数，那么也就是找到答案，删除的时候，
//         删除A或者B前面一部分即可。
// 4. 每次都删除k/2个元素，递归（最多logk次，k是n+m，时间复杂度是log(n+m)，不考虑参数1/2），当k=1的时候，
//     找到解。

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	tot := len(nums1) + len(nums2)
	if tot%2 == 0 { // 两个数组所有个数和是偶数，找到左边一个数和右边一个数
		left := find(nums1, 0, nums2, 0, tot/2)
		right := find(nums1, 0, nums2, 0, tot/2+1)
		return float64(left+right) / 2.0
	} else {
		return float64(find(nums1, 0, nums2, 0, tot/2+1))
	}
}

// find 从i到len(nums1)和从j到len(nums2)这两个区间找到第k个数
func find(nums1 []int, i int, nums2 []int, j, k int) int {
	if len(nums1)-i > len(nums2)-j {
		//如果nums1的区间长度大于nums2，交换一下，也就是始终把nums1变成区间短的那个数组
		return find(nums2, j, nums1, i, k)
	}
	if k == 1 { //k等于1时，找到解，此时从nums1和nums2中取第一个数的较小一个
		if len(nums1) == i {
			return nums2[j] //特殊情况，nums1为空，只能返回第二个数组
		} else {
			return min(nums1[i], nums2[j])
		}
	}
	if len(nums1) == i { //第二个边界，nums1为空，要找的第k个数就是nums2中第k个数
		return nums2[j+k-1] //k从1开始，所以是+k-1
	}
	// 下面就是处理上面步骤3的
	si, sj := min(len(nums1), i+k/2), j+k-k/2
	if nums1[si-1] > nums2[sj-1] {
		return find(nums1, i, nums2, sj, k-(sj-j))
	} else {
		return find(nums1, si, nums2, j, k-(si-i))
	}
}
