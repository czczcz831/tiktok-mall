import axios, { AxiosRequestConfig } from 'axios';

// 创建axios实例
const api = axios.create({
  baseURL: 'https://tiktok-mall-api.czczcz.xyz', // 后端服务地址，根据实际情况修改
  headers: {
    'Content-Type': 'application/json',
  },
});

// 请求拦截器
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token && config.headers) {
      config.headers.Authorization = `${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
api.interceptors.response.use(
  (response) => {
    return response.data;
  },
  async (error) => {
    // 处理401错误，可能是token过期
    if (error.response && error.response.status === 401) {
      const refreshToken = localStorage.getItem('refresh_token');
      // 如果有refresh_token，尝试刷新token
      if (refreshToken) {
        try {
          const res = await refreshTokenApi(refreshToken);
          localStorage.setItem('token', res.token);
          localStorage.setItem('refresh_token', res.refresh_token);
          
          // 重新请求
          const config = error.config;
          config.headers.Authorization = `Bearer ${res.token}`;
          return api(config);
        } catch (err) {
          // 刷新token失败，清除本地存储并跳转到登录页
          localStorage.removeItem('token');
          localStorage.removeItem('refresh_token');
          window.location.href = '/login';
          return Promise.reject(err);
        }
      } else {
        // 没有refresh_token，直接跳转到登录页
        window.location.href = '/login';
      }
    }
    return Promise.reject(error);
  }
);

// 刷新token的方法
const refreshTokenApi = async (refreshToken: string) => {
  const response = await axios.post(
    `${api.defaults.baseURL}/user/refresh_token`,
    {},
    {
      headers: {
        'Refresh-Token': refreshToken,
      },
    }
  );
  return response.data;
};

export default api; 