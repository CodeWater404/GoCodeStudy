package _161_2180

/**
  @author: CodeWater
  @since: 2023/7/6
  @desc: 拆分成最多数目的正偶数之和
**/
func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 == 1 {
		return []int64{}
	}
	var res []int64
	for i := int64(2); i <= finalSum; i += 2 {
		res = append(res, i)
		finalSum -= i
	}
	//最后一种情况：i增长到大于finalSum的情况，退出循环时，这个时候res里面已经把i放进去了，所以这里需要再加上finalSum;如果上述遍历正好finalNum为0，那么加到最后一位上其实是不变的。
	res[len(res)-1] += finalSum
	return res
}
