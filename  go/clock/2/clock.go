package clock

import "fmt"

type Clock int

const minutesInDay = 1440

func New(h, m int) Clock {
	minutes := (h*60 + m) % minutesInDay
	if minutes < 0 {
		minutes = minutesInDay + minutes
	}
	return Clock(minutes)
}

func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, int(c)-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}
