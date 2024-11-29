package casbin

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/api/conf"
	"github.com/czczcz831/tiktok-mall/app/user/biz/dal/mysql"
	"github.com/czczcz831/tiktok-mall/common/utils"

	gormadapter "github.com/casbin/gorm-adapter/v3"
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
)

const (
	CUSTOMER_OBJECT = "customer_obj"
	SELLER_OBJECT   = "seller_obj"
)

func SubjectFromToken(ctx context.Context, c *app.RequestContext) string {
	token := c.GetRequest().Header.Get("Authorization")
	if token == "" {
		return ""
	}

	publicKeyHexString := conf.GetConf().JWT.PublicSecret

	uuid, _, err := utils.VerifyToken(token, publicKeyHexString)

	if err != nil {
		return ""
	}

	return uuid
}

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
	CasbinEnforcer.AddPolicy(ADMIN_ROLE, "*", "*")
	//CustomerRole
	CasbinEnforcer.AddPolicy(CUSTOMER_ROLE, CUSTOMER_OBJECT, "*")
	//SellerRole
	CasbinEnforcer.AddPolicy(SELLER_ROLE, SELLER_OBJECT, "*")
	CasbinEnforcer.AddPolicy(SELLER_ROLE, CUSTOMER_OBJECT, "*")

	//Superuser
	CasbinEnforcer.AddRoleForUser("superuser-uuid", ADMIN_ROLE)

	CasbinEnforcer.SavePolicy()

	if err != nil {
		klog.Fatalf("load policy failed: %v", err)
	}

	CasbinHertzMiddleware, err = casbinHertz.NewCasbinMiddlewareFromEnforcer(CasbinEnforcer, SubjectFromToken)
	if err != nil {
		klog.Fatalf("new casbin middleware failed: %v", err)
	}

}
