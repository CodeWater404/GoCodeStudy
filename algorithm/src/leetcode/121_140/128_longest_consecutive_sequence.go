package _21_140

/**
  @author: CodeWater
  @since: 2024/1/3
  @desc: 128. 最长连续序列
**/

func longestConsecutive(nums []int) int {
	hash := make(map[int]bool)
	for _, v := range nums {
		hash[v] = true
	}
	res := 0
	// 把所有数插入到map中，遍历数组，随机遍历map
	for _, v := range nums {
		//当前元素存在并且前一个数不存在才能遍历
		if hash[v] == true && !hash[v-1] {
			y := v
			//在map中删除掉当前遍历的起始元素，因为对于数组中来说可能会有重复数，然后导致这部分一直
			//重复执行，实际上只需要执行一边即可
			delete(hash, v)
			for hash[y+1] {
				y++
				delete(hash, y)
			}
			res = max(res, y-v+1)
		}
	}
	return res
}
