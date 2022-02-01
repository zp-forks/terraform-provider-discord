package discord

func toUint(i int) uint {
	if i < 0 {
		return 0
	} else {
		return uint(i)
	}
}
