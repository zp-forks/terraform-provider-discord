package discord

import (
	"fmt"
	"strings"
)

func parseTwoIds(id string) (string, string, error) {
	parts := strings.SplitN(id, ":", 2)

	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("unexpected format of ID (%s), expected attribute1:attribute2", id)
	}

	return parts[0], parts[1], nil
}

// Helper function for generating a two part ID
func generateTwoPartId(one string, two string) string {
	return fmt.Sprintf("%s:%s", one, two)
}

func parseThreeIds(id string) (string, string, string, error) {
	parts := strings.SplitN(id, ":", 3)

	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return "", "", "", fmt.Errorf("unexpected format of ID (%s), expected attribute1:attribute2:attriburte3", id)
	}

	return parts[0], parts[1], parts[2], nil
}

func generateThreePartId(one string, two string, three string) string {
	return fmt.Sprintf("%s:%s:%s", one, two, three)
}
