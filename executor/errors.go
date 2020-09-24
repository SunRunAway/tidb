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

package executor

import (
	"github.com/pingcap/parser/terror"
	mysql "github.com/pingcap/tidb/errno"
	"github.com/pingcap/tidb/util/dbterror"
)

// Error instances.
var (
	ErrGetStartTS      = dbterror.NewStd(terror.ClassExecutor, mysql.ErrGetStartTS)
	ErrUnknownPlan     = dbterror.NewStd(terror.ClassExecutor, mysql.ErrUnknownPlan)
	ErrPrepareMulti    = dbterror.NewStd(terror.ClassExecutor, mysql.ErrPrepareMulti)
	ErrPrepareDDL      = dbterror.NewStd(terror.ClassExecutor, mysql.ErrPrepareDDL)
	ErrResultIsEmpty   = dbterror.NewStd(terror.ClassExecutor, mysql.ErrResultIsEmpty)
	ErrBuildExecutor   = dbterror.NewStd(terror.ClassExecutor, mysql.ErrBuildExecutor)
	ErrBatchInsertFail = dbterror.NewStd(terror.ClassExecutor, mysql.ErrBatchInsertFail)

	ErrCantCreateUserWithGrant     = dbterror.NewStd(terror.ClassExecutor, mysql.ErrCantCreateUserWithGrant)
	ErrPasswordNoMatch             = dbterror.NewStd(terror.ClassExecutor, mysql.ErrPasswordNoMatch)
	ErrCannotUser                  = dbterror.NewStd(terror.ClassExecutor, mysql.ErrCannotUser)
	ErrPasswordFormat              = dbterror.NewStd(terror.ClassExecutor, mysql.ErrPasswordFormat)
	ErrCantChangeTxCharacteristics = dbterror.NewStd(terror.ClassExecutor, mysql.ErrCantChangeTxCharacteristics)
	ErrPsManyParam                 = dbterror.NewStd(terror.ClassExecutor, mysql.ErrPsManyParam)
	ErrAdminCheckTable             = dbterror.NewStd(terror.ClassExecutor, mysql.ErrAdminCheckTable)
	ErrDBaccessDenied              = dbterror.NewStd(terror.ClassExecutor, mysql.ErrDBaccessDenied)
	ErrTableaccessDenied           = dbterror.NewStd(terror.ClassExecutor, mysql.ErrTableaccessDenied)
	ErrBadDB                       = dbterror.NewStd(terror.ClassExecutor, mysql.ErrBadDB)
	ErrWrongObject                 = dbterror.NewStd(terror.ClassExecutor, mysql.ErrWrongObject)
	ErrRoleNotGranted              = dbterror.NewStd(terror.ClassPrivilege, mysql.ErrRoleNotGranted)
	ErrDeadlock                    = dbterror.NewStd(terror.ClassExecutor, mysql.ErrLockDeadlock)
	ErrQueryInterrupted            = dbterror.NewStd(terror.ClassExecutor, mysql.ErrQueryInterrupted)

	ErrBRIEBackupFailed  = dbterror.NewStd(terror.ClassExecutor, mysql.ErrBRIEBackupFailed)
	ErrBRIERestoreFailed = dbterror.NewStd(terror.ClassExecutor, mysql.ErrBRIERestoreFailed)
	ErrBRIEImportFailed  = dbterror.NewStd(terror.ClassExecutor, mysql.ErrBRIEImportFailed)
	ErrBRIEExportFailed  = dbterror.NewStd(terror.ClassExecutor, mysql.ErrBRIEExportFailed)
)
