package concurrency

import "sync"

//When to Use Which? Mutex VS RWMutex
//Scenario	Use 							|sync.Mutex		| Use sync.RWMutex
//Read-heavy (e.g., caching, config stores)	| ❌ Slow		| ✅ Faster (many readers at once)
//Write-heavy (e.g., logging, counters) 	| ✅ Better		| ❌ No significant benefit
//Equal read/write operations				| ✅ Good		| ✅ Also works, but minor benefit

// Mutex
type SafeMap struct {
	mu   sync.Mutex
	data map[string]string
}

func (m *SafeMap) Get(key string) string {
	m.mu.Lock() // Blocks both readers and writers
	defer m.mu.Unlock()
	return m.data[key]
}

func (m *SafeMap) Set(key, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

// RWMutex
type SafeMapRW struct {
	mu   sync.RWMutex
	data map[string]string
}

// Read with RLock (multiple readers allowed)
func (m *SafeMapRW) Get(key string) string {
	m.mu.RLock() // Multiple readers can hold this lock at the same time
	defer m.mu.RUnlock()
	return m.data[key]
}

// Write with Lock (only one writer allowed at a time)
func (m *SafeMapRW) Set(key, value string) {
	m.mu.Lock() // Writers block both readers and other writers
	defer m.mu.Unlock()
	m.data[key] = value
}
