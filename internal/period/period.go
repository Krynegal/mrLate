package period

import "time"

type Period struct {
	Hours   uint8
	Minutes uint8
}

func (t *Period) GetDuration() time.Duration {
	return time.Duration(t.Hours)*time.Hour + time.Duration(t.Minutes)*time.Minute
}
