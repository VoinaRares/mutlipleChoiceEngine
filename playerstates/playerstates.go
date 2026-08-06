package playerstates

import (
	"encoding/json"
	"os"
)

type PlayerStates struct {
	// Represents the possible states a player can have and the accumulation of them
	States map[string]int
}

type playerStates struct {
	States []string `json:"states"`
}

func BuildPlayerStates(path string) (PlayerStates, error) {
	initialPlayerStates, err := parseJsonToPlayerStates(path)
	if err != nil {
		return PlayerStates{}, err
	}
	initialStates := make(map[string]int)
	for _, key := range initialPlayerStates.States {
		initialStates[key] = 0
	}
	return PlayerStates{States: initialStates}, nil
}

// We might look into having a specific JsonParser and using that with generics
func parseJsonToPlayerStates(path string) (playerStates, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return playerStates{}, err
	}
	defer jsonFile.Close()
	var playerStates playerStates

	return playerStates, json.NewDecoder(jsonFile).Decode(&playerStates)
}
