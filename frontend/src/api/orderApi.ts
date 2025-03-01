import api from './index';
import { ApiResponse, CreditCard, ChargeResp } from '../types/api';

interface GetOrdersResponse {
  total: number;
  orders: Order[];
}

export interface Order {
  uuid: string;
  user_uuid: string;
  address_uuid: string;
  total: number;
  status: number;
  created_at: number;
  items: OrderItem[];
}

export interface OrderItem {
  product_uuid: string;
  price: number;
  quantity: number;
}

export const getOrders = async (): Promise<ApiResponse<GetOrdersResponse>> => {
  return api.get('/orders');
};

export const payOrder = async (orderUuid: string, creditCard: CreditCard): Promise<ApiResponse<ChargeResp>> => {
  return api.post('/payment/charge', {
    order_uuid: orderUuid,
    credit_card: creditCard
  });
}; 