package queue

/**
  @author: CodeWater
  @since: 2023/4/10
  @desc: queue
		1. 将NewInt定义为int类型
			type NewInt int
**/

//通过别名来扩展别人已经实现的封装结构体slice（数组）
type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
