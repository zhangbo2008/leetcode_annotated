package main

import (
	"fmt"
	"sync"
	"time"
)





var Password = secret{password: "myPassword"}

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

// 通过rwmutex写
func Change(c *secret, pass string) {
	c.RWM.Lock() //写锁的写法就是直接写lock.
	fmt.Println("Change with rwmutex lock")
	time.Sleep(3 * time.Second)
	c.password = pass
	c.RWM.Unlock()
}


func Change2(c *secret, pass string) {
	c.M.Lock()
	fmt.Println("Change with mutex lock")
	time.Sleep(3 * time.Second)
	c.password = pass
	c.M.Unlock()
}

// 通过rwmutex读
func rwMutexShow(c *secret) string {
	c.RWM.RLock()
	fmt.Println("show with rwmutex",time.Now().Second())
	time.Sleep(1 * time.Second)
	defer c.RWM.RUnlock()
	return c.password
}

// 通过mutex读，和rwMutexShow的唯一区别在于锁的方式不同
func mutexShow(c *secret) string {
	c.M.Lock()
	fmt.Println("show with mutex:",time.Now().Second())
	time.Sleep(1 * time.Second)
	defer c.M.Unlock()
	return c.password
}

func main() {
	// 定义一个稍后用于覆盖(重写)的函数
	var show = func(c *secret) string { return "" }

	// 通过变量赋值的方式，选择并重写showFunc函数


	//下面这2行会发现读写锁会快很多.
	//show = rwMutexShow    //因为读写锁里面读的锁会并发
	show = mutexShow


	var wg sync.WaitGroup

	// 激活5个goroutine，每个goroutine都查看
	// 根据选择的函数不同，showFunc()加锁的方式不同
	//首先我们读5遍
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Go Pass:", show(&Password),time.Now().Second())
		}()
	}

	// 激活一个申请写锁的goroutine

	//然后我们写一遍.测速度
	go func() {
		wg.Add(1)
		defer wg.Done()
		Change(&Password, "123456")
	}()


	go func() {
		wg.Add(1)
		defer wg.Done()
		Change2(&Password, "123456")
	}()
	// 阻塞，直到所有wg.Done
	wg.Wait()
}
