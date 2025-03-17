package nocodbgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// listRecordsBuilder is used to build a list query with a fluent API
type listRecordsBuilder struct {
	table *Table

	contextProvider[*listRecordsBuilder]
	filterProvider[*listRecordsBuilder]
	sortProvider[*listRecordsBuilder]
	paginationProvider[*listRecordsBuilder]
	fieldProvider[*listRecordsBuilder]
	shuffleProvider[*listRecordsBuilder]
	viewIDProvider[*listRecordsBuilder]
}

// ListRecords lists records from the table.
func (t *Table) ListRecords() *listRecordsBuilder {
	b := &listRecordsBuilder{
		table: t,
	}

	b.contextProvider = newContextProvider(b)
	b.filterProvider = newFilterProvider(b)
	b.sortProvider = newSortProvider(b)
	b.paginationProvider = newPaginationProvider(b)
	b.fieldProvider = newFieldProvider(b)
	b.shuffleProvider = newShuffleProvider(b)
	b.viewIDProvider = newViewIDProvider(b)

	return b
}

// ListResponse is the response from a list query with pagination information
type ListResponse struct {
	// List contains the records returned by the query
	List []map[string]any `json:"list"`
	// PageInfo contains pagination information
	PageInfo PageInfo `json:"pageInfo"`
}

// PageInfo contains pagination information for list queries
type PageInfo struct {
	// TotalRows is the total number of rows in the table
	TotalRows int `json:"totalRows"`
	// Page is the current page number
	Page int `json:"page"`
	// PageSize is the number of records per page
	PageSize int `json:"pageSize"`
	// IsFirstPage indicates if this is the first page
	IsFirstPage bool `json:"isFirstPage"`
	// IsLastPage indicates if this is the last page
	IsLastPage bool `json:"isLastPage"`
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListResponse.
// It handles both list responses with pagination and single object responses.
func (r *ListResponse) UnmarshalJSON(data []byte) error {
	if r == nil {
		r = &ListResponse{}
		r.List = []map[string]any{}
		r.PageInfo = PageInfo{}
	}

	if len(data) <= 2 {
		return nil
	}

	dataStr := string(data)
	isObject := strings.HasPrefix(dataStr, "{") && strings.HasSuffix(dataStr, "}")
	if dataStr == "{}" || !isObject {
		return nil
	}

	var rawMap map[string]any
	if err := json.Unmarshal(data, &rawMap); err != nil {
		return fmt.Errorf("failed to unmarshal list response: %w", err)
	}

	_, hasList := rawMap["list"]
	_, hasPageInfo := rawMap["pageInfo"]
	if hasList && hasPageInfo && len(rawMap) == 2 {
		// Avoid recursion by using a type alias
		type Alias ListResponse
		var aux Alias

		if err := json.Unmarshal(data, &aux); err != nil {
			return fmt.Errorf("failed to unmarshal list response: %w", err)
		}

		// Copy the data back to r
		*r = ListResponse(aux)
		return nil
	}

	r.List = []map[string]any{rawMap}
	r.PageInfo = PageInfo{
		TotalRows:   1,
		Page:        1,
		PageSize:    1,
		IsFirstPage: true,
		IsLastPage:  true,
	}
	return nil
}

// DecodeInto converts the list response data into a slice of the provided struct type.
// It takes a pointer to a slice of structs as destination and populates it with the data.
// Returns an error if the conversion fails.
func (r ListResponse) DecodeInto(dest any) error {
	return decodeInto(r.List, dest)
}

// Execute finalizes and executes the operation.
func (b *listRecordsBuilder) Execute() (ListResponse, error) {
	query := url.Values{}
	query = b.filterProvider.apply(query)
	query = b.sortProvider.apply(query)
	query = b.paginationProvider.apply(query)
	query = b.fieldProvider.apply(query)
	query = b.shuffleProvider.apply(query)
	query = b.viewIDProvider.apply(query)

	path := fmt.Sprintf("/api/v2/tables/%s/records", b.table.tableID)
	respBody, err := b.table.client.request(b.contextProvider.ctx, http.MethodGet, path, nil, query)
	if err != nil {
		return ListResponse{}, fmt.Errorf("failed to list records: %w", err)
	}

	var response ListResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return ListResponse{}, fmt.Errorf("failed to unmarshal list response: %w", err)
	}

	return response, nil
}
