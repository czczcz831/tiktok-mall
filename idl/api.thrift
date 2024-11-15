namespace go api

struct LoginReq {
    string email (api.body = "email")
    string password (api.body = "password")
    string confirm_password (api.body = "confirm_password")
}

struct RefreshTokenReq {
    string refresh_token (api.header = "Refresh-Token")
}

struct LoginResp {
    string token
    string refresh_token
}

struct RegisterReq {
    string email (api.body = "email")
    string password (api.body = "password")
    string confirm_password (api.body = "confirm_password")
}

struct RegisterResp {
    string user_uuid
}


service UserService {
    LoginResp Login(1: LoginReq req) (api.post="/user/login", api.body="json")
    LoginResp RefreshToken(1: RefreshTokenReq req) (api.post="/user/refresh_token", api.body="json")
    RegisterResp Register(1: RegisterReq req) (api.post="/user/register", api.body="json")
}