package dsc

import "time"

type TransactionManager interface {
	Begin() error

	Commit() error

	Rollback() error
}

//Connection  数据源连接
type Connection interface {
	Config() *Config

	ConnectionPool() chan Connection

	//Close 连接关闭或返回连接池中
	Close() error

	// CloseNow 连接立刻关闭，不返回连接池
	CloseNow() error

	Unwrap(target interface{}) interface{}

	LastUsed() *time.Time

	SetLastUsed(ts *time.Time)

	TransactionManager
}

//ConnectionProvider 数据源连接提供者
type ConnectionProvider interface {
	Get() (Connection, error)

	Config() *Config

	// 获取连接池
	ConnectionPool() chan Connection

	// 尝试初始化连接池
	SpawnConnectionIfNeeded()

	NewConnection() (Connection, error)

	Close() error
}
