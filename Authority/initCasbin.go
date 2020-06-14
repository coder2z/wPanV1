package Authority

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"wPan/v1/Config"
)

var Enforcer *casbin.Enforcer

func InitCasbin() {
	a, _ := gormadapter.NewAdapter("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			Config.DatabaseSetting.User,
			Config.DatabaseSetting.Password,
			Config.DatabaseSetting.Host,
			Config.DatabaseSetting.Name), true)

	Enforcer, _ = casbin.NewEnforcer("Authority/rbac_model.conf", a)
	_ = Enforcer.LoadPolicy()
}
