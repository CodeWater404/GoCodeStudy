package _61_380

import "math/rand"

/*
*

	@author: CodeWater
	@since: 2023/12/23
	@desc: O(1) 时间插入、删除和获取随机元素

*
*/
type RandomizedSet struct {
	hash map[int]int // hash维护数组元素值k和下标v
	nums []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		hash: make(map[int]int),
		nums: make([]int, 0, 200010),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hash[val]; !ok {
		this.nums = append(this.nums, val)
		this.hash[val] = len(this.nums) - 1
		return true
	}
	return false
}

// Remove 删除O（1）实现：把删除的元素交换到最后，然后删除最后一个元素即可。
func (this *RandomizedSet) Remove(val int) bool {
	_, ok := this.hash[val]
	if !ok {
		return false
	}
	y := this.nums[len(this.nums)-1]
	px, py := this.hash[val], this.hash[y]
	this.nums[px], this.nums[py] = y, val
	this.hash[val], this.hash[y] = py, px // 对应的下标也要交换

	delete(this.hash, val)
	this.nums = this.nums[:len(this.nums)-1]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
