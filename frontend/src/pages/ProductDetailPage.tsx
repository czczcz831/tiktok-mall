import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { getProduct } from '../api/productApi';
import { Product } from '../types/api';
import { useCart } from '../contexts/CartContext';

const ProductDetailPage: React.FC = () => {
  const { uuid } = useParams<{ uuid: string }>();
  const [product, setProduct] = useState<Product | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [quantity, setQuantity] = useState<number>(1);
  const navigate = useNavigate();
  const { addToCart } = useCart();

  useEffect(() => {
    if (uuid) {
      fetchProduct(uuid);
    }
  }, [uuid]);

  const fetchProduct = async (productUuid: string) => {
    setLoading(true);
    try {
      const response = await getProduct({ uuid: productUuid });
      setProduct(response.data.product);
    } catch (error) {
      console.error('Failed to fetch product:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleQuantityChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = parseInt(e.target.value);
    if (value > 0 && product && value <= product.stock) {
      setQuantity(value);
    }
  };

  const handleAddToCart = async () => {
    if (!product) return;
    
    try {
      await addToCart(product.uuid, quantity);
      alert('商品已添加到购物车');
    } catch (error) {
      console.error('Failed to add to cart:', error);
      alert('添加到购物车失败');
    }
  };

  if (loading) {
    return <div className="loading">加载中...</div>;
  }

  if (!product) {
    return <div className="error">商品不存在</div>;
  }

  return (
    <div className="product-detail-page">
      <div className="product-detail-container">
        <div className="product-info">
          <h2>{product.name}</h2>
          <p className="description">{product.description}</p>
          <p className="price">¥{product.price / 100}</p>
          <p className="stock">库存: {product.stock}</p>
          
          <div className="quantity-selector">
            <label htmlFor="quantity">数量:</label>
            <input
              type="number"
              id="quantity"
              min="1"
              max={product.stock}
              value={quantity}
              onChange={handleQuantityChange}
              disabled={product.stock <= 0}
            />
          </div>
          
          <div className="product-actions">
            <button
              onClick={handleAddToCart}
              className="add-to-cart"
              disabled={product.stock <= 0}
            >
              {product.stock <= 0 ? '缺货' : '加入购物车'}
            </button>
            <button
              onClick={() => navigate('/cart')}
              className="view-cart"
            >
              查看购物车
            </button>
            <button
              onClick={() => navigate('/')}
              className="back-to-products"
            >
              返回商品列表
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ProductDetailPage; 