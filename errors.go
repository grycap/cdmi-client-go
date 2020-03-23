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
	errBadRequest    = errors.New("The request contains invalid parameters or field names")
	errUnauthorized  = errors.New("The authentication credentials are missing or invalid")
	errForbidden     = errors.New("The client lacks the proper authorization to perform this request")
	errNotFound      = errors.New("The resource was not found at the specified URI")
	errNotAcceptable = errors.New("The server is unable to provide the object in the content type specified in the Accept header")
)

func errorFromCode(errorCode int) error {
	switch errorCode {
	case 400:
		return errBadRequest
	case 401:
		return errUnauthorized
	case 403:
		return errForbidden
	case 404:
		return errNotFound
	case 406:
		return errNotAcceptable
	}

	return nil
}
