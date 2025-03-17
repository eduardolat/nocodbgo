package nocodbgo

import (
	"encoding/json"
	"fmt"
)

// decodeInto converts data from a map or slice of maps into the provided destination struct or slice of structs.
// It uses JSON marshaling and unmarshaling internally to perform the conversion.
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

// structToMap converts a struct into a map[string]any using the struct's JSON tags.
// This is useful when you need to convert a strongly typed struct into a map for API operations.
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

// structsToMaps converts a slice of structs into a slice of maps using JSON tags.
// This is useful when you need to convert a slice of strongly typed structs into a slice of maps for API operations.
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
