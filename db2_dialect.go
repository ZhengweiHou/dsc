package dsc

import (
	"fmt"
	"strings"

	"github.com/viant/toolbox"
)

// "database/sql"
// "fmt"
// "github.com/viant/toolbox"
// "path"
// "strings"
// "time"

//TODO refactor for both better dialect and multi version of the same vendor handling

// tablesSQL              string
const db2TableListSQL = ` SELECT table_name AS name FROM "SYSIBM".TABLES WHERE TABLE_SCHEMA =?`

// sequenceSQL

// schemaSQL
const db2DefaultSchemaSQL = `select current schema AS name from SYSIBM.SYSDUMMY1`

// const db2DefaultSchemaSQL = `SELECT  CURRENT  SCHEMA AS name FROM SYSIBM.DUAL`

// allSchemaSQL
const db2SchemaListSQL = `SELECT NAME AS name FROM "SYSIBM".SYSSCHEMATA`

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

type db2SQLDialect struct {
	*sqlDatastoreDialect
	// DatastoreDialect
}

func newDb2SQLDialect() *db2SQLDialect {
	var result = &db2SQLDialect{}
	sqlDialect := NewSQLDatastoreDialect(db2TableListSQL, "", db2DefaultSchemaSQL, db2SchemaListSQL, db2PrimaryKeySQL, "", "", "", db2TableInfoSQL, 0, result)
	result.sqlDatastoreDialect = sqlDialect
	sqlDialect.DatastoreDialect = result
	return result
}

//GetKeyName returns key/PK columns
func (d db2SQLDialect) GetKeyName(manager Manager, datastore, table string) string {
	if d.keySQL == "" {
		return ""
	}
	SQL := fmt.Sprintf(d.keySQL, table, datastore)
	var records = make([]map[string]interface{}, 0)

	err := manager.ReadAll(&records, SQL, []interface{}{}, nil)
	if err != nil {
		fmt.Printf("ERR:%v\n", err)
		return ""
	}
	var result = make([]string, 0)
	for _, item := range records {
		result = append(result, toolbox.AsString(item["NAME"]))
	}
	return strings.Join(result, ",")
}

//ShowCreateTable returns basic table DDL (this implementation does not check unique and fk constrains)
func (d *db2SQLDialect) ShowCreateTable(manager Manager, table string) (string, error) {
	datastore, err := d.DatastoreDialect.GetCurrentDatastore(manager)
	if err != nil {
		return "", err
	}
	columns, err := d.DatastoreDialect.GetColumns(manager, datastore, table)
	if err != nil {
		return "", fmt.Errorf("unable to get columns for %v.%v, %v", datastore, table, err)
	}
	pkColumns := d.DatastoreDialect.GetKeyName(manager, datastore, table)
	if err != nil {
		return "", fmt.Errorf("unable to get pk key for %v.%v, %v", datastore, table, err)
	}
	var indexPk = map[string]bool{}
	for _, key := range strings.Split(pkColumns, ",") {
		indexPk[key] = true
	}
	var projection = make([]string, 0)
	var keyColumns = make([]string, 0)
	for _, column := range columns {
		var dataType = column.DatabaseTypeName()
		ddlColumn := fmt.Sprintf("%v %v", column.Name(), dataType)

		l, hasl := column.Length()
		p, s, hasps := column.DecimalSize()
		switch dataType {
		case "DECIMAL": // TODO 待补全
			if hasps {
				ddlColumn += fmt.Sprintf("(%v,%v)", p, s)
			}
		case "VARCHAR", "CHAR", "BINARY", "BLOB", "CLOB": // TODO 待补全
			if hasl {
				ddlColumn += fmt.Sprintf("(%v)", l)
			}
		case "DBCLOB", "GRAPHIC", "DECFLOAT":
			// ddlColumn += fmt.Sprintf("(%v CODEUNITS16)", l)
		}

		if nullable, ok := column.Nullable(); ok && !nullable {
			ddlColumn += " NOT NULL "
		}
		projection = append(projection, ddlColumn)
	}
	projection = append(keyColumns, projection...)

	if len(pkColumns) > 0 {
		projection = append(projection, fmt.Sprintf("PRIMARY KEY(%v)", pkColumns))
	}

	csql := fmt.Sprintf("CREATE TABLE %v(\n\t%v);", table, strings.Join(projection, ",\n\t"))

	return csql, nil
}
