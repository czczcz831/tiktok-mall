namespace go eino

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

struct QueryUserOrdersReq{
    string user_uuid
    string query_content
}

struct QueryUserOrdersResp{
    i64 total
    list<Order> orders;
}

service EinoService{
    QueryUserOrdersResp QueryUserOrders(1: QueryUserOrdersReq req)
}


