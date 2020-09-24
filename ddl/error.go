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

package ddl

import (
	"fmt"

	"github.com/pingcap/parser/terror"
	mysql "github.com/pingcap/tidb/errno"
	"github.com/pingcap/tidb/util/dbterror"
)

var (
	// errWorkerClosed means we have already closed the DDL worker.
	errInvalidWorker = dbterror.NewStd(terror.ClassDDL, mysql.ErrInvalidDDLWorker)
	// errNotOwner means we are not owner and can't handle DDL jobs.
	errNotOwner              = dbterror.NewStd(terror.ClassDDL, mysql.ErrNotOwner)
	errCantDecodeRecord      = dbterror.NewStd(terror.ClassDDL, mysql.ErrCantDecodeRecord)
	errInvalidDDLJob         = dbterror.NewStd(terror.ClassDDL, mysql.ErrInvalidDDLJob)
	errCancelledDDLJob       = dbterror.NewStd(terror.ClassDDL, mysql.ErrCancelledDDLJob)
	errFileNotFound          = dbterror.NewStd(terror.ClassDDL, mysql.ErrFileNotFound)
	errRunMultiSchemaChanges = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "multi schema change"))
	errWaitReorgTimeout      = dbterror.NewStd(terror.ClassDDL, mysql.ErrLockWaitTimeout)
	errInvalidStoreVer       = dbterror.NewStd(terror.ClassDDL, mysql.ErrInvalidStoreVersion)
	// ErrRepairTableFail is used to repair tableInfo in repair mode.
	ErrRepairTableFail = dbterror.NewStd(terror.ClassDDL, mysql.ErrRepairTable)

	// We don't support dropping column with index covered now.
	errCantDropColWithIndex                   = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "drop column with index"))
	errUnsupportedAddColumn                   = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "add column"))
	errUnsupportedModifyColumn                = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "modify column: %s"))
	errUnsupportedModifyCharset               = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "modify %s"))
	errUnsupportedModifyCollation             = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "modifying collation from %s to %s"))
	errUnsupportedPKHandle                    = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "drop integer primary key"))
	errUnsupportedCharset                     = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "charset %s and collate %s"))
	errUnsupportedShardRowIDBits              = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "shard_row_id_bits for table with primary key as row id"))
	errUnsupportedAlterTableWithValidation    = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, "ALTER TABLE WITH VALIDATION is currently unsupported")
	errUnsupportedAlterTableWithoutValidation = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, "ALTER TABLE WITHOUT VALIDATION is currently unsupported")
	errBlobKeyWithoutLength                   = dbterror.NewStd(terror.ClassDDL, mysql.ErrBlobKeyWithoutLength)
	errKeyPart0                               = dbterror.NewStd(terror.ClassDDL, mysql.ErrKeyPart0)
	errIncorrectPrefixKey                     = dbterror.NewStd(terror.ClassDDL, mysql.ErrWrongSubKey)
	errTooLongKey                             = dbterror.NewStd(terror.ClassDDL, mysql.ErrTooLongKey)
	errKeyColumnDoesNotExits                  = dbterror.NewStd(terror.ClassDDL, mysql.ErrKeyColumnDoesNotExits)
	errUnknownTypeLength                      = dbterror.NewStd(terror.ClassDDL, mysql.ErrUnknownTypeLength)
	errUnknownFractionLength                  = dbterror.NewStd(terror.ClassDDL, mysql.ErrUnknownFractionLength)
	errInvalidDDLJobVersion                   = dbterror.NewStd(terror.ClassDDL, mysql.ErrInvalidDDLJobVersion)
	errInvalidUseOfNull                       = dbterror.NewStd(terror.ClassDDL, mysql.ErrInvalidUseOfNull)
	errTooManyFields                          = dbterror.NewStd(terror.ClassDDL, mysql.ErrTooManyFields)
	errInvalidSplitRegionRanges               = dbterror.NewStd(terror.ClassDDL, mysql.ErrInvalidSplitRegionRanges)
	errReorgPanic                             = dbterror.NewStd(terror.ClassDDL, mysql.ErrReorgPanic)
	errFkColumnCannotDrop                     = dbterror.NewStd(terror.ClassDDL, mysql.ErrFkColumnCannotDrop)
	errFKIncompatibleColumns                  = dbterror.NewStd(terror.ClassDDL, mysql.ErrFKIncompatibleColumns)

	errOnlyOnRangeListPartition = dbterror.NewStd(terror.ClassDDL, mysql.ErrOnlyOnRangeListPartition)
	// errWrongKeyColumn is for table column cannot be indexed.
	errWrongKeyColumn = dbterror.NewStd(terror.ClassDDL, mysql.ErrWrongKeyColumn)
	// errWrongFKOptionForGeneratedColumn is for wrong foreign key reference option on generated columns.
	errWrongFKOptionForGeneratedColumn = dbterror.NewStd(terror.ClassDDL, mysql.ErrWrongFKOptionForGeneratedColumn)
	// ErrUnsupportedOnGeneratedColumn is for unsupported actions on generated columns.
	ErrUnsupportedOnGeneratedColumn = dbterror.NewStd(terror.ClassDDL, mysql.ErrUnsupportedOnGeneratedColumn)
	// errGeneratedColumnNonPrior forbids to refer generated column non prior to it.
	errGeneratedColumnNonPrior = dbterror.NewStd(terror.ClassDDL, mysql.ErrGeneratedColumnNonPrior)
	// errDependentByGeneratedColumn forbids to delete columns which are dependent by generated columns.
	errDependentByGeneratedColumn = dbterror.NewStd(terror.ClassDDL, mysql.ErrDependentByGeneratedColumn)
	// errJSONUsedAsKey forbids to use JSON as key or index.
	errJSONUsedAsKey = dbterror.NewStd(terror.ClassDDL, mysql.ErrJSONUsedAsKey)
	// errBlobCantHaveDefault forbids to give not null default value to TEXT/BLOB/JSON.
	errBlobCantHaveDefault = dbterror.NewStd(terror.ClassDDL, mysql.ErrBlobCantHaveDefault)
	errTooLongIndexComment = dbterror.NewStd(terror.ClassDDL, mysql.ErrTooLongIndexComment)
	// ErrInvalidDefaultValue returns for invalid default value for columns.
	ErrInvalidDefaultValue = dbterror.NewStd(terror.ClassDDL, mysql.ErrInvalidDefault)
	// ErrGeneratedColumnRefAutoInc forbids to refer generated columns to auto-increment columns .
	ErrGeneratedColumnRefAutoInc = dbterror.NewStd(terror.ClassDDL, mysql.ErrGeneratedColumnRefAutoInc)
	// ErrUnsupportedAddPartition returns for does not support add partitions.
	ErrUnsupportedAddPartition = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "add partitions"))
	// ErrUnsupportedCoalescePartition returns for does not support coalesce partitions.
	ErrUnsupportedCoalescePartition   = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "coalesce partitions"))
	errUnsupportedReorganizePartition = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "reorganize partition"))
	errUnsupportedCheckPartition      = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "check partition"))
	errUnsupportedOptimizePartition   = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "optimize partition"))
	errUnsupportedRebuildPartition    = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "rebuild partition"))
	errUnsupportedRemovePartition     = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "remove partitioning"))
	errUnsupportedRepairPartition     = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "repair partition"))
	errUnsupportedExchangePartition   = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "exchange partition"))
	// ErrGeneratedColumnFunctionIsNotAllowed returns for unsupported functions for generated columns.
	ErrGeneratedColumnFunctionIsNotAllowed = dbterror.NewStd(terror.ClassDDL, mysql.ErrGeneratedColumnFunctionIsNotAllowed)
	// ErrUnsupportedPartitionByRangeColumns returns for does unsupported partition by range columns.
	ErrUnsupportedPartitionByRangeColumns = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "partition by range columns"))
	errUnsupportedCreatePartition         = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "partition type, treat as normal table"))
	errTablePartitionDisabled             = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, "Partitions are ignored because Table Partition is disabled, please set 'tidb_enable_table_partition' if you need to need to enable it")
	errUnsupportedIndexType               = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "index type"))

	// ErrDupKeyName returns for duplicated key name
	ErrDupKeyName = dbterror.NewStd(terror.ClassDDL, mysql.ErrDupKeyName)
	// ErrInvalidDDLState returns for invalid ddl model object state.
	ErrInvalidDDLState = dbterror.New(terror.ClassDDL, mysql.ErrInvalidDDLState, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrInvalidDDLState]))
	// ErrUnsupportedModifyPrimaryKey returns an error when add or drop the primary key.
	// It's exported for testing.
	ErrUnsupportedModifyPrimaryKey = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "%s primary key"))
	// ErrPKIndexCantBeInvisible return an error when primary key is invisible index
	ErrPKIndexCantBeInvisible = dbterror.NewStd(terror.ClassDDL, mysql.ErrPKIndexCantBeInvisible)

	// ErrColumnBadNull returns for a bad null value.
	ErrColumnBadNull = dbterror.NewStd(terror.ClassDDL, mysql.ErrBadNull)
	// ErrBadField forbids to refer to unknown column.
	ErrBadField = dbterror.NewStd(terror.ClassDDL, mysql.ErrBadField)
	// ErrCantRemoveAllFields returns for deleting all columns.
	ErrCantRemoveAllFields = dbterror.NewStd(terror.ClassDDL, mysql.ErrCantRemoveAllFields)
	// ErrCantDropFieldOrKey returns for dropping a non-existent field or key.
	ErrCantDropFieldOrKey = dbterror.NewStd(terror.ClassDDL, mysql.ErrCantDropFieldOrKey)
	// ErrInvalidOnUpdate returns for invalid ON UPDATE clause.
	ErrInvalidOnUpdate = dbterror.NewStd(terror.ClassDDL, mysql.ErrInvalidOnUpdate)
	// ErrTooLongIdent returns for too long name of database/table/column/index.
	ErrTooLongIdent = dbterror.NewStd(terror.ClassDDL, mysql.ErrTooLongIdent)
	// ErrWrongDBName returns for wrong database name.
	ErrWrongDBName = dbterror.NewStd(terror.ClassDDL, mysql.ErrWrongDBName)
	// ErrWrongTableName returns for wrong table name.
	ErrWrongTableName = dbterror.NewStd(terror.ClassDDL, mysql.ErrWrongTableName)
	// ErrWrongColumnName returns for wrong column name.
	ErrWrongColumnName = dbterror.NewStd(terror.ClassDDL, mysql.ErrWrongColumnName)
	// ErrInvalidGroupFuncUse returns for using invalid group functions.
	ErrInvalidGroupFuncUse = dbterror.NewStd(terror.ClassDDL, mysql.ErrInvalidGroupFuncUse)
	// ErrTableMustHaveColumns returns for missing column when creating a table.
	ErrTableMustHaveColumns = dbterror.NewStd(terror.ClassDDL, mysql.ErrTableMustHaveColumns)
	// ErrWrongNameForIndex returns for wrong index name.
	ErrWrongNameForIndex = dbterror.NewStd(terror.ClassDDL, mysql.ErrWrongNameForIndex)
	// ErrUnknownCharacterSet returns unknown character set.
	ErrUnknownCharacterSet = dbterror.NewStd(terror.ClassDDL, mysql.ErrUnknownCharacterSet)
	// ErrUnknownCollation returns unknown collation.
	ErrUnknownCollation = dbterror.NewStd(terror.ClassDDL, mysql.ErrUnknownCollation)
	// ErrCollationCharsetMismatch returns when collation not match the charset.
	ErrCollationCharsetMismatch = dbterror.NewStd(terror.ClassDDL, mysql.ErrCollationCharsetMismatch)
	// ErrConflictingDeclarations return conflict declarations.
	ErrConflictingDeclarations = dbterror.New(terror.ClassDDL, mysql.ErrConflictingDeclarations, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrConflictingDeclarations], "CHARACTER SET ", "%s", "CHARACTER SET ", "%s"))
	// ErrPrimaryCantHaveNull returns All parts of a PRIMARY KEY must be NOT NULL; if you need NULL in a key, use UNIQUE instead
	ErrPrimaryCantHaveNull = dbterror.NewStd(terror.ClassDDL, mysql.ErrPrimaryCantHaveNull)
	// ErrErrorOnRename returns error for wrong database name in alter table rename
	ErrErrorOnRename = dbterror.NewStd(terror.ClassDDL, mysql.ErrErrorOnRename)
	// ErrViewSelectClause returns error for create view with select into clause
	ErrViewSelectClause = dbterror.NewStd(terror.ClassDDL, mysql.ErrViewSelectClause)

	// ErrNotAllowedTypeInPartition returns not allowed type error when creating table partition with unsupported expression type.
	ErrNotAllowedTypeInPartition = dbterror.NewStd(terror.ClassDDL, mysql.ErrFieldTypeNotAllowedAsPartitionField)
	// ErrPartitionMgmtOnNonpartitioned returns it's not a partition table.
	ErrPartitionMgmtOnNonpartitioned = dbterror.NewStd(terror.ClassDDL, mysql.ErrPartitionMgmtOnNonpartitioned)
	// ErrDropPartitionNonExistent returns error in list of partition.
	ErrDropPartitionNonExistent = dbterror.NewStd(terror.ClassDDL, mysql.ErrDropPartitionNonExistent)
	// ErrSameNamePartition returns duplicate partition name.
	ErrSameNamePartition = dbterror.NewStd(terror.ClassDDL, mysql.ErrSameNamePartition)
	// ErrRangeNotIncreasing returns values less than value must be strictly increasing for each partition.
	ErrRangeNotIncreasing = dbterror.NewStd(terror.ClassDDL, mysql.ErrRangeNotIncreasing)
	// ErrPartitionMaxvalue returns maxvalue can only be used in last partition definition.
	ErrPartitionMaxvalue = dbterror.NewStd(terror.ClassDDL, mysql.ErrPartitionMaxvalue)
	// ErrDropLastPartition returns cannot remove all partitions, use drop table instead.
	ErrDropLastPartition = dbterror.NewStd(terror.ClassDDL, mysql.ErrDropLastPartition)
	// ErrTooManyPartitions returns too many partitions were defined.
	ErrTooManyPartitions = dbterror.NewStd(terror.ClassDDL, mysql.ErrTooManyPartitions)
	// ErrPartitionFunctionIsNotAllowed returns this partition function is not allowed.
	ErrPartitionFunctionIsNotAllowed = dbterror.NewStd(terror.ClassDDL, mysql.ErrPartitionFunctionIsNotAllowed)
	// ErrPartitionFuncNotAllowed returns partition function returns the wrong type.
	ErrPartitionFuncNotAllowed = dbterror.NewStd(terror.ClassDDL, mysql.ErrPartitionFuncNotAllowed)
	// ErrUniqueKeyNeedAllFieldsInPf returns must include all columns in the table's partitioning function.
	ErrUniqueKeyNeedAllFieldsInPf = dbterror.NewStd(terror.ClassDDL, mysql.ErrUniqueKeyNeedAllFieldsInPf)
	errWrongExprInPartitionFunc   = dbterror.NewStd(terror.ClassDDL, mysql.ErrWrongExprInPartitionFunc)
	// ErrWarnDataTruncated returns data truncated error.
	ErrWarnDataTruncated = dbterror.NewStd(terror.ClassDDL, mysql.WarnDataTruncated)
	// ErrCoalesceOnlyOnHashPartition returns coalesce partition can only be used on hash/key partitions.
	ErrCoalesceOnlyOnHashPartition = dbterror.NewStd(terror.ClassDDL, mysql.ErrCoalesceOnlyOnHashPartition)
	// ErrViewWrongList returns create view must include all columns in the select clause
	ErrViewWrongList = dbterror.NewStd(terror.ClassDDL, mysql.ErrViewWrongList)
	// ErrAlterOperationNotSupported returns when alter operations is not supported.
	ErrAlterOperationNotSupported = dbterror.NewStd(terror.ClassDDL, mysql.ErrAlterOperationNotSupportedReason)
	// ErrWrongObject returns for wrong object.
	ErrWrongObject = dbterror.NewStd(terror.ClassDDL, mysql.ErrWrongObject)
	// ErrTableCantHandleFt returns FULLTEXT keys are not supported by table type
	ErrTableCantHandleFt = dbterror.NewStd(terror.ClassDDL, mysql.ErrTableCantHandleFt)
	// ErrFieldNotFoundPart returns an error when 'partition by columns' are not found in table columns.
	ErrFieldNotFoundPart = dbterror.NewStd(terror.ClassDDL, mysql.ErrFieldNotFoundPart)
	// ErrWrongTypeColumnValue returns 'Partition column values of incorrect type'
	ErrWrongTypeColumnValue = dbterror.NewStd(terror.ClassDDL, mysql.ErrWrongTypeColumnValue)
	// ErrFunctionalIndexPrimaryKey returns 'The primary key cannot be a functional index'
	ErrFunctionalIndexPrimaryKey = dbterror.NewStd(terror.ClassDDL, mysql.ErrFunctionalIndexPrimaryKey)
	// ErrFunctionalIndexOnField returns 'Functional index on a column is not supported. Consider using a regular index instead'
	ErrFunctionalIndexOnField = dbterror.NewStd(terror.ClassDDL, mysql.ErrFunctionalIndexOnField)
	// ErrInvalidAutoRandom returns when auto_random is used incorrectly.
	ErrInvalidAutoRandom = dbterror.NewStd(terror.ClassDDL, mysql.ErrInvalidAutoRandom)
	// ErrUnsupportedConstraintCheck returns when use ADD CONSTRAINT CHECK
	ErrUnsupportedConstraintCheck = dbterror.NewStd(terror.ClassDDL, mysql.ErrUnsupportedConstraintCheck)

	// ErrSequenceRunOut returns when the sequence has been run out.
	ErrSequenceRunOut = dbterror.NewStd(terror.ClassDDL, mysql.ErrSequenceRunOut)
	// ErrSequenceInvalidData returns when sequence values are conflicting.
	ErrSequenceInvalidData = dbterror.NewStd(terror.ClassDDL, mysql.ErrSequenceInvalidData)
	// ErrSequenceAccessFail returns when sequences are not able to access.
	ErrSequenceAccessFail = dbterror.NewStd(terror.ClassDDL, mysql.ErrSequenceAccessFail)
	// ErrNotSequence returns when object is not a sequence.
	ErrNotSequence = dbterror.NewStd(terror.ClassDDL, mysql.ErrNotSequence)
	// ErrUnknownSequence returns when drop / alter unknown sequence.
	ErrUnknownSequence = dbterror.NewStd(terror.ClassDDL, mysql.ErrUnknownSequence)
	// ErrSequenceUnsupportedTableOption returns when unsupported table option exists in sequence.
	ErrSequenceUnsupportedTableOption = dbterror.NewStd(terror.ClassDDL, mysql.ErrSequenceUnsupportedTableOption)
	// ErrColumnTypeUnsupportedNextValue is returned when sequence next value is assigned to unsupported column type.
	ErrColumnTypeUnsupportedNextValue = dbterror.NewStd(terror.ClassDDL, mysql.ErrColumnTypeUnsupportedNextValue)
	// ErrAddColumnWithSequenceAsDefault is returned when the new added column with sequence's nextval as it's default value.
	ErrAddColumnWithSequenceAsDefault = dbterror.NewStd(terror.ClassDDL, mysql.ErrAddColumnWithSequenceAsDefault)
	// ErrUnsupportedExpressionIndex is returned when create an expression index without allow-expression-index.
	ErrUnsupportedExpressionIndex = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "creating expression index without allow-expression-index in config"))
	// ErrPartitionExchangePartTable is returned when exchange table partition with another table is partitioned.
	ErrPartitionExchangePartTable = dbterror.NewStd(terror.ClassDDL, mysql.ErrPartitionExchangePartTable)
	// ErrTablesDifferentMetadata is returned when exchanges tables is not compatible.
	ErrTablesDifferentMetadata = dbterror.NewStd(terror.ClassDDL, mysql.ErrTablesDifferentMetadata)
	// ErrRowDoesNotMatchPartition is returned when the row record of exchange table does not match the partition rule.
	ErrRowDoesNotMatchPartition = dbterror.NewStd(terror.ClassDDL, mysql.ErrRowDoesNotMatchPartition)
	// ErrPartitionExchangeForeignKey is returned when exchanged normal table has foreign keys.
	ErrPartitionExchangeForeignKey = dbterror.NewStd(terror.ClassDDL, mysql.ErrPartitionExchangeForeignKey)
	// ErrCheckNoSuchTable is returned when exchanged normal table is view or sequence.
	ErrCheckNoSuchTable         = dbterror.NewStd(terror.ClassDDL, mysql.ErrCheckNoSuchTable)
	errUnsupportedPartitionType = dbterror.New(terror.ClassDDL, mysql.ErrUnsupportedDDLOperation, fmt.Sprintf(mysql.MySQLErrName[mysql.ErrUnsupportedDDLOperation], "partition type of table %s when exchanging partition"))
	// ErrPartitionExchangeDifferentOption is returned when attribute does not match between partition table and normal table.
	ErrPartitionExchangeDifferentOption = dbterror.NewStd(terror.ClassDDL, mysql.ErrPartitionExchangeDifferentOption)
	// ErrTableOptionUnionUnsupported is returned when create/alter table with union option.
	ErrTableOptionUnionUnsupported = dbterror.NewStd(terror.ClassDDL, mysql.ErrTableOptionUnionUnsupported)
	// ErrTableOptionInsertMethodUnsupported is returned when create/alter table with insert method option.
	ErrTableOptionInsertMethodUnsupported = dbterror.NewStd(terror.ClassDDL, mysql.ErrTableOptionInsertMethodUnsupported)

	// ErrInvalidPlacementSpec is returned when add/alter an invalid placement rule
	ErrInvalidPlacementSpec = dbterror.NewStd(terror.ClassDDL, mysql.ErrInvalidPlacementSpec)
)
