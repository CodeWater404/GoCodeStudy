package _41_860

/**
  @author: CodeWater
  @since: 2023/7/22
  @desc:  860. 柠檬水找零
**/
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, v := range bills {
		if v == 5 {
			five++
		} else if v == 10 {
			if five <= 0 {
				return false
			} else {
				five--
				ten++
			}
		} else if v == 20 {
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
