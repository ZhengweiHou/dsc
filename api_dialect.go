package dsc

type DatastoreDialect interface {
	GetDatastores(manager Manager) ([]string, error)

	GetTables(manager Manager, datastore string) ([]string, error)

	GetCurrentDatastore(manager Manager) (string, error)

	GetKeyName(manager Manager, datastore, table string) string

	GetColumns(manager Manager, datastore, table string) ([]Column, error)

	GetColumnsByColumnNames(manager Manager, datastore, tableName string, columnNames []string) ([]Column, error)

	CanPersistBatch() bool
}
