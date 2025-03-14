namespace go auth

service AuthService {
    DeliveryResp DeliverTokenByRPC(1: DeliverTokenReq req)
    DeliveryResp RefeshTokenByRPC(1: RefeshTokenReq req)
}

struct RefeshTokenReq {
    1: string refresh_token
}

struct DeliverTokenReq {
    1: string user_uuid
}

struct DeliveryResp {
    1: string token
    2: string refresh_token
    3: i64 token_expire_after
    4: i64 refresh_token_expire_after
}