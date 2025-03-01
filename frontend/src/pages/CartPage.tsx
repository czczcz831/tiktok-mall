import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useCart } from '../contexts/CartContext';
import { useAuth } from '../contexts/AuthContext';
import { getProduct } from '../api/productApi';
import { Product } from '../types/api';

interface CartItemWithDetails {
  product_uuid: string;
  quantity: number;
  product: Product | null;
}

const CartPage: React.FC = () => {
  const { items, total, clearCart, fetchCart } = useCart();
  const { isAuthenticated } = useAuth();
  const [cartItems, setCartItems] = useState<CartItemWithDetails[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const navigate = useNavigate();

  useEffect(() => {
    if (isAuthenticated) {
      fetchCart();
    } else {
      navigate('/login');
    }
  }, [isAuthenticated]);

  useEffect(() => {
    const fetchProductDetails = async () => {
      setLoading(true);
      const itemsWithDetails = await Promise.all(
        items.map(async (item) => {
          try {
            const response = await getProduct({ uuid: item.product_uuid });
            return {
              ...item,
              product: response.data.product
            };
          } catch (error) {
            console.error(`Failed to fetch product ${item.product_uuid}:`, error);
            return {
              ...item,
              product: null
            };
          }
        })
      );
      setCartItems(itemsWithDetails);
      setLoading(false);
    };

    if (items.length > 0) {
      fetchProductDetails();
    } else {
      setCartItems([]);
      setLoading(false);
    }
  }, [items]);

  const handleClearCart = async () => {
    if (window.confirm('确定要清空购物车吗？')) {
      try {
        await clearCart();
      } catch (error) {
        console.error('Failed to clear cart:', error);
        alert('清空购物车失败');
      }
    }
  };

  const handleCheckout = () => {
    navigate('/checkout');
  };

  if (!isAuthenticated) {
    return null; // 已经在useEffect中重定向到登录页
  }

  if (loading) {
    return <div className="loading">加载中...</div>;
  }

  return (
    <div className="cart-page">
      <h2>购物车</h2>
      
      {cartItems.length === 0 ? (
        <div className="empty-cart">
          <p>购物车是空的</p>
          <button onClick={() => navigate('/')}>去购物</button>
        </div>
      ) : (
        <>
          <div className="cart-items">
            {cartItems.map((item) => (
              <div key={item.product_uuid} className="cart-item">
                {item.product ? (
                  <>
                    <div className="item-details">
                      <h3>{item.product.name}</h3>
                      <p className="price">单价: ¥{item.product.price / 100}</p>
                      <p className="quantity">数量: {item.quantity}</p>
                      <p className="subtotal">小计: ¥{(item.product.price * item.quantity) / 100}</p>
                    </div>
                    <div className="item-actions">
                      <button onClick={() => navigate(`/product/${item.product_uuid}`)}>
                        查看商品
                      </button>
                    </div>
                  </>
                ) : (
                  <div className="item-error">商品信息加载失败</div>
                )}
              </div>
            ))}
          </div>
          
          <div className="cart-summary">
            <div className="cart-total">
              <span>总计:</span>
              <span className="total-price">
                ¥{cartItems.reduce((sum, item) => item.product ? sum + (item.product.price * item.quantity) / 100 : sum, 0).toFixed(2)}
              </span>
            </div>
            
            <div className="cart-actions">
              <button onClick={handleClearCart} className="clear-cart">
                清空购物车
              </button>
              <button onClick={handleCheckout} className="checkout">
                去结算
              </button>
            </div>
          </div>
        </>
      )}
    </div>
  );
};

export default CartPage; 