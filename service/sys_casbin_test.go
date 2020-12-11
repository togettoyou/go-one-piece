package service

import (
	"go-one-server/model"
	"go-one-server/util/conf"
	"testing"
)

func TestCasbin(t *testing.T) {
	conf.Setup()
	model.Setup()
	model.DB().Exec("SHOW TABLES")
	Casbin()
}
