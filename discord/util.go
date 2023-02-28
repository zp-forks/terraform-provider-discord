package discord

import (
	"errors"
	"fmt"
	"hash/crc32"
	"strings"
)

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

// Helper function for generating a two part ID
func generateTwoPartId(one string, two string) string {
	return fmt.Sprintf("%s:%s", one, two)
}

// helper function for parsing a two part ID
func parseTwoPartId(id string) (string, string, error) {
	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		return "", "", errors.New(fmt.Sprintf("Unable to parse ID, length of returned value is different than 2, got %d", len(parts)))
	}

	return parts[0], parts[1], nil
}

func contains[T comparable](array []T, value T) bool {
	for _, elem := range array {
		if elem == value {
			return true
		}
	}

	return false
}
