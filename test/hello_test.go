package test

import (
	"fmt"
	"hzw/dsc"
	"testing"
)

func Test_hello1(t *testing.T) {

	conf := &dsc.Config{
		DriverName: "go_ibm_db",
		Descriptor: "HOSTNAME=localhost;DATABASE=testdb;PORT=50003;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=DBSYNCTEST",
	}

	mf, _ := dsc.GetManagerFactory(conf.DriverName)
	maneger, _ := mf.Create(conf)

	result, error := maneger.Execute("select * from T", nil)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	afct, error := result.RowsAffected()
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	fmt.Printf("afct:%v", afct)

}
