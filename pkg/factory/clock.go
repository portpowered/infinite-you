package factory

import "time"

// Clock is the runtime time source used by replay-sensitive factory paths.
type Clock interface {
	Now() time.Time
}

// LogicalClock is a clock that can align itself to the current engine tick.
type LogicalClock interface {
	Clock
	SetTick(tick int)
}

// RealClock reads the host wall clock.
type RealClock struct{}

var _ Clock = RealClock{}

// Now returns the current wall-clock time.
func (RealClock) Now() time.Time {
	return time.Now()
}

// EnsureClock returns a real clock when the supplied clock is nil.
func EnsureClock(clock Clock) Clock {
	if clock == nil {
		return RealClock{}
	}
	return clock
}
