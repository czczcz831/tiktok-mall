import React, { useState, useEffect, useCallback, useRef } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { useCart } from '../contexts/CartContext';
import { getOrders, payOrder, Order } from '../api/orderApi';
import { getProduct } from '../api/productApi';
import { Product, CreditCard } from '../types/api';
import '../styles/OrderListPage.css';
// Material UI 图标导入
import ShoppingBagIcon from '@mui/icons-material/ShoppingBag';
import CreditCardIcon from '@mui/icons-material/CreditCard';
import LocalOfferIcon from '@mui/icons-material/LocalOffer';
import AccessTimeIcon from '@mui/icons-material/AccessTime';
import ListAltIcon from '@mui/icons-material/ListAlt';
import StorefrontIcon from '@mui/icons-material/Storefront';
import CloseIcon from '@mui/icons-material/Close';
import WarningIcon from '@mui/icons-material/Warning';
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import ErrorIcon from '@mui/icons-material/Error';
import ShoppingCartIcon from '@mui/icons-material/ShoppingCart';
import AddShoppingCartIcon from '@mui/icons-material/AddShoppingCart';
import DescriptionIcon from '@mui/icons-material/Description';

// 订单状态常量
const ORDER_STATUS = {
  UNPAID: 0,
  PAID: 1,
  CANCELLED: -1
};

// 订单状态文本
const ORDER_STATUS_TEXT = {
  [ORDER_STATUS.UNPAID]: '待支付',
  [ORDER_STATUS.PAID]: '已支付',
  [ORDER_STATUS.CANCELLED]: '已取消'
};

// 订单状态样式类
const ORDER_STATUS_CLASS = {
  [ORDER_STATUS.UNPAID]: 'status-unpaid',
  [ORDER_STATUS.PAID]: 'status-paid',
  [ORDER_STATUS.CANCELLED]: 'status-cancelled'
};

interface OrderItemWithProduct extends Product {
  quantity: number;
  price: number;
  description: string; // 添加描述字段，替代图片
}

interface OrderWithProducts extends Order {
  productDetails: OrderItemWithProduct[];
}

const OrderListPage: React.FC = () => {
  const { isAuthenticated } = useAuth();
  const { clearCart } = useCart();
  const navigate = useNavigate();
  const [orders, setOrders] = useState<OrderWithProducts[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [loadingProducts, setLoadingProducts] = useState<boolean>(false);
  const [error, setError] = useState<string>('');
  const [isPaymentModalOpen, setIsPaymentModalOpen] = useState<boolean>(false);
  const [selectedOrderUuid, setSelectedOrderUuid] = useState<string>('');
  const [paymentSuccess, setPaymentSuccess] = useState<boolean>(false);
  const [transactionUuid, setTransactionUuid] = useState<string>('');
  const [loadedProductsMap, setLoadedProductsMap] = useState<Record<string, boolean>>({});
  const observerRef = useRef<IntersectionObserver | null>(null);
  
  // 支付表单状态
  const [creditCardNumber, setCreditCardNumber] = useState<string>('');
  const [creditCardCvv, setCreditCardCvv] = useState<string>('');
  const [creditCardExpMonth, setCreditCardExpMonth] = useState<string>('');
  const [creditCardExpYear, setCreditCardExpYear] = useState<string>('');
  const [paymentError, setPaymentError] = useState<string>('');
  const [paymentLoading, setPaymentLoading] = useState<boolean>(false);

  const fetchOrders = async () => {
    setLoading(true);
    setError('');
    try {
      const response = await getOrders();
      
      if (response.code === 0 && response.data) {
        // 初始化订单数据，但不包含产品详情
        const ordersWithoutProducts = response.data.orders.map(order => ({
          ...order,
          productDetails: []
        }));
        setOrders(ordersWithoutProducts);
      } else {
        setError('获取订单失败: ' + response.msg);
      }
    } catch (error) {
      console.error('Failed to fetch orders:', error);
      setError('获取订单失败，请稍后再试');
    } finally {
      setLoading(false);
    }
  };

  const fetchProductDetails = async (order: OrderWithProducts) => {
    if (loadedProductsMap[order.uuid]) return; // 如果已经加载过，直接返回
    
    setLoadingProducts(true);
    try {
      const productDetailsPromises = order.items.map(async (item) => {
        try {
          const productResponse = await getProduct({ uuid: item.product_uuid });
          if (productResponse.code === 0 && productResponse.data && productResponse.data.product) {
            return {
              ...productResponse.data.product,
              quantity: item.quantity,
              price: item.price,
              description: productResponse.data.product.description || '暂无描述'
            };
          }
          return null;
        } catch (error) {
          console.error(`Failed to fetch product ${item.product_uuid}:`, error);
          return null;
        }
      });

      const productDetailsResults = await Promise.all(productDetailsPromises);
      const validProductDetails = productDetailsResults.filter((product): product is OrderItemWithProduct => product !== null);

      setOrders(prevOrders => 
        prevOrders.map(prevOrder => 
          prevOrder.uuid === order.uuid 
            ? { ...prevOrder, productDetails: validProductDetails }
            : prevOrder
        )
      );

      setLoadedProductsMap(prev => ({
        ...prev,
        [order.uuid]: true
      }));
    } catch (error) {
      console.error('Failed to fetch product details:', error);
    } finally {
      setLoadingProducts(false);
    }
  };

  const setupIntersectionObserver = useCallback(() => {
    if (observerRef.current) {
      observerRef.current.disconnect();
    }

    observerRef.current = new IntersectionObserver(
      (entries) => {
        entries.forEach(entry => {
          if (entry.isIntersecting) {
            const orderUuid = entry.target.getAttribute('data-order-uuid');
            if (orderUuid) {
              const order = orders.find(o => o.uuid === orderUuid);
              if (order && !loadedProductsMap[orderUuid]) {
                fetchProductDetails(order);
              }
            }
          }
        });
      },
      {
        root: null,
        rootMargin: '50px',
        threshold: 0.1
      }
    );

    // 为每个订单添加观察
    document.querySelectorAll('.order-card').forEach(orderCard => {
      observerRef.current?.observe(orderCard);
    });
  }, [orders, loadedProductsMap]);

  useEffect(() => {
    if (!isAuthenticated) {
      navigate('/login');
      return;
    }
    
    fetchOrders();
    
    const shouldClearCart = localStorage.getItem('clear_cart_after_payment');
    if (shouldClearCart === 'true') {
      clearCart().then(() => {
        localStorage.removeItem('clear_cart_after_payment');
      }).catch(error => {
        console.error('Failed to clear cart:', error);
      });
    }
  }, [isAuthenticated, navigate]);

  useEffect(() => {
    if (!loading && orders.length > 0) {
      setupIntersectionObserver();
    }
    return () => {
      observerRef.current?.disconnect();
    };
  }, [loading, orders, setupIntersectionObserver]);

  const handlePayClick = (orderUuid: string) => {
    setSelectedOrderUuid(orderUuid);
    setIsPaymentModalOpen(true);
    // 重置支付表单
    setCreditCardNumber('');
    setCreditCardCvv('');
    setCreditCardExpMonth('');
    setCreditCardExpYear('');
    setPaymentError('');
  };

  const handlePaymentSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setPaymentError('');
    setPaymentLoading(true);

    if (!creditCardNumber || !creditCardCvv || !creditCardExpMonth || !creditCardExpYear) {
      setPaymentError('请填写所有支付信息');
      setPaymentLoading(false);
      return;
    }

    try {
      const creditCard: CreditCard = {
        credit_card_number: creditCardNumber,
        credit_card_cvv: parseInt(creditCardCvv),
        credit_card_exp_month: parseInt(creditCardExpMonth),
        credit_card_exp_year: parseInt(creditCardExpYear)
      };

      const response = await payOrder(selectedOrderUuid, creditCard);

      if (response.code === 0) {
        // 设置交易流水号和支付成功状态
        setTransactionUuid(response.data.transaction_uuid);
        setPaymentSuccess(true);
        // 不立即关闭支付窗口，等用户在成功页面确认后再关闭
      } else {
        setPaymentError('支付失败(Mock随机失败，不是BUG): ' + response.msg);
      }
    } catch (error) {
      console.error('Payment failed:', error);
      setPaymentError('支付处理出错，请稍后再试');
    } finally {
      setPaymentLoading(false);
    }
  };

  const formatDate = (timestamp: number) => {
    const date = new Date(timestamp * 1000);
    return date.toLocaleString('zh-CN', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit'
    });
  };

  // 生成渐变色背景，为商品提供视觉区分
  const getProductColorClass = (index: number) => {
    const colorClasses = [
      'product-color-1',
      'product-color-2',
      'product-color-3',
      'product-color-4',
      'product-color-5'
    ];
    return colorClasses[index % colorClasses.length];
  };

  if (loading) {
    return (
      <div className="order-list-page">
        <div className="loading">
          <div className="loading-spinner"></div>
          <span>正在加载订单数据...</span>
        </div>
      </div>
    );
  }

  return (
    <div className="order-list-page">
      <h2><ShoppingBagIcon sx={{ verticalAlign: 'middle', marginRight: '0.5rem' }} />我的订单</h2>
      {error && <div className="error-message">{error}</div>}
      
      {orders.length === 0 ? (
        <div className="empty-orders">
          <ListAltIcon sx={{ fontSize: 60, color: '#ccc', marginBottom: '1rem' }} />
          <p>您还没有订单</p>
          <button onClick={() => navigate('/')}>
            <StorefrontIcon sx={{ marginRight: '0.5rem' }} />
            去购物
          </button>
        </div>
      ) : (
        <div className="orders-container">
          {orders.map((order, orderIndex) => (
            <div 
              key={order.uuid} 
              className="order-card"
              data-order-uuid={order.uuid}
            >
              <div className="order-header">
                <div className="order-info">
                  <span className="order-id">订单号: {order.uuid}</span>
                  <span className="order-date">
                    <AccessTimeIcon sx={{ fontSize: '0.9rem', verticalAlign: 'middle', marginRight: '0.3rem' }} />
                    下单时间: {formatDate(order.created_at)}
                  </span>
                </div>
                <div className={`order-status ${ORDER_STATUS_CLASS[order.status]}`}>
                  {order.status === ORDER_STATUS.PAID && <CheckCircleIcon sx={{ fontSize: '0.9rem', marginRight: '0.3rem' }} />}
                  {order.status === ORDER_STATUS.UNPAID && <WarningIcon sx={{ fontSize: '0.9rem', marginRight: '0.3rem' }} />}
                  {order.status === ORDER_STATUS.CANCELLED && <ErrorIcon sx={{ fontSize: '0.9rem', marginRight: '0.3rem' }} />}
                  {ORDER_STATUS_TEXT[order.status]}
                </div>
              </div>
              
              <div className="order-items">
                {!loadedProductsMap[order.uuid] ? (
                  <div className="product-loading-container">
                    <div className="product-loading">
                      <div className="loading-spinner"></div>
                      <span>正在加载商品信息...</span>
                    </div>
                  </div>
                ) : (
                  order.productDetails.map((product, index) => (
                    <div key={index} className={`order-item ${getProductColorClass(index)}`}>
                      <div className="product-icon">
                        <ShoppingCartIcon />
                      </div>
                      <div className="product-info">
                        <h4>
                          <AddShoppingCartIcon sx={{ fontSize: '1.1rem', marginRight: '0.5rem' }} />
                          {product.name}
                        </h4>
                        <p className="product-description">
                          <DescriptionIcon sx={{ fontSize: '0.9rem', verticalAlign: 'middle', marginRight: '0.5rem' }} />
                          {product.description}
                        </p>
                        <div className="product-details">
                          <p className="product-price">
                            <LocalOfferIcon sx={{ fontSize: '0.9rem', verticalAlign: 'middle', marginRight: '0.2rem', color: '#ff4a6b' }} />
                            单价: ¥{(product.price / 100).toFixed(2)}
                          </p>
                          <p className="product-quantity">数量: {product.quantity}</p>
                          <p className="product-subtotal">小计: ¥{((product.price * product.quantity) / 100).toFixed(2)}</p>
                        </div>
                      </div>
                    </div>
                  ))
                )}
              </div>
              
              <div className="order-footer">
                <div className="order-total">
                  共{order.items.reduce((sum, item) => sum + item.quantity, 0)}件商品，
                  总计: <span className="total-price">¥{(order.total / 100).toFixed(2)}</span>
                </div>
                
                {order.status === ORDER_STATUS.UNPAID && (
                  <button 
                    className="pay-button"
                    onClick={() => handlePayClick(order.uuid)}
                  >
                    立即支付
                  </button>
                )}
              </div>
            </div>
          ))}
        </div>
      )}
      
      {/* 支付弹窗 */}
      {isPaymentModalOpen && (
        <div className="payment-modal-overlay">
          <div className="payment-modal">
            {paymentSuccess ? (
              <div className="payment-success">
                <CheckCircleIcon sx={{ fontSize: 60, color: '#52c41a', marginBottom: '1rem' }} />
                <h3>支付成功！</h3>
                <p>您的订单已完成支付</p>
                <p className="transaction-id">交易流水号: {transactionUuid}</p>
                <button 
                  onClick={() => {
                    setIsPaymentModalOpen(false);
                    setPaymentSuccess(false);
                    fetchOrders(); // 刷新订单列表
                  }}
                  className="confirm-button"
                >
                  确定
                </button>
              </div>
            ) : (
              <>
                <h3><CreditCardIcon sx={{ verticalAlign: 'middle', marginRight: '0.5rem' }} />订单支付</h3>
                {paymentError && <div className="error-message">{paymentError}</div>}
                
                <form onSubmit={handlePaymentSubmit}>
                  <div className="form-group">
                    <label htmlFor="creditCardNumber">信用卡号</label>
                    <input
                      type="text"
                      id="creditCardNumber"
                      value={creditCardNumber}
                      onChange={(e) => setCreditCardNumber(e.target.value)}
                      required
                      placeholder="请输入16位卡号"
                      maxLength={16}
                    />
                  </div>
                  <div className="form-group">
                    <label htmlFor="creditCardCvv">CVV安全码</label>
                    <input
                      type="text"
                      id="creditCardCvv"
                      value={creditCardCvv}
                      onChange={(e) => setCreditCardCvv(e.target.value)}
                      required
                      placeholder="请输入3位CVV码"
                      maxLength={3}
                    />
                  </div>
                  <div className="form-row">
                    <div className="form-group">
                      <label htmlFor="creditCardExpMonth">过期月份</label>
                      <input
                        type="text"
                        id="creditCardExpMonth"
                        value={creditCardExpMonth}
                        onChange={(e) => setCreditCardExpMonth(e.target.value)}
                        placeholder="MM"
                        required
                        maxLength={2}
                      />
                    </div>
                    <div className="form-group">
                      <label htmlFor="creditCardExpYear">过期年份</label>
                      <input
                        type="text"
                        id="creditCardExpYear"
                        value={creditCardExpYear}
                        onChange={(e) => setCreditCardExpYear(e.target.value)}
                        placeholder="YYYY"
                        required
                        maxLength={4}
                      />
                    </div>
                  </div>
                  
                  <div className="modal-actions">
                    <button 
                      type="button" 
                      className="cancel-button"
                      onClick={() => setIsPaymentModalOpen(false)}
                    >
                      <CloseIcon sx={{ marginRight: '0.3rem', fontSize: '1rem' }} />
                      取消
                    </button>
                    <button 
                      type="submit" 
                      className="confirm-button"
                      disabled={paymentLoading}
                    >
                      {paymentLoading ? '处理中...' : '确认支付'}
                    </button>
                  </div>
                </form>
              </>
            )}
          </div>
        </div>
      )}
    </div>
  );
};

export default OrderListPage; 