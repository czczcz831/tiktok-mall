// User相关类型
export interface LoginReq {
  email: string;
  password: string;
}

export interface RefreshTokenReq {
  refresh_token: string;
}

// 用户角色枚举
export enum UserRole {
  Customer = "Customer",
  Seller = "Seller",
  Admin = "Admin",
}

// 用户信息接口
export interface UserInfo {
  uuid: string;
  email: string;
  role: UserRole;
}

// getUserInfo API类型
export interface GetUserInfoReq {}

export interface GetUserInfoResp {
  user: UserInfo;
}

// 通用API响应格式
export interface ApiResponse<T> {
  code: number;
  msg: string;
  data: T;
}

export interface LoginResp {
  token: string;
  refresh_token: string;
}

export interface RegisterReq {
  email: string;
  password: string;
  confirm_password: string;
}

export interface RegisterResp {
  user_uuid: string;
}

export interface LogoutReq {}

export interface LogoutResp {
  ok: boolean;
}

// Product相关类型
export interface Product {
  uuid: string;
  name: string;
  description: string;
  price: number;
  stock: number;
}

export interface CreateProductReq {
  name: string;
  description: string;
  price: number;
  stock: number;
}

export interface CreateProductResp {
  product: Product;
}

export interface UpdateProductReq {
  product: Product;
}

export interface UpdateProductResp {
  product: Product;
}

export interface DeleteProductReq {
  uuid: string;
}

export interface DeleteProductResp {
  uuid: string;
}

export interface GetProductReq {
  uuid: string;
}

export interface GetProductResp {
  product: Product;
}

export interface GetProductListReq {
  page: number;
  limit: number;
  name?: string;
  min_price?: number;
  max_price?: number;
}

// 产品列表响应数据结构
export interface ProductListData {
  total: number;
  products: Product[];
}

// 匹配API文档的产品列表响应类型
export interface GetProductListResp extends ProductListData {}

// Cart相关类型
export interface CartItem {
  product_uuid: string;
  quantity: number;
}

export interface AddProductToCartReq {
  item: CartItem;
}

export interface AddProductToCartResp {
  item: CartItem;
}

export interface ClearCartReq {}

export interface ClearCartResp {
  user_uuid: string;
}

export interface GetCartReq {}

export interface GetCartResp {
  total: number;
  items: CartItem[];
}

// Checkout相关类型
export interface Address {
  uuid: string;
  street_address: string;
  city: string;
  state: string;
  country: string;
  zip_code: number;
}

export interface OrderItem {
  product_uuid: string;
  quantity: number;
}

export interface CreateAddressReq {
  street_address: string;
  city: string;
  state: string;
  country: string;
  zip_code: number;
}

export interface CreateAddressResp {
  address: Address;
}

export interface UpdateAddressReq {
  address: Address;
}

export interface UpdateAddressResp {
  address: Address;
}

export interface DeleteAddressReq {
  uuid: string;
}

export interface DeleteAddressResp {
  uuid: string;
}

export interface GetAddressReq {}

export interface GetAddressResp {
  addresses: Address[];
}

export interface CheckoutReq {
  first_name: string;
  last_name: string;
  email: string;
  address_uuid: string;
  items: OrderItem[];
}

export interface CheckoutResp {
  order_uuid: string;
}

// Payment相关类型
export interface CreditCard {
  credit_card_number: string;
  credit_card_cvv: number;
  credit_card_exp_month: number;
  credit_card_exp_year: number;
}

export interface ChargeReq {
  order_uuid: string;
  credit_card: CreditCard;
}

export interface ChargeResp {
  transaction_uuid: string;
}

// AI助手相关类型
export interface AIChatReq {
  content: string;
}

export interface AIChatResp {
  reply: string;
} 