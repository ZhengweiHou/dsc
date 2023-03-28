package dsc

import (
// "database/sql"
// "fmt"
// "github.com/viant/toolbox"
// "path"
// "strings"
// "time"
)

//TODO refactor for both better dialect and multi version of the same vendor handling

// tablesSQL              string
const db2TableListSQL = ` SELECT table_name AS name FROM "SYSIBM".TABLES WHERE TABLE_SCHEMA =?`

// sequenceSQL

// schemaSQL
// const db2DefaultSchemaSQL = `select current schema AS name from SYSIBM.SYSDUMMY1`
const db2DefaultSchemaSQL = `SELECT  CURRENT  SCHEMA AS name FROM SYSIBM.DUAL`

// allSchemaSQL
const db2SchemaListSql = `SELECT NAME AS name FROM "SYSIBM".SYSSCHEMATA`

// keySQL
const db2PrimaryKeySQL = `SELECT COLNAME  AS name FROM SYSCAT.KEYCOLUSE  
WHERE TABNAME ='%v' AND TABSCHEMA ='%v'`

// disableForeignKeyCheck 关闭外键检查
// enableForeignKeyCheck 启用外键检查
// autoIncrementSQL

// tableInfoSQL
const db2TableInfoSQL = ` SELECT
COLUMN_NAME,
DATA_TYPE,
CHARACTER_MAXIMUM_LENGTH AS data_type_length,
NUMERIC_PRECISION,
NUMERIC_SCALE,
IS_NULLABLE
FROM "SYSIBM".COLUMNS  
WHERE TABLE_NAME ='%s' AND TABLE_SCHEMA ='%s'
ORDER  BY ORDINAL_POSITION`

// type db2SQLDialect struct {
// 	*sqlDatastoreDialect
// }

type db2Dialect struct {
	DatastoreDialect
}

func newDb2Dialect() db2Dialect {
	var result = db2Dialect{}
	sqlDialect := NewSQLDatastoreDialect(db2TableListSQL, "", db2DefaultSchemaSQL, db2SchemaListSql, db2PrimaryKeySQL, "", "", "", db2TableInfoSQL, 0, result)
	result.DatastoreDialect = sqlDialect
	sqlDialect.DatastoreDialect = result
	return result
}
