package queue

/**
  @author: CodeWater
  @since: 2023/4/10
  @desc: queue
		1. 将NewInt定义为int类型
			type NewInt int
**/

//通过别名来扩展别人已经实现的封装结构体slice（数组）
//type Queue []int //这里只支持int类型
type Queue []interface{} //interface支持任何类型

//func (q *Queue) Push(v int) { //参数限定类型
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
	//函数内部限定参数的类型.main函数中如过push别的类型进来只能在运行时才知道报错，无法在编译时就知道
	//*q = append(*q , v.(int))
}

//func (q *Queue) Pop() int {
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	//return head //1.原类型返回

	return head.(int) //2.强制把interface转换成int返回
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
