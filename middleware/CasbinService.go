package middleware

import (
	"gorm.io/driver/postgres"
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

type CasbinService struct {
	Enforcer *casbin.Enforcer
}

func NewCasbinService(dsn string) *CasbinService {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 初始化 Adapter，自动建表
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatalf("初始化 Casbin Adapter 失败: %v", err)
	}

	// 使用基础模型 model.conf（见下）
	enforcer, err := casbin.NewEnforcer("config/casbin_model.conf", adapter)
	if err != nil {
		log.Fatalf("初始化 Casbin Enforcer 失败: %v", err)
	}

	// 同步加载策略
	if err := enforcer.LoadPolicy(); err != nil {
		log.Fatalf("加载策略失败: %v", err)
	}

	return &CasbinService{Enforcer: enforcer}
}
