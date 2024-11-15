namespace go user

service UserService {
    RegisterResp register(1: RegisterReq req)
    LoginResp login(1: LoginReq req)
}

struct RegisterReq {
    1: string email
    2: string password
}

struct RegisterResp {
    1: string user_uuid
}

struct LoginReq {
    1: string email
    2: string password
}

struct LoginResp {
    1: string user_uuid
}