import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { getOrders, payOrder, Order } from '../api/orderApi';
import { getProduct } from '../api/productApi';
import { Product, CreditCard } from '../types/api';
import '../styles/OrderListPage.css';

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
  image_url: string; // 改为必选字段，确保在使用前赋值
}

interface OrderWithProducts extends Order {
  productDetails: OrderItemWithProduct[];
}

// 默认商品图片
const DEFAULT_PRODUCT_IMAGE = '/placeholder-product.svg'; 

const OrderListPage: React.FC = () => {
  const { isAuthenticated } = useAuth();
  const navigate = useNavigate();
  const [orders, setOrders] = useState<OrderWithProducts[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>('');
  const [isPaymentModalOpen, setIsPaymentModalOpen] = useState<boolean>(false);
  const [selectedOrderUuid, setSelectedOrderUuid] = useState<string>('');
  
  // 支付表单状态
  const [creditCardNumber, setCreditCardNumber] = useState<string>('');
  const [creditCardCvv, setCreditCardCvv] = useState<string>('');
  const [creditCardExpMonth, setCreditCardExpMonth] = useState<string>('');
  const [creditCardExpYear, setCreditCardExpYear] = useState<string>('');
  const [paymentError, setPaymentError] = useState<string>('');
  const [paymentLoading, setPaymentLoading] = useState<boolean>(false);

  useEffect(() => {
    if (!isAuthenticated) {
      navigate('/login');
      return;
    }
    
    fetchOrders();
  }, [isAuthenticated, navigate]);

  const fetchOrders = async () => {
    setLoading(true);
    setError('');
    try {
      const response = await getOrders();
      
      if (response.code === 0 && response.data) {
        // 获取每个订单中商品的详细信息
        const ordersWithProducts = await Promise.all(
          response.data.orders.map(async (order) => {
            const productDetailsPromises = order.items.map(async (item) => {
              try {
                const productResponse = await getProduct({ uuid: item.product_uuid });
                if (productResponse.code === 0 && productResponse.data && productResponse.data.product) {
                  // 创建包含必要字段的OrderItemWithProduct对象
                  const productWithDetails: OrderItemWithProduct = {
                    ...productResponse.data.product,
                    quantity: item.quantity,
                    price: item.price,
                    image_url: (productResponse.data.product as any).image_url || DEFAULT_PRODUCT_IMAGE
                  };
                  return productWithDetails;
                }
                return null;
              } catch (error) {
                console.error(`Failed to fetch product ${item.product_uuid}:`, error);
                return null;
              }
            });
            
            // 等待所有商品信息获取完成
            const productDetailsResults = await Promise.all(productDetailsPromises);
            
            // 过滤掉null值并确保类型正确
            const validProductDetails: OrderItemWithProduct[] = productDetailsResults
              .filter((product): product is OrderItemWithProduct => product !== null);
            
            // 创建包含商品详情的订单对象
            const orderWithProducts: OrderWithProducts = {
              ...order,
              productDetails: validProductDetails
            };
            
            return orderWithProducts;
          })
        );
        
        setOrders(ordersWithProducts);
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
        // 关闭支付窗口并刷新订单列表
        setIsPaymentModalOpen(false);
        fetchOrders();
      } else {
        setPaymentError('支付失败: ' + response.msg);
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

  if (loading) {
    return <div className="loading">加载中...</div>;
  }

  return (
    <div className="order-list-page">
      <h2>我的订单</h2>
      {error && <div className="error-message">{error}</div>}
      
      {orders.length === 0 ? (
        <div className="empty-orders">
          <p>您还没有订单</p>
          <button onClick={() => navigate('/')}>去购物</button>
        </div>
      ) : (
        <div className="orders-container">
          {orders.map((order) => (
            <div key={order.uuid} className="order-card">
              <div className="order-header">
                <div className="order-info">
                  <span className="order-id">订单号: {order.uuid}</span>
                  <span className="order-date">下单时间: {formatDate(order.created_at)}</span>
                </div>
                <div className={`order-status ${ORDER_STATUS_CLASS[order.status]}`}>
                  {ORDER_STATUS_TEXT[order.status]}
                </div>
              </div>
              
              <div className="order-items">
                {order.productDetails.map((product, index) => (
                  <div key={index} className="order-item">
                    <img 
                      src={product.image_url} 
                      alt={product.name} 
                      className="product-image"
                      onError={(e) => {
                        (e.target as HTMLImageElement).src = DEFAULT_PRODUCT_IMAGE;
                      }}
                    />
                    <div className="product-info">
                      <h4>{product.name}</h4>
                      <p className="product-price">单价: ¥{(product.price / 100).toFixed(2)}</p>
                      <p className="product-quantity">数量: {product.quantity}</p>
                    </div>
                  </div>
                ))}
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
            <h3>订单支付</h3>
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
                />
              </div>
              <div className="form-group">
                <label htmlFor="creditCardCvv">CVV</label>
                <input
                  type="text"
                  id="creditCardCvv"
                  value={creditCardCvv}
                  onChange={(e) => setCreditCardCvv(e.target.value)}
                  required
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
                  />
                </div>
              </div>
              
              <div className="modal-actions">
                <button 
                  type="button" 
                  className="cancel-button"
                  onClick={() => setIsPaymentModalOpen(false)}
                >
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
          </div>
        </div>
      )}
    </div>
  );
};

export default OrderListPage; 