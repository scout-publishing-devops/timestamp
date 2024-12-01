package timestamp

import (
	"errors"
	"time"
)

// ToEpoch converts a given time.Time to an epoch timestamp in seconds.
func ToEpoch(t time.Time) int64 {
	return t.Unix()
}

// FromEpoch converts an epoch timestamp in seconds to a time.Time.
func FromEpoch(epoch int64) time.Time {
	return time.Unix(epoch, 0)
}

// EpochInDays calculates the epoch timestamp for the current time plus the given number of days.
func EpochInDays(days int) int64 {
	now := time.Now()
	futureTime := now.AddDate(0, 0, days)
	return ToEpoch(futureTime)
}

// CompareEpoch compares two epoch timestamps based on the operator.
// Operators supported: "<", ">", "<=", ">=".
func CompareEpoch(epoch1, epoch2 int64, operator string) bool {
	switch operator {
	case "<":
		return epoch1 < epoch2
	case ">":
		return epoch1 > epoch2
	case "<=":
		return epoch1 <= epoch2
	case ">=":
		return epoch1 >= epoch2
	default:
		return false
	}
}

// TimeRange checks if a given epoch timestamp is within a range [start, end].
func TimeRange(epoch, start, end int64) bool {
	return epoch >= start && epoch <= end
}

// EpochForDate calculates the epoch timestamp for a specific date and time.
// Expects date as "YYYY-MM-DD" and time as "HH:MM:SS" in 24-hour format.
// Returns epoch timestamp or an error if the format is incorrect.
func EpochForDate(date, timeStr string) (int64, error) {
	layout := "2006-01-02 15:04:05"
	dateTime := date + " " + timeStr
	parsedTime, err := time.Parse(layout, dateTime)
	if err != nil {
		return 0, errors.New("invalid date or time format, use 'YYYY-MM-DD' and 'HH:MM:SS'")
	}
	return ToEpoch(parsedTime), nil
}
