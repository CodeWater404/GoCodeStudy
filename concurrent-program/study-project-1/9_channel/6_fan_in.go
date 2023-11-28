package main

import "reflect"

/**
  @author: CodeWater
  @since: 2023/11/28
  @desc: 任务编排——channel：扇入模式
	有多个源 Channel 输入、一个目的 Channel 输出
**/

// fanInReflect 反射实现
func fanInReflect(chans ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		// 构建select case slice
		var cases []reflect.SelectCase
		for _, c := range chans {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		// 循环，从cases中选择一个可用的
		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok { // 此channel已经close
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			out <- v.Interface()
		}
	}()
	return out
}

// fanInRec 递归+二分实现
func fanInRec(chans ...<-chan interface{}) <-chan interface{} {
	switch len(chans) {
	case 0:
		c := make(chan interface{})
		close(c)
		return c
	case 1:
		return chans[0]
	case 2:
		return mergeTwo(chans[0], chans[1])
	default:
		m := len(chans) / 2
		return mergeTwo(fanInRec(chans[:m]...), fanInRec(chans[m:]...))
	}
}

// mergeTwo 将两个 Channel 合并成一个 Channel，是扇入形式的一种特例（只处理两个 Channel）
func mergeTwo(a, b <-chan interface{}) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		for a != nil || b != nil { // 只要还有可读的chan
			select {
			case v, ok := <-a:
				if !ok { // a已关闭，设置为nil
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok { // b已关闭，设置为nil
					b = nil
					continue
				}
				c <- v
			}

		}
	}()
	return c
}
