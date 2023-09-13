package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations() (LocationsResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationsResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationsResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationsResp{}, err
	}

	locations := LocationsResp{}

	err = json.Unmarshal(data, &locations)

	if err != nil {
		return LocationsResp{}, err
	}

	return locations, nil

}
