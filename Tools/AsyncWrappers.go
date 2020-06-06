package Tools

import "sync"

// PerformWithMutex will lock mutex, then perform `f` function and when it ends - it will unlock mutex
// You should call this method with `go`
func PerformWithMutex(mx *sync.Mutex, f func()) {
	mx.Lock()
	f()
	mx.Unlock()
}

// PerformWithChan execute `f` and then sends result to channel
func PerformWithChan(ch chan interface{}, f func() interface{}) {
	ch <- f()
}