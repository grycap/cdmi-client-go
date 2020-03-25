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

import "errors"

var (
	// ErrBadRequest error thrown when request contains invalid parameters or field names (HTTP status code 400).
	ErrBadRequest = errors.New("The request contains invalid parameters or field names")

	// ErrUnauthorized error thrown when authentication credentials are missing or invalid (HTTP status code 401).
	ErrUnauthorized = errors.New("The authentication credentials are missing or invalid")

	// ErrForbidden error thrown when client lacks the proper authorization to perform this request (HTTP status code 403).
	ErrForbidden = errors.New("The client lacks the proper authorization to perform this request")

	// ErrNotFound error thrown when resource was not found at the specified URI (HTTP status code 404).
	ErrNotFound = errors.New("The resource was not found at the specified URI")

	// ErrNotAcceptable error thrown when server is unable to provide the object in the content type specified in the Accept header (HTTP status code 406).
	ErrNotAcceptable = errors.New("The server is unable to provide the object in the content type specified in the Accept header")

	// ErrConflict error thrown when operation conflicts with a non-CDMI access protocol lock or has caused a state transition error on the server (HTTP status code 409).
	ErrConflict = errors.New("The operation conflicts with a non-CDMI access protocol lock or has caused a state transition error on the server")

	// ErrInternalServer error thrown when server fails (HTTP status code 500).
	ErrInternalServer = errors.New("Internal server error")
)

func errorFromCode(errorCode int) error {
	switch errorCode {
	case 400:
		return ErrBadRequest
	case 401:
		return ErrUnauthorized
	case 403:
		return ErrForbidden
	case 404:
		return ErrNotFound
	case 406:
		return ErrNotAcceptable
	case 409:
		return ErrConflict
	case 500:
		return ErrInternalServer
	}

	return nil
}
