package main

import (
	"flag"
	"os/exec"
	"strconv"
	"strings"

	"tiktok-audio/lib"
	"tiktok-audio/types"
)

var CommandList = map[string][]string{
	"monitor":   {"hyprctl", "monitors", "-j"},
	"workspace": {"hyprctl", "workspaces", "-j"},
	"audio":     {"pactl", "-f", "json", "list", "sink-inputs"},
}

func main() {
	reset := flag.Bool("r", false, "Reset all audios to 100% both sides")
	flag.Parse()

	audios := lib.GetResources[types.Audio](CommandList["audio"])
	if *reset {
		for _, audio := range audios {
			exec.Command("pactl", "set-sink-input-volume", strconv.Itoa(audio.ID), lib.FullAudio, lib.FullAudio).Output()
		}
		return
	}

	monitors := lib.GetResources[types.Monitor](CommandList["monitor"])
	workspaces := lib.GetResources[types.Workspace](CommandList["workspace"])
	leftMonitorID := lib.GetLeftMonitorID(monitors)

	lefties := []int{}
	righties := []int{}

	for _, workspace := range workspaces {
		media := workspace.LastWindowTitle

		matchedAudios := []types.Audio{}
		for _, audio := range audios {
			if strings.HasPrefix(media, audio.Properties.Media) {
				matchedAudios = append(matchedAudios, audio)
			}
		}

		for _, matched := range matchedAudios {
			if workspace.MonitorID == leftMonitorID {
				lefties = append(lefties, matched.ID)
			} else {
				righties = append(righties, matched.ID)
			}
		}
	}

	lib.SetAudioBalance(lefties, righties)
}
