package _641_2660

/**
  @author: CodeWater
  @since: 2023/7/5
  @desc: K 件物品的最大和
**/
//贪心
func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	if k <= numOnes {
		return k
	} else if k <= numOnes+numZeros {
		return numOnes
	} else {
		return numOnes - (k - numOnes - numZeros)
	}
}
