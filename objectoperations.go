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
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// CreateObject creates a new object with the content of ...
func (c *Client) CreateObject(objectPath string, data io.Reader, createContainer bool) error {
	objectPath = strings.Trim(objectPath, " /")
	endpoint, _ := url.Parse(c.Endpoint.String())

	endpoint.Path = path.Join(endpoint.Path, objectPath)

	if createContainer {
		// Check if parent folder exists
		if _, err := c.ReadContainer(path.Dir(objectPath)); err != nil {
			err = c.CreateContainer(path.Dir(objectPath), true)
			if err != nil {
				return err
			}
		}
	}

	req, err := http.NewRequest("PUT", endpoint.String(), data)
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

// DeleteObject deletes an object
func (c *Client) DeleteObject(objectPath string) error {
	endpoint, _ := url.Parse(c.Endpoint.String())
	endpoint.Path = path.Join(endpoint.Path, objectPath)

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

// GetObject downloads an object
func (c *Client) GetObject(objectPath string) (io.ReadCloser, error) {
	endpoint, _ := url.Parse(c.Endpoint.String())
	endpoint.Path = path.Join(endpoint.Path, objectPath)

	req, err := http.NewRequest("GET", endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Error making the request: %v", err)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if statusErr := errorFromCode(res.StatusCode); statusErr != nil {
		return nil, statusErr
	}

	return res.Body, nil
}
