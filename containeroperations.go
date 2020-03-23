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
	"net/http"
	"path"
)

type readContainerResponse struct {
	Children []string `json:"children"`
}

// CreateContainer creates a new container
func (c Client) CreateContainer(containerPath string, recursive bool) error {
	// TODO
	return nil
}

// ReadContainer checks if a container exists and returns a slice with its children
func (c Client) ReadContainer(containerPath string) ([]string, error) {
	c.Endpoint.Path = path.Join(c.Endpoint.Path, containerPath)

	req, _ := http.NewRequest("GET", c.Endpoint.String(), nil)
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
