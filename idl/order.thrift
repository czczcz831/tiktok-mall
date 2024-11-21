namespace go order

include "checkout.thrift"


struct Order {
    string uuid
    string user_uuid
    i64 total
    bool is_paid
    i64 created_at
    list<checkout.OrderItem> items
}

struct CreateOrderReq {
    string user_uuid
    checkout.Address address
    list<checkout.OrderItem> items
}

struct CreateOrderResp{
    Order order;
}

struct MarkOrderPaidReq {
    string uuid
}

struct MarkOrderPaidResp {
    Order order
}