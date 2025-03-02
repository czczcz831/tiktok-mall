// Code generated by hertz generator.

package api

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/czczcz831/tiktok-mall/app/api/biz/dal/casbin"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

// 更新商品需要角色拥有SELLER_OBJECT权限

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _refreshtokenMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}

// 更新商品需要角色拥有SELLER_OBJECT权限
func _updateproductMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		casbin.CasbinHertzMiddleware.RequiresPermissions(casbin.SELLER_OBJECT),
	}
}

func _getproductlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _productMw() []app.HandlerFunc {
	// your code...
	return nil
}

// 创建商品需要角色拥有SELLER_OBJECT权限
func _createproductMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		casbin.CasbinHertzMiddleware.RequiresPermissions(casbin.SELLER_OBJECT),
	}
}

// 删除商品需要角色拥有SELLER_OBJECT权限
func _deleteproductMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		casbin.CasbinHertzMiddleware.RequiresPermissions(casbin.SELLER_OBJECT),
	}
}

func _getproductMw() []app.HandlerFunc {
	// your code...
	return nil
}

// 购物车组需要角色拥有CUSTOMER_OBJECT权限
func _cartMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		casbin.CasbinHertzMiddleware.RequiresPermissions(casbin.CUSTOMER_OBJECT),
	}
}

func _clearcartMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		casbin.CasbinHertzMiddleware.RequiresPermissions(casbin.CUSTOMER_OBJECT),
	}
}

func _getcartMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		casbin.CasbinHertzMiddleware.RequiresPermissions(casbin.CUSTOMER_OBJECT),
	}
}

func _addproducttocartMw() []app.HandlerFunc {
	// your code...
	return nil
}

// 结算需要角色拥有CUSTOMER_OBJECT权限
func _checkoutMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		casbin.CasbinHertzMiddleware.RequiresPermissions(casbin.CUSTOMER_OBJECT),
	}
}

// 结算组需要角色拥有CUSTOMER_OBJECT权限
func _checkout0Mw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		casbin.CasbinHertzMiddleware.RequiresPermissions(casbin.CUSTOMER_OBJECT),
	}
}

func _chargeMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _checkout1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateaddressMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _credit_cardMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createcreditcardMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getcreditcardMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deletecreditcardMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatecreditcardMw() []app.HandlerFunc {
	// your code...
	return nil
}

// 地址组需要角色拥有CUSTOMER_OBJECT权限
func _addressMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		casbin.CasbinHertzMiddleware.RequiresPermissions(casbin.CUSTOMER_OBJECT),
	}
}

func _createaddressMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getaddressMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deleteaddressMw() []app.HandlerFunc {
	// your code...
	return nil
}

// 支付需要角色拥有CUSTOMER_OBJECT权限
func _paymentMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		casbin.CasbinHertzMiddleware.RequiresPermissions(casbin.CUSTOMER_OBJECT),
	}
}

// 登出需要角色拥有CUSTOMER_OBJECT权限
func _logoutMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		casbin.CasbinHertzMiddleware.RequiresPermissions(casbin.CUSTOMER_OBJECT),
	}
}

// 获取用户订单需要角色拥有CUSTOMER_OBJECT权限
func _getuserordersMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		casbin.CasbinHertzMiddleware.RequiresPermissions(casbin.CUSTOMER_OBJECT),
	}
}

// 获取用户信息需要角色拥有CUSTOMER_OBJECT权限
func _getuserinfoMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		casbin.CasbinHertzMiddleware.RequiresPermissions(casbin.CUSTOMER_OBJECT),
	}
}
