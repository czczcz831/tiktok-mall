import React, { useState, useEffect, useRef } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { getProduct, getProductList } from '../api/productApi';
import { Product } from '../types/api';
import { useCart } from '../contexts/CartContext';
import '../styles/ProductDetail.css';
// Material UI 图标导入
import ShoppingCartIcon from '@mui/icons-material/ShoppingCart';
import AddShoppingCartIcon from '@mui/icons-material/AddShoppingCart';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import RemoveIcon from '@mui/icons-material/Remove';
import AddIcon from '@mui/icons-material/Add';
import WarningIcon from '@mui/icons-material/Warning';
import Inventory2Icon from '@mui/icons-material/Inventory2';
import LocalOfferIcon from '@mui/icons-material/LocalOffer';
import CheckCircleOutlineIcon from '@mui/icons-material/CheckCircleOutline';

const ProductDetailPage: React.FC = () => {
  const { uuid } = useParams<{ uuid: string }>();
  const [product, setProduct] = useState<Product | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [quantity, setQuantity] = useState<number>(1);
  const [relatedProducts, setRelatedProducts] = useState<Product[]>([]);
  const [addingToCart, setAddingToCart] = useState<boolean>(false);
  const navigate = useNavigate();
  const { addToCart } = useCart();
  const navbarCartRef = useRef<HTMLElement | null>(null);

  useEffect(() => {
    if (uuid) {
      fetchProduct(uuid);
      window.scrollTo({ top: 0, behavior: 'smooth' });
    }
    
    // 获取导航栏购物车图标的引用
    const navbarCart = document.querySelector('.navbar .cart-icon') as HTMLElement;
    if (navbarCart) {
      navbarCartRef.current = navbarCart;
    }
  }, [uuid]);

  const fetchProduct = async (productUuid: string) => {
    setLoading(true);
    try {
      const response = await getProduct({ uuid: productUuid });
      setProduct(response.data.product);
      fetchRelatedProducts();
    } catch (error) {
      console.error('Failed to fetch product:', error);
    } finally {
      setLoading(false);
    }
  };

  const fetchRelatedProducts = async () => {
    try {
      const response = await getProductList({ page: 1, limit: 4 });
      setRelatedProducts(response.data.products.filter(p => p.uuid !== uuid).slice(0, 3));
    } catch (error) {
      console.error('Failed to fetch related products:', error);
    }
  };

  const decreaseQuantity = () => {
    if (quantity > 1) {
      setQuantity(quantity - 1);
    }
  };

  const increaseQuantity = () => {
    if (product && quantity < product.stock) {
      setQuantity(quantity + 1);
    }
  };

  const handleQuantityChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = parseInt(e.target.value);
    if (value > 0 && product && value <= product.stock) {
      setQuantity(value);
    }
  };

  // 显示一个通知
  const showNotification = (message: string, isError = false) => {
    const notification = document.createElement('div');
    notification.className = `cart-notification ${isError ? 'error' : ''}`;
    notification.textContent = message;
    document.body.appendChild(notification);
    
    setTimeout(() => {
      notification.classList.add('show');
    }, 100);
    
    setTimeout(() => {
      notification.classList.remove('show');
      setTimeout(() => {
        document.body.removeChild(notification);
      }, 300);
    }, 2000);
  };

  // 创建飞向购物车的动画
  const createFlyToCartAnimation = (startElement: HTMLElement) => {
    // 获取起始元素的位置
    const startRect = startElement.getBoundingClientRect();
    const startX = startRect.left + startRect.width / 2;
    const startY = startRect.top + startRect.height / 2;
    
    // 获取目标元素（导航栏的购物车图标）的位置
    let endX = window.innerWidth - 20; // 默认右上角
    let endY = 20;
    
    if (navbarCartRef.current) {
      const cartRect = navbarCartRef.current.getBoundingClientRect();
      endX = cartRect.left + cartRect.width / 2;
      endY = cartRect.top + cartRect.height / 2;
    }
    
    // 计算需要移动的距离
    const distanceX = endX - startX;
    const distanceY = endY - startY;
    
    // 创建飞向购物车的元素
    const flyElement = document.createElement('div');
    flyElement.className = 'fly-to-cart';
    flyElement.innerHTML = `+${quantity}`;
    flyElement.style.setProperty('--end-x', `${distanceX}px`);
    flyElement.style.setProperty('--end-y', `${distanceY}px`);
    flyElement.style.left = `${startX}px`;
    flyElement.style.top = `${startY}px`;
    document.body.appendChild(flyElement);
    
    // 动画结束后移除元素
    flyElement.addEventListener('animationend', () => {
      document.body.removeChild(flyElement);
      
      // 添加购物车图标的动画效果
      if (navbarCartRef.current) {
        navbarCartRef.current.classList.add('cart-bounce');
        setTimeout(() => {
          navbarCartRef.current?.classList.remove('cart-bounce');
        }, 500);
      }
    });
  };

  const handleAddToCart = async () => {
    if (!product) return;
    
    setAddingToCart(true);
    try {
      await addToCart(product.uuid, quantity);
      
      // 获取添加到购物车按钮
      const button = document.querySelector('.add-to-cart') as HTMLElement;
      if (button) {
        // 创建飞向购物车的动画
        createFlyToCartAnimation(button);
      }
      
      // 显示成功提示
      showNotification(`已将 ${quantity} 件商品添加到购物车`);
    } catch (error) {
      console.error('Failed to add to cart:', error);
      showNotification('添加到购物车失败', true);
    } finally {
      setAddingToCart(false);
    }
  };

  // 获取背景渐变色
  const getGradientBackground = (index: number = 0) => {
    const gradients = [
      'linear-gradient(120deg, #f6d365 0%, #fda085 100%)',
      'linear-gradient(120deg, #84fab0 0%, #8fd3f4 100%)',
      'linear-gradient(120deg, #a1c4fd 0%, #c2e9fb 100%)',
      'linear-gradient(120deg, #f093fb 0%, #f5576c 100%)'
    ];
    return gradients[index % gradients.length];
  };

  if (loading) {
    return (
      <div className="loading">
        <div className="loading-spinner"></div>
        <span>正在加载商品详情...</span>
      </div>
    );
  }

  if (!product) {
    return <div className="error">商品不存在</div>;
  }

  return (
    <div className="product-detail-page">
      <div className="product-detail-container">
        <div className="product-info">
          <div className="product-header">
            <h2>{product.name}</h2>
            {product.stock <= 5 && product.stock > 0 && (
              <div className="product-tag limited">
                <WarningIcon />
                <span>库存紧张</span>
              </div>
            )}
            {product.stock <= 0 && (
              <div className="product-tag sold-out">
                <Inventory2Icon />
                <span>已售罄</span>
              </div>
            )}
            {product.stock > 5 && (
              <div className="product-tag in-stock">
                <CheckCircleOutlineIcon />
                <span>有货</span>
              </div>
            )}
          </div>
          
          <div className="product-detail-card" style={{ background: getGradientBackground() }}>
            <div className="price-section">
              <LocalOfferIcon className="price-icon" />
              <p className="price premium-price">¥{product.price / 100}</p>
            </div>
            <p className="stock">库存: {product.stock}</p>
          </div>

          <div className="product-description">
            <h3>商品描述</h3>
            <p className="description">{product.description}</p>
          </div>
          
          <div className="quantity-selector">
            <label htmlFor="quantity">数量:</label>
            <div className="quantity-controls">
              <button 
                className="quantity-btn decrease" 
                onClick={decreaseQuantity}
                disabled={quantity <= 1 || product.stock <= 0}
                aria-label="减少数量"
              >
                <RemoveIcon className="quantity-icon" />
              </button>
              <input
                type="number"
                id="quantity"
                min="1"
                max={product.stock}
                value={quantity}
                onChange={handleQuantityChange}
                disabled={product.stock <= 0}
                aria-label="商品数量"
              />
              <button 
                className="quantity-btn increase" 
                onClick={increaseQuantity}
                disabled={quantity >= product.stock || product.stock <= 0}
                aria-label="增加数量"
              >
                <AddIcon className="quantity-icon" />
              </button>
            </div>
          </div>
          
          <div className="product-actions">
            <button
              onClick={handleAddToCart}
              className={`add-to-cart ${addingToCart ? 'adding' : ''}`}
              disabled={product.stock <= 0 || addingToCart}
            >
              <AddShoppingCartIcon />
              <span>{product.stock <= 0 ? '缺货' : addingToCart ? '添加中...' : '加入购物车'}</span>
            </button>
            <button
              onClick={() => navigate('/cart')}
              className="view-cart"
            >
              <ShoppingCartIcon />
              <span>查看购物车</span>
            </button>
            <button
              onClick={() => navigate('/products')}
              className="back-to-products"
            >
              <ArrowBackIcon />
              <span>返回商品列表</span>
            </button>
          </div>
        </div>
      </div>
      
      {relatedProducts.length > 0 && (
        <div className="recommended-products">
          <h3>你可能还喜欢</h3>
          <div className="related-products-grid">
            {relatedProducts.map((relatedProduct, index) => (
              <div 
                key={relatedProduct.uuid} 
                className="related-product-card"
                onClick={() => navigate(`/product/${relatedProduct.uuid}`)}
                style={{ background: getGradientBackground(index + 1) }}
              >
                <div className="related-product-info">
                  <h4>{relatedProduct.name}</h4>
                  <p className="description">{relatedProduct.description}</p>
                  <div className="price-tag">
                    <LocalOfferIcon className="price-icon" />
                    <p className="price">¥{relatedProduct.price / 100}</p>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>
      )}
    </div>
  );
};

export default ProductDetailPage; 