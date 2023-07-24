package _60_780

/**
  @author: CodeWater
  @since: 2023/7/24
  @desc: 宝石与石头
**/
func numJewelsInStones(jewels string, stones string) int {
	hash := make(map[rune]bool)
	for _, j := range jewels {
		hash[j] = true
	}
	res := 0
	for _, s := range stones {
		if hash[s] {
			res++
		}
	}
	return res
}
