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

package infoschema

import (
	"github.com/pingcap/parser/terror"
	mysql "github.com/pingcap/tidb/errno"
	"github.com/pingcap/tidb/util/dbterror"
)

var (
	// ErrDatabaseExists returns for database already exists.
	ErrDatabaseExists = dbterror.NewStd(terror.ClassSchema, mysql.ErrDBCreateExists)
	// ErrDatabaseDropExists returns for dropping a non-existent database.
	ErrDatabaseDropExists = dbterror.NewStd(terror.ClassSchema, mysql.ErrDBDropExists)
	// ErrAccessDenied return when the user doesn't have the permission to access the table.
	ErrAccessDenied = dbterror.NewStd(terror.ClassSchema, mysql.ErrAccessDenied)
	// ErrDatabaseNotExists returns for database not exists.
	ErrDatabaseNotExists = dbterror.NewStd(terror.ClassSchema, mysql.ErrBadDB)
	// ErrTableExists returns for table already exists.
	ErrTableExists = dbterror.NewStd(terror.ClassSchema, mysql.ErrTableExists)
	// ErrTableDropExists returns for dropping a non-existent table.
	ErrTableDropExists = dbterror.NewStd(terror.ClassSchema, mysql.ErrBadTable)
	// ErrSequenceDropExists returns for dropping a non-exist sequence.
	ErrSequenceDropExists = dbterror.NewStd(terror.ClassSchema, mysql.ErrUnknownSequence)
	// ErrColumnNotExists returns for column not exists.
	ErrColumnNotExists = dbterror.NewStd(terror.ClassSchema, mysql.ErrBadField)
	// ErrColumnExists returns for column already exists.
	ErrColumnExists = dbterror.NewStd(terror.ClassSchema, mysql.ErrDupFieldName)
	// ErrKeyNameDuplicate returns for index duplicate when rename index.
	ErrKeyNameDuplicate = dbterror.NewStd(terror.ClassSchema, mysql.ErrDupKeyName)
	// ErrNonuniqTable returns when none unique tables errors.
	ErrNonuniqTable = dbterror.NewStd(terror.ClassSchema, mysql.ErrNonuniqTable)
	// ErrMultiplePriKey returns for multiple primary keys.
	ErrMultiplePriKey = dbterror.NewStd(terror.ClassSchema, mysql.ErrMultiplePriKey)
	// ErrTooManyKeyParts returns for too many key parts.
	ErrTooManyKeyParts = dbterror.NewStd(terror.ClassSchema, mysql.ErrTooManyKeyParts)
	// ErrForeignKeyNotExists returns for foreign key not exists.
	ErrForeignKeyNotExists = dbterror.NewStd(terror.ClassSchema, mysql.ErrCantDropFieldOrKey)
	// ErrTableNotLockedForWrite returns for write tables when only hold the table read lock.
	ErrTableNotLockedForWrite = dbterror.NewStd(terror.ClassSchema, mysql.ErrTableNotLockedForWrite)
	// ErrTableNotLocked returns when session has explicitly lock tables, then visit unlocked table will return this error.
	ErrTableNotLocked = dbterror.NewStd(terror.ClassSchema, mysql.ErrTableNotLocked)
	// ErrTableNotExists returns for table not exists.
	ErrTableNotExists = dbterror.NewStd(terror.ClassSchema, mysql.ErrNoSuchTable)
	// ErrKeyNotExists returns for index not exists.
	ErrKeyNotExists = dbterror.NewStd(terror.ClassSchema, mysql.ErrKeyDoesNotExist)
	// ErrCannotAddForeign returns for foreign key exists.
	ErrCannotAddForeign = dbterror.NewStd(terror.ClassSchema, mysql.ErrCannotAddForeign)
	// ErrForeignKeyNotMatch returns for foreign key not match.
	ErrForeignKeyNotMatch = dbterror.NewStd(terror.ClassSchema, mysql.ErrWrongFkDef)
	// ErrIndexExists returns for index already exists.
	ErrIndexExists = dbterror.NewStd(terror.ClassSchema, mysql.ErrDupIndex)
	// ErrUserDropExists returns for dropping a non-existent user.
	ErrUserDropExists = dbterror.NewStd(terror.ClassSchema, mysql.ErrBadUser)
	// ErrUserAlreadyExists return for creating a existent user.
	ErrUserAlreadyExists = dbterror.NewStd(terror.ClassSchema, mysql.ErrUserAlreadyExists)
	// ErrTableLocked returns when the table was locked by other session.
	ErrTableLocked = dbterror.NewStd(terror.ClassSchema, mysql.ErrTableLocked)
	// ErrWrongObject returns when the table/view/sequence is not the expected object.
	ErrWrongObject = dbterror.NewStd(terror.ClassSchema, mysql.ErrWrongObject)
)
