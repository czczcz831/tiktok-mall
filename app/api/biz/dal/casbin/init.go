package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/czczcz831/tiktok-mall/app/api/biz/dal/mysql"

	"context"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/cloudwego/hertz/pkg/app"
	casbinHertz "github.com/hertz-contrib/casbin"
)

var (
	CasbinEnforcer        *casbin.Enforcer
	CasbinHertzMiddleware *casbinHertz.Middleware
)

const (
	ADMIN_ROLE    = "Admin"
	CUSTOMER_ROLE = "Customer"
	SELLER_ROLE   = "Seller"
	BANNED_ROLE   = "Banned"
)

const (
	CUSTOMER_OBJECT = "customer_obj"
	SELLER_OBJECT   = "seller_obj"
)

func Init() {
	a, err := gormadapter.NewAdapterByDB(mysql.DB)
	if err != nil {
		hlog.Fatalf("new casbin enforcer failed: %v", err)
	}

	CasbinEnforcer, err = casbin.NewEnforcer("casbin.conf", a)
	if err != nil {
		hlog.Fatalf("new casbin enforcer failed: %v", err)
	}
	CasbinEnforcer.EnableAutoSave(true)

	err = CasbinEnforcer.LoadPolicy()
	if err != nil {
		hlog.Fatalf("load policy failed: %v", err)
	}

	// 初始化
	// AdminRole
	CasbinEnforcer.AddPolicy(ADMIN_ROLE, ".*") // Admin 可以访问所有资源
	//CustomerRole
	CasbinEnforcer.AddPolicy(CUSTOMER_ROLE, CUSTOMER_OBJECT) // Customer 只能访问 customer_obj
	// SellerRole
	CasbinEnforcer.AddPolicy(SELLER_ROLE, SELLER_OBJECT)   // Seller 只能访问 seller_obj
	CasbinEnforcer.AddPolicy(SELLER_ROLE, CUSTOMER_OBJECT) // Seller 也可以访问 customer_obj

	// Superuser
	CasbinEnforcer.AddRoleForUser("0000000000000000000", ADMIN_ROLE)

	CasbinEnforcer.SavePolicy()

	if err != nil {
		hlog.Fatalf("load policy failed: %v", err)
	}

	CasbinHertzMiddleware, err = casbinHertz.NewCasbinMiddlewareFromEnforcer(CasbinEnforcer, func(ctx context.Context, c *app.RequestContext) string {
		uuid, ok := c.Get("uuid")
		if !ok {
			return ""
		}
		return uuid.(string)
	})
	if err != nil {
		hlog.Fatalf("new casbin middleware failed: %v", err)
	}
}
