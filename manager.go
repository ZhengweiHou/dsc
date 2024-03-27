package dsc

import "database/sql"

type AbstractManager struct {
	Manager
	config             *Config
	connectionProvider ConnectionProvider
}

func (m *AbstractManager) Config() *Config {
	return m.config
}

func (m *AbstractManager) ConnectionProvider() ConnectionProvider {
	return m.connectionProvider
}

func (m *AbstractManager) Execute(sql string, sqlParameters ...interface{}) (result sql.Result, err error) {
	var connection Connection
	connection, err = m.Manager.ConnectionProvider().Get()
	if err != nil {
		return nil, err
	}
	defer connection.Close()
	return m.Manager.ExecuteOnConnection(connection, sql, sqlParameters)
}

func (m *AbstractManager) ReadAllWithHandler(query string, queryParameters []interface{}, readingHandler func(scanner Scanner) (toContinue bool, err error)) error {
	connection, err := m.Manager.ConnectionProvider().Get()
	if err != nil {
		return err
	}
	defer connection.Close()
	return m.Manager.ReadAllWithHandlerOnConnection(connection, query, queryParameters, readingHandler)
}

// AbstractManager不实现 xxxOnConnection， 由子类实现，被上面的方法调用
// func (m *AbstractManager) ExecuteOnConnection(connection Connection, sql string, parameters []interface{}) (sql.Result, error)
// func (m *AbstractManager) ReadAllWithHandlerOnConnection(connection Connection, query string, parameters []interface{}, readingHandler func(scanner Scanner) (toContinue bool, err error)) error
