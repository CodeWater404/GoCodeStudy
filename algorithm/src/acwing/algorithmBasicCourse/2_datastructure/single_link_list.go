package __datastructure

import "fmt"

/**
  @author: CodeWater
  @since: 2023/12/15
  @desc: 单链表（数组模拟）
**/

const N int = 100010

// head 表示头结点的下标
// e[i] 表示节点i的值
// ne[i] 表示节点i的next指针是多少,也就是节点i指向哪一个节点，存的是下一个节点的下标（数组模拟链表）。
// idx 存储当前已经用到了哪个点
var (
	head, idx int
	e, ne     [N]int
)

// new 初始化
func new() {
	head, idx = -1, 0
}

// add_to_head 将x插到头节点
func add_to_head(x int) {
	e[idx] = x     //先存下节点值
	ne[idx] = head // head指向头节点的指针付给当前节点
	head = idx     // head指向当前节点
	idx++          // 新增一个点
}

// add 将x插到下标是k的节点后面
func add(k, x int) {
	e[idx] = x
	ne[idx] = ne[k]
	ne[k] = idx
	idx++
}

// remove 将下标是k的节点后面一个点删除
func remove(k int) {
	ne[k] = ne[ne[k]]
}

func main() {
	var m int
	fmt.Scan(&m)
	new()

	for ; m > 0; m-- {
		var k, x int
		var op byte
		fmt.Scanf("%c", &op)
		if op == 'H' {
			fmt.Scanf("%d", &x)
			add_to_head(x)
		} else if op == 'I' {
			fmt.Scanf("%d%d", &k, &x)
			add(k-1, x)
		} else {
			fmt.Scanf("%d", &k)
			if k == 0 {
				head = ne[head]
			} else {
				remove(k - 1)
			}
		}
	}

	for i := head; i != -1; i = ne[i] {
		fmt.Printf("%d ", e[i])
	}
	fmt.Println()
}
