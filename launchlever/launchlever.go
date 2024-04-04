package launchlever

import (
	"encoding/json"
	"log"
)

type Toggle struct {
	Name        string
	Description string
	Status      string
}

func LaunchLever(jsonData string) func(string) (Toggle, bool) {
	var toggles []Toggle
	err := json.Unmarshal([]byte(jsonData), &toggles)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return func(toggleName string) (Toggle, bool) {
		for _, val := range toggles {
			if val.Name == toggleName {
				return val, true
			}
		}
		return Toggle{}, false // Return false if no toggle found
	}
}
