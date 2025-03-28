package nocodbgo

import "net/url"

// viewIDProvider provides a reusable set of methods for building query with support for view selection
// using the "viewId" query parameter.
//
// It is designed to be embedded in builder types to provide consistent view selection capabilities.
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
type viewIDProvider[T any] struct {
	builder   T
	rawViewID string
}

// newViewIDProvider creates a new viewIDProvider instance with the given builder.
func newViewIDProvider[T any](builder T) viewIDProvider[T] {
	return viewIDProvider[T]{
		builder: builder,
	}
}

// apply takes the url.Values and adds the "viewId" query parameter to it with the value
// that has been set on the viewIDProvider instance.
//
// It returns a new copy of the provided url.Values with the "viewId" query parameter added.
func (v *viewIDProvider[T]) apply(query url.Values) url.Values {
	if query == nil || v.rawViewID == "" {
		return query
	}

	query.Set("viewId", v.rawViewID)
	return query
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
func (v *viewIDProvider[T]) WithViewId(viewId string) T {
	if viewId != "" {
		v.rawViewID = viewId
	}
	return v.builder
}
