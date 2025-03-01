import api from './index';
import {
  CreateProductReq,
  CreateProductResp,
  UpdateProductReq,
  UpdateProductResp,
  DeleteProductReq,
  DeleteProductResp,
  GetProductReq,
  GetProductResp,
  GetProductListReq,
  GetProductListResp,
  ApiResponse,
  ProductListData
} from '../types/api';

// 创建商品
export const createProduct = async (data: CreateProductReq): Promise<ApiResponse<CreateProductResp>> => {
  // 将价格转换为分
  const convertedData = {
    ...data,
    price: Math.round(data.price * 100)
  };
  return api.post('/product', convertedData);
};

// 更新商品
export const updateProduct = async (data: UpdateProductReq): Promise<ApiResponse<UpdateProductResp>> => {
  // 将价格转换为分
  const convertedData = {
    ...data,
    product: {
      ...data.product,
      price: Math.round(data.product.price * 100)
    }
  };
  return api.put('/product', convertedData);
};

// 删除商品
export const deleteProduct = async (data: DeleteProductReq): Promise<ApiResponse<DeleteProductResp>> => {
  return api.delete(`/product/${data.uuid}`);
};

// 获取单个商品
export const getProduct = async (data: GetProductReq): Promise<ApiResponse<GetProductResp>> => {
  return api.get(`/product/${data.uuid}`);
};

// 获取商品列表
export const getProductList = async (data: GetProductListReq): Promise<ApiResponse<GetProductListResp>> => {
  const { page, limit, name, min_price, max_price } = data;
  let url = `/product?page=${page}&limit=${limit}`;
  
  if (name) {
    url += `&name=${name}`;
  }
  
  if (min_price !== undefined) {
    const minPriceInCents = Math.floor(min_price * 100);
    url += `&min_price=${minPriceInCents}`;
  }
  
  if (max_price !== undefined) {
    const maxPriceInCents = Math.ceil(max_price * 100);
    url += `&max_price=${maxPriceInCents}`;
  }
  
  return api.get(url);
}; 