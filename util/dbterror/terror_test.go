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

package dbterror_test

import (
	"testing"

	"github.com/pingcap/parser/terror"
	"github.com/pingcap/tidb/config"
	mysql "github.com/pingcap/tidb/errno"
	"github.com/pingcap/tidb/kv"
	"github.com/pingcap/tidb/util/dbterror"
	"github.com/stretchr/testify/assert"
)

func TestRedact(t *testing.T) {
	errorsNeedRedact := []struct {
		error *dbterror.Error
		pos   []int
	}{
		{kv.ErrKeyExists, []int{0, 1}},
	}

	oriCfg := config.GetGlobalConfig()
	defer config.StoreGlobalConfig(oriCfg)

	newCfg := *oriCfg
	newCfg.EnableRedactLog = 0
	config.StoreGlobalConfig(&newCfg)
	for _, cs := range errorsNeedRedact {
		args := make([]interface{}, len(cs.pos))
		for i := range args {
			args[i] = "sensitiveData"
		}
		err := cs.error.FastGenByArgs(args...)
		assert.Contains(t, err.Error(), "sensitiveData")
	}

	newCfg = *oriCfg
	newCfg.EnableRedactLog = 1
	config.StoreGlobalConfig(&newCfg)
	for _, cs := range errorsNeedRedact {
		args := make([]interface{}, len(cs.pos))
		for i := range args {
			args[i] = "sensitiveData"
		}
		err := cs.error.FastGenByArgs(args...)
		assert.NotContains(t, err.Error(), "sensitiveData")
		assert.Contains(t, err.Error(), "?")
	}
}

func TestTError(t *testing.T) {
	oriCfg := config.GetGlobalConfig()
	defer config.StoreGlobalConfig(oriCfg)

	newCfg := *oriCfg
	newCfg.EnableRedactLog = 0
	config.StoreGlobalConfig(&newCfg)

	e := dbterror.NewStd(terror.ClassKV, mysql.ErrDupEntry)
	err := e.FastGenByArgs("sensitive", "data")
	assert.Contains(t, err.Error(), "sensitive")
	assert.Contains(t, err.Error(), "data")
	err = e.GenWithStackByArgs("sensitive", "data")
	assert.Contains(t, err.Error(), "sensitive")
	assert.Contains(t, err.Error(), "data")

	newCfg = *oriCfg
	newCfg.EnableRedactLog = 1
	config.StoreGlobalConfig(&newCfg)

	err = e.FastGenByArgs("sensitive", "data")
	assert.NotContains(t, err.Error(), "sensitive")
	assert.NotContains(t, err.Error(), "data")
	assert.Contains(t, err.Error(), "?")
	err = e.GenWithStackByArgs("sensitive", "data")
	assert.NotContains(t, err.Error(), "sensitive")
	assert.NotContains(t, err.Error(), "data")
	assert.Contains(t, err.Error(), "?")

	assert.True(t, e.Equal(terror.ClassKV.New(mysql.ErrDupEntry, mysql.MySQLErrName[mysql.ErrDupEntry])))
}
