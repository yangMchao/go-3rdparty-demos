package chapter_02_ants

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/panjf2000/ants/v2"
)

func add(d int) {
	sum := 0
	for i := 0; i < d; i++ {
		sum += i
	}
}

func Test_Ants(t *testing.T) {
	var wg sync.WaitGroup
	now := time.Now()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		ants.Submit(func() {
			add(10000000000)
			wg.Done()
		})
	}
	wg.Wait()
	fmt.Println(time.Since(now))
	now = time.Now()
	for i := 0; i < 5; i++ {
		add(10000000000)
	}
	fmt.Println(time.Since(now))
}

//NewPoolWithFunc

func TestNewPoolWithFunc(t *testing.T) {
	var sum int32
	var wg sync.WaitGroup

	// 创建一个函数池
	pool, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		n := i.(int32)
		atomic.AddInt32(&sum, n)
		wg.Done()
	})
	defer pool.Release()

	// 提交带不同参数的任务
	for i := int32(1); i <= 1000; i++ {
		wg.Add(1)
		_ = pool.Invoke(i)
	}

	wg.Wait()
	fmt.Printf("总和：%d\n", sum)
}

// NewPoolWithFuncGeneric
func TestNewPoolWithFuncGeneric(t *testing.T) {
	var sum int32
	var wg sync.WaitGroup

	// 创建一个泛型函数池
	pool, _ := ants.NewPoolWithFuncGeneric(10, func(i int32) {
		atomic.AddInt32(&sum, i)
		wg.Done()
	})
	defer pool.Release()

	// 提交带类型安全的任务
	for i := int32(1); i <= 1000; i++ {
		wg.Add(1)
		_ = pool.Invoke(i)
	}

	wg.Wait()
	fmt.Printf("总和：%d\n", sum)
}
