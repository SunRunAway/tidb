// Copyright 2017 PingCAP, Inc.
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

package expression

import (
	"github.com/pingcap/parser/terror"
	mysql "github.com/pingcap/tidb/errno"
	"github.com/pingcap/tidb/sessionctx"
	"github.com/pingcap/tidb/types"
	"github.com/pingcap/tidb/util/dbterror"
)

// Error instances.
var (
	// All the exported errors are defined here:
	ErrIncorrectParameterCount     = dbterror.NewStd(terror.ClassExpression, mysql.ErrWrongParamcountToNativeFct)
	ErrDivisionByZero              = dbterror.NewStd(terror.ClassExpression, mysql.ErrDivisionByZero)
	ErrRegexp                      = dbterror.NewStd(terror.ClassExpression, mysql.ErrRegexp)
	ErrOperandColumns              = dbterror.NewStd(terror.ClassExpression, mysql.ErrOperandColumns)
	ErrCutValueGroupConcat         = dbterror.NewStd(terror.ClassExpression, mysql.ErrCutValueGroupConcat)
	ErrFunctionsNoopImpl           = dbterror.New(terror.ClassExpression, mysql.ErrNotSupportedYet, "function %s has only noop implementation in tidb now, use tidb_enable_noop_functions to enable these functions")
	ErrInvalidArgumentForLogarithm = dbterror.NewStd(terror.ClassExpression, mysql.ErrInvalidArgumentForLogarithm)
	ErrIncorrectType               = dbterror.NewStd(terror.ClassExpression, mysql.ErrIncorrectType)

	// All the un-exported errors are defined here:
	errFunctionNotExists             = dbterror.NewStd(terror.ClassExpression, mysql.ErrSpDoesNotExist)
	errZlibZData                     = dbterror.NewStd(terror.ClassExpression, mysql.ErrZlibZData)
	errZlibZBuf                      = dbterror.NewStd(terror.ClassExpression, mysql.ErrZlibZBuf)
	errIncorrectArgs                 = dbterror.NewStd(terror.ClassExpression, mysql.ErrWrongArguments)
	errUnknownCharacterSet           = dbterror.NewStd(terror.ClassExpression, mysql.ErrUnknownCharacterSet)
	errDefaultValue                  = dbterror.New(terror.ClassExpression, mysql.ErrInvalidDefault, "invalid default value")
	errDeprecatedSyntaxNoReplacement = dbterror.NewStd(terror.ClassExpression, mysql.ErrWarnDeprecatedSyntaxNoReplacement)
	errBadField                      = dbterror.NewStd(terror.ClassExpression, mysql.ErrBadField)
	errWarnAllowedPacketOverflowed   = dbterror.NewStd(terror.ClassExpression, mysql.ErrWarnAllowedPacketOverflowed)
	errWarnOptionIgnored             = dbterror.NewStd(terror.ClassExpression, mysql.WarnOptionIgnored)
	errTruncatedWrongValue           = dbterror.NewStd(terror.ClassExpression, mysql.ErrTruncatedWrongValue)
	errUnknownLocale                 = dbterror.NewStd(terror.ClassExpression, mysql.ErrUnknownLocale)
	errNonUniq                       = dbterror.NewStd(terror.ClassExpression, mysql.ErrNonUniq)

	// Sequence usage privilege check.
	errSequenceAccessDenied = dbterror.NewStd(terror.ClassExpression, mysql.ErrTableaccessDenied)
)

// handleInvalidTimeError reports error or warning depend on the context.
func handleInvalidTimeError(ctx sessionctx.Context, err error) error {
	if err == nil || !(types.ErrWrongValue.Equal(err) ||
		types.ErrTruncatedWrongVal.Equal(err) || types.ErrInvalidWeekModeFormat.Equal(err) ||
		types.ErrDatetimeFunctionOverflow.Equal(err)) {
		return err
	}
	sc := ctx.GetSessionVars().StmtCtx
	err = sc.HandleTruncate(err)
	if ctx.GetSessionVars().StrictSQLMode && (sc.InInsertStmt || sc.InUpdateStmt || sc.InDeleteStmt) {
		return err
	}
	return nil
}

// handleDivisionByZeroError reports error or warning depend on the context.
func handleDivisionByZeroError(ctx sessionctx.Context) error {
	sc := ctx.GetSessionVars().StmtCtx
	if sc.InInsertStmt || sc.InUpdateStmt || sc.InDeleteStmt {
		if !ctx.GetSessionVars().SQLMode.HasErrorForDivisionByZeroMode() {
			return nil
		}
		if ctx.GetSessionVars().StrictSQLMode && !sc.DividedByZeroAsWarning {
			return ErrDivisionByZero
		}
	}
	sc.AppendWarning(ErrDivisionByZero)
	return nil
}
