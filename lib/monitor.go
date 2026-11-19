package lib

import "tiktok-audio/types"

func GetLeftMonitorID(monitors []types.Monitor) int {
	leftID := -1
	rightID := -1

	leftX := 0

	for _, monitor := range monitors {
		if leftID == -1 && rightID == -1 {
			leftID = monitor.ID
			rightID = monitor.ID

			leftX = monitor.X
		}

		if monitor.X < leftX {
			leftID = monitor.ID
			leftX = monitor.X
		}
	}

	return leftID
}
