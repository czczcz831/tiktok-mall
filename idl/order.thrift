namespace go order

struct OrderItem{
    string product_uuid
    i64 price
    i64 quantity
}

struct Order {
    string uuid
    string user_uuid
    string address_uuid
    i64 total
    i32 status
    i64 created_at
    list<OrderItem> items
}

struct CreateOrderReq {
    string user_uuid
    string address_uuid
    i64 total
    list<OrderItem> items
}

struct UpdateOrderAddressReq {
    string uuid
    string user_uuid
    string address_uuid
}

struct UpdateOrderAddressResp {
    string address_uuid
}

struct CreateOrderResp{
    Order order;
}

struct GetOrderReq {
    string uuid
}

struct GetOrderResp {
    Order order
}

struct MarkOrderPaidReq {
    string uuid
}

struct MarkOrderPaidResp {
    Order order
}

struct GetUserOrdersReq{
    string user_uuid
}

struct GetUserOrdersResp{
    i64 total
    list<Order> orders
}

service OrderService{
    CreateOrderResp CreateOrder(1: CreateOrderReq req)
    UpdateOrderAddressResp UpdateOrderAddress(1: UpdateOrderAddressReq req)
    MarkOrderPaidResp MarkOrderPaid(1: MarkOrderPaidReq req)
    GetOrderResp GetOrder(1: GetOrderReq req)
    GetUserOrdersResp GetUserOrders(1: GetUserOrdersReq req)
}
