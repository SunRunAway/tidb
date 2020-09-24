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

package variable

import (
	"github.com/pingcap/parser/terror"
	mysql "github.com/pingcap/tidb/errno"
	"github.com/pingcap/tidb/util/dbterror"
)

// Error instances.
var (
	errCantGetValidID              = dbterror.NewStd(terror.ClassVariable, mysql.ErrCantGetValidID)
	errWarnDeprecatedSyntax        = dbterror.NewStd(terror.ClassVariable, mysql.ErrWarnDeprecatedSyntax)
	ErrCantSetToNull               = dbterror.NewStd(terror.ClassVariable, mysql.ErrCantSetToNull)
	ErrSnapshotTooOld              = dbterror.NewStd(terror.ClassVariable, mysql.ErrSnapshotTooOld)
	ErrUnsupportedValueForVar      = dbterror.NewStd(terror.ClassVariable, mysql.ErrUnsupportedValueForVar)
	ErrUnknownSystemVar            = dbterror.NewStd(terror.ClassVariable, mysql.ErrUnknownSystemVariable)
	ErrIncorrectScope              = dbterror.NewStd(terror.ClassVariable, mysql.ErrIncorrectGlobalLocalVar)
	ErrUnknownTimeZone             = dbterror.NewStd(terror.ClassVariable, mysql.ErrUnknownTimeZone)
	ErrReadOnly                    = dbterror.NewStd(terror.ClassVariable, mysql.ErrVariableIsReadonly)
	ErrWrongValueForVar            = dbterror.NewStd(terror.ClassVariable, mysql.ErrWrongValueForVar)
	ErrWrongTypeForVar             = dbterror.NewStd(terror.ClassVariable, mysql.ErrWrongTypeForVar)
	ErrTruncatedWrongValue         = dbterror.NewStd(terror.ClassVariable, mysql.ErrTruncatedWrongValue)
	ErrMaxPreparedStmtCountReached = dbterror.NewStd(terror.ClassVariable, mysql.ErrMaxPreparedStmtCountReached)
	ErrUnsupportedIsolationLevel   = dbterror.NewStd(terror.ClassVariable, mysql.ErrUnsupportedIsolationLevel)
)
