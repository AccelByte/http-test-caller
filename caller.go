/*
 * Copyright 2018 AccelByte Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package caller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

// HTTPTestCaller is a wrapper for testing HTTP handler,
// it will pass the request to the handler and record
// the response using httptest.ResponseRecorder.
// If the response type is set, it will try to unmarshal
// the response body to the type
type HTTPTestCaller struct {
	handler  http.Handler
	request  *http.Request
	response interface{}
}

// Call creates the caller with handler specified,
// use this function to start the caller
func Call(handler http.Handler) *HTTPTestCaller {
	return &HTTPTestCaller{handler: handler}
}

// To receives a request to be sent to the handler.
// The second parameter is a convenience if you want to pass
// directly from http.NewRequest() or gorequest.MakeRequest()
func (caller *HTTPTestCaller) To(request *http.Request, _ ...interface{}) *HTTPTestCaller {
	caller.request = request
	return caller
}

// Read specifies the expected type of response returned by
// the handler being tested
func (caller *HTTPTestCaller) Read(response interface{}) *HTTPTestCaller {
	caller.response = response
	return caller
}

// Execute calls the handler with the request and records the response
// in httptest.ResponseRecorder.
// If the expected response type exists, it will try to unmarshal the
// response body to the type
func (caller *HTTPTestCaller) Execute() (*httptest.ResponseRecorder, interface{}, error) {
	recorder := httptest.NewRecorder()
	caller.handler.ServeHTTP(recorder, caller.request)

	var err error
	if caller.response != nil {
		err = json.Unmarshal(recorder.Body.Bytes(), &caller.response)
	}

	return recorder, caller.response, err
}
