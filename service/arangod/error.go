//
// DISCLAIMER
//
// Copyright 2017 ArangoDB GmbH, Cologne, Germany
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
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//
// Author Ewout Prangsma
//

package arangod

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	maskAny              = errors.WithStack
	ConditionFailedError = errors.New("Condition failed")
	KeyNotFoundError     = errors.New("Key not found")
)

type StatusError struct {
	StatusCode int
}

func (e StatusError) Error() string {
	return fmt.Sprintf("Status %d", e.StatusCode)
}

func IsStatusError(err error) (int, bool) {
	err = errors.Cause(err)
	if serr, ok := err.(StatusError); ok {
		return serr.StatusCode, true
	}
	return 0, false
}

func IsConditionFailed(err error) bool {
	return errors.Cause(err) == ConditionFailedError
}

func IsKeyNotFound(err error) bool {
	return errors.Cause(err) == KeyNotFoundError
}

func IsPreconditionFailed(err error) bool {
	statusCode, ok := IsStatusError(err)
	return ok && statusCode == 412
}

// NotLeaderError indicates the response of an agent when it is
// not the leader of the agency.
type NotLeaderError struct {
	Leader string // Endpoint of the current leader
}

// Error implements error.
func (e NotLeaderError) Error() string {
	return "not the leader"
}

// IsNotLeader returns true if the given error is (or is caused by) a NotLeaderError.
func IsNotLeader(err error) (string, bool) {
	nlErr, ok := errors.Cause(err).(NotLeaderError)
	if ok {
		return nlErr.Leader, true
	}
	return "", false
}
