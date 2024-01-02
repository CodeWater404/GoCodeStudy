package _01_220

/*
*

	@author: CodeWater
	@since: 2024/1/2
	@desc:  202. 快乐数

*
*/

func isHappy(n int) bool {
	//快慢指针，快的走两步（两次平方和），慢的走一步
	fast, slow := get(n), n
	for fast != slow {
		fast = get(get(fast))
		slow = get(slow)
	}
	// 有两种情况：
	// 是快乐数：最后两个指针都会变成1，重合在一起
	// 不是快乐数：相当于一个环，快慢指针总会在某个数处相遇
	// 即：不管是不是快乐数，快慢指针最终都会汇聚到一个数上，我们只需要判断这个数是不是1即可
	return fast == 1
}

// get x每一位上的数平方和
func get(x int) (res int) {
	res = 0
	//每一位的变化范围都是0-810之间（9*9*10位，n最大小于9^10，也就是每一位都是9）
	for x > 0 {
		res += (x % 10) * (x % 10)
		x /= 10
	}
	return
}
