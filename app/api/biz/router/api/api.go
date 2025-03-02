// Code generated by hertz generator. DO NOT EDIT.

package api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	api "github.com/czczcz831/tiktok-mall/app/api/biz/handler/api"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.DELETE("/cart", append(_clearcartMw(), api.ClearCart)...)
	root.GET("/cart", append(_getcartMw(), api.GetCart)...)
	root.POST("/checkout", append(_checkoutMw(), api.Checkout)...)
	root.GET("/orders", append(_getuserordersMw(), api.GetUserOrders)...)
	root.PUT("/product", append(_updateproductMw(), api.UpdateProduct)...)
	root.GET("/product", append(_getproductlistMw(), api.GetProductList)...)
	root.GET("/user", append(_getuserinfoMw(), api.GetUserInfo)...)
	{
		_cart := root.Group("/cart", _cartMw()...)
		_cart.POST("/add_product", append(_addproducttocartMw(), api.AddProductToCart)...)
	}
	{
		_checkout0 := root.Group("/checkout", _checkout0Mw()...)
		_checkout0.PUT("/address", append(_updateaddressMw(), api.UpdateAddress)...)
		_checkout0.GET("/address", append(_getaddressMw(), api.GetAddress)...)
		_checkout0.POST("/address", append(_createaddressMw(), api.CreateAddress)...)
		_address := _checkout0.Group("/address", _addressMw()...)
		_address.DELETE("/:uuid", append(_deleteaddressMw(), api.DeleteAddress)...)
	}
	{
		_eino := root.Group("/eino", _einoMw()...)
		_eino.POST("/chat", append(_callassistantagentMw(), api.CallAssistantAgent)...)
	}
	{
		_payment := root.Group("/payment", _paymentMw()...)
		_payment.POST("/charge", append(_chargeMw(), api.Charge)...)
	}
	root.POST("/product", append(_createproductMw(), api.CreateProduct)...)
	_product := root.Group("/product", _productMw()...)
	_product.DELETE("/:uuid", append(_deleteproductMw(), api.DeleteProduct)...)
	_product.GET("/:uuid", append(_getproductMw(), api.GetProduct)...)
	{
		_user := root.Group("/user", _userMw()...)
		_user.POST("/login", append(_loginMw(), api.Login)...)
		_user.POST("/logout", append(_logoutMw(), api.Logout)...)
		_user.POST("/refresh_token", append(_refreshtokenMw(), api.RefreshToken)...)
		_user.POST("/register", append(_registerMw(), api.Register)...)
	}
}
