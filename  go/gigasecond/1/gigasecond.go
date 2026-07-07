// Package gigasecond converts given time to gigasecond later
package gigasecond

import "time"

const gigaSecond time.Duration = time.Second * 1_000 * 1_000_000

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaSecond)
}
