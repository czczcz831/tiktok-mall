import React, { useState, useEffect, useCallback } from 'react';
import { useNavigate } from 'react-router-dom';
import { useCart } from '../contexts/CartContext';
import { getProduct } from '../api/productApi';
import { CartItem, Product } from '../types/api';
import '../styles/CartPage.css';
// Material UI 图标
import ShoppingCartIcon from '@mui/icons-material/ShoppingCart';
import DeleteIcon from '@mui/icons-material/Delete';
import AddIcon from '@mui/icons-material/Add';
import RemoveIcon from '@mui/icons-material/Remove';
import ShoppingBagIcon from '@mui/icons-material/ShoppingBag';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';

interface CartItemWithProduct extends CartItem {
  product: Product;
}

const CartPage: React.FC = () => {
  const navigate = useNavigate();
  const { items, total, clearCart, fetchCart, addToCart } = useCart();
  const [cartItemsWithProducts, setCartItemsWithProducts] = useState<CartItemWithProduct[]>([]);
  const [loading, setLoading] = useState(true);
  const [isInitialLoad, setIsInitialLoad] = useState(true);
  const [calculatedTotal, setCalculatedTotal] = useState(0);

  // 获取购物车数据的函数
  const loadCartWithProducts = useCallback(async () => {
    setLoading(true);
    try {
      // 只在初始加载时调用fetchCart，避免形成循环
      if (isInitialLoad) {
        await fetchCart();
        setIsInitialLoad(false);
      }
      
      // 获取每个商品的详细信息
      if (items.length > 0) {
        const itemsWithProducts = await Promise.all(
          items.map(async (item) => {
            const response = await getProduct({ uuid: item.product_uuid });
            return {
              ...item,
              product: response.data.product
            };
          })
        );
        setCartItemsWithProducts(itemsWithProducts);
        
        // 计算总价 = 每件商品价格 * 数量 之和
        const newTotal = itemsWithProducts.reduce((sum, item) => {
          return sum + (item.product.price * item.quantity);
        }, 0);
        
        setCalculatedTotal(newTotal);
      } else {
        setCartItemsWithProducts([]);
        setCalculatedTotal(0);
      }
    } catch (error) {
      console.error('Error loading cart:', error);
    } finally {
      setLoading(false);
    }
  }, [items, fetchCart, isInitialLoad]);

  // 初始加载和items变化时更新数据
  useEffect(() => {
    loadCartWithProducts();
  }, [loadCartWithProducts]);

  // 获取背景渐变色
  const getGradientBackground = (index: number) => {
    const gradients = [
      'linear-gradient(120deg, #f6d365 0%, #fda085 100%)',
      'linear-gradient(120deg, #84fab0 0%, #8fd3f4 100%)',
      'linear-gradient(120deg, #a1c4fd 0%, #c2e9fb 100%)',
      'linear-gradient(120deg, #f093fb 0%, #f5576c 100%)'
    ];
    return gradients[index % gradients.length];
  };

  // 使用延迟添加到购物车，防止请求过快
  const addToCartAfterDelay = async (productUuid: string, quantity: number) => {
    await new Promise(resolve => setTimeout(resolve, 300));
    await addToCart(productUuid, quantity);
  };

  const handleQuantityChange = async (productUuid: string, quantity: number) => {
    if (quantity < 1) return;
    
    // 找到当前项目
    const item = cartItemsWithProducts.find(item => item.product_uuid === productUuid);
    if (!item) return;
    
    // 删除并重新添加到购物车，以更新数量
    try {
      await clearCart();
      
      // 重新添加所有商品，但将当前商品的数量更新
      for (const cartItem of cartItemsWithProducts) {
        if (cartItem.product_uuid === productUuid) {
          // 使用新数量
          await addToCartAfterDelay(cartItem.product_uuid, quantity);
        } else {
          // 保持原数量
          await addToCartAfterDelay(cartItem.product_uuid, cartItem.quantity);
        }
      }
      
      // 重新加载购物车
      await fetchCart();
      // 不需要调用loadCartWithProducts，因为items更新会触发它
    } catch (error) {
      console.error('Error updating quantity:', error);
    }
  };

  const handleRemove = async (productUuid: string) => {
    try {
      // 因为API没有提供单独移除商品的方法，我们需要清空购物车然后重新添加其他商品
      await clearCart();
      
      // 重新添加除了要删除的商品之外的所有商品
      for (const cartItem of cartItemsWithProducts) {
        if (cartItem.product_uuid !== productUuid) {
          await addToCartAfterDelay(cartItem.product_uuid, cartItem.quantity);
        }
      }
      
      // 重新加载购物车
      await fetchCart();
      // 不需要调用loadCartWithProducts，因为items更新会触发它
    } catch (error) {
      console.error('Error removing item:', error);
    }
  };

  const handleCheckout = () => {
    navigate('/checkout');
  };

  if (loading) {
    return (
      <div className="loading">
        <div className="loading-spinner"></div>
        <p>正在加载购物车...</p>
      </div>
    );
  }

  return (
    <div className="cart-page">
      <div className="cart-header">
        <h1><ShoppingCartIcon /> 购物车</h1>
        <p className="item-count">总计 {cartItemsWithProducts.length} 件商品</p>
      </div>

      {cartItemsWithProducts.length === 0 ? (
        <div className="empty-cart">
          <ShoppingBagIcon className="empty-icon" />
          <p>购物车是空的</p>
          <button onClick={() => navigate('/products')} className="continue-shopping">
            <ArrowBackIcon /> 继续购物
          </button>
        </div>
      ) : (
        <div className="cart-content">
          <div className="cart-items">
            {cartItemsWithProducts.map((item, index) => (
              <div 
                className="cart-item" 
                key={item.product_uuid}
                style={{ background: getGradientBackground(index) }}
              >
                <div className="item-info">
                  <h3>{item.product.name}</h3>
                  <p className="item-price">¥{item.product.price / 100} × {item.quantity}</p>
                </div>
                
                <div className="item-actions">
                  <div className="quantity-control">
                    <button 
                      onClick={() => handleQuantityChange(item.product_uuid, item.quantity - 1)}
                      disabled={item.quantity <= 1}
                      className="quantity-btn"
                    >
                      <RemoveIcon />
                    </button>
                    <span>{item.quantity}</span>
                    <button 
                      onClick={() => handleQuantityChange(item.product_uuid, item.quantity + 1)}
                      className="quantity-btn"
                    >
                      <AddIcon />
                    </button>
                  </div>
                  
                  <button 
                    onClick={() => handleRemove(item.product_uuid)}
                    className="remove-btn"
                  >
                    <DeleteIcon />
                  </button>
                </div>
              </div>
            ))}
          </div>
          
          <div className="cart-summary">
            <h2>订单摘要</h2>
            <div className="summary-item">
              <span>商品数量:</span>
              <span>{cartItemsWithProducts.reduce((sum, item) => sum + item.quantity, 0)} 件</span>
            </div>
            <div className="summary-item">
              <span>小计:</span>
              <span>¥{calculatedTotal / 100}</span>
            </div>
            <div className="summary-item total">
              <span>总计:</span>
              <span>¥{calculatedTotal / 100}</span>
            </div>
            <button onClick={handleCheckout} className="checkout-btn">
              去结账
            </button>
            <button onClick={() => navigate('/products')} className="continue-shopping-btn">
              <ArrowBackIcon /> 继续购物
            </button>
          </div>
        </div>
      )}
    </div>
  );
};

export default CartPage; 