import api from './index';
import { 
  LoginReq, 
  LoginResp, 
  LogoutReq, 
  LogoutResp, 
  RegisterReq, 
  RegisterResp,
  RefreshTokenReq,
  ApiResponse
} from '../types/api';

// 登录
export const login = async (data: LoginReq): Promise<ApiResponse<LoginResp>> => {
  return api.post('/user/login', data);
};

// 登出
export const logout = async (data: LogoutReq): Promise<ApiResponse<LogoutResp>> => {
  return api.post('/user/logout', data);
};

// 刷新token
export const refreshToken = async (data: RefreshTokenReq): Promise<ApiResponse<LoginResp>> => {
  return api.post('/user/refresh_token', {}, {
    headers: {
      'Refresh-Token': data.refresh_token
    }
  });
};

// 注册
export const register = async (data: RegisterReq): Promise<ApiResponse<RegisterResp>> => {
  return api.post('/user/register', data);
}; 