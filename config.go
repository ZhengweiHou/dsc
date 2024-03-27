package dsc

import (
	"sync"

	"github.com/viant/toolbox/cred"
)

//Config represent datastore config.
type Config struct {
	URL                 string
	DriverName          string
	PoolSize            int
	MaxPoolSize         int
	Descriptor          string
	Parameters          map[string]interface{}
	Credentials         string
	MaxRequestPerSecond int
	cred                string
	username            string
	password            string
	dsnDescriptor       string
	lock                *sync.Mutex
	race                uint32
	initRun             bool
	CredConfig          *cred.Config `json:"-"`
}
