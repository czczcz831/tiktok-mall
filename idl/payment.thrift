namespace go payment

struct CreditCard{
    string credit_card_number
    i64 credit_card_cvv
    i64 credit_card_exp_month
    i64 credit_card_exp_year
}


struct ChargeReq {
    string user_uuid
    string order_uuid
    i64 amount
    CreditCard credit_card
}

struct ChargeResp {
  string transaction_uuid    
  bool success
}

struct CancelChargeReq{
    string user_uuid
    string transaction_uuid    
}

struct CancelChargeResp{
    bool success
}

service PaymentService {
    ChargeResp Charge(1: ChargeReq req)
    CancelChargeResp CancelCharge(1: CancelChargeReq req)
}
