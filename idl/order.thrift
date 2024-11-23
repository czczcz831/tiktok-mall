namespace go order

struct OrderItem{
    string product_uuid
    i64 price
    i64 quantity
}

struct Order {
    string uuid
    string user_uuid
    i64 total
    bool is_paid
    i64 created_at
    list<OrderItem> items
}

struct CreateOrderReq {
    string user_uuid
    i64 total
    list<OrderItem> items
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

service OrderService{
    CreateOrderResp CreateOrder(1: CreateOrderReq req)
    MarkOrderPaidResp MarkOrderPaid(1: MarkOrderPaidReq req)
}