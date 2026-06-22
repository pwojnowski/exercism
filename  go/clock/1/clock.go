package clock

import "fmt"

type Clock struct {
	minutes int
}

func New(h, m int) Clock {
	h += m / 60
	h %= 24
	if h < 0 {
		h += 24
	}
	m %= 60
	return Clock{h*60 + m}
}

const minutesInDay = 1440

func (c Clock) Add(m int) Clock {
	m %= minutesInDay
	m = c.minutes + m
	if m < 0 {
		m = minutesInDay + m
	}
	m %= minutesInDay
	return Clock{m}
}

func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

func (c Clock) String() string {
	hours := (c.minutes / 60) % 24
	minutes := c.minutes % 60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
