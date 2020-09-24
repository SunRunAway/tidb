// Copyright 2016 PingCAP, Inc.
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

package tikv

import (
	"github.com/pingcap/errors"
	"github.com/pingcap/kvproto/pkg/kvrpcpb"
	"github.com/pingcap/kvproto/pkg/pdpb"
	"github.com/pingcap/parser/terror"
	mysql "github.com/pingcap/tidb/errno"
	"github.com/pingcap/tidb/util/dbterror"
)

var (
	// ErrBodyMissing response body is missing error
	ErrBodyMissing = errors.New("response body is missing")
	// When TiDB is closing and send request to tikv fail, do not retry, return this error.
	errTiDBShuttingDown = errors.New("tidb server shutting down")
)

// mismatchClusterID represents the message that the cluster ID of the PD client does not match the PD.
const mismatchClusterID = "mismatch cluster id"

// MySQL error instances.
var (
	ErrTiKVServerTimeout           = dbterror.NewStd(terror.ClassTiKV, mysql.ErrTiKVServerTimeout)
	ErrResolveLockTimeout          = dbterror.NewStd(terror.ClassTiKV, mysql.ErrResolveLockTimeout)
	ErrPDServerTimeout             = dbterror.NewStd(terror.ClassTiKV, mysql.ErrPDServerTimeout)
	ErrRegionUnavailable           = dbterror.NewStd(terror.ClassTiKV, mysql.ErrRegionUnavailable)
	ErrTiKVServerBusy              = dbterror.NewStd(terror.ClassTiKV, mysql.ErrTiKVServerBusy)
	ErrTiKVStaleCommand            = dbterror.NewStd(terror.ClassTiKV, mysql.ErrTiKVStaleCommand)
	ErrTiKVMaxTimestampNotSynced   = dbterror.NewStd(terror.ClassTiKV, mysql.ErrTiKVMaxTimestampNotSynced)
	ErrGCTooEarly                  = dbterror.NewStd(terror.ClassTiKV, mysql.ErrGCTooEarly)
	ErrQueryInterrupted            = dbterror.NewStd(terror.ClassTiKV, mysql.ErrQueryInterrupted)
	ErrLockAcquireFailAndNoWaitSet = dbterror.NewStd(terror.ClassTiKV, mysql.ErrLockAcquireFailAndNoWaitSet)
	ErrLockWaitTimeout             = dbterror.NewStd(terror.ClassTiKV, mysql.ErrLockWaitTimeout)
	ErrTokenLimit                  = dbterror.NewStd(terror.ClassTiKV, mysql.ErrTiKVStoreLimit)
	ErrLockExpire                  = dbterror.NewStd(terror.ClassTiKV, mysql.ErrLockExpire)
	ErrUnknown                     = dbterror.NewStd(terror.ClassTiKV, mysql.ErrUnknown)
)

// Registers error returned from TiKV.
var (
	_ = dbterror.NewStd(terror.ClassTiKV, mysql.ErrDataOutOfRange)
	_ = dbterror.NewStd(terror.ClassTiKV, mysql.ErrTruncatedWrongValue)
	_ = dbterror.NewStd(terror.ClassTiKV, mysql.ErrDivisionByZero)
)

// ErrDeadlock wraps *kvrpcpb.Deadlock to implement the error interface.
// It also marks if the deadlock is retryable.
type ErrDeadlock struct {
	*kvrpcpb.Deadlock
	IsRetryable bool
}

func (d *ErrDeadlock) Error() string {
	return d.Deadlock.String()
}

// PDError wraps *pdpb.Error to implement the error interface.
type PDError struct {
	Err *pdpb.Error
}

func (d *PDError) Error() string {
	return d.Err.String()
}
