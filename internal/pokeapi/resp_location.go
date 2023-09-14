package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(url *string) (LocationsResp, error) {
	endpoint := "/location-area/?offset=0&limit=20"
	fullURL := baseURL + endpoint
	if url != nil {
		fullURL = *url
	}

	locations := LocationsResp{}
	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit")

		err := json.Unmarshal(data, &locations)

		if err != nil {
			return LocationsResp{}, err
		}

		return locations, nil
	}
	fmt.Println("cache miss")

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

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationsResp{}, err
	}

	err = json.Unmarshal(data, &locations)

	if err != nil {
		return LocationsResp{}, err
	}

	c.cache.Add(fullURL, data)

	return locations, nil

}
