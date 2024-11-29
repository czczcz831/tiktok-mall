namespace go checkout

struct Address{
    string uuid
    string user_uuid
    string street_address
    string city
    string state
    string country
    i64 zip_code
}


struct OrderItem{
    string product_uuid
    i64 price
    i64 quantity
}

#Address#
struct CreateAddressReq{
    string user_uuid
    string street_address
    string city
    string state
    string country
    i64 zip_code
}

struct CreateAddressResp{
    Address address
}

struct UpdateAddressReq{
    Address address
}

struct UpdateAddressResp{
    Address address
}

struct DeleteAddressReq{
    string uuid
}

struct DeleteAddressResp{
    string uuid
}

struct GetAddressReq{
    string user_uuid
}

struct GetAddressResp{
    list<Address> addresses
}


#Checkout#
struct CheckoutReq{
    string user_uuid
    string first_name
    string last_name
    string email
    string address_uuid
    list<OrderItem> items
}

struct CheckoutResp{
    string order_uuid
}

service CheckoutService {
    #Address
    CreateAddressResp CreateAddress(1: CreateAddressReq req)
    UpdateAddressResp UpdateAddress(1: UpdateAddressReq req)
    DeleteAddressResp DeleteAddress(1: DeleteAddressReq req)
    GetAddressResp GetAddress(1: GetAddressReq req)

    #Checkout
    CheckoutResp Checkout(1: CheckoutReq req)
}
