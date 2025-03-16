package nocodbgo

import (
	"encoding/json"
	"testing"
)

func TestListResponseUnmarshalJSON(t *testing.T) {
	// Test case 1: Standard list response with pagination
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
	if standardResp.PageInfo.TotalRows != 2 {
		t.Errorf("Expected TotalRows=2 in standard response, got %d", standardResp.PageInfo.TotalRows)
	}
	if standardResp.PageInfo.Page != 1 {
		t.Errorf("Expected Page=1 in standard response, got %d", standardResp.PageInfo.Page)
	}

	// Test case 2: Single object response
	singleJSON := `{
		"Id": 3,
		"Title": "Single Record"
	}`

	var singleResp ListResponse
	err = json.Unmarshal([]byte(singleJSON), &singleResp)
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

	// Test case 3: Empty array response
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
	err = json.Unmarshal([]byte(emptyJSON), &emptyResp)
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
}
