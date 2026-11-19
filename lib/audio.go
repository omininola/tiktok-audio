package lib

import (
	"os/exec"
	"strconv"
)

const (
	FullAudio  = "100%"
	FadedAudio = "25%"
)

func SetAudioBalance(left []int, right []int) {
	for _, idx := range left {
		RunAudioCommand(idx, FullAudio, FadedAudio)
	}
	for _, idx := range right {
		RunAudioCommand(idx, FadedAudio, FullAudio)
	}
}

func RunAudioCommand(idx int, full string, faded string) {
	exec.Command("pactl", "set-sink-input-volume", strconv.Itoa(idx), full, faded).Output()
}
