package dsc

import "reflect"

//Column represents TableColumn type interface (compabible with *sql.ColumnType
type Column interface {
	ColumnType
	// Name returns the name or alias of the TableColumn.
	Name() string
}

type ColumnType interface {
	// Length returns the TableColumn type length for variable length TableColumn types such
	// as text and binary field types. If the type length is unbounded the value will
	// be math.MaxInt64 (any database limits will still apply).
	// If the TableColumn type is not variable length, such as an int, or if not supported
	// by the driver ok is false.
	Length() (length int64, ok bool)

	// DecimalSize returns the scale and precision of a decimal type.
	// If not applicable or if not supported ok is false.
	DecimalSize() (precision, scale int64, ok bool)

	// ScanType returns a Go type suitable for scanning into using Rows.Scan.
	// If a driver does not support this property ScanType will return
	// the type of an empty interface.
	ScanType() reflect.Type

	// Nullable returns whether the TableColumn may be null.
	// If a driver does not support this property ok will be false.
	Nullable() (nullable, ok bool)

	// DatabaseTypeName returns the database system name of the TableColumn type. If an empty
	// string is returned the driver type name is not supported.
	// Consult your driver documentation for a list of driver data types. Length specifiers
	// are not included.
	// Common type include "VARCHAR", "TEXT", "NVARCHAR", "DECIMAL", "BOOL", "INT", "BIGINT".
	DatabaseTypeName() string
}
