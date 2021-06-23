package hive

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

func TestNewPool(t *testing.T) {
	NewPool(10000)
	defer Release()
	var runTimes int32 = 10000000
	var i int32
	var wg sync.WaitGroup

	for i = 0; i < runTimes; i++ {
		syncCalculateSum := func() {
			demoFunc(i)
			wg.Done()
		}
		wg.Add(1)
		_ = Submit(syncCalculateSum)
		//wg.Add(1)
		//go func() {
		//	demoFunc(i)
		//	wg.Done()
		//}()
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", Running())
	fmt.Printf("finish all tasks.\n")

}

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
}

func demoFunc(a int32) {
	fmt.Println(a)
}

func TestCap(t *testing.T) {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Println(n)
}
