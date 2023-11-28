package main

/**
  @author: CodeWater
  @since: 2023/11/28
  @desc: 任务编排——channel：扇出模式
	有一个源 Channel 输入、多个目的 Channel 输出
**/

func fanOut(ch <-chan interface{}, out []chan interface{}, async bool) {
	go func() {
		defer func() { // 退出时关闭所有的输出chan
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		for v := range ch { // 从输入chan中读取数据
			v := v
			for i := 0; i < len(out); i++ {
				i := i
				if async {
					go func() {
						out[i] <- v // 放入到输出chan中，异步方式
					}()
				} else {
					out[i] <- v // 放入到输出chan中，同步方式
				}
			}
		}
	}()

}
