import React, { useState, useEffect, useRef, useCallback } from 'react';
import { useNavigate } from 'react-router-dom';
import { getProductList } from '../api/productApi';
import { Product, GetProductListReq } from '../types/api';
import { useCart } from '../contexts/CartContext';
import '../styles/ProductList.css';
// Material UI 图标导入
import SearchIcon from '@mui/icons-material/Search';
import AddShoppingCartIcon from '@mui/icons-material/AddShoppingCart';
import VisibilityIcon from '@mui/icons-material/Visibility';
import Inventory2Icon from '@mui/icons-material/Inventory2';
import LocalOfferIcon from '@mui/icons-material/LocalOffer';
import WarningIcon from '@mui/icons-material/Warning';
import FilterListIcon from '@mui/icons-material/FilterList';
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp';

const ProductListPage: React.FC = () => {
  const [products, setProducts] = useState<Product[]>([]);
  const [total, setTotal] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(true);
  const [page, setPage] = useState<number>(1);
  const [limit] = useState<number>(10);
  const [searchName, setSearchName] = useState<string>('');
  const [minPrice, setMinPrice] = useState<string>('');
  const [maxPrice, setMaxPrice] = useState<string>('');
  const [hasMore, setHasMore] = useState<boolean>(true);
  const [isFiltersVisible, setIsFiltersVisible] = useState<boolean>(true);
  const [loadingMore, setLoadingMore] = useState<boolean>(false);
  const [showScrollTop, setShowScrollTop] = useState<boolean>(false);
  const [fetchedPages, setFetchedPages] = useState<Set<number>>(new Set([1]));
  const navigate = useNavigate();
  const { addToCart, clearCart } = useCart();
  const navbarCartRef = useRef<HTMLElement | null>(null);
  const observer = useRef<IntersectionObserver | null>(null);
  
  const lastProductRef = useCallback((node: HTMLDivElement | null) => {
    if (loading || loadingMore) return;
    if (observer.current) observer.current.disconnect();
    
    observer.current = new IntersectionObserver(entries => {
      if (entries[0].isIntersecting && hasMore) {
        loadMoreProducts();
      }
    });
    
    if (node) observer.current.observe(node);
  }, [loading, loadingMore, hasMore]);

  useEffect(() => {
    // 监听滚动位置，用于显示/隐藏回到顶部按钮
    const handleScroll = () => {
      if (window.scrollY > 300) {
        setShowScrollTop(true);
      } else {
        setShowScrollTop(false);
      }
    };

    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  useEffect(() => {
    // 检查是否需要清空购物车（从支付成功页面跳转过来）
    const shouldClearCart = localStorage.getItem('clear_cart_after_payment');
    if (shouldClearCart === 'true') {
      clearCart().then(() => {
        localStorage.removeItem('clear_cart_after_payment');
      }).catch(error => {
        console.error('Failed to clear cart:', error);
      });
    }
    
    fetchProducts(true);
    // 获取导航栏购物车图标的引用
    const navbarCart = document.querySelector('.navbar .cart-icon') as HTMLElement;
    if (navbarCart) {
      navbarCartRef.current = navbarCart;
    }
  }, []);

  const fetchProducts = async (reset: boolean = false) => {
    // 如果是重置，则清空所有状态并加载第一页
    if (reset) {
      setLoading(true);
      setPage(1);
      setFetchedPages(new Set([1]));
      
      try {
        const req: GetProductListReq = {
          page: 1,
          limit,
          name: searchName || undefined,
          min_price: minPrice ? parseFloat(minPrice) : undefined,
          max_price: maxPrice ? parseFloat(maxPrice) : undefined
        };
        
        const response = await getProductList(req);
        setProducts(response.data.products);
        setTotal(response.data.total);
        setHasMore(response.data.products.length === limit && response.data.products.length < response.data.total);
      } catch (error) {
        console.error('Failed to fetch products:', error);
      } finally {
        setLoading(false);
      }
    } 
    // 否则是加载更多，加载下一页
    else {
      if (loadingMore) return; // 防止重复请求
      
      const nextPage = page + 1;
      
      // 检查是否已经获取过这一页
      if (fetchedPages.has(nextPage)) return;
      
      setLoadingMore(true);
      
      try {
        const req: GetProductListReq = {
          page: nextPage,
          limit,
          name: searchName || undefined,
          min_price: minPrice ? parseFloat(minPrice) : undefined,
          max_price: maxPrice ? parseFloat(maxPrice) : undefined
        };
        
        const response = await getProductList(req);
        
        // 确保没有重复的商品
        const newProducts = response.data.products.filter(
          newProduct => !products.some(existingProduct => existingProduct.uuid === newProduct.uuid)
        );
        
        if (newProducts.length > 0) {
          setProducts(prev => [...prev, ...newProducts]);
          // 更新已获取页面集合
          setFetchedPages(prev => {
            const updated = new Set(prev);
            updated.add(nextPage);
            return updated;
          });
          setPage(nextPage);
        }
        
        setTotal(response.data.total);
        setHasMore(response.data.products.length === limit && products.length + newProducts.length < response.data.total);
      } catch (error) {
        console.error('Failed to fetch more products:', error);
      } finally {
        setLoadingMore(false);
      }
    }
  };

  const loadMoreProducts = () => {
    if (products.length < total && !loadingMore && hasMore) {
      fetchProducts(false);
    }
  };

  const handleSearch = () => {
    fetchProducts(true);
  };

  const scrollToTop = () => {
    window.scrollTo({ top: 0, behavior: 'smooth' });
  };

  const toggleFilters = () => {
    setIsFiltersVisible(!isFiltersVisible);
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
    flyElement.innerHTML = '+1';
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

  const handleAddToCart = async (productUuid: string, event: React.MouseEvent) => {
    event.stopPropagation(); // 阻止事件冒泡到卡片点击
    const button = event.currentTarget as HTMLElement;
    
    try {
      // 创建飞向购物车的动画（在API调用前就开始动画，提升体验）
      createFlyToCartAnimation(button);
      
      // 显示成功提示
      showNotification('商品已添加到购物车');
      
      // API调用放在动画之后执行，优化用户体验
      await addToCart(productUuid, 1);
    } catch (error) {
      console.error('Failed to add to cart:', error);
      showNotification('添加到购物车失败', true);
    }
  };
  
  const handleCardClick = (productUuid: string) => {
    navigate(`/product/${productUuid}`);
  };

  // 获取卡片的渐变背景色
  const getCardGradient = (index: number) => {
    const gradients = [
      'linear-gradient(135deg, #ffedbc, #ffdbed)',
      'linear-gradient(135deg, #e3f1fc, #e6f9e9)',
      'linear-gradient(135deg, #ffeae6, #e6f8ff)',
      'linear-gradient(135deg, #e8e6ff, #fff6e6)',
      'linear-gradient(135deg, #f6e6ff, #e6ffd9)',
      'linear-gradient(135deg, #ffe6f6, #e6ffff)',
    ];
    return gradients[index % gradients.length];
  };

  return (
    <div className="product-list-page">
      <h2 className="section-title">精选商品</h2>
      
      <div className="filters-container">
        <button className="filter-toggle" onClick={toggleFilters}>
          <FilterListIcon />
          <span>{isFiltersVisible ? '隐藏筛选' : '显示筛选'}</span>
        </button>
        
        {isFiltersVisible && (
          <div className="search-filters">
            <div className="search-input">
              <SearchIcon className="search-icon" />
              <input
                type="text"
                placeholder="搜索商品名称"
                value={searchName}
                onChange={(e) => setSearchName(e.target.value)}
              />
            </div>
            <div className="price-filters">
              <input
                type="number"
                placeholder="最低价格"
                value={minPrice}
                onChange={(e) => setMinPrice(e.target.value)}
              />
              <span>-</span>
              <input
                type="number"
                placeholder="最高价格"
                value={maxPrice}
                onChange={(e) => setMaxPrice(e.target.value)}
              />
            </div>
            <button onClick={handleSearch} className="search-button">
              <SearchIcon />
              <span>搜索</span>
            </button>
          </div>
        )}
      </div>

      {loading ? (
        <div className="loading">
          <div className="loading-spinner"></div>
          <span>正在加载精选商品...</span>
        </div>
      ) : (
        <>
          {products.length > 0 ? (
            <div className="products-grid">
              {products.map((product, index) => (
                <div 
                  key={product.uuid} 
                  className="product-card" 
                  onClick={() => handleCardClick(product.uuid)}
                  style={{ background: getCardGradient(index) }}
                  ref={index === products.length - 1 ? lastProductRef : null}
                >
                  <div className="product-card-content">
                    <div className="product-card-header">
                      <h3>{product.name}</h3>
                      {product.stock <= 5 && product.stock > 0 && (
                        <div className="product-tag limited">
                          <WarningIcon fontSize="small" />
                          <span>库存紧张</span>
                        </div>
                      )}
                      {product.stock <= 0 && (
                        <div className="product-tag sold-out">
                          <Inventory2Icon fontSize="small" />
                          <span>已售罄</span>
                        </div>
                      )}
                    </div>
                    <p className="description">{product.description}</p>
                    <div className="product-card-info">
                      <div className="price-tag">
                        <LocalOfferIcon className="price-icon" />
                        <p className="price">¥{product.price / 100}</p>
                      </div>
                      <p className="stock">库存: {product.stock}</p>
                    </div>
                  </div>
                  <div className="product-actions">
                    <button
                      onClick={(e) => handleAddToCart(product.uuid, e)}
                      className="add-to-cart full-width"
                      disabled={product.stock <= 0}
                    >
                      <AddShoppingCartIcon />
                      <span>{product.stock <= 0 ? '缺货' : '加入购物车'}</span>
                    </button>
                  </div>
                </div>
              ))}
            </div>
          ) : (
            <div className="no-products">没有找到商品</div>
          )}
          
          {loadingMore && (
            <div className="loading-more">
              <div className="loading-spinner"></div>
              <span>加载更多商品...</span>
            </div>
          )}
          
          {!hasMore && products.length > 0 && (
            <div className="end-message">
              <span>已经到底啦，没有更多商品了~</span>
            </div>
          )}
        </>
      )}
      
      {showScrollTop && (
        <button className="scroll-top-button" onClick={scrollToTop}>
          <KeyboardArrowUpIcon />
        </button>
      )}
    </div>
  );
};

export default ProductListPage; 