package nocodbgo

import "net/url"

// viewSelectable provides a reusable set of methods for building query with support for view selection
// using the "viewId" query parameter.
//
// It is designed to be embedded in builder types to provide consistent view selection capabilities.
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
type viewSelectable[T any] struct {
	builder T
	viewId  string
}

// newViewSelectable creates a new viewSelectable instance with the given builder.
func newViewSelectable[T any](builder T) viewSelectable[T] {
	return viewSelectable[T]{
		builder: builder,
	}
}

// apply takes the url.Values and adds the "viewId" query parameter to it with the value
// that has been set on the viewSelectable instance.
func (v *viewSelectable[T]) apply(query url.Values) {
	if v.viewId != "" {
		query.Set("viewId", v.viewId)
	}
}

// WithViewId specifies the view identifier to fetch records that are currently visible within a specific view.
//
// When set, the API retrieves records in the order they are displayed if the SORT option is enabled within that view.
// If you also specify sort parameters, they will take precedence over any sorting configuration defined in the view.
// If you specify where parameters, they will be applied over the filtering configuration defined in the view.
//
// By default, all fields (including those disabled within the view) are included in the response.
// To explicitly specify which fields to include or exclude, use the ReturnFields method.
//
// Example:
//
//	// Get records from the "DashboardView" view
//	query = query.WithViewId("DashboardView")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
func (v *viewSelectable[T]) WithViewId(viewId string) T {
	if viewId != "" {
		v.viewId = viewId
	}
	return v.builder
}
