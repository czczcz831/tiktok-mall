namespace go api

struct BaseResp {
    string status_message = "" 
    i32 status_code = 0
}

struct LoginReq {
    BaseResp base_resp
    string email (api.body = "email")
    string password (api.body = "password")
    string confirm_password (api.body = "confirm_password")
}

struct LoginResp {
    string token
    string refresh_token
}

service UserService {
    LoginResp login(1: LoginReq req) (api.post="/user/login", api.body="json")
}