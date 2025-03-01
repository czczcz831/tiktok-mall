import api from './index';
import {
  AddProductToCartReq,
  AddProductToCartResp,
  ApiResponse,
  ClearCartReq,
  ClearCartResp,
  GetCartReq,
  GetCartResp
} from '../types/api';

// 添加商品到购物车
export const addProductToCart = async (data: AddProductToCartReq): Promise<ApiResponse<AddProductToCartResp>> => {
  return api.post('/cart/add_product', data);
};

// 清空购物车
export const clearCart = async (data: ClearCartReq): Promise<ApiResponse<ClearCartResp>> => {
  return api.delete('/cart');
};

// 获取购物车
export const getCart = async (data: GetCartReq): Promise<ApiResponse<GetCartResp>> => {
  return api.get('/cart');
}; 