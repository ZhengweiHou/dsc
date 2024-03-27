package dsc

import "database/sql"

//Manager
type Manager interface {
	Config() *Config

	ConnectionProvider() ConnectionProvider

	Execute(sql string, parameters ...interface{}) (sql.Result, error)

	ExecuteOnConnection(connection Connection, sql string, parameters []interface{}) (sql.Result, error)

	ReadAllWithHandler(query string, parameters []interface{}, readingHandler func(scanner Scanner) (toContinue bool, err error)) error

	ReadAllWithHandlerOnConnection(connection Connection, query string, parameters []interface{}, readingHandler func(scanner Scanner) (toContinue bool, err error)) error

	// TODO 数据持久化
	// PersistData(connection Connection, data interface{}, table string, keySetter KeySetter, sqlProvider func(item interface{}) *ParametrizedSQL) (int, error)

}

//ManagerFactory manager工厂，用于创建manager
type ManagerFactory interface {
	//Creates manager, takes config pointer.
	Create(config *Config) (Manager, error)

	//Creates manager from url, can url needs to point to Config JSON.
	CreateFromURL(url string) (Manager, error)
}

//ManagerRegistry manager注册器，用于程序数据源管理
type ManagerRegistry interface {
	Get(name string) Manager

	Register(name string, manager Manager)
}
