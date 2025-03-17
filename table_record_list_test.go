package nocodbgo

import (
	"encoding/json"
	"testing"
)

func TestListResponseUnmarshalJSON(t *testing.T) {
	t.Run("StandardListResponse", func(t *testing.T) {
		// Test standard list response with pagination
		standardJSON := `{
			"list": [
				{
					"Id": 1,
					"Title": "Record 1"
				},
				{
					"Id": 2,
					"Title": "Record 2"
				}
			],
			"pageInfo": {
				"totalRows": 2,
				"page": 1,
				"pageSize": 10,
				"isFirstPage": true,
				"isLastPage": true
			}
		}`

		var standardResp ListResponse
		err := json.Unmarshal([]byte(standardJSON), &standardResp)
		if err != nil {
			t.Fatalf("Failed to unmarshal standard list response: %v", err)
		}

		// Verify standard response
		if len(standardResp.List) != 2 {
			t.Errorf("Expected 2 records in standard response, got %d", len(standardResp.List))
		}
		if standardResp.List[1]["Title"] != "Record 2" {
			t.Errorf("Expected Title='Record 2' in standard response, got %v", standardResp.List[1]["Title"])
		}
		if standardResp.PageInfo.TotalRows != 2 {
			t.Errorf("Expected TotalRows=2 in standard response, got %d", standardResp.PageInfo.TotalRows)
		}
		if standardResp.PageInfo.Page != 1 {
			t.Errorf("Expected Page=1 in standard response, got %d", standardResp.PageInfo.Page)
		}
	})

	t.Run("SingleObjectResponse", func(t *testing.T) {
		// Test single object response
		singleJSON := `{
			"Id": 3,
			"Title": "Single Record"
		}`

		var singleResp ListResponse
		err := json.Unmarshal([]byte(singleJSON), &singleResp)
		if err != nil {
			t.Fatalf("Failed to unmarshal single object response: %v", err)
		}

		// Verify single object response
		if len(singleResp.List) != 1 {
			t.Errorf("Expected 1 record in single object response, got %d", len(singleResp.List))
		}
		if singleResp.List[0]["Id"].(float64) != 3 {
			t.Errorf("Expected Id=3 in single object response, got %v", singleResp.List[0]["Id"])
		}
		if singleResp.List[0]["Title"].(string) != "Single Record" {
			t.Errorf("Expected Title='Single Record' in single object response, got %v", singleResp.List[0]["Title"])
		}
		if singleResp.PageInfo.TotalRows != 1 {
			t.Errorf("Expected TotalRows=1 in single object response, got %d", singleResp.PageInfo.TotalRows)
		}
		if singleResp.PageInfo.Page != 1 {
			t.Errorf("Expected Page=1 in single object response, got %d", singleResp.PageInfo.Page)
		}
		if !singleResp.PageInfo.IsFirstPage || !singleResp.PageInfo.IsLastPage {
			t.Errorf("Expected IsFirstPage=true and IsLastPage=true in single object response")
		}
	})

	t.Run("EmptyListResponse", func(t *testing.T) {
		// Test empty array in list field
		emptyJSON := `{
			"list": [],
			"pageInfo": {
				"totalRows": 0,
				"page": 1,
				"pageSize": 10,
				"isFirstPage": true,
				"isLastPage": true
			}
		}`

		var emptyResp ListResponse
		err := json.Unmarshal([]byte(emptyJSON), &emptyResp)
		if err != nil {
			t.Fatalf("Failed to unmarshal empty array response: %v", err)
		}

		// Verify empty array response
		if len(emptyResp.List) != 0 {
			t.Errorf("Expected 0 records in empty array response, got %d", len(emptyResp.List))
		}
		if emptyResp.PageInfo.TotalRows != 0 {
			t.Errorf("Expected TotalRows=0 in empty array response, got %d", emptyResp.PageInfo.TotalRows)
		}
	})

	t.Run("EmptyObjectResponse", func(t *testing.T) {
		// Test empty object
		emptyObjJSON := `{}`

		var emptyObjResp ListResponse
		err := json.Unmarshal([]byte(emptyObjJSON), &emptyObjResp)
		if err != nil {
			t.Fatalf("Failed to unmarshal empty JSON object: %v", err)
		}

		// After unmarshaling an empty object, the struct fields should remain zero-initialized
		if emptyObjResp.List != nil {
			t.Errorf("Expected nil List after unmarshaling empty object, got %v", emptyObjResp.List)
		}
		if emptyObjResp.PageInfo != (PageInfo{}) {
			t.Errorf("Expected zero PageInfo after unmarshaling empty object, got %v", emptyObjResp.PageInfo)
		}
	})

	t.Run("EmptyArrayResponse", func(t *testing.T) {
		// Test empty array
		emptyArrayJSON := `[]`

		var emptyArrayResp ListResponse
		err := json.Unmarshal([]byte(emptyArrayJSON), &emptyArrayResp)
		if err != nil {
			t.Fatalf("Failed to unmarshal empty JSON array: %v", err)
		}

		// After unmarshaling an empty array, the struct fields should remain zero-initialized
		if emptyArrayResp.List != nil {
			t.Errorf("Expected nil List after unmarshaling empty array, got %v", emptyArrayResp.List)
		}
		if emptyArrayResp.PageInfo != (PageInfo{}) {
			t.Errorf("Expected zero PageInfo after unmarshaling empty array, got %v", emptyArrayResp.PageInfo)
		}
	})

	t.Run("ShortDataResponse", func(t *testing.T) {
		// Test data with length <= 2
		shortJSON := `1`

		var shortResp ListResponse
		err := json.Unmarshal([]byte(shortJSON), &shortResp)
		// This would normally cause an error in the standard JSON parser, but our custom function
		// should handle it gracefully if it gets through
		if err == nil {
			if shortResp.List != nil {
				t.Errorf("Expected nil List after unmarshaling short data, got %v", shortResp.List)
			}
		}
	})

	t.Run("NonObjectResponse", func(t *testing.T) {
		// Test non-object data
		nullJSON := `null`

		var nullResp ListResponse
		err := json.Unmarshal([]byte(nullJSON), &nullResp)
		if err != nil {
			t.Fatalf("Failed to unmarshal JSON null: %v", err)
		}

		// After unmarshaling null, the struct fields should remain zero-initialized
		if nullResp.List != nil {
			t.Errorf("Expected nil List after unmarshaling null, got %v", nullResp.List)
		}
		if nullResp.PageInfo != (PageInfo{}) {
			t.Errorf("Expected zero PageInfo after unmarshaling null, got %v", nullResp.PageInfo)
		}
	})

	t.Run("MalformedObjectResponse", func(t *testing.T) {
		// Test object that doesn't start with { or end with }
		malformedJSON := `{"id": 1`

		var malformedResp ListResponse
		err := json.Unmarshal([]byte(malformedJSON), &malformedResp)
		// This should result in an error from the standard JSON parser
		if err == nil {
			t.Errorf("Expected error when unmarshaling malformed JSON object")
		}
	})

	t.Run("ListWithoutPageInfo", func(t *testing.T) {
		// Test list without pageInfo
		listNoPageJSON := `{
			"list": [
				{
					"Id": 1,
					"Title": "Record 1"
				}
			]
		}`

		var listNoPageResp ListResponse
		err := json.Unmarshal([]byte(listNoPageJSON), &listNoPageResp)
		if err != nil {
			t.Fatalf("Failed to unmarshal list without pageInfo: %v", err)
		}

		// Should be treated as a single object
		if len(listNoPageResp.List) != 1 {
			t.Errorf("Expected 1 record in list response, got %d", len(listNoPageResp.List))
		}
		// Should have default page info for a single record
		if listNoPageResp.PageInfo.TotalRows != 1 {
			t.Errorf("Expected TotalRows=1, got %d", listNoPageResp.PageInfo.TotalRows)
		}
	})
}
