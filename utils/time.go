package utils

import "time"

// ToUTC converts a *time.Time to UTC, returning nil if the input is nil.
func ToUTC(t *time.Time) *time.Time {
	if t == nil {
		return nil
	}
	utcTime := t.UTC()
	return &utcTime
}
