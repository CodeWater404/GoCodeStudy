package main

import (
	"fmt"
	"sync"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/11/20
  @desc: 死锁例子
	去派出所开证明，派出所要求物业先证明我是本物业的业主，但是，物业要我提供派出所的证明，
	才能给我开物业证明，结果就陷入了死锁状态。派出所和物业看成两个 goroutine，
	派出所证明和物业证明是两个资源，双方都持有自己的资源而要求对方的资源，而且自己的资源
	自己持有，不可剥夺。
**/

func main() {
	// 派出所证明
	var psCertificate sync.Mutex
	// 物业证明
	var properCertificate sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2) // 需要派出所和物业都处理

	// 派出所开证明
	go func() {
		defer wg.Done() // 派出所开证明完成

		psCertificate.Lock()
		defer psCertificate.Unlock()

		// 检查材料
		time.Sleep(5 * time.Second)
		// 请求物业证明
		properCertificate.Lock()
		properCertificate.Unlock()
	}()

	// 物业开证明
	go func() {
		defer wg.Done() // 物业开证明完成

		properCertificate.Lock()
		defer properCertificate.Unlock()

		// 检查材料
		time.Sleep(5 * time.Second)
		// 请求派出所证明
		psCertificate.Lock()
		psCertificate.Unlock()
	}()

	wg.Wait()
	fmt.Println("证明办理完成...")
	/*解决方案：
	可以引入一个第三方的锁，大家都依赖这个锁进行业务处理，比如现在政府推行的一站
	式政务服务中心。或者是解决持有等待问题，物业不需要看到派出所的证明才给开物业证
	明，等等
	*/
}
