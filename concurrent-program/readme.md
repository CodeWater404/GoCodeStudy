# concurrent-program
`Golang`并发编程学习，案列来自网络，如有侵权，联系删除。


# Go并发技巧

## 命令集合
1. `go tool compile -S file.go`：查看汇编代码命令
2. `go run -race counter.go`：帮助我们自动发现程序有没有 `data race `（数据竞争）的问题。
3. `go vet copy.go`:vet 工具，把检查写在 Makefile 文件中，在持续集成的时候跑一跑，这样可以及时发现**死锁**问题，及时修复
4. 

## race detector
Go 提供的一个检测并发访问共享资源是否有问题的工具，它可以帮助我们自动发现程序有没有 `data race `（数据竞争）的问题。

* 使用： 
在编译（`compile`）、测试（`test`）或者运行（`run`）Go 代码的时候，加上 `race` 参数，就有可能发现并发问题（编译的时候加上race不能发现data race,但是编译的时候可以开启race参数，这样编译后的程序在运行时就可以data race问题了。但是，绝对不要把带race参数编译的程序部署到线上。）。比如`go run -race counter.go`，就会输出警告信息, 会告诉你哪个 goroutine 在哪一行对哪个变量**有写**操作，同时，哪个`goroutine`
  在哪一行对哪个变量**有读**操作，就是这些并发的读写访问，引起了 `data race`。
* 缺陷：虽然这个工具使用起来很方便，但是，因为它的实现方式，**只能通过真正对实际地址进行
  读写访问的时候才能探测，所以它并不能在编译的时候发现 data race 的问题。而且，在
  运行的时候，只有在触发了 data race 之后，才能检测到**，如果**碰巧没有触发**（比如一个
  data race 问题只能在 2 月 14 号零点或者 11 月 11 号零点才出现），**是检测不出来的**。而且，把开启了 race 的程序部署在线上，还是比较*影响性能*的。

总结一下，通过在编译的时候插入一些指令，在运行时通过这些插入的指令检测并发读写
从而发现 data race 问题，就是这个工具的实现机制。(`go tool compile -
race -S counter.go`)
> 相关案例：
> 比如 `Docker issue` 37583、35517、32826、30696等、`kubernetes issue`
72361、71617等，都是后来发现的 data race 而采用互斥锁 Mutex 进行修复的。



## CAS
CAS 指令将给定的值和一个内存地址中的值进行比较，如果它们是同一个值，就使用新值替换内存地址中的值，这个操作是原子性的。

## go:linkname
go:linkname 是 Go 语言的一个编译器指令，它可以将一个 Go 语言的标识符链接到另一个标识符。这个指令通常用于实现一些底层的功能，例如访问和修改运行时系统的内部状态。

go:linkname 指令的语法是：
```go
//go:linkname localname importpath.name
```
其中，`localname` 是本地的标识符，`importpath.name` 是要链接的标识符。这个指令会将 localname 链接到 importpath.name，这样在代码中就可以通过 localname 来访问和修改 importpath.name 的值。

需要注意的是，go:linkname 指令**会破坏 Go 语言的封装性，因为它可以访问和修改任何标识符，包括私有的标识符**。因此，除非你非常清楚你在做什么，否则不应该使用这个指令。

## Mutex
* `state`:
    ![img.png](./attach/img-3.png)
* `race.Enabled`:
    ![img.png](./attach/img-4.png)
  这行代码是用于`数据竞态检测`的。如果启用了数据竞态检测，那么就调用 `race.Acquire` 函数来记录这个互斥锁的获取操作。这个函数的参数是互斥锁的地址，转换为 unsafe.Pointer 类型。数据竞态（Race Condition）是多线程或多进程环境中的一个问题，当两个或更多的并发操作访问和修改同一块内存区域，且操作的顺序会影响结果时，就可能发生数据竞态。
    `数据竞态检测（Race Detection）`是一种技术，用于检测程序中是否存在数据竞态。在 Go 语言中，可以通过在 go 命令后添加` -race `参数来启用数据竞态检测，例如 `go run -race myprogram.go`。_启用数据竞态检测后，Go 运行时会监控内存访问，以检测并报告数据竞态。_
    在代码中，race.Enabled 是一个布尔值，表示是否启用了数据竞态检测。race.Acquire 函数用于记录一个内存访问操作，这个操作是获取互斥锁。如果存在数据竞态，那么 race.Acquire 函数会报告一个错误。
    > go语言底层库：
![img.png](./attach/img-5.png)
     Go 语言的 runtime 包是一个特殊的包，它的很多函数都是在 Go 语言的运行时系统中实现的，这些实现通常是用 C 或者汇编语言编写的，而不是 Go 语言。
    例如，`runtime_canSpin` 函数，它的声明在 runtime 包的 Go 语言源代码中，但是它的实现在运行时系统的 C 语言源代码中。
    这是因为 runtime 包的函数需要执行一些底层的操作，例如**操作系统调用、内存管理、垃圾回收、goroutine 调度**等，这些操作需要直接访问和操作硬件和操作系统，所以它们不能用 Go 语言实现，必须用更底层的语言实现。
    这也是为什么在 runtime 包的 Go 语言源代码中看不到这些函数的实现的原因。如果想看这些函数的实现，你需要查看 Go 语言的运行时系统的源代码，这些源代码通常在 Go 语言的源代码库的 `src/runtime` 目录下。

* `Unlock` 方法可以被任意的 goroutine 调用释放锁，即使是没持有这个互斥锁的 goroutine，也可以进行这个操作。这是因为，Mutex 本身并没有包含持有这把锁的goroutine 的信息，所以，Unlock 也不会对此进行检查。Mutex 的这个设计一直保至今。所以，我们在使用 Mutex 的时候，必须要保证 goroutine 尽可能不去释放自己未持有的锁，一定要遵循“谁申请，谁释放”的原则。在真实的实践中，我们使用互斥锁的时候，很少在一个方法中单独申请锁，而在另外一个方法中单独释放锁，一般都会在同一个方法中获取锁和释放锁。

### 互斥锁的几种状态
在并发编程中，互斥锁（Mutex）是一种常用的同步机制，用于保护共享资源的访问。互斥锁有几种状态，包括加锁、唤醒和饥饿：
* 加锁（Locked）：当一个线程成功获取到互斥锁时，我们说互斥锁处于加锁状态。在这种状态下，其他试图获取锁的线程将会被阻塞，
直到锁被释放
* 唤醒（Woken）：当一个等待锁的线程被唤醒（即被通知可以尝试获取锁）时，我们说互斥锁处于唤醒状态。这通常发生
在锁被释放时，系统会从等待队列中唤醒一个或多个线程来尝试获取锁。
* 饥饿（Starving）：当一个线程长时间等待获取锁，但总是被其他线程抢先获取，我们说这个线程处于饥饿状态。为了防止饥饿，一些
互斥锁的实现会提供公平锁机制，即按照线程到达的顺序分配锁，这样可以保证每个线程最终都能获取到锁。在`3_tryLock_demo.go`代码中，mutexLocked、
mutexWoken 和 mutexStarving 是用来表示互斥锁的这三种状态的常量。它们的值是通过位操作得到的，这样可以在一个 int32 变量中
同时存储多个状态。

### 常见四种错误使用场景
1. Lock/Unlock 不是成对出现，就意味着会出现死锁的情况，或者是因为 Unlock 一个未加锁的 Mutex 而导致 panic。
2. 误用是 Copy 已使用的 Mutex。Mutex 是一个有状态的对象，它的 state 字段记录这个锁的状态。如果你要复制一个已经加锁的 Mutex 给一个新的变量，那么新的刚初始化的变量居然被加锁了，这显然不符合你的期望，因为你期望的是一个零值的 Mutex。关键是在并发环境下，你根本不知道要复制的 Mutex 状态是什么，因为要复制的 Mutex 是由其它 goroutine 并发访问的，状态可能总是在变化。例子：
    ```go
    type Counter struct {
    sync.Mutex
    Count int
    }
    func main() {
    var c Counter
    c.Lock()
    defer c.Unlock()
    c.Count++
    foo(c) // 复制锁
    }
    // 这里Counter的参数是通过复制的方式传入的
    func foo(c Counter) {
    c.Lock()
    defer c.Unlock()
    fmt.Println("in foo") 
    }
    ```
   可以使用 vet 工具，把检查写在 Makefile 文件中，在持续集成的时候跑一跑，这样可以及时发现问题，及时修复死锁：`go vet copy.go`
3. 可重入锁：当一个线程获取锁时，如果没有其它线程拥有这个锁，那么，这个线程就成功获取到这个锁。之后，如果其它线程再请求这个锁，就会处于阻塞等待的状态。但是，如果拥有这把 锁的线程再请求这把锁的话，不会阻塞，而是成功返回，所以叫可重入锁（有时候也叫做 递归锁）。只要你拥有这把锁，你可以可着劲儿地调用，比如通过递归实现一些算法，调 用者不会阻塞或者死锁。**Mutex 不是可重入的锁。**因为 Mutex 的实现中没有记录哪个 goroutine 拥有这把锁。理论上，任何goroutine 都可以随意地 Unlock 这把锁，所以没办法计算重入条件。所以，一旦误用 Mutex 的重入，就会导致报错。例子：
   ```go
    func foo(l sync.Locker) {
    fmt.Println("in foo")
    l.Lock()
    bar(l)
    l.Unlock()
    }
    func bar(l sync.Locker) {
    l.Lock()
    fmt.Println("in bar")
    l.Unlock()
    }
    func main() {
    l := &sync.Mutex{}
    foo(l)
    }
   ```
   虽然标准库 Mutex 不是可重入锁，但是我们就自己实现一个。
   1. 方案一：通过 hacker 的方式获取到 goroutine id，记录下获取锁的 goroutine id，它可以实现 Locker 接口。 
   2. 方案二：调用 Lock/Unlock 方法时，由 goroutine 提供一个 token，用来标识它自 己，而不是我们通过 hacker 的方式获取到 goroutine id，但是，这样一来，就不满足Locker 接口了。
4. 死锁。死锁产生的必要条件。如果你想避免死锁，只要破坏这四个条件中的一个
   或者几个，就可以了。
   1. 互斥： 至少一个资源是被排他性独享的，其他线程必须处于等待状态，直到资源被释 放。
   2. 持有和等待：goroutine 持有一个资源，并且还在请求其它 goroutine 持有的资源，也就是咱们常说的“吃着碗里，看着锅里”的意思。
   3. 不可剥夺：资源只能由持有它的 goroutine 来释放。 
   4. 环路等待：一般来说，存在一组等待进程，P={P1，P2，…，PN}，P1 等待 P2 持有的资源，P2 等待 P3 持有的资源，依此类推，最后是 PN 等待 P1 持有的资源，这就形成了一个环路等待的死结。

### mutex总结
![Alt text](./attach/image.png)
![Alt text](./attach/image-1.png)
![Alt text](./attach/image-2.png)


## RWMutex读写锁
标准库中的 RWMutex 是一个 reader/writer 互斥锁。RWMutex 在某一时刻只能由**任意数量的 reader 持有**，或者是**只被单个的 writer 持有**。

RWMutex 的方法也很少，总共有 5 个:
1. `Lock/Unlock`：**写**操作时调用的方法。如果锁已经被 reader 或者 writer 持有，那么，
Lock 方法会一直阻塞，直到能获取到锁；Unlock 则是配对的释放锁的方法。
2. `RLock/RUnlock`：**读**操作时调用的方法。如果锁已经被 writer 持有的话，RLock 方法
会一直阻塞，直到能获取到锁，否则就直接返回；而 RUnlock 是 reader 释放锁的方
法。
3. `RLocker`：这个方法的作用是**为读操作返回一个 Locker 接口的对象**。它的 Lock 方法会
调用 RWMutex 的 RLock 方法，它的 Unlock 方法会调用 RWMutex 的 RUnlock 方
法
> RWMutex 的零值是未加锁的状态，所以，当你使用 RWMutex 的时候，无论是声明变
量，还是嵌入到其它 struct 中，**都不必显式地初始化**.
> 遇到可以明确区分 reader 和 writer goroutine 的场景，且有大量的并发读、少量
的并发写，并且有强烈的性能需求，你就可以考虑使用读写锁 RWMutex 替换 Mutex
>

### RWMutex 的实现原理
RWMutex 是很常见的并发原语，很多编程语言的库都提供了类似的并发类型。RWMutex
一般都是**基于互斥锁、条件变量（condition variables）或者信号量（semaphores）**等
并发原语来实现。Go 标准库中的 RWMutex 是基于 Mutex 实现的。
`readers-writers `问题一般有三类，基于对读和写操作的优先级，读写锁的设计和实现也分
成三类:
1. `Read-preferring`：读优先的设计可以提供很高的并发性，但是，在竞争激烈的情况下
可能会导致**写饥饿**。这是因为，如果有大量的读，这种设计会导致只有所有的读都释放
了锁之后，写才可能获取到锁。
2. `Write-preferring`：写优先的设计意味着，如果已经有一个 writer 在等待请求锁的
话，它会阻止新来的请求锁的 reader 获取到锁，所以优先保障 writer。当然，如果有
一些 reader 已经请求了锁的话，新请求的 writer 也会等待已经存在的 reader 都释放
锁之后才能获取。所以，写优先级设计中的优先权是针对新来的请求而言的。这种设计
主要避免了 writer 的饥饿问题,但是可能会导致**读饥饿**。
3. 不指定优先级：这种设计比较简单，不区分 reader 和 writer 优先级，某些场景下这种
不指定优先级的设计反而更有效，因为第一类优先级会导致写饥饿，第二类优先级可能会导致读饥饿，这种不指定优先级的访问不再区分读写，大家都是同一个优先级，解决了饥饿的问题。
> Go 标准库中的 RWMutex 设计是 Write-preferring 方案。一个正在阻塞的 Lock 调用
会排除新的 reader 请求到锁。
>
RWMutex 包含一个 Mutex，以及四个辅助字段 writerSem、readerSem、readerCount
和 readerWait：
```go
type RWMutex struct {
w Mutex // 互斥锁解决多个writer的竞争
writerSem uint32 // writer信号量
readerSem uint32 // reader信号量
readerCount int32 // reader的数量
readerWait int32 // writer等待完成的reader的数量
}
const rwmutexMaxReaders = 1 << 30
```
1. 字段 w：为 writer 的竞争锁而设计；
2. 字段 readerCount：记录当前 reader 的数量（以及是否有 writer 竞争锁）；
   * 没有 writer 竞争或持有锁时，readerCount 和我们正常理解的 reader 的计数是一样
   的；
   * 但是，如果有 writer 竞争锁或者持有锁时，那么，readerCount 不仅仅承担着 reader
   的计数功能，还能够标识当前是否有 writer 竞争或持有锁，在这种情况下，请求锁的
   reader 的处理变成阻塞等待锁的释放。
3. readerWait：记录 writer 请求锁时需要等待 read 完成的 reader 的数量；
4. writerSem 和 readerSem：都是为了阻塞设计的信号量。
5. 常量 rwmutexMaxReaders，定义了最大的 reader 数量。

#### RWMutex 的 RLock、RUlock、rUnlockSlow 方法
![img.png](./attach/img.png)
![img.png](./attach/img_1.png)
1.  `Add` 的返回值还有另外一个含义。如果它是**负值**，就表示当前**有 writer 竞争锁**，在这种情况下，还会调用 rUnlockSlow 方法，检查是不是reader 都释放读锁了，如果读锁都释放了，那么可以唤醒请求写锁的 writer 了。当一个或者多个 reader 持有锁的时候，竞争锁的 writer 会等待这些 reader 释放完，才可能持有这把锁。


#### RWMutex 的 Lock、Unlock、unlockSlow 方法
1. Lock:
![img.png](./attach/img_2.png)
RWMutex 是一个多 writer 多 reader 的读写锁，所以同时可能有多个 writer 和 reader。
那么，为了避免 writer 之间的竞争，RWMutex 就会使用一个 Mutex 来保证 writer 的互
斥。<br/>
一旦一个 writer 获得了内部的互斥锁，就会反转 readerCount 字段，把它从原来的正整
数 readerCount(>=0) 修改为负数（readerCount-rwmutexMaxReaders），让这个字段
保持两个含义（既保存了 reader 的数量，又表示当前有 writer）。<br/>
如果 readerCount 不是 0，就说明当前有持有读锁的 reader，RWMutex 需要把这个当
前 readerCount 赋值给 readerWait 字段保存下来， 同时，这个 writer 进入
阻塞等待状态。
每当一个 reader 释放读锁的时候（调用 RUnlock 方法时），readerWait 字段就减 1，直
到所有的活跃的 reader 都释放了读锁，才会唤醒这个 writer。
2. Unlock: 
![img.png](./attach/img_3.png)
当一个 writer 释放锁的时候，它会再次反转 readerCount 字段。这里的反转方法就是给它增加
rwmutexMaxReaders 这个常数值。<br/>
既然 writer 要释放锁了，那么就需要唤醒之后新来的 reader，不必再阻塞它们了，让它们
开开心心地继续执行就好了。<br/>
在 RWMutex 的 Unlock 返回之前，需要把内部的互斥锁释放。释放完毕后，其他的
writer 才可以继续竞争这把锁。


### RWMutex 3个踩坑点
1. 不可复制.
前面刚刚说过，RWMutex 是由一个互斥锁和四个辅助字段组成的。我们很容易想到，互斥锁是不可复制的，再加上四个有状态的字段，RWMutex 就更加不能复制使用了。不能复制的原因和互斥锁一样。一旦读写锁被使用，它的字段就会记录它当前的一些状态。这个时候你去复制这把锁，就会把它的状态也给复制过来。但是，原来的锁在释放的时候，并不会修改你复制出来的这个读写锁，这就会导致复制出来的读写锁的状态不对，可能永远无法释放锁.那该怎么办呢？其实，解决方案也和互斥锁一样。你可以借助 vet 工具，在变量赋值、函数传参、函数返回值、遍历数据、struct 初始化等时，检查是否有读写锁隐式复制的情景。
2. 重入导致死锁. 三种情况：
   * 因为读写锁内部基于互斥锁实现对 writer 的并发访问，而互斥锁本身是有重入问题的，所以，writer 重入调用 Lock 的时候，就会出现死锁的现象.例子：[2_reentrant_deat_lock.go](study-project-1%2F2_RWMutex%2F2_reentrant_deat_lock.go)
    * 有活跃 reader 的时候，writer 会等待，如果我们在 reader 的读操作时调用 writer 的写操作（它会调用 Lock 方法），那么，这个 reader和 writer 就会形成互相依赖的死锁状态。Reader 想等待 writer 完成后再释放锁，而writer 需要这个 reader 释放锁之后，才能不阻塞地继续执行。这是一个读写锁常见的死锁场景。 
    * 当一个 writer 请求锁的时候，如果已经有一些活跃的 reader，它会等待这些活跃的reader 完成，才有可能获取到锁，但是，如果之后活跃的 reader 再依赖新的 reader 的话，这些新的 reader 就会等待 writer 释放锁之后才能继续执行，这就形成了一个环形依赖： writer 依赖活跃的 reader -> 活跃的 reader 依赖新来的 reader -> 新来的 reader依赖 writer。例子：[3_reetrant_n_factorial.go](study-project-1%2F2_RWMutex%2F3_reetrant_n_factorial.go)
     ![img.png](./attach/img_4.png)   
   > 所以，使用读写锁最需要注意的一点就是尽量避免重入，重入带来的死锁非常隐蔽，而且
   难以诊断。
   > 
3. 释放未加锁的 RWMutex
   和互斥锁一样，Lock 和 Unlock 的调用总是成对出现的，RLock 和 RUnlock 的调用也必
   须成对出现。Lock 和 RLock 多余的调用会导致锁没有被释放，可能会出现死锁，而Unlock 和 RUnlock 多余的调用会导致 panic.

### RWMutex 总结
![img.png](./attach/img_5.png)
![img.png](./attach/img_6.png)

## WaitGroup协同等待，任务编排利器
WaitGroup 很简单，就是 package sync 用来做任务编排的一个并发原语。它要解决的就是并发 - 等待的问题： 现在有一个 goroutine A 在检查点（checkpoint）等待一组 goroutine 全部完成，如果在执行任务的这些 goroutine 还没全部完成，那么 goroutine A 就会阻塞在检查点，直到所有 goroutine 都完成后才能继续执行。

Go 标准库中的 WaitGroup 提供了三个方法:
1. `Add`，用来设置 WaitGroup 的计数值；
2. `Done`，用来将 WaitGroup 的计数值减 1，其实就是调用了 Add(-1)；
3. `Wait`，调用这个方法的 goroutine 会一直阻塞，直到 WaitGroup 的计数值变为 0

基本示例：[1_concurrent_count.go](study-project-1%2F3_WaitGroup%2F1_concurrent_count.go)

### WaitGroup 的实现原理
WaitGroup 的数据结构。它包括了一个 noCopy 的辅助字段，一个state 记录 WaitGroup 状态的数组。
1. `noCopy` 的辅助字段，主要就是辅助 vet 工具检查是否通过 copy 赋值这个 WaitGroup
实例。我会在后面和你详细分析这个字段；
2. `state`，一个具有复合意义的字段，包含 WaitGroup 的计数、阻塞在检查点的 waiter
数和信号量。<br/>
WaitGroup 的数据结构定义以及 state 信息的获取方法如下：
```go
//1.21.4 go版本已经修改了，新增sema  uint32字段
type WaitGroup struct {
// 避免复制使用的一个技巧，可以告诉vet工具违反了复制使用的规则
noCopy noCopy
// 64bit(8bytes)的值分成两段，高32bit是计数值，低32bit是waiter的计数
// 另外32bit是用作信号量的
// 因为64bit值的原子操作需要64bit对齐，但是32bit编译器不支持，所以数组中的元素在不同的
// 总之，会找到对齐的那64bit作为state，其余的32bit做信号量
state1 [3]uint32
}
```

#### Add 方法
Add 方法主要操作的是 state 的计数部分。你可以为计数值增加一个 delta 值，内部通过原子操作把这个值加到计数值上。需要注意的是，这个 delta 也可以是个负数，相当于为计数值减去一个值，Done 方法内部其实就是通过Add(-1) 实现的。
![img.png](./attach/img_7.png)

#### Done 方法
Done 方法其实就是调用了 Add(-1)。它会把计数值减 1，如果计数值变为 0，就会唤醒所有阻塞在 Wait 方法上的 goroutine。
```go
func (wg *WaitGroup) Done() {
	wg.Add(-1)
}
```

#### Wait 方法
Wait 方法的实现逻辑是：不断检查 state 的值。如果其中的计数值变为了 0，那么说明所有的任务已完成，调用者不必再等待，直接返回。如果计数值大于 0，说明此时还有任务没完成，那么调用者就变成了等待者，需要加入 waiter 队列，并且阻塞住自己。
![img.png](./attach/img_8.png)


### WaitGroup 的常见错误

#### 计数器设置为负值
WaitGroup 的计数器的值必须大于等于 0。我们在更改这个计数值的时候，WaitGroup 会先做检查，如果计数值被设置为负数，就会导致 panic。一般情况下，有两种方法会导致计数器设置为负数：
1. 调用 Add 的时候传递一个负数。如果你能保证当前的计数器加上这个负数后还是大于等于 0 的话，也没有问题，否则就会导致 panic。
    ```go
    func main() {
    var wg sync.WaitGroup
    wg.Add(10)
    wg.Add(-10)//将-10作为参数调用Add，计数值被设置为0
    wg.Add(-1)//将-1作为参数调用Add，如果加上-1计数值就会变为负数。这是不对的，所以会触发panic
    }
    ```
2. 调用 Done 方法的次数过多，超过了 WaitGroup 的计数值。<br/>
   使用 WaitGroup 的正确姿势是，预先确定好 WaitGroup 的计数值，然后调用相同次数的 Done 完成相应的任务。比如，在 WaitGroup 变量声明之后，就立即设置它的计数值，或者在 goroutine 启动之前增加 1，然后在 goroutine 中调用 Done。
    ```go
    func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    wg.Done()
    wg.Done()
    }
    ```
   
#### 不期望的 Add 时机
在使用 WaitGroup 的时候，你一定要遵循的原则就是，等所有的 Add 方法调用之后再调用 Wait，否则就可能导致 panic 或者不期望的结果。
```go
func main() {
var wg sync.WaitGroup
go dosomething(100, &wg) // 启动第一个goroutine
go dosomething(110, &wg) // 启动第二个goroutine
go dosomething(120, &wg) // 启动第三个goroutine
go dosomething(130, &wg) // 启动第四个goroutine
wg.Wait() // 主goroutine等待完成
fmt.Println("Done")
}
func dosomething(millisecs time.Duration, wg *sync.WaitGroup) {
duration := millisecs * time.Millisecond
time.Sleep(duration) // 故意sleep一段时间
wg.Add(1)
fmt.Println("后台执行, duration:", duration)
wg.Done()
}
```
在这个例子中，我们原本设想的是，等四个 goroutine 都执行完毕后输出 Done 的信息，
但是它的错误之处在于，**将 WaitGroup.Add 方法的调用放在了子 gorotuine 中**。等主
goorutine 调用 Wait 的时候，因为四个任务 goroutine 一开始都休眠，所以*可能
WaitGroup 的 Add 方法还没有被调用，WaitGroup 的计数还是 0，所以它并没有等待四
个子 goroutine 执行完毕才继续执行，而是立刻执行了下一步。*
导致这个错误的原因是，没有遵循先完成所有的 Add 之后才 Wait。要解决这个问题，一
个方法是，预先设置计数值,在启动协程之前设置。<br/>
另一种修复是在启动子 goroutine 之前才调用 Add：
```go
func main() {
var wg sync.WaitGroup
dosomething(100, &wg) // 调用方法，把计数值加1，并启动任务goroutine
dosomething(110, &wg) // 调用方法，把计数值加1，并启动任务goroutine
dosomething(120, &wg) // 调用方法，把计数值加1，并启动任务goroutine
dosomething(130, &wg) // 调用方法，把计数值加1，并启动任务goroutine
wg.Wait() // 主goroutine等待，代码逻辑保证了四次Add(1)都已经执行完了
fmt.Println("Done")
}
func dosomething(millisecs time.Duration, wg *sync.WaitGroup) {
wg.Add(1) // 计数值加1，再启动goroutine
go func() {
duration := millisecs * time.Millisecond
time.Sleep(duration)
fmt.Println("后台执行, duration:", duration)
wg.Done()
}()
}
```
> 无论是怎么修复，都要保证所有的 Add 方法是在 Wait 方法之前被调用的。
 

#### 前一个 Wait 还没结束就重用 WaitGroup
“前一个 Wait 还没结束就重用 WaitGroup”这一点似乎不太好理解，我借用田径比赛的
例子和你解释下吧。在田径比赛的百米小组赛中，需要把选手分成几组，一组选手比赛完
之后，就可以进行下一组了。为了确保两组比赛时间上没有冲突，我们在模型化这个场景
的时候，可以使用 WaitGroup。<br/>
WaitGroup 等一组比赛的所有选手都跑完后 5 分钟，才开始下一组比赛。下一组比赛还可
以使用这个 WaitGroup 来控制，因为 WaitGroup 是可以重用的。只要 WaitGroup 的计
数值恢复到零值的状态，那么它就可以被看作是新创建的 WaitGroup，被重复使用.看一个例子，初始设置 WaitGroup 的计数值为 1，启动一个 goroutine 先调 用 Done 方法，接着就调用 Add 方法，Add 方法有可能和主 goroutine 并发执行。
```go
func main() {
var wg sync.WaitGroup
wg.Add(1)
go func() {
time.Sleep(time.Millisecond)
wg.Done() // 计数器减1
wg.Add(1) // 计数值加1
}()
wg.Wait() // 主goroutine等待，有可能在done结束之后和add并发执行。
}
```
在这个例子中，done虽然让 WaitGroup 的计数恢复到 0，但是主goroutine有个 waiter在等待，如果等待 Wait 的主goroutine，刚被唤醒就和 Add 调用的子goroutine有并发执行的冲突，所以就会出现 panic。
> WaitGroup 虽然可以重用，但是是有一个前提的，那就是必须等到上一轮的
Wait 完成之后，才能重用 WaitGroup 执行下一轮的 Add/Wait，如果你在 Wait 还没执
行完的时候就调用下一轮 Add 方法，就有可能出现 panic。


### noCopy：辅助 vet 检查
我们刚刚在学习 WaitGroup 的数据结构时，提到了里面有一个 noCopy 字段。你还记得它的作用吗？其实，它就是指示 vet 工具在做检查的时候，这个数据结构不能做值复制使用。更严谨地说，是不能在第一次使用之后复制使用 ( must not be copied after first use).<br/>
vet 会对实现 Locker 接口的数据类型做静态检查，一旦代码中有复制使用这种数据类型的情况，就会发出警告。
通过给 WaitGroup 添加一个 noCopy 字段，我们就可以为 WaitGroup
实现 Locker 接口，这样 vet 工具就可以做复制检查了。而且因为 noCopy 字段是未输出
类型，所以 WaitGroup 不会暴露 Lock/Unlock 方法。
> 如果你想要自己定义的数据结构不被复制使用，或者说，不能通过 vet 工具检查出复制使
用的报警，就可以通过嵌入 noCopy 这个数据类型来实现。

例子： 
```go
type TestStruct struct {
Wait sync.WaitGroup
}
func main() {
w := sync.WaitGroup{}
w.Add(1)
t := &TestStruct{
Wait: w,
}
t.Wait.Done()
fmt.Println("Finished")
}
```
这段代码最大的一个问题，就是第 9 行 copy 了 WaitGroup 的实例 w。虽然这段代码能执行成功，但确实是违反了 WaitGroup 使用之后不要复制的规则。在项目中，我们可以通过 vet 工具检查出这样的错误。

### 如何避免 WaitGroup 的常见错误
只需要尽量保证下面 5 点就可以了：
1. 不重用 WaitGroup。新建一个 WaitGroup 不会带来多大的资源开销，重用反而更容易出错。
2. 保证所有的 Add 方法调用都在 Wait 之前。
3. 不传递负数给 Add 方法，只通过 Done 来给计数值减 1。
4. 不做多余的 Done 方法调用，保证 Add 的计数值和 Done 方法调用的数量是一样的。
5. 不遗漏 Done 方法的调用，否则会导致 Wait hang 住无法返回。

### WaitGroup 总结
![img.png](./attach/img_9.png)
![img.png](./attach/img_10.png)

## Cond:条件变量
Go 标准库提供 Cond 原语的目的是，为等待 / 通知场景下的并发问题提供支持。Cond 通
常应用于等待某个条件的一组 goroutine，等条件变为 true 的时候，其中一个 goroutine 或者所有的 goroutine 都会被唤醒执行。<br/>
顾名思义，Cond 是和某个条件相关，这个条件需要一组 goroutine 协作共同完成，在条 件还没有满足的时候，所有等待这个条件的 goroutine 都会被阻塞住，只有这一组goroutine 通过协作达到了这个条件，等待的 goroutine 才可能继续进行下去。<br/>
那这里等待的条件是什么呢？等待的条件，可以是某个变量达到了某个阈值或者某个时间点，也可以是一组变量分别都达到了某个阈值，还可以是某个对象的状态满足了特定的条 件。总结来讲，等待的条件是一种可以用来计算结果是 true 还是 false 的条件.<br/>
使用 Cond 的场景比较少，因为一旦遇到需要使用 Cond 的场景，我们更多地会使用 Channel 的方式去实现，因为那才是更地道的 Go 语言的写法。

### Cond 的基本用法
标准库中的 Cond 并发原语初始化的时候，需要关联一个 Locker 接口的实例，一般我们 使用 Mutex 或者 RWMutex。
首先，Cond 关联的 Locker 实例可以通过 c.L 访问，它内部维护着一个先入先出的等待队
列。 然后，我们分别看下它的三个方法 Broadcast、Signal 和 Wait 方法。
1. `Signal` 方法，允许调用者 Caller 唤醒一个等待此 Cond 的 goroutine。如果此时没有等待的 goroutine，显然无需通知 waiter；如果 Cond 等待队列中有一个或者多个等待的 goroutine，则需要从等待队列中移除第一个 goroutine 并把它唤醒。在其他编程语言中，比如 Java 语言中，Signal 方法也被叫做 notify 方法。调用 Signal 方法时，不强求你一定要持有 c.L 的锁。
2. `Broadcast` 方法，允许调用者 Caller 唤醒所有等待此 Cond 的 goroutine。如果此时没有等待的 goroutine，显然无需通知 waiter；如果 Cond 等待队列中有一个或者多个等待的 goroutine，则清空所有等待的 goroutine，并全部唤醒。在其他编程语言中，比如 Java 语言中，Broadcast 方法也被叫做 notifyAll 方法。同样地，调用 Broadcast 方法时，也不强求你一定持有 c.L 的锁。
3. `Wait` 方法，会把调用者 Caller 放入 Cond 的等待队列中并阻塞，直到被 Signal 或者Broadcast 的方法从等待队列中移除并唤醒。调用 Wait 方法时**必须要持有** c.L 的锁。<br/>

例子：[1_cond_example.go](study-project-1%2F4_Cond%2F1_cond_example.go)

### Cond 的实现原理
其实，Cond 的实现非常简单，或者说复杂的逻辑已经被 Locker 或者 runtime 的等待队
列实现了。
```go
type Cond struct {
    noCopy noCopy
    // 当观察或者修改等待条件的时候需要加锁
    L Locker
    // 等待队列
    notify notifyList
    checker copyChecker
}
func NewCond(l Locker) *Cond {
    return &Cond{L: l}
}

func (c *Cond) Wait() {
    c.checker.check()
    // 增加到等待队列中
    t := runtime_notifyListAdd(&c.notify)
    c.L.Unlock()
    // 阻塞休眠直到被唤醒
    runtime_notifyListWait(&c.notify, t)
    c.L.Lock()
}
func (c *Cond) Signal() {
    c.checker.check()
    runtime_notifyListNotifyOne(&c.notify)
}
func (c *Cond) Broadcast() {
    c.checker.check()
    runtime_notifyListNotifyAll(&c.notify)
}
```
runtime_notifyListXXX 是运行时实现的方法，实现了一个等待 / 通知的队列。

copyChecker 是一个辅助结构，可以在运行时检查 Cond 是否被复制使用。

Signal 和 Broadcast 只涉及到 notifyList 数据结构，不涉及到锁。

Wait 把调用者加入到等待队列时会释放锁，在被唤醒之后还会请求锁。在阻塞休眠期间，
调用者是不持有锁的，这样能让其他 goroutine 有机会检查或者更新等待变量。

### Cond 的常见错误
1. 是调用 Wait 的时候没有加锁.
    ```go
    func main() {
        c := sync.NewCond(&sync.Mutex{})
        var ready int
        for i := 0; i < 10; i++ {
            go func(i int) {
                time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
                // 加锁更改等待条件
                c.L.Lock()
                ready++
                c.L.Unlock()
                log.Printf("运动员#%d 已准备就绪\n", i)
                // 广播唤醒所有的等待者
                c.Broadcast()
            }(i)
        }
        // c.L.Lock() //没加上锁
        for ready != 10 {
            c.Wait()
            log.Println("裁判员被唤醒一次")
        }
        // c.L.Unlock()
        //所有的运动员是否就绪
        log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")
    }
    ```
   出现这个问题的原因在于，cond.Wait 方法的实现是，把当前调用者加入到 notify 队列之中后会释放锁（如果不释放锁，其他 Wait 的调用者就没有机会加入到 notify 队列中了），然后一直等待；等调用者被唤醒之后，又会去争抢这把锁。如果调用 Wait 之前不加锁的话，就**有可能 Unlock 一个未加锁的 Locker**。所以切记，调用 cond.Wait 方法之前一定要加锁。
2. 只调用了一次 Wait，没有检查等待条件是否满足，结果条件没满足，程序就继续执行了。出现这个问题的原因在于，误以为 Cond 的使用，就像WaitGroup 那样调用一下 Wait 方法等待那么简单。比如：
    ```go
    func main() {
         c := sync.NewCond(&sync.Mutex{})
         var ready int
		 for i := 0; i < 10; i++ {
            go func(i int) {
                time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
                // 加锁更改等待条件
                c.L.Lock()
                ready++
                c.L.Unlock()
                log.Printf("运动员#%d 已准备就绪\n", i)
                // 广播唤醒所有的等待者
                c.Broadcast()
            }(i)
         }
         c.L.Lock()
         // for ready != 10 {
         c.Wait()
         log.Println("裁判员被唤醒一次")
         // }
         c.L.Unlock()
         //所有的运动员是否就绪
         log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")
    }
    ``` 
      > 一定要记住，waiter goroutine 被唤醒不等于等待条件被满足，只是有
      goroutine 把它唤醒了而已，等待条件有可能已经满足了，也有可能不满足，我们需要进
      一步检查。你也可以理解为，等待者被唤醒，只是得到了一次检查的机会而已。
  
### Cond 总结
如果你想在使用 Cond 的时候避免犯错，只要时刻记住调用 **cond.Wait 方法之前一定要加锁**，以及 **waiter goroutine 被唤醒不等于等待条件被满足**这两个知识点。<br/>
Cond 有三点特性是 Channel 无法替代的：
1. Cond 和一个 Locker 关联，可以利用这个 Locker 对相关的依赖条件更改提供保护。
2. Cond 可以同时支持 Signal 和 Broadcast 方法，而 Channel 只能同时支持其中一种。
3. Cond 的 Broadcast 方法可以被重复调用。等待条件再次变成不满足的状态后，我们又可以调用 Broadcast 再次唤醒等待的 goroutine。这也是 Channel 不能支持的，
4. Channel 被 close 掉了之后不支持再 open。

WaitGroup 和 Cond 是有本质上的区别的：
* WaitGroup 是主 goroutine 等待确定数量的子goroutine 完成任务；而 Cond 是等待某个条件满足，这个条件的修改可以被任意多的goroutine 更新，
* 而且 Cond 的 Wait 不关心也不知道其他 goroutine 的数量，只关心等待条件。而且 Cond 还有单个通知的机制，也就是 Signal 方法。
![img.png](./attach/img_11.png)