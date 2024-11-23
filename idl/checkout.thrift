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

struct CreditCard{
    string uuid
    string user_uuid
    string credit_card_number
    i64 credit_card_cvv
    i64 credit_card_exp_month
    i64 credit_card_exp_year
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



#CreditCard#
struct CreateCreditCardReq{
    string user_uuid
    string credit_card_number
    i64 credit_card_cvv
    i64 credit_card_exp_month
    i64 credit_card_exp_year
}

struct CreateCreditCardResp{
    CreditCard credit_card
}

struct UpdateCreditCardReq{
    CreditCard credit_card
}

struct UpdateCreditCardResp{
    CreditCard credit_card
}

struct DeleteCreditCardReq{
    string uuid
}

struct DeleteCreditCardResp{
    string uuid
}

struct GetCreditCardReq{
    string user_uuid
}

struct GetCreditCardResp{
    list<CreditCard> credit_cards
}

#Checkout#
struct CheckoutReq{
    string user_uuid
    string first_name
    string last_name
    string email
    string address_uuid
    string credit_card_uuid
    list<OrderItem> items
}

struct CheckoutResp{
    string order_uuid
}

struct ChargeReq{
    string order_uuid
    string payment_uuid
}

struct ChargeResp{
    string transaction_uuid
}

service CheckoutService {
    #Address
    CreateAddressResp CreateAddress(1: CreateAddressReq req)
    UpdateAddressResp UpdateAddress(1: UpdateAddressReq req)
    DeleteAddressResp DeleteAddress(1: DeleteAddressReq req)
    GetAddressResp GetAddress(1: GetAddressReq req)

    #CreditCard
    CreateCreditCardResp CreateCreditCard(1: CreateCreditCardReq req)
    UpdateCreditCardResp UpdateCreditCard(1: UpdateCreditCardReq req)
    DeleteCreditCardResp DeleteCreditCard(1: DeleteCreditCardReq req)
    GetCreditCardResp GetCreditCard(1: GetCreditCardReq req)

    #Checkout
    CheckoutResp Checkout(1: CheckoutReq req)

    #Charge
    ChargeResp Charge(1: ChargeReq req)
}