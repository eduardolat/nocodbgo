package nocodbgo

import (
	"encoding/json"
	"fmt"
)

// decodeInto converts a map or slice of maps into a struct or slice of structs
func decodeInto(data any, dest any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	if err := json.Unmarshal(jsonData, dest); err != nil {
		return fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return nil
}

// structToMap converts a struct into a map[string]any using JSON tags
func structToMap(data any) (map[string]any, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal struct: %w", err)
	}

	var result map[string]any
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal into map: %w", err)
	}

	return result, nil
}

// structsToMaps converts a slice of structs into a slice of maps using JSON tags
func structsToMaps(data any) ([]map[string]any, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal structs: %w", err)
	}

	var result []map[string]any
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal into maps: %w", err)
	}

	return result, nil
}
