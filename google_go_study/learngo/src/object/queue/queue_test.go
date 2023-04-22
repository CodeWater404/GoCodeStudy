package queue

import "fmt"

/**
  @author: CodeWater
  @since: 2023/4/22
  @desc: 测试的另外一个特殊例子：example
		专门用于生成文档的
		必须要有这句，才会运行(输出例子结果的):
		//output:
		如果啥都不行运行会报错，拿到的结果和想要的结果（正确的运行示例，这时候加上去再次运行就能通过了）
**/

func ExampleQueue_Pop() {
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	//output:
	//1
	//2
	//false
	//3
	//true
}
