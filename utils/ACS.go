package utils

import (
	"BlockApp/db"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v3"
)

var Enforcer *casbin.Enforcer

func init() {
	adapter, _ := gormadapter.NewAdapterByDBWithCustomTable(db.PgsqlDB, nil)
	// 通过mysql适配器新建一个enforcer
	Enforcer, _ = casbin.NewEnforcer("config/keymatch2_model.conf", adapter)
	// 日志记录
	Enforcer.EnableLog(true)
}
