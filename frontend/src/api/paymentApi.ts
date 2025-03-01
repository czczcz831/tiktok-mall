import api from './index';
import {
  ChargeReq,
  ChargeResp,
  ApiResponse
} from '../types/api';

// 支付
export const charge = async (data: ChargeReq): Promise<ApiResponse<ChargeResp>> => {
  return api.post('/payment/charge', data);
}; 