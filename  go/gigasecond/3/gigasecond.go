package gigasecond

import "time"

const gigaSecond time.Duration = time.Second * 1e9

func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaSecond)
}
