package _81_200

/**
  @author: CodeWater
  @since: 2024/2/22
  @desc: 190. 颠倒二进制位
**/

func reverseBits(n uint32) uint32 {
	res := uint32(0)
	for i := 0; i < 32; i++ {
		// 依次取得低位，然后加上这一位。
		res = res*2 + (n >> i & 1)
	}
	return res
}
