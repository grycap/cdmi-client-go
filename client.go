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
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
)

// Client represents a CDMI client
type Client struct {
	Endpoint   url.URL
	HTTPClient *http.Client
}

// New creates a new CDMI client
func New(endpoint url.URL, token string, verify bool) *Client {

	var transport http.RoundTripper
	transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !verify},
	}

	if token != "" {
		transport = &customTransport{
			transport: transport,
			token:     token,
		}
	}

	return &Client{
		Endpoint:   endpoint,
		HTTPClient: &http.Client{Transport: transport},
	}
}

// Struct to decorate a transport with a token
type customTransport struct {
	transport http.RoundTripper
	token     string
}

// RoundTrip function to implement the RoundTripper interface adding the token to request's headers
func (ct *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ct.token))
	return ct.transport.RoundTrip(req)
}
