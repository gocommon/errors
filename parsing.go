// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errors

import (
	"encoding/json"
	"fmt"
)

// ParseError represents a parsing error
type ParseError struct {
	code    int32
	Name    string
	In      string
	Value   string
	Reason  error
	message string
}

func (e *ParseError) Error() string {
	return e.message
}

// Code returns the http status code for this error
func (e *ParseError) Code() int32 {
	return e.code
}

// MarshalJSON implements the JSON encoding interface
func (e ParseError) MarshalJSON() ([]byte, error) {
	var reason string
	if e.Reason != nil {
		reason = e.Reason.Error()
	}
	return json.Marshal(map[string]interface{}{
		"code":    e.code,
		"message": e.message,
		"in":      e.In,
		"name":    e.Name,
		"value":   e.Value,
		"reason":  reason,
	})
}

const (
	parseErrorTemplContent     = `从 %q 解析 %s %s 失败，因为 %s`
	parseErrorTemplContentNoIn = `从 %q 解析 %s 失败，因为 %s`
)

// NewParseError creates a new parse error
func NewParseError(name, in, value string, reason error) *ParseError {
	var msg string
	if in == "" {
		msg = fmt.Sprintf(parseErrorTemplContentNoIn, value, name, reason)
	} else {
		msg = fmt.Sprintf(parseErrorTemplContent, value, in, name, reason)
	}
	return &ParseError{
		code:    400,
		Name:    name,
		In:      in,
		Value:   value,
		Reason:  reason,
		message: msg,
	}
}
