package demo

import (
	"runtime"
	"sync"
)

func LearnLock() {
	mp := make(map[int]int)
	runtime.GOMAXPROCS(4)
	mu := &sync.Mutex{}
	go func() {
		mu.Lock()
		mp[1] = 1
		mp[1] = 2
		mp[1] = 3
		mp[1] = 4
		mp[1] = 5
		mp[1] = 6
		mu.Unlock()
	}()
	go func() {
		mu.Lock()
		mp[1] = 1
		mp[1] = 2
		mp[1] = 3
		mp[1] = 4
		mp[1] = 5
		mp[1] = 6
		mu.Unlock()
	}()
	go func() {
		mu.Lock()
		mp[1] = 1
		mp[1] = 2
		mp[1] = 3
		mp[1] = 4
		mp[1] = 5
		mp[1] = 6
		mu.Unlock()
	}()
	go func() {
		mu.Lock()
		mp[1] = 1
		mp[1] = 2
		mp[1] = 3
		mp[1] = 4
		mp[1] = 5
		mp[1] = 6
		mu.Unlock()
	}()
	mu.Lock()
	print(mp[1])
	mu.Unlock()
	mu.Lock()
	print(mp[1])
	mu.Unlock()
	mu.Lock()
	print(mp[1])
	mu.Unlock()
	mu.Lock()
	print(mp[1])
	mu.Unlock()
	mu.Lock()
	print(mp[1])
	mu.Unlock()
	mu.Lock()
	print(mp[1])
	mu.Unlock()
	mu.Lock()
	print(mp[1])
	mu.Unlock()
	mu.Lock()
	print(mp[1])
	mu.Unlock()
	mu.Lock()
	print(mp[1])
	mu.Unlock()
	mu.Lock()
	print(mp[1])
	mu.Unlock()
	mu.Lock()
	print(mp[1])
	mu.Unlock()
	mu.Lock()
	print(mp[1])
	mu.Unlock()
	mu.Lock()
	print(mp[1])
	mu.Unlock()
	mu.Lock()
	print(mp[1])
	mu.Unlock()
	mu.Lock()
	print(mp[1])
	mu.Unlock()
	mu.Lock()
	print(mp[1])
	mu.Unlock()
}

