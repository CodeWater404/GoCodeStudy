package _1_80

/**
  @author: CodeWater
  @since: 2024/2/28
  @desc: $
**/

func plusOne(digits []int) []int {
	reverse(digits) // 把数组最末尾的个位放到数组开头
	t := 1          // 加1操作
	for i, x := range digits {
		// 这里面对数组每个元素进行重新更新
		t += x
		//当前第i为就是t%10的结果
		digits[i] = t % 10
		//更新t
		t /= 10
	}
	//t还不等于0，说明有进位，直接添加到数组末尾
	if t != 0 {
		digits = append(digits, t)
	}
	//再反转一遍，高位放到数组开头
	reverse(digits)
	return digits
}

func reverse(digits []int) {
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
}
