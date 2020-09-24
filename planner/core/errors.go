// Copyright 2018 PingCAP, Inc.
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

package core

import (
	"github.com/pingcap/parser/terror"
	mysql "github.com/pingcap/tidb/errno"
	"github.com/pingcap/tidb/util/dbterror"
)

// error definitions.
var (
	ErrUnsupportedType                 = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrUnsupportedType)
	ErrAnalyzeMissIndex                = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrAnalyzeMissIndex)
	ErrWrongParamCount                 = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWrongParamCount)
	ErrSchemaChanged                   = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrSchemaChanged)
	ErrTablenameNotAllowedHere         = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrTablenameNotAllowedHere)
	ErrNotSupportedYet                 = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrNotSupportedYet)
	ErrWrongUsage                      = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWrongUsage)
	ErrUnknown                         = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrUnknown)
	ErrUnknownTable                    = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrUnknownTable)
	ErrNoSuchTable                     = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrNoSuchTable)
	ErrWrongArguments                  = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWrongArguments)
	ErrWrongNumberOfColumnsInSelect    = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWrongNumberOfColumnsInSelect)
	ErrBadGeneratedColumn              = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrBadGeneratedColumn)
	ErrFieldNotInGroupBy               = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrFieldNotInGroupBy)
	ErrBadTable                        = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrBadTable)
	ErrKeyDoesNotExist                 = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrKeyDoesNotExist)
	ErrOperandColumns                  = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrOperandColumns)
	ErrInvalidGroupFuncUse             = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrInvalidGroupFuncUse)
	ErrIllegalReference                = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrIllegalReference)
	ErrNoDB                            = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrNoDB)
	ErrUnknownExplainFormat            = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrUnknownExplainFormat)
	ErrWrongGroupField                 = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWrongGroupField)
	ErrDupFieldName                    = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrDupFieldName)
	ErrNonUpdatableTable               = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrNonUpdatableTable)
	ErrInternal                        = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrInternal)
	ErrNonUniqTable                    = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrNonuniqTable)
	ErrWindowInvalidWindowFuncUse      = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWindowInvalidWindowFuncUse)
	ErrWindowInvalidWindowFuncAliasUse = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWindowInvalidWindowFuncAliasUse)
	ErrWindowNoSuchWindow              = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWindowNoSuchWindow)
	ErrWindowCircularityInWindowGraph  = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWindowCircularityInWindowGraph)
	ErrWindowNoChildPartitioning       = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWindowNoChildPartitioning)
	ErrWindowNoInherentFrame           = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWindowNoInherentFrame)
	ErrWindowNoRedefineOrderBy         = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWindowNoRedefineOrderBy)
	ErrWindowDuplicateName             = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWindowDuplicateName)
	ErrPartitionClauseOnNonpartitioned = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrPartitionClauseOnNonpartitioned)
	ErrWindowFrameStartIllegal         = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWindowFrameStartIllegal)
	ErrWindowFrameEndIllegal           = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWindowFrameEndIllegal)
	ErrWindowFrameIllegal              = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWindowFrameIllegal)
	ErrWindowRangeFrameOrderType       = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWindowRangeFrameOrderType)
	ErrWindowRangeFrameTemporalType    = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWindowRangeFrameTemporalType)
	ErrWindowRangeFrameNumericType     = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWindowRangeFrameNumericType)
	ErrWindowRangeBoundNotConstant     = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWindowRangeBoundNotConstant)
	ErrWindowRowsIntervalUse           = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWindowRowsIntervalUse)
	ErrWindowFunctionIgnoresFrame      = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWindowFunctionIgnoresFrame)
	ErrUnsupportedOnGeneratedColumn    = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrUnsupportedOnGeneratedColumn)
	ErrPrivilegeCheckFail              = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrPrivilegeCheckFail)
	ErrInvalidWildCard                 = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrInvalidWildCard)
	ErrMixOfGroupFuncAndFields         = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrMixOfGroupFuncAndFieldsIncompatible)
	errTooBigPrecision                 = dbterror.NewStd(terror.ClassExpression, mysql.ErrTooBigPrecision)
	ErrDBaccessDenied                  = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrDBaccessDenied)
	ErrTableaccessDenied               = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrTableaccessDenied)
	ErrSpecificAccessDenied            = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrSpecificAccessDenied)
	ErrViewNoExplain                   = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrViewNoExplain)
	ErrWrongValueCountOnRow            = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrWrongValueCountOnRow)
	ErrViewInvalid                     = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrViewInvalid)
	ErrNoSuchThread                    = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrNoSuchThread)
	ErrUnknownColumn                   = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrBadField)
	ErrCartesianProductUnsupported     = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrCartesianProductUnsupported)
	ErrStmtNotFound                    = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrPreparedStmtNotFound)
	ErrAmbiguous                       = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrNonUniq)
	// Since we cannot know if user logged in with a password, use message of ErrAccessDeniedNoPassword instead
	ErrAccessDenied = dbterror.NewStd(terror.ClassOptimizer, mysql.ErrAccessDenied)
)
