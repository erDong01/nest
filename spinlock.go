package nest

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type spinLock uint32

func (s1 *spinLock) Lock() {
	for !atomic.CompareAndSwapUint32((*uint32)(s1), 0, 1) {
		runtime.Gosched()
	}

}

func (s1 *spinLock) Unlock() {
	atomic.StoreUint32((*uint32)(s1), 0)

}

func NewSpinLock() sync.Locker {
	return new(spinLock)
}
