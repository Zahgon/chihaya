// Package stop implements a pattern for shutting down a group of processes.
package stop

import (
	"sync"
)

// Channel is used to return zero or more errors asynchronously. Call Done()
// once to pass errors to the Channel.
type Channel chan []error

// Result is a receive-only version of Channel. Call Wait() once to receive any
// returned errors.
type Result <-chan []error

// Done adds zero or more errors to the Channel and closes it, indicating the
// caller has finished stopping. It should be called exactly once.
func (ch Channel) Done(errs ...error) { _ = "STUB: not implemented"; return }

// Result converts a Channel to a Result.
func (ch Channel) Result() <-chan []error {
	_ = "STUB: not implemented"

	// Wait blocks until Done() is called on the underlying Channel and returns any
	// errors. It should be called exactly once.
	return nil
}

func (r Result) Wait() []error {
	_ = "STUB: not implemented"

	// AlreadyStopped is a closed error channel to be used by Funcs when
	// an element was already stopped.
	return nil
}

var AlreadyStopped Result

// AlreadyStoppedFunc is a Func that returns AlreadyStopped.
var AlreadyStoppedFunc = func() Result { return AlreadyStopped }

func init() {
	closeMe := make(Channel)
	close(closeMe)
	AlreadyStopped = closeMe.Result()
}

// Stopper is an interface that allows a clean shutdown.
type Stopper interface {
	// Stop returns a channel that indicates whether the stop was
	// successful.
	//
	// The channel can either return one error or be closed.
	// Closing the channel signals a clean shutdown.
	// Stop() should return immediately and perform the actual shutdown in a
	// separate goroutine.
	Stop() Result
}

// Func is a function that can be used to provide a clean shutdown.
type Func func() Result

// Group is a collection of Stoppers that can be stopped all at once.
type Group struct {
	stoppables []Func
	sync.Mutex
}

// NewGroup allocates a new Group.
func NewGroup() *Group { _ = "STUB: not implemented"; return nil }

// Add appends a Stopper to the Group.
func (cg *Group) Add(toAdd Stopper) { _ = "STUB: not implemented"; return }

// AddFunc appends a Func to the Group.
func (cg *Group) AddFunc(toAddFunc Func) { _ = "STUB: not implemented"; return }

// Stop stops all members of the Group.
//
// Stopping will be done in a concurrent fashion.
// The slice of errors returned contains all errors returned by stopping the
// members.
func (cg *Group) Stop() Result { _ = "STUB: not implemented"; return *new(Result) }
