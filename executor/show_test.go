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

package executor_test

import (
	"fmt"
	"strconv"
	"time"

	. "github.com/pingcap/check"
	"github.com/pingcap/errors"
	"github.com/pingcap/parser/auth"
	"github.com/pingcap/parser/model"
	"github.com/pingcap/parser/mysql"
	"github.com/pingcap/tidb/domain"
	"github.com/pingcap/tidb/meta/autoid"
	plannercore "github.com/pingcap/tidb/planner/core"
	"github.com/pingcap/tidb/session"
	"github.com/pingcap/tidb/sessionctx"
	"github.com/pingcap/tidb/sessionctx/variable"
	"github.com/pingcap/tidb/util/testkit"
	"github.com/pingcap/tidb/util/testutil"
	"golang.org/x/net/context"
)

func (s *testSuite) TestShow(c *C) {
	tk := testkit.NewTestKit(c, s.store)
	tk.MustExec("use test")

	testSQL := `drop table if exists show_test`
	tk.MustExec(testSQL)
	testSQL = `create table SHOW_test (id int PRIMARY KEY AUTO_INCREMENT, c1 int comment "c1_comment", c2 int, c3 int default 1, c4 text, c5 boolean, key idx_wide_c4(c3, c4(10))) ENGINE=InnoDB AUTO_INCREMENT=28934 DEFAULT CHARSET=utf8 COMMENT "table_comment";`
	tk.MustExec(testSQL)

	testSQL = "show columns from show_test;"
	result := tk.MustQuery(testSQL)
	c.Check(result.Rows(), HasLen, 6)

	testSQL = "show create table show_test;"
	result = tk.MustQuery(testSQL)
	c.Check(result.Rows(), HasLen, 1)
	row := result.Rows()[0]
	// For issue https://github.com/pingcap/tidb/issues/1061
	expectedRow := []interface{}{
		"SHOW_test", "CREATE TABLE `SHOW_test` (\n  `id` int(11) NOT NULL AUTO_INCREMENT,\n  `c1` int(11) DEFAULT NULL COMMENT 'c1_comment',\n  `c2` int(11) DEFAULT NULL,\n  `c3` int(11) DEFAULT '1',\n  `c4` text DEFAULT NULL,\n  `c5` tinyint(1) DEFAULT NULL,\n  PRIMARY KEY (`id`),\n  KEY `idx_wide_c4` (`c3`,`c4`(10))\n) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=28934 COMMENT='table_comment'"}
	for i, r := range row {
		c.Check(r, Equals, expectedRow[i])
	}

	// For issue https://github.com/pingcap/tidb/issues/1918
	testSQL = `create table ptest(
		a int primary key,
		b double NOT NULL DEFAULT 2.0,
		c varchar(10) NOT NULL,
		d time unique,
		e timestamp NULL,
		f timestamp
	);`
	tk.MustExec(testSQL)
	testSQL = "show create table ptest;"
	result = tk.MustQuery(testSQL)
	c.Check(result.Rows(), HasLen, 1)
	row = result.Rows()[0]
	expectedRow = []interface{}{
		"ptest", "CREATE TABLE `ptest` (\n  `a` int(11) NOT NULL,\n  `b` double NOT NULL DEFAULT '2.0',\n  `c` varchar(10) NOT NULL,\n  `d` time DEFAULT NULL,\n  `e` timestamp NULL DEFAULT NULL,\n  `f` timestamp NULL DEFAULT NULL,\n  PRIMARY KEY (`a`),\n  UNIQUE KEY `d` (`d`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin"}
	for i, r := range row {
		c.Check(r, Equals, expectedRow[i])
	}

	// Issue #4684.
	tk.MustExec("drop table if exists `t1`")
	testSQL = "create table `t1` (" +
		"`c1` tinyint unsigned default null," +
		"`c2` smallint unsigned default null," +
		"`c3` mediumint unsigned default null," +
		"`c4` int unsigned default null," +
		"`c5` bigint unsigned default null);`"

	tk.MustExec(testSQL)
	testSQL = "show create table t1"
	result = tk.MustQuery(testSQL)
	c.Check(result.Rows(), HasLen, 1)
	row = result.Rows()[0]
	expectedRow = []interface{}{
		"t1", "CREATE TABLE `t1` (\n" +
			"  `c1` tinyint(3) UNSIGNED DEFAULT NULL,\n" +
			"  `c2` smallint(5) UNSIGNED DEFAULT NULL,\n" +
			"  `c3` mediumint(8) UNSIGNED DEFAULT NULL,\n" +
			"  `c4` int(10) UNSIGNED DEFAULT NULL,\n" +
			"  `c5` bigint(20) UNSIGNED DEFAULT NULL\n" +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin"}
	for i, r := range row {
		c.Check(r, Equals, expectedRow[i])
	}

	// Issue #7665
	tk.MustExec("drop table if exists `decimalschema`")
	testSQL = "create table `decimalschema` (`c1` decimal);"
	tk.MustExec(testSQL)
	testSQL = "show create table decimalschema"
	result = tk.MustQuery(testSQL)
	c.Check(result.Rows(), HasLen, 1)
	row = result.Rows()[0]
	expectedRow = []interface{}{
		"decimalschema", "CREATE TABLE `decimalschema` (\n" +
			"  `c1` decimal(11,0) DEFAULT NULL\n" +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin"}
	for i, r := range row {
		c.Check(r, Equals, expectedRow[i])
	}

	tk.MustExec("drop table if exists `decimalschema`")
	testSQL = "create table `decimalschema` (`c1` decimal(15));"
	tk.MustExec(testSQL)
	testSQL = "show create table decimalschema"
	result = tk.MustQuery(testSQL)
	c.Check(result.Rows(), HasLen, 1)
	row = result.Rows()[0]
	expectedRow = []interface{}{
		"decimalschema", "CREATE TABLE `decimalschema` (\n" +
			"  `c1` decimal(15,0) DEFAULT NULL\n" +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin"}
	for i, r := range row {
		c.Check(r, Equals, expectedRow[i])
	}

	testSQL = "SHOW VARIABLES LIKE 'character_set_results';"
	result = tk.MustQuery(testSQL)
	c.Check(result.Rows(), HasLen, 1)

	// Test case for index type and comment
	tk.MustExec(`create table show_index (id int, c int, primary key (id), index cIdx using hash (c) comment "index_comment_for_cIdx");`)
	tk.MustExec(`create index idx1 on show_index (id) using hash;`)
	tk.MustExec(`create index idx2 on show_index (id) comment 'idx';`)
	tk.MustExec(`create index idx3 on show_index (id) using hash comment 'idx';`)
	tk.MustExec(`alter table show_index add index idx4 (id) using btree comment 'idx';`)
	tk.MustExec(`create index idx5 using hash on show_index (id) using btree comment 'idx';`)
	tk.MustExec(`create index idx6 using hash on show_index (id);`)
	tk.MustExec(`create index idx7 on show_index (id);`)
	testSQL = "SHOW index from show_index;"
	tk.MustQuery(testSQL).Check(testutil.RowsWithSep("|",
		"show_index|0|PRIMARY|1|id|A|0|<nil>|<nil>||BTREE||",
		"show_index|1|cIdx|1|c|A|0|<nil>|<nil>|YES|HASH||index_comment_for_cIdx",
		"show_index|1|idx1|1|id|A|0|<nil>|<nil>|YES|HASH||",
		"show_index|1|idx2|1|id|A|0|<nil>|<nil>|YES|BTREE||idx",
		"show_index|1|idx3|1|id|A|0|<nil>|<nil>|YES|HASH||idx",
		"show_index|1|idx4|1|id|A|0|<nil>|<nil>|YES|BTREE||idx",
		"show_index|1|idx5|1|id|A|0|<nil>|<nil>|YES|BTREE||idx",
		"show_index|1|idx6|1|id|A|0|<nil>|<nil>|YES|HASH||",
		"show_index|1|idx7|1|id|A|0|<nil>|<nil>|YES|BTREE||",
	))

	// For show like with escape
	testSQL = `show tables like 'SHOW\_test'`
	result = tk.MustQuery(testSQL)
	rows := result.Rows()
	c.Check(rows, HasLen, 1)
	c.Check(rows[0], DeepEquals, []interface{}{"SHOW_test"})

	var ss stats
	variable.RegisterStatistics(ss)
	testSQL = "show status like 'character_set_results';"
	result = tk.MustQuery(testSQL)
	c.Check(result.Rows(), NotNil)

	tk.MustQuery("SHOW PROCEDURE STATUS WHERE Db='test'").Check(testkit.Rows())
	tk.MustQuery("SHOW TRIGGERS WHERE `Trigger` ='test'").Check(testkit.Rows())
	tk.MustQuery("SHOW PROCESSLIST;").Check(testkit.Rows())
	tk.MustQuery("SHOW FULL PROCESSLIST;").Check(testkit.Rows())
	tk.MustQuery("SHOW EVENTS WHERE Db = 'test'").Check(testkit.Rows())
	tk.MustQuery("SHOW PLUGINS").Check(testkit.Rows())
	tk.MustQuery("SHOW PROFILES").Check(testkit.Rows())

	// +-------------+--------------------+--------------+------------------+-------------------+
	// | File        | Position           | Binlog_Do_DB | Binlog_Ignore_DB | Executed_Gtid_Set |
	// +-------------+--------------------+--------------+------------------+-------------------+
	// | tidb-binlog | 400668057259474944 |              |                  |                   |
	// +-------------+--------------------+--------------+------------------+-------------------+
	result = tk.MustQuery("SHOW MASTER STATUS")
	c.Check(result.Rows(), HasLen, 1)
	row = result.Rows()[0]
	c.Check(row, HasLen, 5)
	c.Assert(row[1].(string) != "0", IsTrue)

	tk.MustQuery("SHOW PRIVILEGES")

	// Test show create database
	testSQL = `create database show_test_DB`
	tk.MustExec(testSQL)
	testSQL = "show create database show_test_DB;"
	tk.MustQuery(testSQL).Check(testutil.RowsWithSep("|",
		"show_test_DB|CREATE DATABASE `show_test_DB` /* !40100 DEFAULT CHARACTER SET utf8mb4 */",
	))

	tk.MustExec("use show_test_DB")
	result = tk.MustQuery("SHOW index from show_index from test where Column_name = 'c'")
	c.Check(result.Rows(), HasLen, 1)

	// Test show full columns
	// for issue https://github.com/pingcap/tidb/issues/4224
	tk.MustExec(`drop table if exists show_test_comment`)
	tk.MustExec(`create table show_test_comment (id int not null default 0 comment "show_test_comment_id")`)
	tk.MustQuery(`show full columns from show_test_comment`).Check(testutil.RowsWithSep("|",
		"id|int(11)|<nil>|NO||0||select,insert,update,references|show_test_comment_id",
	))

	// Test show create table with AUTO_INCREMENT option
	// for issue https://github.com/pingcap/tidb/issues/3747
	tk.MustExec(`drop table if exists show_auto_increment`)
	tk.MustExec(`create table show_auto_increment (id int key auto_increment) auto_increment=4`)
	tk.MustQuery(`show create table show_auto_increment`).Check(testutil.RowsWithSep("|",
		""+
			"show_auto_increment CREATE TABLE `show_auto_increment` (\n"+
			"  `id` int(11) NOT NULL AUTO_INCREMENT,\n"+
			"  PRIMARY KEY (`id`)\n"+
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin AUTO_INCREMENT=4",
	))
	// for issue https://github.com/pingcap/tidb/issues/4678
	autoIDStep := autoid.GetStep()
	tk.MustExec("insert into show_auto_increment values(20)")
	autoID := autoIDStep + 21
	tk.MustQuery(`show create table show_auto_increment`).Check(testutil.RowsWithSep("|",
		""+
			"show_auto_increment CREATE TABLE `show_auto_increment` (\n"+
			"  `id` int(11) NOT NULL AUTO_INCREMENT,\n"+
			"  PRIMARY KEY (`id`)\n"+
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin AUTO_INCREMENT="+strconv.Itoa(int(autoID)),
	))
	tk.MustExec(`drop table show_auto_increment`)
	tk.MustExec(`create table show_auto_increment (id int primary key auto_increment)`)
	tk.MustQuery(`show create table show_auto_increment`).Check(testutil.RowsWithSep("|",
		""+
			"show_auto_increment CREATE TABLE `show_auto_increment` (\n"+
			"  `id` int(11) NOT NULL AUTO_INCREMENT,\n"+
			"  PRIMARY KEY (`id`)\n"+
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin",
	))
	tk.MustExec("insert into show_auto_increment values(10)")
	autoID = autoIDStep + 11
	tk.MustQuery(`show create table show_auto_increment`).Check(testutil.RowsWithSep("|",
		""+
			"show_auto_increment CREATE TABLE `show_auto_increment` (\n"+
			"  `id` int(11) NOT NULL AUTO_INCREMENT,\n"+
			"  PRIMARY KEY (`id`)\n"+
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin AUTO_INCREMENT="+strconv.Itoa(int(autoID)),
	))

	// Test show table with column's comment contain escape character
	// for issue https://github.com/pingcap/tidb/issues/4411
	tk.MustExec(`drop table if exists show_escape_character`)
	tk.MustExec(`create table show_escape_character(id int comment 'a\rb\nc\td\0ef')`)
	tk.MustQuery(`show create table show_escape_character`).Check(testutil.RowsWithSep("|",
		""+
			"show_escape_character CREATE TABLE `show_escape_character` (\n"+
			"  `id` int(11) DEFAULT NULL COMMENT 'a\\rb\\nc	d\\0ef'\n"+
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin",
	))

	// for issue https://github.com/pingcap/tidb/issues/4424
	tk.MustExec("drop table if exists show_test")
	testSQL = `create table show_test(
		a varchar(10) COMMENT 'a\nb\rc\td\0e'
	) COMMENT='a\nb\rc\td\0e';`
	tk.MustExec(testSQL)
	testSQL = "show create table show_test;"
	result = tk.MustQuery(testSQL)
	c.Check(result.Rows(), HasLen, 1)
	row = result.Rows()[0]
	expectedRow = []interface{}{
		"show_test", "CREATE TABLE `show_test` (\n  `a` varchar(10) DEFAULT NULL COMMENT 'a\\nb\\rc	d\\0e'\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='a\\nb\\rc	d\\0e'"}
	for i, r := range row {
		c.Check(r, Equals, expectedRow[i])
	}

	// for issue https://github.com/pingcap/tidb/issues/4425
	tk.MustExec("drop table if exists show_test")
	testSQL = `create table show_test(
		a varchar(10) DEFAULT 'a\nb\rc\td\0e'
	);`
	tk.MustExec(testSQL)
	testSQL = "show create table show_test;"
	result = tk.MustQuery(testSQL)
	c.Check(result.Rows(), HasLen, 1)
	row = result.Rows()[0]
	expectedRow = []interface{}{
		"show_test", "CREATE TABLE `show_test` (\n  `a` varchar(10) DEFAULT 'a\\nb\\rc	d\\0e'\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin"}
	for i, r := range row {
		c.Check(r, Equals, expectedRow[i])
	}

	// for issue https://github.com/pingcap/tidb/issues/4426
	tk.MustExec("drop table if exists show_test")
	testSQL = `create table show_test(
		a bit(1),
		b bit(32) DEFAULT 0b0,
		c bit(1) DEFAULT 0b1,
		d bit(10) DEFAULT 0b1010
	);`
	tk.MustExec(testSQL)
	testSQL = "show create table show_test;"
	result = tk.MustQuery(testSQL)
	c.Check(result.Rows(), HasLen, 1)
	row = result.Rows()[0]
	expectedRow = []interface{}{
		"show_test", "CREATE TABLE `show_test` (\n  `a` bit(1) DEFAULT NULL,\n  `b` bit(32) DEFAULT b'0',\n  `c` bit(1) DEFAULT b'1',\n  `d` bit(10) DEFAULT b'1010'\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin"}
	for i, r := range row {
		c.Check(r, Equals, expectedRow[i])
	}

	// for issue #4255
	result = tk.MustQuery(`show function status like '%'`)
	result.Check(result.Rows())
	result = tk.MustQuery(`show plugins like '%'`)
	result.Check(result.Rows())

	// for issue #4740
	testSQL = `drop table if exists t`
	tk.MustExec(testSQL)
	testSQL = `create table t (a int1, b int2, c int3, d int4, e int8)`
	tk.MustExec(testSQL)
	testSQL = `show create table t;`
	result = tk.MustQuery(testSQL)
	c.Check(result.Rows(), HasLen, 1)
	row = result.Rows()[0]
	expectedRow = []interface{}{
		"t",
		"CREATE TABLE `t` (\n" +
			"  `a` tinyint(4) DEFAULT NULL,\n" +
			"  `b` smallint(6) DEFAULT NULL,\n" +
			"  `c` mediumint(9) DEFAULT NULL,\n" +
			"  `d` int(11) DEFAULT NULL,\n" +
			"  `e` bigint(20) DEFAULT NULL\n" +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin",
	}
	for i, r := range row {
		c.Check(r, Equals, expectedRow[i])
	}

	// Test get default collate for a specified charset.
	tk.MustExec(`drop table if exists t`)
	tk.MustExec(`create table t (a int) default charset=utf8mb4`)
	tk.MustQuery(`show create table t`).Check(testutil.RowsWithSep("|",
		"t CREATE TABLE `t` (\n"+
			"  `a` int(11) DEFAULT NULL\n"+
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin",
	))

	// Test range partition
	tk.MustExec(`drop table if exists t`)
	tk.MustExec(`CREATE TABLE t (a int) PARTITION BY RANGE(a) (
 	PARTITION p0 VALUES LESS THAN (10),
 	PARTITION p1 VALUES LESS THAN (20),
 	PARTITION p2 VALUES LESS THAN (MAXVALUE))`)
	tk.MustQuery("show create table t").Check(testutil.RowsWithSep("|",
		"t CREATE TABLE `t` (\n"+
			"  `a` int(11) DEFAULT NULL\n"+
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin"+"\nPARTITION BY RANGE ( `a` ) (\n  PARTITION p0 VALUES LESS THAN (10),\n  PARTITION p1 VALUES LESS THAN (20),\n  PARTITION p2 VALUES LESS THAN (MAXVALUE)\n)",
	))

	tk.MustExec(`drop table if exists t`)
	_, err := tk.Exec(`CREATE TABLE t (x int, y char) PARTITION BY RANGE(y) (
 	PARTITION p0 VALUES LESS THAN (10),
 	PARTITION p1 VALUES LESS THAN (20),
 	PARTITION p2 VALUES LESS THAN (MAXVALUE))`)
	c.Assert(err, NotNil)

	// Test range columns partition
	tk.MustExec(`drop table if exists t`)
	tk.MustExec(`CREATE TABLE t (a int, b int, c char, d int) PARTITION BY RANGE COLUMNS(a,d,c) (
 	PARTITION p0 VALUES LESS THAN (5,10,'ggg'),
 	PARTITION p1 VALUES LESS THAN (10,20,'mmm'),
 	PARTITION p2 VALUES LESS THAN (15,30,'sss'),
        PARTITION p3 VALUES LESS THAN (50,MAXVALUE,MAXVALUE))`)
	tk.MustQuery("show create table t").Check(testutil.RowsWithSep("|",
		"t CREATE TABLE `t` (\n"+
			"  `a` int(11) DEFAULT NULL,\n"+
			"  `b` int(11) DEFAULT NULL,\n"+
			"  `c` char(1) DEFAULT NULL,\n"+
			"  `d` int(11) DEFAULT NULL\n"+
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin"+"\nPARTITION BY RANGE COLUMNS(a,d,c) (\n  PARTITION p0 VALUES LESS THAN (5,10,\"ggg\"),\n  PARTITION p1 VALUES LESS THAN (10,20,\"mmm\"),\n  PARTITION p2 VALUES LESS THAN (15,30,\"sss\"),\n  PARTITION p3 VALUES LESS THAN (50,MAXVALUE,MAXVALUE)\n)",
	))

	// Test show create table compression type.
	tk.MustExec(`drop table if exists t1`)
	tk.MustExec(`CREATE TABLE t1 (c1 INT) COMPRESSION="zlib";`)
	tk.MustQuery("show create table t1").Check(testutil.RowsWithSep("|",
		"t1 CREATE TABLE `t1` (\n"+
			"  `c1` int(11) DEFAULT NULL\n"+
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMPRESSION='zlib'",
	))

	// Test show create table year type
	tk.MustExec(`drop table if exists t`)
	tk.MustExec(`create table t(y year unsigned signed zerofill zerofill, x int, primary key(y));`)
	tk.MustQuery(`show create table t`).Check(testutil.RowsWithSep("|",
		"t CREATE TABLE `t` (\n"+
			"  `y` year NOT NULL,\n"+
			"  `x` int(11) DEFAULT NULL,\n"+
			"  PRIMARY KEY (`y`)\n"+
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin"))

	// Test show create table with zerofill flag
	tk.MustExec(`drop table if exists t`)
	tk.MustExec(`create table t(id int primary key, val tinyint(10) zerofill);`)
	tk.MustQuery(`show create table t`).Check(testutil.RowsWithSep("|",
		"t CREATE TABLE `t` (\n"+
			"  `id` int(11) NOT NULL,\n"+
			"  `val` tinyint(10) UNSIGNED ZEROFILL DEFAULT NULL,\n"+
			"  PRIMARY KEY (`id`)\n"+
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin"))
}

func (s *testSuite) TestShowVisibility(c *C) {
	tk := testkit.NewTestKit(c, s.store)
	tk.MustExec("create database showdatabase")
	tk.MustExec("use showdatabase")
	tk.MustExec("create table t1 (id int)")
	tk.MustExec("create table t2 (id int)")
	tk.MustExec(`create user 'show'@'%'`)
	tk.MustExec(`flush privileges`)

	tk1 := testkit.NewTestKit(c, s.store)
	se, err := session.CreateSession4Test(s.store)
	c.Assert(err, IsNil)
	c.Assert(se.Auth(&auth.UserIdentity{Username: "show", Hostname: "%"}, nil, nil), IsTrue)
	tk1.Se = se

	// No ShowDatabases privilege, this user would see nothing except INFORMATION_SCHEMA.
	tk.MustQuery("show databases").Check(testkit.Rows("INFORMATION_SCHEMA"))

	// After grant, the user can see the database.
	tk.MustExec(`grant select on showdatabase.t1 to 'show'@'%'`)
	tk.MustExec(`flush privileges`)
	tk1.MustQuery("show databases").Check(testkit.Rows("INFORMATION_SCHEMA", "showdatabase"))

	// The user can see t1 but not t2.
	tk1.MustExec("use showdatabase")
	tk1.MustQuery("show tables").Check(testkit.Rows("t1"))

	// After revoke, show database result should be just except INFORMATION_SCHEMA.
	tk.MustExec(`revoke select on showdatabase.t1 from 'show'@'%'`)
	tk.MustExec(`flush privileges`)
	tk1.MustQuery("show databases").Check(testkit.Rows("INFORMATION_SCHEMA"))

	// Grant any global privilege would make show databases available.
	tk.MustExec(`grant CREATE on *.* to 'show'@'%'`)
	tk.MustExec(`flush privileges`)
	rows := tk1.MustQuery("show databases").Rows()
	c.Assert(len(rows), GreaterEqual, 2) // At least INFORMATION_SCHEMA and showdatabase

	tk.MustExec(`drop user 'show'@'%'`)
	tk.MustExec("drop database showdatabase")
}

// mockSessionManager is a mocked session manager that wraps one session
// it returns only this session's current process info as processlist for test.
type mockSessionManager struct {
	session.Session
}

// Kill implements the SessionManager.Kill interface.
func (msm *mockSessionManager) Kill(cid uint64, query bool) {
}

type stats struct {
}

func (s stats) GetScope(status string) variable.ScopeFlag { return variable.DefaultStatusVarScopeFlag }

func (s stats) Stats(vars *variable.SessionVars) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	var a, b interface{}
	b = "123"
	m["test_interface_nil"] = a
	m["test_interface"] = b
	m["test_interface_slice"] = []interface{}{"a", "b", "c"}
	return m, nil
}

func (s *testSuite) TestShowWarnings(c *C) {
	tk := testkit.NewTestKit(c, s.store)
	tk.MustExec("use test")
	testSQL := `create table if not exists show_warnings (a int)`
	tk.MustExec(testSQL)
	tk.MustExec("set @@sql_mode=''")
	tk.MustExec("insert show_warnings values ('a')")
	c.Assert(tk.Se.GetSessionVars().StmtCtx.WarningCount(), Equals, uint16(1))
	tk.MustQuery("show warnings").Check(testutil.RowsWithSep("|", "Warning|1265|Data Truncated"))
	c.Assert(tk.Se.GetSessionVars().StmtCtx.WarningCount(), Equals, uint16(0))
	tk.MustQuery("show warnings").Check(testutil.RowsWithSep("|", "Warning|1265|Data Truncated"))
	c.Assert(tk.Se.GetSessionVars().StmtCtx.WarningCount(), Equals, uint16(0))

	// Test Warning level 'Error'
	testSQL = `create table show_warnings (a int)`
	tk.Exec(testSQL)
	c.Assert(tk.Se.GetSessionVars().StmtCtx.WarningCount(), Equals, uint16(1))
	tk.MustQuery("show warnings").Check(testutil.RowsWithSep("|", "Error|1050|Table 'test.show_warnings' already exists"))
	tk.MustQuery("select @@error_count").Check(testutil.RowsWithSep("|", "1"))

	// Test Warning level 'Note'
	testSQL = `create table show_warnings_2 (a int)`
	tk.MustExec(testSQL)
	testSQL = `create table if not exists show_warnings_2 like show_warnings`
	tk.Exec(testSQL)
	c.Assert(tk.Se.GetSessionVars().StmtCtx.WarningCount(), Equals, uint16(1))
	tk.MustQuery("show warnings").Check(testutil.RowsWithSep("|", "Note|1050|Table 'test.show_warnings_2' already exists"))
	tk.MustQuery("select @@warning_count").Check(testutil.RowsWithSep("|", "1"))
	tk.MustQuery("select @@warning_count").Check(testutil.RowsWithSep("|", "0"))
}

func (s *testSuite) TestShowErrors(c *C) {
	tk := testkit.NewTestKit(c, s.store)
	tk.MustExec("use test")
	testSQL := `create table if not exists show_errors (a int)`
	tk.MustExec(testSQL)
	testSQL = `create table show_errors (a int)`
	tk.Exec(testSQL)

	tk.MustQuery("show errors").Check(testutil.RowsWithSep("|", "Error|1050|Table 'test.show_errors' already exists"))
}

func (s *testSuite) TestIssue3641(c *C) {
	tk := testkit.NewTestKit(c, s.store)
	_, err := tk.Exec("show tables;")
	c.Assert(err.Error(), Equals, plannercore.ErrNoDB.Error())
	_, err = tk.Exec("show table status;")
	c.Assert(err.Error(), Equals, plannercore.ErrNoDB.Error())
}

// TestShow2 is moved from session_test
func (s *testSuite) TestShow2(c *C) {
	tk := testkit.NewTestKit(c, s.store)
	tk.MustExec("use test")

	tk.MustExec("set global autocommit=0")
	tk1 := testkit.NewTestKit(c, s.store)
	tk1.MustQuery("show global variables where variable_name = 'autocommit'").Check(testkit.Rows("autocommit 0"))
	tk.MustExec("set global autocommit = 1")
	tk2 := testkit.NewTestKit(c, s.store)
	// TODO: In MySQL, the result is "autocommit ON".
	tk2.MustQuery("show global variables where variable_name = 'autocommit'").Check(testkit.Rows("autocommit 1"))

	// TODO: Specifying the charset for national char/varchar should not be supported.
	tk.MustExec("drop table if exists test_full_column")
	tk.MustExec(`create table test_full_column(
					c_int int,
					c_float float,
					c_bit bit,
					c_bool bool,
					c_char char(1) charset ascii collate ascii_bin,
					c_nchar national char(1) charset ascii collate ascii_bin,
					c_binary binary,
					c_varchar varchar(1) charset ascii collate ascii_bin,
					c_nvarchar national varchar(1) charset ascii collate ascii_bin,
					c_varbinary varbinary(1),
					c_year year,
					c_date date,
					c_time time,
					c_datetime datetime,
					c_timestamp timestamp,
					c_blob blob,
					c_tinyblob tinyblob,
					c_mediumblob mediumblob,
					c_longblob longblob,
					c_text text charset ascii collate ascii_bin,
					c_tinytext tinytext charset ascii collate ascii_bin,
					c_mediumtext mediumtext charset ascii collate ascii_bin,
					c_longtext longtext charset ascii collate ascii_bin,
					c_json json,
					c_enum enum('1') charset ascii collate ascii_bin,
					c_set set('1') charset ascii collate ascii_bin
				);`)

	tk.MustQuery(`show full columns from test_full_column`).Check(testkit.Rows(
		"" +
			"c_int int(11) <nil> YES  <nil>  select,insert,update,references ]\n" +
			"[c_float float <nil> YES  <nil>  select,insert,update,references ]\n" +
			"[c_bit bit(1) <nil> YES  <nil>  select,insert,update,references ]\n" +
			"[c_bool tinyint(1) <nil> YES  <nil>  select,insert,update,references ]\n" +
			"[c_char char(1) ascii_bin YES  <nil>  select,insert,update,references ]\n" +
			"[c_nchar char(1) ascii_bin YES  <nil>  select,insert,update,references ]\n" +
			"[c_binary binary(1) <nil> YES  <nil>  select,insert,update,references ]\n" +
			"[c_varchar varchar(1) ascii_bin YES  <nil>  select,insert,update,references ]\n" +
			"[c_nvarchar varchar(1) ascii_bin YES  <nil>  select,insert,update,references ]\n" +
			"[c_varbinary varbinary(1) <nil> YES  <nil>  select,insert,update,references ]\n" +
			"[c_year year <nil> YES  <nil>  select,insert,update,references ]\n" +
			"[c_date date <nil> YES  <nil>  select,insert,update,references ]\n" +
			"[c_time time <nil> YES  <nil>  select,insert,update,references ]\n" +
			"[c_datetime datetime <nil> YES  <nil>  select,insert,update,references ]\n" +
			"[c_timestamp timestamp <nil> YES  <nil>  select,insert,update,references ]\n" +
			"[c_blob blob <nil> YES  <nil>  select,insert,update,references ]\n" +
			"[c_tinyblob tinyblob <nil> YES  <nil>  select,insert,update,references ]\n" +
			"[c_mediumblob mediumblob <nil> YES  <nil>  select,insert,update,references ]\n" +
			"[c_longblob longblob <nil> YES  <nil>  select,insert,update,references ]\n" +
			"[c_text text ascii_bin YES  <nil>  select,insert,update,references ]\n" +
			"[c_tinytext tinytext ascii_bin YES  <nil>  select,insert,update,references ]\n" +
			"[c_mediumtext mediumtext ascii_bin YES  <nil>  select,insert,update,references ]\n" +
			"[c_longtext longtext ascii_bin YES  <nil>  select,insert,update,references ]\n" +
			"[c_json json <nil> YES  <nil>  select,insert,update,references ]\n" +
			"[c_enum enum('1') ascii_bin YES  <nil>  select,insert,update,references ]\n" +
			"[c_set set('1') ascii_bin YES  <nil>  select,insert,update,references "))

	tk.MustExec("drop table if exists test_full_column")

	tk.MustExec("drop table if exists t")
	tk.MustExec(`create table if not exists t (c int) comment '注释'`)
	tk.MustQuery(`show columns from t`).Check(testutil.RowsWithSep(",", "c,int(11),YES,,<nil>,"))
	tk.MustQuery("show collation where Charset = 'utf8' and Collation = 'utf8_bin'").Check(testutil.RowsWithSep(",", "utf8_bin,utf8,83,,Yes,1"))
	tk.MustQuery("show tables").Check(testkit.Rows("t"))
	tk.MustQuery("show full tables").Check(testkit.Rows("t BASE TABLE"))
	ctx := tk.Se.(sessionctx.Context)
	is := domain.GetDomain(ctx).InfoSchema()
	tblInfo, err := is.TableByName(model.NewCIStr("test"), model.NewCIStr("t"))
	c.Assert(err, IsNil)
	createTime := model.TSConvert2Time(tblInfo.Meta().UpdateTS).String()
	r := tk.MustQuery("show table status from test like 't'")
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	r.Check(testkit.Rows(fmt.Sprintf("t InnoDB 10 Compact 100 100 100 100 100 100 100 %s %s %s utf8_general_ci   注释", createTime, timeStr, timeStr)))

	tk.Se.Auth(&auth.UserIdentity{Username: "root", Hostname: "%"}, nil, []byte("012345678901234567890"))

	tk.MustQuery("show databases like 'test'").Check(testkit.Rows("test"))

	tk.MustExec(`grant all on *.* to 'root'@'%'`)
	tk.MustQuery("show grants").Check(testkit.Rows(`GRANT ALL PRIVILEGES ON *.* TO 'root'@'%'`))

	tk.MustQuery("show grants for current_user()").Check(testkit.Rows(`GRANT ALL PRIVILEGES ON *.* TO 'root'@'%'`))
	tk.MustQuery("show grants for current_user").Check(testkit.Rows(`GRANT ALL PRIVILEGES ON *.* TO 'root'@'%'`))
}

func (s *testSuite) TestCollation(c *C) {
	tk := testkit.NewTestKit(c, s.store)
	tk.MustExec("use test")

	rs, err := tk.Exec("show collation;")
	c.Assert(err, IsNil)
	fields := rs.Fields()
	c.Assert(fields[0].Column.Tp, Equals, mysql.TypeVarchar)
	c.Assert(fields[1].Column.Tp, Equals, mysql.TypeVarchar)
	c.Assert(fields[2].Column.Tp, Equals, mysql.TypeLonglong)
	c.Assert(fields[3].Column.Tp, Equals, mysql.TypeVarchar)
	c.Assert(fields[4].Column.Tp, Equals, mysql.TypeVarchar)
	c.Assert(fields[5].Column.Tp, Equals, mysql.TypeLonglong)
}

func (s *testSuite) TestShowTableStatus(c *C) {
	tk := testkit.NewTestKit(c, s.store)

	tk.MustExec("use test")
	tk.MustExec("set @@session.tidb_enable_table_partition=1;")
	tk.MustExec(`drop table if exists t;`)
	tk.MustExec(`create table t(a bigint);`)

	// It's not easy to test the result contents because every time the test runs, "Create_time" changed.
	tk.MustExec("show table status;")
	rs, err := tk.Exec("show table status;")
	c.Assert(errors.ErrorStack(err), Equals, "")
	c.Assert(rs, NotNil)
	rows, err := session.GetRows4Test(context.Background(), tk.Se, rs)
	c.Assert(errors.ErrorStack(err), Equals, "")
	err = rs.Close()
	c.Assert(errors.ErrorStack(err), Equals, "")

	for i := range rows {
		row := rows[i]
		c.Assert(row.GetString(0), Equals, "t")
		c.Assert(row.GetString(1), Equals, "InnoDB")
		c.Assert(row.GetInt64(2), Equals, int64(10))
		c.Assert(row.GetString(3), Equals, "Compact")
	}
	tk.MustExec(`drop table if exists tp;`)
	tk.MustExec(`create table tp (a int)
 		partition by range(a)
 		( partition p0 values less than (10),
		  partition p1 values less than (20),
		  partition p2 values less than (maxvalue)
  		);`)
	rs, err = tk.Exec("show table status from test like 'tp';")
	c.Assert(errors.ErrorStack(err), Equals, "")
	rows, err = session.GetRows4Test(context.Background(), tk.Se, rs)
	c.Assert(errors.ErrorStack(err), Equals, "")
	c.Assert(rows[0].GetString(16), Equals, "partitioned")
}

func (s *testSuite) TestShowSlow(c *C) {
	tk := testkit.NewTestKit(c, s.store)
	// The test result is volatile, because
	// 1. Slow queries is stored in domain, which may be affected by other tests.
	// 2. Collecting slow queries is a asynchronous process, check immediately may not get the expected result.
	// 3. Make slow query like "select sleep(1)" would slow the CI.
	// So, we just cover the code but do not check the result.
	tk.MustQuery(`admin show slow recent 3`)
	tk.MustQuery(`admin show slow top 3`)
	tk.MustQuery(`admin show slow top internal 3`)
	tk.MustQuery(`admin show slow top all 3`)
}

func (s *testSuite) TestShowOpenTables(c *C) {
	tk := testkit.NewTestKit(c, s.store)
	tk.MustQuery("show open tables")
	tk.MustQuery("show open tables in test")
}

func (s *testSuite) TestShowCreateTable(c *C) {
	tk := testkit.NewTestKit(c, s.store)
	tk.MustExec("use test")
	tk.MustExec(`drop table if exists different_charset`)
	tk.MustExec(`create table different_charset(ch1 varchar(10) charset utf8, ch2 varchar(10) charset binary);`)
	tk.MustQuery(`show create table different_charset`).Check(testutil.RowsWithSep("|",
		""+
			"different_charset CREATE TABLE `different_charset` (\n"+
			"  `ch1` varchar(10) CHARSET utf8 COLLATE utf8_bin DEFAULT NULL,\n"+
			"  `ch2` varbinary(10) DEFAULT NULL\n"+
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin",
	))

	tk.MustExec("CREATE TABLE `log` (" +
		"`LOG_ID` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT," +
		"`ROUND_ID` bigint(20) UNSIGNED NOT NULL," +
		"`USER_ID` int(10) UNSIGNED NOT NULL," +
		"`USER_IP` int(10) UNSIGNED DEFAULT NULL," +
		"`END_TIME` datetime NOT NULL," +
		"`USER_TYPE` int(11) DEFAULT NULL," +
		"`APP_ID` int(11) DEFAULT NULL," +
		"PRIMARY KEY (`LOG_ID`,`END_TIME`)," +
		"KEY `IDX_EndTime` (`END_TIME`)," +
		"KEY `IDX_RoundId` (`ROUND_ID`)," +
		"KEY `IDX_UserId_EndTime` (`USER_ID`,`END_TIME`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=505488 " +
		"PARTITION BY RANGE ( month(`end_time`) ) (" +
		"PARTITION p1 VALUES LESS THAN (2)," +
		"PARTITION p2 VALUES LESS THAN (3)," +
		"PARTITION p3 VALUES LESS THAN (4)," +
		"PARTITION p4 VALUES LESS THAN (5)," +
		"PARTITION p5 VALUES LESS THAN (6)," +
		"PARTITION p6 VALUES LESS THAN (7)," +
		"PARTITION p7 VALUES LESS THAN (8)," +
		"PARTITION p8 VALUES LESS THAN (9)," +
		"PARTITION p9 VALUES LESS THAN (10)," +
		"PARTITION p10 VALUES LESS THAN (11)," +
		"PARTITION p11 VALUES LESS THAN (12)," +
		"PARTITION p12 VALUES LESS THAN (MAXVALUE))")
	tk.MustQuery("show create table log").Check(testutil.RowsWithSep("|",
		"log CREATE TABLE `log` (\n"+
			"  `LOG_ID` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,\n"+
			"  `ROUND_ID` bigint(20) UNSIGNED NOT NULL,\n"+
			"  `USER_ID` int(10) UNSIGNED NOT NULL,\n"+
			"  `USER_IP` int(10) UNSIGNED DEFAULT NULL,\n"+
			"  `END_TIME` datetime NOT NULL,\n"+
			"  `USER_TYPE` int(11) DEFAULT NULL,\n"+
			"  `APP_ID` int(11) DEFAULT NULL,\n"+
			"  PRIMARY KEY (`LOG_ID`,`END_TIME`),\n"+
			"  KEY `IDX_EndTime` (`END_TIME`),\n"+
			"  KEY `IDX_RoundId` (`ROUND_ID`),\n"+
			"  KEY `IDX_UserId_EndTime` (`USER_ID`,`END_TIME`)\n"+
			") ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=505488\n"+
			"PARTITION BY RANGE ( month(`end_time`) ) (\n"+
			"  PARTITION p1 VALUES LESS THAN (2),\n"+
			"  PARTITION p2 VALUES LESS THAN (3),\n"+
			"  PARTITION p3 VALUES LESS THAN (4),\n"+
			"  PARTITION p4 VALUES LESS THAN (5),\n"+
			"  PARTITION p5 VALUES LESS THAN (6),\n"+
			"  PARTITION p6 VALUES LESS THAN (7),\n"+
			"  PARTITION p7 VALUES LESS THAN (8),\n"+
			"  PARTITION p8 VALUES LESS THAN (9),\n"+
			"  PARTITION p9 VALUES LESS THAN (10),\n"+
			"  PARTITION p10 VALUES LESS THAN (11),\n"+
			"  PARTITION p11 VALUES LESS THAN (12),\n"+
			"  PARTITION p12 VALUES LESS THAN (MAXVALUE)\n"+
			")"))

	tk.MustExec("create table t (a int, b int) shard_row_id_bits = 4 pre_split_regions=3;")
	tk.MustQuery("show create table `t`").Check(testutil.RowsWithSep("|",
		""+
			"t CREATE TABLE `t` (\n"+
			"  `a` int(11) DEFAULT NULL,\n"+
			"  `b` int(11) DEFAULT NULL\n"+
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin/*!90000 SHARD_ROW_ID_BITS=4 PRE_SPLIT_REGIONS=3 */",
	))
	tk.MustExec("drop table t")
}

func (s *testSuite) TestShowEscape(c *C) {
	tk := testkit.NewTestKit(c, s.store)

	tk.MustExec("use test")
	tk.MustExec("drop table if exists `t``abl\"e`")
	tk.MustExec("create table `t``abl\"e`(`c``olum\"n` int(11) primary key)")
	tk.MustQuery("show create table `t``abl\"e`").Check(testutil.RowsWithSep("|",
		""+
			"t`abl\"e CREATE TABLE `t``abl\"e` (\n"+
			"  `c``olum\"n` int(11) NOT NULL,\n"+
			"  PRIMARY KEY (`c``olum\"n`)\n"+
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin",
	))

	// ANSI_QUOTES will change the SHOW output
	tk.MustExec("set @old_sql_mode=@@sql_mode")
	tk.MustExec("set sql_mode=ansi_quotes")
	tk.MustQuery("show create table \"t`abl\"\"e\"").Check(testutil.RowsWithSep("|",
		""+
			"t`abl\"e CREATE TABLE \"t`abl\"\"e\" (\n"+
			"  \"c`olum\"\"n\" int(11) NOT NULL,\n"+
			"  PRIMARY KEY (\"c`olum\"\"n\")\n"+
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin",
	))

	tk.MustExec("rename table \"t`abl\"\"e\" to t")
	tk.MustExec("set sql_mode=@old_sql_mode")
}
