import React, { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { addProductToCart, clearCart, getCart } from '../api/cartApi';
import { CartItem, AddProductToCartReq } from '../types/api';
import { useAuth } from './AuthContext';

interface CartContextType {
  items: CartItem[];
  total: number;
  isLoading: boolean;
  addToCart: (productUuid: string, quantity: number) => Promise<void>;
  clearCart: () => Promise<void>;
  fetchCart: () => Promise<void>;
}

const CartContext = createContext<CartContextType | undefined>(undefined);

export const useCart = () => {
  const context = useContext(CartContext);
  if (!context) {
    throw new Error('useCart must be used within a CartProvider');
  }
  return context;
};

interface CartProviderProps {
  children: ReactNode;
}

export const CartProvider: React.FC<CartProviderProps> = ({ children }) => {
  const [items, setItems] = useState<CartItem[]>([]);
  const [total, setTotal] = useState<number>(0);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const { isAuthenticated } = useAuth();

  // 当用户登录状态改变时，获取购物车
  useEffect(() => {
    if (isAuthenticated) {
      fetchCart();
    } else {
      setItems([]);
      setTotal(0);
    }
  }, [isAuthenticated]);

  const fetchCart = async () => {
    if (!isAuthenticated) return;
    
    setIsLoading(true);
    try {
      const response = await getCart({});
      setItems(response.data.items);
      setTotal(response.data.total);
    } catch (error) {
      console.error('Failed to fetch cart:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const addToCart = async (productUuid: string, quantity: number) => {
    setIsLoading(true);
    try {
      const req: AddProductToCartReq = {
        item: {
          product_uuid: productUuid,
          quantity: quantity
        }
      };
      const response = await addProductToCart(req);
      await fetchCart(); // 重新获取购物车
    } catch (error) {
      console.error('Failed to add to cart:', error);
      throw error;
    } finally {
      setIsLoading(false);
    }
  };

  const clearCartItems = async () => {
    setIsLoading(true);
    try {
      const response = await clearCart({});
      setItems([]);
      setTotal(0);
    } catch (error) {
      console.error('Failed to clear cart:', error);
      throw error;
    } finally {
      setIsLoading(false);
    }
  };

  const value = {
    items,
    total,
    isLoading,
    addToCart,
    clearCart: clearCartItems,
    fetchCart
  };

  return <CartContext.Provider value={value}>{children}</CartContext.Provider>;
}; 