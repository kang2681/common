package drivers

import (
	"testing"
)

func TestNewMysqlWithConfig(t *testing.T) {
	conf := MysqlConfig{
		Username: "root",
		Password: "123456",
		DBName:   "ztlog",
	}
	_, err := NewMysqlWithConfig(&conf)
	if err != nil {
		t.Error(err)
	}
}
