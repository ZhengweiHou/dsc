package dsc

import (
	"fmt"
	"sync"
)

var managerFactories_mux *sync.RWMutex = &sync.RWMutex{}
var managerFactories = make(map[string]ManagerFactory)

// RegisterManagerFactory 注册ManagerFactory
func RegisterManagerFactory(driver string, factory ManagerFactory) {
	managerFactories_mux.Lock()
	defer managerFactories_mux.Unlock()
	managerFactories[driver] = factory
}

// GetManagerFactory 获取ManagerFactory
func GetManagerFactory(driver string) (ManagerFactory, error) {
	managerFactories_mux.RLock()
	result, ok := managerFactories[driver]
	managerFactories_mux.RUnlock()

	if ok {
		return result, nil
	}

	return nil, fmt.Errorf("failed to lookup manager factory for '%v' ", driver)
}
