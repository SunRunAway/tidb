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

package types

import (
	"github.com/pingcap/parser/terror"
	parser_types "github.com/pingcap/parser/types"
	mysql "github.com/pingcap/tidb/errno"
	"github.com/pingcap/tidb/util/dbterror"
)

// const strings for ErrWrongValue
const (
	DateTimeStr = "datetime"
	TimeStr     = "time"
)

var (
	// ErrInvalidDefault is returned when meet a invalid default value.
	ErrInvalidDefault = parser_types.ErrInvalidDefault
	// ErrDataTooLong is returned when converts a string value that is longer than field type length.
	ErrDataTooLong = dbterror.NewStd(terror.ClassTypes, mysql.ErrDataTooLong)
	// ErrIllegalValueForType is returned when value of type is illegal.
	ErrIllegalValueForType = dbterror.NewStd(terror.ClassTypes, mysql.ErrIllegalValueForType)
	// ErrTruncated is returned when data has been truncated during conversion.
	ErrTruncated = dbterror.NewStd(terror.ClassTypes, mysql.WarnDataTruncated)
	// ErrOverflow is returned when data is out of range for a field type.
	ErrOverflow = dbterror.NewStd(terror.ClassTypes, mysql.ErrDataOutOfRange)
	// ErrDivByZero is return when do division by 0.
	ErrDivByZero = dbterror.NewStd(terror.ClassTypes, mysql.ErrDivisionByZero)
	// ErrTooBigDisplayWidth is return when display width out of range for column.
	ErrTooBigDisplayWidth = dbterror.NewStd(terror.ClassTypes, mysql.ErrTooBigDisplaywidth)
	// ErrTooBigFieldLength is return when column length too big for column.
	ErrTooBigFieldLength = dbterror.NewStd(terror.ClassTypes, mysql.ErrTooBigFieldlength)
	// ErrTooBigSet is returned when too many strings for column.
	ErrTooBigSet = dbterror.NewStd(terror.ClassTypes, mysql.ErrTooBigSet)
	// ErrTooBigScale is returned when type DECIMAL/NUMERIC scale is bigger than mysql.MaxDecimalScale.
	ErrTooBigScale = dbterror.NewStd(terror.ClassTypes, mysql.ErrTooBigScale)
	// ErrTooBigPrecision is returned when type DECIMAL/NUMERIC precision is bigger than mysql.MaxDecimalWidth
	ErrTooBigPrecision = dbterror.NewStd(terror.ClassTypes, mysql.ErrTooBigPrecision)
	// ErrBadNumber is return when parsing an invalid binary decimal number.
	ErrBadNumber = dbterror.NewStd(terror.ClassTypes, mysql.ErrBadNumber)
	// ErrInvalidFieldSize is returned when the precision of a column is out of range.
	ErrInvalidFieldSize = dbterror.NewStd(terror.ClassTypes, mysql.ErrInvalidFieldSize)
	// ErrMBiggerThanD is returned when precision less than the scale.
	ErrMBiggerThanD = dbterror.NewStd(terror.ClassTypes, mysql.ErrMBiggerThanD)
	// ErrWarnDataOutOfRange is returned when the value in a numeric column that is outside the permissible range of the column data type.
	// See https://dev.mysql.com/doc/refman/5.5/en/out-of-range-and-overflow.html for details
	ErrWarnDataOutOfRange = dbterror.NewStd(terror.ClassTypes, mysql.ErrWarnDataOutOfRange)
	// ErrDuplicatedValueInType is returned when enum column has duplicated value.
	ErrDuplicatedValueInType = dbterror.NewStd(terror.ClassTypes, mysql.ErrDuplicatedValueInType)
	// ErrDatetimeFunctionOverflow is returned when the calculation in datetime function cause overflow.
	ErrDatetimeFunctionOverflow = dbterror.NewStd(terror.ClassTypes, mysql.ErrDatetimeFunctionOverflow)
	// ErrCastAsSignedOverflow is returned when positive out-of-range integer, and convert to it's negative complement.
	ErrCastAsSignedOverflow = dbterror.NewStd(terror.ClassTypes, mysql.ErrCastAsSignedOverflow)
	// ErrCastNegIntAsUnsigned is returned when a negative integer be casted to an unsigned int.
	ErrCastNegIntAsUnsigned = dbterror.NewStd(terror.ClassTypes, mysql.ErrCastNegIntAsUnsigned)
	// ErrInvalidYearFormat is returned when the input is not a valid year format.
	ErrInvalidYearFormat = dbterror.NewStd(terror.ClassTypes, mysql.ErrInvalidYearFormat)
	// ErrInvalidYear is returned when the input value is not a valid year.
	ErrInvalidYear = dbterror.NewStd(terror.ClassTypes, mysql.ErrInvalidYear)
	// ErrTruncatedWrongVal is returned when data has been truncated during conversion.
	ErrTruncatedWrongVal = dbterror.NewStd(terror.ClassTypes, mysql.ErrTruncatedWrongValue)
	// ErrInvalidWeekModeFormat is returned when the week mode is wrong.
	ErrInvalidWeekModeFormat = dbterror.NewStd(terror.ClassTypes, mysql.ErrInvalidWeekModeFormat)
	// ErrWrongValue is returned when the input value is in wrong format.
	ErrWrongValue = dbterror.NewStd(terror.ClassTypes, mysql.ErrTruncatedWrongValue)
)
