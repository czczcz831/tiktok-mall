#*********User*********#
struct LoginReq {
    string email (api.body = "email")
    string password (api.body = "password")
    string confirm_password (api.body = "confirm_password")
}

struct RefreshTokenReq {
    string refresh_token (api.header = "Refresh-Token")
}

struct LoginResp {
    string token
    string refresh_token
}

struct RegisterReq {
    string email (api.body = "email")
    string password (api.body = "password")
    string confirm_password (api.body = "confirm_password")
}

struct RegisterResp {
    string user_uuid
}


service UserService {
    LoginResp Login(1: LoginReq req) (api.post="/user/login", api.body="json")
    LoginResp RefreshToken(1: RefreshTokenReq req) (api.post="/user/refresh_token", api.body="json")
    RegisterResp Register(1: RegisterReq req) (api.post="/user/register", api.body="json")
}

#*********Product*********#

struct Product{
    string uuid (api.body = "uuid")
    string name (api.body = "name")
    string description (api.body = "description")
    i64 price (api.body = "price")
    i64 stock (api.body = "stock")
}

struct CreateProductReq{
    string name (api.body = "name")
    string description (api.body = "description")
    i64 price (api.body = "price")
    i64 stock (api.body = "stock")
}

struct CreateProductResp{
    Product product 
}

struct UpdateProductReq{
    Product product  
}

struct UpdateProductResp{
    Product product 
}

struct DeleteProductReq{
    string uuid (api.path = "uuid")
}

struct DeleteProductResp{
    string uuid
}

struct GetProductReq{
    string uuid (api.path = "uuid")
}

struct GetProductResp{
    Product product
}

struct GetProductListReq{
    i32 page (api.query = "page")
    i32 limit (api.query = "limit")
    optional string name (api.query = "name")
    optional i64 min_price (api.query = "min_price")
    optional i64 max_price (api.query = "max_price")
}

struct GetProductListResp{
    i64 total
    list<Product> products
}

service ProductService {
    CreateProductResp CreateProduct(1: CreateProductReq req) (api.post="/product", api.body="json")
    UpdateProductResp UpdateProduct(1: UpdateProductReq req) (api.put="/product", api.body="json")
    DeleteProductResp DeleteProduct(1: DeleteProductReq req) (api.delete="/product/:uuid" ,)
    GetProductResp GetProduct(1: GetProductReq req) (api.get="/product/:uuid")
    GetProductListResp GetProductList(1: GetProductListReq req) (api.get="/product")
}

#*********Cart*********#

struct CartItem {
    string user_uuid (api.body = "user_uuid")
    string product_uuid (api.body = "product_uuid")
    i64 quantity (api.body = "quantity")
}

struct AddProductToCartReq {
    CartItem item
}

struct AddProductToCartResp {
    CartItem item
}

struct ClearCartReq {
    string user_uuid (api.path = "user_uuid")
}

struct ClearCartResp {
    string user_uuid
}

struct GetCartReq {
    string user_uuid (api.path = "user_uuid")
}

struct GetCartResp {
    i64 total
    list<CartItem> items
}

service CartService {
    AddProductToCartResp AddProductToCart(1: AddProductToCartReq req) (api.post="/cart/add_product", api.body="json")
    ClearCartResp ClearCart(1: ClearCartReq req) (api.delete="/cart/:user_uuid")
    GetCartResp GetCart(1: GetCartReq req) (api.get="/cart/:user_uuid")
}


#*********Checkout*********#

struct Address {
    string uuid (api.body = "uuid")
    string user_uuid (api.body = "user_uuid")
    string street_address (api.body = "street_address")
    string city (api.body = "city")
    string state (api.body = "state")
    string country (api.body = "country")
    i64 zip_code (api.body = "zip_code")
}


struct OrderItem {
    string product_uuid (api.body = "product_uuid")
    i64 price (api.body = "price")
    i64 quantity (api.body = "quantity")
}

struct CreateAddressReq {
    string user_uuid (api.body = "user_uuid")
    string street_address (api.body = "street_address")
    string city (api.body = "city")
    string state (api.body = "state")
    string country (api.body = "country")
    i64 zip_code (api.body = "zip_code")
}

struct CreateAddressResp {
    Address address
}

struct UpdateAddressReq {
    Address address
}

struct UpdateAddressResp {
    Address address
}

struct DeleteAddressReq {
    string uuid (api.path = "uuid")
}

struct DeleteAddressResp {
    string uuid
}

struct GetAddressReq {
    string user_uuid (api.path = "user_uuid")
}

struct GetAddressResp {
    list<Address> addresses
}

struct CheckoutReq {
    string user_uuid (api.body = "user_uuid")
    string first_name (api.body = "first_name")
    string last_name (api.body = "last_name")
    string email (api.body = "email")
    string address_uuid (api.body = "address_uuid")
    string credit_card_uuid (api.body = "credit_card_uuid")
    list<OrderItem> items (api.body = "items")
}

struct CheckoutResp {
    string order_uuid
}

service CheckoutService {
    CreateAddressResp CreateAddress(1: CreateAddressReq req) (api.post="/checkout/address", api.body="json")
    UpdateAddressResp UpdateAddress(1: UpdateAddressReq req) (api.put="/checkout/address", api.body="json")
    DeleteAddressResp DeleteAddress(1: DeleteAddressReq req) (api.delete="/checkout/address/:uuid")
    GetAddressResp GetAddress(1: GetAddressReq req) (api.get="/checkout/address/:user_uuid")

    CheckoutResp Checkout(1: CheckoutReq req) (api.post="/checkout", api.body="json")
}

#************PAYMENT***********8

struct CreditCard {
    string uuid (api.body = "uuid")
    string user_uuid (api.body = "user_uuid")
    string credit_card_number (api.body = "credit_card_number")
    i64 credit_card_cvv (api.body = "credit_card_cvv")
    i64 credit_card_exp_month (api.body = "credit_card_exp_month")
    i64 credit_card_exp_year (api.body = "credit_card_exp_year")
}

struct ChargeReq {
    string order_uuid (api.body = "order_uuid")
    CreditCard credit_card (api.body = "credit_card") 
}

struct ChargeResp {
    string transaction_uuid
}

service PaymentService{
    ChargeResp Charge(1: ChargeReq req) (api.post="/payment/charge", api.body="json")
}
