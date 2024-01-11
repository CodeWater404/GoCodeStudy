package __datastructure

import "fmt"

/**
  @author: CodeWater
  @since: 2024/1/11
  @desc: 双链表
**/

const N int = 100010

var (
	m, idx  int
	e, l, r [N]int
)

// insert 在节点a的右边插入一个数x(题目在第几个节点插入是从1开始计算，但是这里的插入是从0开始)
func insert(a, x int) {
	e[idx] = x
	l[idx] = a
	r[idx] = r[a]
	l[r[idx]] = idx
	r[a] = idx
	idx++
}

// remove 删除节点a
func remove(a int) {
	r[l[a]] = r[a]
	l[r[a]] = l[a]
}

func main() {
	fmt.Scan(&m)
	// 初始化：0是左端点，1是右端点
	r[0], l[1] = 1, 0
	idx = 2

	for ; m > 0; m-- {
		var op string
		fmt.Scan(&op)
		var k, x int
		if op == "L" {
			fmt.Scan(&x)
			insert(0, x)
		} else if op == "R" {
			fmt.Scan(&x)
			insert(l[1], x)
		} else if op == "D" {
			fmt.Scan(&k)
			remove(k + 1)
		} else if op == "IL" {
			fmt.Scan(&k, &x)
			insert(l[k+1], x)
		} else {
			fmt.Scan(&k, &x)
			insert(k+1, x)
		}
	}

	for i := r[0]; i != 1; i = r[i] {
		fmt.Printf("%d ", e[i])
	}

}
