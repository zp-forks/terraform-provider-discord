package discord

import "hash/crc32"

func Hashcode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

func contains[T comparable](array []T, value T) bool {
	for _, elem := range array {
		if elem == value {
			return true
		}
	}

	return false
}

// BoolPtr is a helper routine that allocates a new bool value to store v and
// returns a pointer to it.
func BoolPtr(v bool) *bool { return &v }

// Bool is a helper routine that accepts a bool pointer and returns a value
// to it.
func Bool(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

// IntPtr is a helper routine that allocates a new int value to store v and
// returns a pointer to it.
func IntPtr(v int) *int { return &v }

// Int is a helper routine that accepts a int pointer and returns a value
// to it.
func Int(v *int) int {
	if v != nil {
		return *v
	}
	return 0
}

// Int64 is a helper routine that accepts an int64 pointer and returns a
// value to it.
func Int64(v *int64) int64 {
	if v != nil {
		return *v
	}
	return 0
}

// Int64Ptr is a helper routine that allocates a new int64 value to store v
// and returns a pointer to it.
func Int64Ptr(v int64) *int64 { return &v }
