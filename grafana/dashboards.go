package grafana

import (
	"fmt"
	"net/http"
)

// ListDashboards returns a list of dashboards.
func (c *Client) ListDashboards() ([]Dashboard, error) {

	resp, err := c.doRequest(&request{
		method:   http.MethodGet,
		endpoint: DashboardSearchEndpoint,
	})
	if err != nil {
		return nil, err
	}

	// Decode the response into a []User response object.
	var dashboards []Dashboard
	if err := decodeResponse(resp, &dashboards); err != nil {
		// Read the body of the request, ignore the error since we are already in the error state.
		return nil, fmt.Errorf("decoding response from request to failed, err -> %v", err)
	}

	// Check if we didn't get a result and return an error if true.
	if dashboards == nil || len(dashboards) <= 0 {
		return make([]Dashboard, 0), nil
	}

	return dashboards, nil
}