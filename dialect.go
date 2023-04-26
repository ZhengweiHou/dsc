package dsc

import (
	"errors"
)

var errUnsupportedOperation = errors.New("unsupported operation")

type DefaultDialect struct{}

func (d DefaultDialect) GetDatastores(manager Manager) ([]string, error) {
	return nil, nil
}

func (d DefaultDialect) GetTables(manager Manager, datastore string) ([]string, error) {
	return nil, nil
}

func (d DefaultDialect) DropTable(manager Manager, datastore string, table string) error {
	return nil
}

func (d DefaultDialect) IsKeyCheckSwitchSessionLevel() bool {
	return true
}

func (d DefaultDialect) CreateTable(manager Manager, datastore string, table string, options interface{}) error {
	return nil
}

func (d DefaultDialect) CanCreateDatastore(manager Manager) bool {
	return false
}

func (d DefaultDialect) GetColumns(manager Manager, datastore, table string) ([]Column, error) {
	return []Column{}, nil
}

func (d DefaultDialect) GetColumnsByColumnNames(manager Manager, datastore, tableName string, columnNames []string) ([]Column, error) {
	return []Column{}, nil
}

func (d DefaultDialect) CreateDatastore(manager Manager, datastore string) error {
	return errUnsupportedOperation
}

func (d DefaultDialect) CanDropDatastore(manager Manager) bool {
	return false
}

func (d DefaultDialect) DropDatastore(manager Manager, datastore string) error {
	return errUnsupportedOperation
}

func (d DefaultDialect) GetCurrentDatastore(manager Manager) (string, error) {
	return "", nil
}

func (d DefaultDialect) BulkInsertType() string {
	return ""
}

func (d DefaultDialect) GetSequence(manager Manager, name string) (int64, error) {
	return 0, errUnsupportedOperation
}

func (d DefaultDialect) GetKeyName(manager Manager, datastore, table string) string {
	return ""
}

func (d DefaultDialect) IsAutoincrement(manager Manager, datastore, table string) bool {
	return false
}

func (d DefaultDialect) CanPersistBatch() bool {
	return false
}

func (d DefaultDialect) Init(manager Manager, connection Connection) error {
	return nil
}

//DisableForeignKeyCheck disables fk check
func (d DefaultDialect) DisableForeignKeyCheck(manager Manager, connection Connection) error {
	return nil
}

//DisableForeignKeyCheck disables fk check
func (d DefaultDialect) EnableForeignKeyCheck(manager Manager, connection Connection) error {
	return nil
}

func (d DefaultDialect) NormalizeSQL(SQL string) string {
	return SQL
}

func (d DefaultDialect) ShowCreateTable(manager Manager, table string) (string, error) {
	return "", errors.New("unsupported")
}

func (d DefaultDialect) CanHandleTransaction() bool {
	return false
}

//EachTable iterates each datastore table
func (d DefaultDialect) EachTable(manager Manager, handler func(table string) error) error {
	dbname, err := d.GetCurrentDatastore(manager)
	if err != nil {
		return err
	}
	tables, err := d.GetTables(manager, dbname)
	if err != nil {
		return err
	}
	for _, table := range tables {
		if err := handler(table); err != nil {
			return err
		}
	}
	return err
}

func (d DefaultDialect) Ping(manager Manager) error {
	return nil
}

//NewDefaultDialect crates a defulat dialect. DefaultDialect can be used as a embeddable struct (super class).
func NewDefaultDialect() DatastoreDialect {
	return &DefaultDialect{}
}
