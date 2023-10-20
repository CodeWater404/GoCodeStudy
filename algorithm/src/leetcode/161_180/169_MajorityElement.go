package _61_180

/**
  @author: CodeWater
  @since: 2023/10/20
  @desc:多数元素
**/

func majorityElement(nums []int) int {
	// r表示当前存的数；c表示当前数的数量（多少个）
	r, c := 0, 0
	for _, v := range nums {
		/*思路：
		  1. 当当前存的数r和当前遍历到的数v相同的时候，c数量+1
		  2. 当当前存的数r和当前遍历到的数v不同的时候，c数量-1
		  3. 当c存的数量等于0，说明r当前没有数存着，把当前遍历到的数v当作新存下的数
		  遍历完之后，r就是最后的多数元素。
		  正确性：如果存在一个数出现次数大于n/2的元素，那么在遍历过程中用他和另外的元素做抵消，最后一定会存下。
		*/
		if c == 0 {
			r = v
			c = 1
		} else if r == v {
			c++
		} else {
			c--
		}
	}
	return r
}
