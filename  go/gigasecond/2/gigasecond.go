package gigasecond

import "time"

const gigaSecond time.Duration = time.Second * 1_000 * 1_000_000

func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaSecond)
}
