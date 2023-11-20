package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/11/20
  @desc: 获取goroutineId
  方案一：使用runtime包中的函数
	通过 runtime.Stack 方法获取栈帧信息，栈帧信息里包含 goroutine id。常见的报错如下，可以根据这个格式获取到：
	goroutine 1 [semacquire]:
	sync.runtime_SemacquireMutex(0xc00001206c,0x55cc00,0xl)
		/usr/local/go/src/runtime/sema.go:71 +0x47
**/

// foo 和 bar 函数都会调用 locker 的 Lock 方法，foo拥有锁，然后在bar中又请求这一把锁，重入锁了。go没有提供重入锁，所以会死锁。
// 这里为了演示获取goroutine id
func foo(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	defer l.Unlock()
}

func bar(l sync.Locker) {
	fmt.Printf("===>GoID foo: %d\n", GoID())
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}

// GoID 获取goroutine id
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// 得到id字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine"))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id:%v", err))
	}
	return id
}

func main() {
	var mu sync.Mutex
	/*传递指针类型：
	在 Go 语言中，当接口类型的变量被赋值为某个具体类型的值时，接口内部会存储这个值的类型信息和值。如果赋值的是一个非接口类型的变量，
	那么接口会存储这个变量的值；如果赋值的是一个指针，那么接口会存储这个指针。在你的代码中，foo 函数的参数是 sync.Locker 类型，
	这是一个接口类型。sync.Mutex 结构体实现了这个接口，所以你可以将一个 sync.Mutex 变量或者一个 *sync.Mutex 变量赋值给 sync.Locker 类型的变量。
	然而，sync.Mutex 的 Lock 和 Unlock 方法都是指针接收者，这意味着你只能通过 *sync.Mutex 类型的变量来调用这两个方法。所以，
	如果你将一个 sync.Mutex 变量赋值给 sync.Locker 类型的变量，那么你将无法通过这个 sync.Locker 变量来调用 Lock 和 Unlock 方法。
	因此，你需要使用 & 运算符获取 mu 变量的地址，然后将这个地址赋值给 sync.Locker 类型的变量。这样，你就可以通过这个 sync.Locker
	变量来调用 Lock 和 Unlock 方法了。
	*/
	go foo(&mu)

	// 看一下主线程的goroutine id
	fmt.Printf("===>GoID main: %d\n", GoID())

	// 等待一会，让foo先执行
	time.Sleep(3 * time.Second)
}
