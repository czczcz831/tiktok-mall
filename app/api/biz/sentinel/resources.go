package sentinel

const (
	// UserService
	APIUserLogin        = "POST:/user/login"
	APIUserRefreshToken = "POST:/user/refresh_token"
	APIUserRegister     = "POST:/user/register"

	// ProductService
	APIProductCreate  = "POST:/product"
	APIProductUpdate  = "PUT:/product"
	APIProductDelete  = "DELETE:/product/:uuid"
	APIProductGet     = "GET:/product/:uuid"
	APIProductGetList = "GET:/product"

	// CartService
	APICartAddProduct = "POST:/cart/add_product"
	APICartClear      = "DELETE:/cart/:user_uuid"
	APICartGet        = "GET:/cart/:user_uuid"

	// CheckoutService
	APICheckoutCreateAddress = "POST:/checkout/address"
	APICheckoutUpdateAddress = "PUT:/checkout/address"
	APICheckoutDeleteAddress = "DELETE:/checkout/address/:uuid"
	APICheckoutGetAddress    = "GET:/checkout/address/:user_uuid"
	APICheckout              = "POST:/checkout"

	// PaymentService
	APIPaymentCharge = "POST:/payment/charge"
)

// Control clients in PROJECT_DIR/client/*, only works in api.
const (
	// UserService
	RpcCallUserLogin    = "user:login"
	RpcCallUserRegister = "user:register"

	//TODO: More to be added

)
