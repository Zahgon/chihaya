// Package bytepool implements an optimization for byte slice allocations
// by caching those slices for later reuse.
package bytepool

import "sync"

// BytePool is a cached pool of reusable byte slices.
type BytePool struct {
	sync.Pool
}

// New allocates a new BytePool with slices of equal length and capacity.
func New(length int) *BytePool { _ = "STUB: not implemented"; return nil }

// Get returns a byte slice from the pool.
func (bp *BytePool) Get() *[]byte { _ = "STUB: not implemented"; return nil }

// Put returns a byte slice to the pool.
func (bp *BytePool) Put(b *[]byte) {
	_ = "STUB: not implemented"

	// Zero out the bytes.
	// This specific expression is optimized by the compiler:
	// https://github.com/golang/go/issues/5373.
	return
}
