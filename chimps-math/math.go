package chimps_math

import "time"

// Min Min
func Min(a, b int) (c int) {
	if a >= b {
		return b
	}
	return a
}

// Max Max
func Max(a, b int) (c int) {
	if a > b {
		return a
	}
	return b
}

// Min Min
func MinDuration(a, b time.Duration) (c time.Duration) {
	if a >= b {
		return b
	}
	return a
}

// Max Max
func MaxDuration(a, b time.Duration) (c time.Duration) {
	if a > b {
		return a
	}
	return b
}

// MinInt64 MinInt64
func MinInt64(a, b int64) (c int64) {
	if a >= b {
		return b
	}
	return a
}

// MaxInt64 MaxInt64
func MaxInt64(a, b int64) (c int64) {
	if a > b {
		return a
	}
	return b
}