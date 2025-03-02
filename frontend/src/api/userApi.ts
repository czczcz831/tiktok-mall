import api from './index';
import { 
  LoginReq, 
  LoginResp, 
  LogoutReq, 
  LogoutResp, 
  RegisterReq, 
  RegisterResp,
  RefreshTokenReq,
  ApiResponse,
  GetUserInfoReq,
  GetUserInfoResp,
  UserRole
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

// 获取用户信息
export const getUserInfo = async (data: GetUserInfoReq): Promise<ApiResponse<GetUserInfoResp>> => {
  // 注意：由于响应拦截器已经返回response.data
  // 因此这里的responseData已经是API返回的数据对象，而不是axios的响应对象
  const responseData = await api.get('/user') as any;
  
  // 调试输出
  console.log('用户信息API响应:', JSON.stringify(responseData));
  
  // 检查用户角色
  let userRole = UserRole.Customer;
  if (responseData.data && Array.isArray(responseData.data.roles)) {
    // 调试输出
    console.log('用户角色列表:', responseData.data.roles);
    
    // 如果角色数组中不包含Customer角色，则视为管理员
    const isAdmin = !responseData.data.roles.includes('Customer');
    userRole = isAdmin ? UserRole.Admin : UserRole.Customer;
    
    // 调试输出
    console.log('是否为管理员:', isAdmin, '设置角色为:', userRole);
  }
  
  // 构造符合前端接口的响应数据
  const adaptedResponse: ApiResponse<GetUserInfoResp> = {
    code: responseData.code,
    msg: responseData.msg,
    data: {
      user: {
        uuid: '', // 后端未返回uuid，暂时使用空字符串
        email: responseData.data.email,
        role: userRole
      }
    }
  };
  
  return adaptedResponse;
}; 