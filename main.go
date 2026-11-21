package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strconv"

	"tiktok-audio/lib"
	"tiktok-audio/types"
)

var CommandList = map[string][]string{
	"monitor": {"hyprctl", "monitors", "-j"},
	"client":  {"hyprctl", "clients", "-j"},
	"audio":   {"pactl", "-f", "json", "list", "sink-inputs"},
}

func main() {
	reset := flag.Bool("r", false, "Reset all audios to 100% both sides")
	caos := flag.Bool("c", false, "Total caos, volumes and sides are random")
	flag.Parse()

	audios := lib.GetResources[types.Audio](CommandList["audio"])
	if *reset {
		for _, audio := range audios {
			exec.Command("pactl", "set-sink-input-volume", strconv.Itoa(audio.ID), lib.FullAudio, lib.FullAudio).Output()
		}
		return
	}

	if *caos {
		min := 25
		max := 100
		for _, audio := range audios {
			leftEarVol := lib.GetRandomValue(min, max)
			rightEarVol := lib.GetRandomValue(min, max)
			exec.Command("pactl", "set-sink-input-volume", strconv.Itoa(audio.ID), fmt.Sprintf("%d%%", leftEarVol), fmt.Sprintf("%d%%", rightEarVol)).Output()
		}
		return
	}

	monitors := lib.GetResources[types.Monitor](CommandList["monitor"])
	clients := lib.GetResources[types.Client](CommandList["client"])
	leftMonitorID := lib.GetLeftMonitorID(monitors)

	lefties := []int{}
	righties := []int{}

	// FIX: Some clients share the same PID even if they are in different windows, this cause the
	// loop to be blocked by audios within the same PID, so only the first one is appended to the lists
	// and this happens on both sides (lefties and righties), this cause the audio of the client to
	// only go to the right because it is the last one setted
	// *Also, the block happens purely by list order
	//
	// NOTE: I have searched for information inside audio json but it does not feature anything that differs
	// one audio from another, so there is no way of finding out which audios link to which clients, a
	// solution would be checking if there are repeated PIDs and actually do nothing with them, or map
	// them purely by luck (a left client audio could be mapped to the right (which sounds fun))
	for _, client := range clients {
		for _, audio := range audios {
			if audio.Properties.PID == strconv.Itoa(client.PID) {
				if client.Monitor == leftMonitorID {
					lefties = append(lefties, audio.ID)
				} else {
					righties = append(righties, audio.ID)
				}
				break
			}
		}
	}

	lib.SetAudioBalance(lefties, righties)
}
