package nocodbgo

import (
	"encoding/json"
	"fmt"
)

// decodeInto is a helper function to decode data into a destination struct
func decodeInto(data any, dest any) error {
	// Convert the data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	// Unmarshal the JSON into the destination
	if err := json.Unmarshal(jsonData, dest); err != nil {
		return fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return nil
}
