package lib

import "math/rand/v2"

func GetRandomValue(min int, max int) int {
	return rand.IntN(max-min+1) + min
}
