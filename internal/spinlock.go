package internal

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type spinLock uint32

const maxBackoff = 64

func (this *spinLock) Lock() {
	backoff := 1
	for !atomic.CompareAndSwapUint32((*uint32)(this), 0, 1) {
		for i := 0; i < maxBackoff; i++ {
			runtime.Gosched()
		}
		if backoff < maxBackoff {
			backoff <<= 1
		}
	}
}

func (this *spinLock) Unlock() {
	atomic.StoreUint32((*uint32)(this), 0)
}

func NewSpinLock() sync.Locker {
	return new(spinLock)
}
