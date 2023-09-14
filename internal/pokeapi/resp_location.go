package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	locationArea := LocationArea{}
	data, ok := c.cache.Get(fullURL)
	if ok {

		err := json.Unmarshal(data, &locationArea)

		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	err = json.Unmarshal(data, &locationArea)

	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, data)

	return locationArea, nil
}

func (c *Client) ListLocationAreas(url *string) (LocationAreasResp, error) {
	endpoint := "/location-area/?offset=0&limit=20"
	fullURL := baseURL + endpoint
	if url != nil {
		fullURL = *url
	}

	locations := LocationAreasResp{}
	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit")

		err := json.Unmarshal(data, &locations)

		if err != nil {
			return LocationAreasResp{}, err
		}

		return locations, nil
	}
	fmt.Println("cache miss")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	err = json.Unmarshal(data, &locations)

	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullURL, data)

	return locations, nil

}
