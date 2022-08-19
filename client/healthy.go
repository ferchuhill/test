package client

import "net/http"

// IsHealthy Response is the type used to receive the status of the /health url
type IsHealthyResponse struct {
	Status string `json:"status"`
}

// IsHealthy tests if the service is alive and responds accordingly
func (c *Client) IsHealthy() bool {
	data := IsHealthyResponse{}
	err := c.CallAPI(http.MethodGet, "/health", nil, &data)
	return (err == nil && data.Status == "up")

}
