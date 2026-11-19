package lib

import (
	"encoding/json"
	"log"
	"os/exec"
)

func GetResources[T any](command []string) []T {
	cmd := exec.Command(command[0], command[1:]...)

	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("[ERROR] Command failed: %v\n", err)
	}

	var resources []T
	if err := json.Unmarshal(output, &resources); err != nil {
		log.Fatalf("[ERROR] Failed to parse JSON: %v\n", err)
	}

	return resources
}
