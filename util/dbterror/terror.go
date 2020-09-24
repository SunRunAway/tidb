// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package dbterror

import (
	"github.com/pingcap/parser/terror"
	"github.com/pingcap/tidb/config"
	"github.com/pingcap/tidb/errno"
)

func New(class terror.ErrClass, code terror.ErrCode, message string) *Error {
	return &Error{
		tError:          class.New(code, message),
		redactPositions: errno.NeedRedact[uint16(code)],
	}
}

func NewStd(class terror.ErrClass, code terror.ErrCode) *Error {
	return &Error{
		tError:          class.New(code, errno.MySQLErrName[uint16(code)]),
		redactPositions: errno.NeedRedact[uint16(code)],
	}
}

type tError = terror.Error

type Error struct {
	*tError
	redactPositions []int
}

// GenWithStackByArgs generates a new *Error with the same class and code, and new arguments.
func (e *Error) GenWithStackByArgs(args ...interface{}) error {
	e.redactErrorArg(args, e.redactPositions)
	return e.tError.GenWithStackByArgs(args...)
}

// FastGen generates a new *Error with the same class and code, and a new arguments.
func (e *Error) FastGenByArgs(args ...interface{}) error {
	e.redactErrorArg(args, e.redactPositions)
	return e.tError.GenWithStackByArgs(args...)
}

// Equal checks if err is equal to e.
func (e *Error) Equal(err error) bool {
	if underlying, ok := err.(*Error); ok {
		return e.tError.Equal(underlying.tError)
	}
	return e.tError.Equal(err)
}

// Cause implement the Cause interface.
func (e *Error) Cause() error {
	return e.tError
}

func (e *Error) redactErrorArg(args []interface{}, position []int) {
	if config.RedactLogEnabled() {
		for _, pos := range position {
			if len(args) > pos {
				args[pos] = "?"
			}
		}
	}
}
