// Package clock implements a clock that handles times without dates.
// It allows adding and subtracting minutes to it.
// Two clocks that represent the same time are equals to each other.
package clock

import (
	"fmt"
)

// Clock holds hours and minutes
type Clock struct {
	hours   int
	minutes int
}

// New creates a Clock
func New(hours, minutes int) Clock {
	hours, minutes = reduce(hours, minutes)
	return Clock{hours, minutes}
}

// Add adds minutes
func (c Clock) Add(minutes int) Clock {
	return Clock{c.hours, c.minutes + minutes}
}

// Subtract subtracts minutes
func (c Clock) Subtract(minutes int) Clock {
	return Clock{c.hours, c.minutes - minutes}
}

func (c Clock) String() string {
	hours, minutes := reduce(c.hours, c.minutes)
	return fmt.Sprintf("%.2d:%.2d", hours, minutes)
}

func reduce(h, m int) (int, int) {
	minutes := m % 60
	hours := (h + m/60) % 24

	if minutes < 0 {
		hours--
		minutes = 60 + minutes
	}

	if hours < 0 {
		hours = 24 + hours
	}

	return hours, minutes
}
