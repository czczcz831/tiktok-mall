namespace go payment

include "checkout.thrift"

struct ChargeReq {
    string user_uuid
    string order_uuid
    i64 amount
    checkout.CreditCard credit_card
}
