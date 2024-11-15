namespace go product

struct Product{
    string uuid
    string name
    string description
    i64 price
}

struct CreateProductReq{
    string name
    string description
    i64 price
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
    string uuid
}

struct DeleteProductResp{
    string uuid
}

struct GetProductReq{
    string uuid
}

struct GetProductResp{
    Product product
}

struct GetProductListReq{
    i32 page
    i32 limit
    optional string name
    optional i64 min_price
    optional i64 max_price
}

struct GetProductListResp{
    i64 total
    list<Product> products
}

service ProductService {
    CreateProductResp CreateProduct(1: CreateProductReq req)
    UpdateProductResp UpdateProduct(1: UpdateProductReq req)
    DeleteProductResp DeleteProduct(1: DeleteProductReq req)
    GetProductResp GetProduct(1: GetProductReq req)
    GetProductListResp GetProductList(1: GetProductListReq req)
}