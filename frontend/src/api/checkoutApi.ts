import api from './index';
import {
  CreateAddressReq,
  CreateAddressResp,
  UpdateAddressReq,
  UpdateAddressResp,
  DeleteAddressReq,
  DeleteAddressResp,
  GetAddressReq,
  GetAddressResp,
  CheckoutReq,
  CheckoutResp,
  ApiResponse
} from '../types/api';

// 创建地址
export const createAddress = async (data: CreateAddressReq): Promise<ApiResponse<CreateAddressResp>> => {
  return api.post('/checkout/address', data);
};

// 更新地址
export const updateAddress = async (data: UpdateAddressReq): Promise<ApiResponse<UpdateAddressResp>> => {
  return api.put('/checkout/address', data);
};

// 删除地址
export const deleteAddress = async (data: DeleteAddressReq): Promise<ApiResponse<DeleteAddressResp>> => {
  return api.delete(`/checkout/address/${data.uuid}`);
};

// 获取地址列表
export const getAddress = async (data: GetAddressReq): Promise<ApiResponse<GetAddressResp>> => {
  return api.get('/checkout/address');
};

// 结账
export const checkout = async (data: CheckoutReq): Promise<ApiResponse<CheckoutResp>> => {
  return api.post('/checkout', data);
}; 