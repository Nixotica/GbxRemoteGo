package transport

import "sync/atomic"

// Define initial handler value.
const initialHandler = 0x80000000

// Use atomic for safe concurrent updates.
var currentHandler uint32 = initialHandler

// GetNextHandler increments the handler and wraps if needed
func GetNextHandler() uint32 {
	// Atomically increment the handler
	next := atomic.AddUint32(&currentHandler, 1)

	// If we exceed the max, wrap to initialHandler 
	if next == 0x00000000 {
		atomic.StoreUint32(&currentHandler, initialHandler)
		return initialHandler
	}

	return next
}