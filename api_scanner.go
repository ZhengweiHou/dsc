package dsc

//Scanner represents a datastore data scanner. This abstraction provides the ability to convert and assign datastore record of data to provided destination
type Scanner interface {
	//Returns all columns specified in select statement.
	Columns() ([]string, error)

	//ColumnTypes return column types
	ColumnTypes() ([]ColumnType, error)

	//Scans datastore record data to convert and assign it to provided destinations, a destination needs to be pointer.
	Scan(dest ...interface{}) error
}
