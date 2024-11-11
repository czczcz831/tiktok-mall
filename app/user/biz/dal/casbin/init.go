package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/user/biz/dal/mysql"

	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var (
	CasbinEnforcer *casbin.Enforcer
)

func Init() {
	a, err := gormadapter.NewAdapterByDB(mysql.DB)

	if err != nil {
		klog.Fatalf("new casbin enforcer failed: %v", err)
	}

	CasbinEnforcer, err = casbin.NewEnforcer("casbin.conf", a)
	if err != nil {
		klog.Fatalf("new casbin enforcer failed: %v", err)
	}
	CasbinEnforcer.EnableAutoSave(true)

	err = CasbinEnforcer.LoadPolicy()
	if err != nil {
		klog.Fatalf("load policy failed: %v", err)
	}

	//初始化
	//AdminRole
	CasbinEnforcer.AddPolicy("AdminRole", "*", "*")
	//CustomerRole
	CasbinEnforcer.AddPolicy("CustomerRole", "/api/v1/customer/.*", "*")
	//SellerRole
	CasbinEnforcer.AddPolicy("SellerRole", "/api/v1/seller/.*", "*")

	//Superuser
	CasbinEnforcer.AddRoleForUser("superuser@admin.com", "AdminRole")

	CasbinEnforcer.SavePolicy()

	if err != nil {
		klog.Fatalf("load policy failed: %v", err)
	}

}
