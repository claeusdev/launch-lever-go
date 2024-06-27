package launchlever

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type Store struct {
	Toggles ToggleStore `json:"toggles"`
}

type ToggleStore map[string]Toggle

type Toggle struct {
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

func NewStore(file string) (*Store, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading json file: %+v", err))
	}

	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	var config ToggleStore
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		fmt.Println(err)
	}

	return &Store{config}, nil
}

func (s *Store) GetToggles() ToggleStore {
	return s.Toggles
}

func LaunchLever(jsonData string) func(n string) bool {
	toggleStore, err := NewStore(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	return func(toggleName string) bool {
		for key, val := range toggleStore.GetToggles() {
			fmt.Println(key, val, toggleStore.GetToggles())
			if key == toggleName {
				return val.Enabled
			}
		}
		return false // Return false if no toggle found
	}
}
