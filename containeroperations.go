// Copyright (C) GRyCAP - I3M - UPV
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cdmi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type readContainerResponse struct {
	Children []string `json:"children"`
}

// CreateContainer creates a new container
func (c *Client) CreateContainer(containerPath string, recursive bool) error {
	containerPath = strings.Trim(containerPath, " /")
	endpoint, _ := url.Parse(c.Endpoint.String())

	endpoint.Path = fmt.Sprintf("%s/", path.Join(endpoint.Path, containerPath))

	if recursive {
		// Check if parent folder exists
		if _, err := c.ReadContainer(path.Dir(containerPath)); err != nil {
			err = c.CreateContainer(path.Dir(containerPath), recursive)
			if err != nil {
				return err
			}
		}
	}

	req, err := http.NewRequest("PUT", endpoint.String(), nil)
	if err != nil {
		return fmt.Errorf("Error making the request: %v", err)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	if statusErr := errorFromCode(res.StatusCode); statusErr != nil {
		return statusErr
	}

	return nil
}

// ReadContainer checks if a container exists and returns a slice with its children
func (c *Client) ReadContainer(containerPath string) ([]string, error) {
	endpoint, _ := url.Parse(c.Endpoint.String())
	endpoint.Path = path.Join(endpoint.Path, containerPath)

	req, err := http.NewRequest("GET", endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Error making the request: %v", err)
	}
	req.Header.Add(VersionHeader, Version)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if statusErr := errorFromCode(res.StatusCode); statusErr != nil {
		return nil, statusErr
	}

	readContainerResponse := &readContainerResponse{}

	json.NewDecoder(res.Body).Decode(readContainerResponse)

	return readContainerResponse.Children, nil
}

// DeleteContainer deletes a container with its children
func (c *Client) DeleteContainer(containerPath string) error {
	endpoint, _ := url.Parse(c.Endpoint.String())
	endpoint.Path = path.Join(endpoint.Path, containerPath)

	req, err := http.NewRequest("DELETE", endpoint.String(), nil)
	if err != nil {
		return fmt.Errorf("Error making the request: %v", err)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	if statusErr := errorFromCode(res.StatusCode); statusErr != nil {
		return statusErr
	}

	return nil
}
