package pratices

import (
	"bytes"
	"fmt"
	"sync"
)

//Prefer sync.Pool for short-lived object reuse.
//Use sync.Map instead of a mutex when high-concurrency read-heavy workloads.

func Optimizing() {
	buf := pool.Get().(*bytes.Buffer)
	buf.WriteString("Hello, Pool!")
	fmt.Println(buf.String())
	buf.Reset()
	pool.Put(buf)
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}
