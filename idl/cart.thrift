namespace go cart

###########Export#####################
struct CartItem {
    string user_uuid
    string product_uuid
    i32 quantity
}
######################################

struct AddProductToCartReq {
    CartItem item
}

struct AddProductToCartResp {
    CartItem item
}

struct ClearCartReq {
    string user_uuid
}

struct ClearCartResp {
    string user_uuid
}

struct GetCartReq {
    string user_uuid
}

struct GetCartResp {
    list<CartItem> items
}

service CartService {
    AddProductToCartResp AddProductToCart(1: AddProductToCartReq req)
    ClearCartResp ClearCart(1: ClearCartReq req)
    GetCartResp GetCart(1: GetCartReq req)
}