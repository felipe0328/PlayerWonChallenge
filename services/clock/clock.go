package clock

import "time"

type IClock interface {
	Now() time.Time
	Parse(string, string) (time.Time, error)
}

type Clock struct{}

func (c *Clock) Now() time.Time {
	return time.Now()
}

func (c *Clock) Parse(layout string, value string) (time.Time, error) {
	return time.Parse(layout, value)
}
