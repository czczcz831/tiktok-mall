import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useCart } from '../contexts/CartContext';
import { useAuth } from '../contexts/AuthContext';
import { getAddress, createAddress, checkout } from '../api/checkoutApi';
import { charge } from '../api/paymentApi';
import { getProduct } from '../api/productApi';
import { Address, OrderItem, CreditCard, Product } from '../types/api';
import '../styles/CheckoutPage.css';
// Material UI 图标导入
import LocalOfferIcon from '@mui/icons-material/LocalOffer';
import ShoppingCartIcon from '@mui/icons-material/ShoppingCart';
import CreditCardIcon from '@mui/icons-material/CreditCard';
import HomeIcon from '@mui/icons-material/Home';
import PersonIcon from '@mui/icons-material/Person';
import EmailIcon from '@mui/icons-material/Email';
import AddCircleOutlineIcon from '@mui/icons-material/AddCircleOutline';
import CheckCircleOutlineIcon from '@mui/icons-material/CheckCircleOutline';

interface CartItemWithDetails {
  product_uuid: string;
  quantity: number;
  product: Product | null;
}

const CheckoutPage: React.FC = () => {
  const { items, total, clearCart } = useCart();
  const { isAuthenticated } = useAuth();
  const navigate = useNavigate();
  
  const [addresses, setAddresses] = useState<Address[]>([]);
  const [selectedAddressUuid, setSelectedAddressUuid] = useState<string>('');
  const [showAddressForm, setShowAddressForm] = useState<boolean>(false);
  const [orderPlaced, setOrderPlaced] = useState<boolean>(false);
  const [orderUuid, setOrderUuid] = useState<string>('');
  const [cartItemsWithDetails, setCartItemsWithDetails] = useState<CartItemWithDetails[]>([]);
  const [calculatedTotal, setCalculatedTotal] = useState<number>(0);
  const [transactionUuid, setTransactionUuid] = useState<string>('');
  const [navigationPath, setNavigationPath] = useState<string>('');
  
  // 个人信息
  const [firstName, setFirstName] = useState<string>('');
  const [lastName, setLastName] = useState<string>('');
  const [email, setEmail] = useState<string>('');
  
  // 新地址信息
  const [streetAddress, setStreetAddress] = useState<string>('');
  const [city, setCity] = useState<string>('');
  const [state, setState] = useState<string>('');
  const [country, setCountry] = useState<string>('');
  const [zipCode, setZipCode] = useState<string>('');
  
  // 信用卡信息
  const [creditCardNumber, setCreditCardNumber] = useState<string>('');
  const [creditCardCvv, setCreditCardCvv] = useState<string>('');
  const [creditCardExpMonth, setCreditCardExpMonth] = useState<string>('');
  const [creditCardExpYear, setCreditCardExpYear] = useState<string>('');
  
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>('');
  const [paymentSuccess, setPaymentSuccess] = useState<boolean>(false);

  useEffect(() => {
    if (!isAuthenticated) {
      navigate('/login');
      return;
    }
    
    if (items.length === 0) {
      navigate('/cart');
      return;
    }
    
    fetchAddresses();
    fetchProductDetails();
  }, [isAuthenticated, items]);

  const fetchProductDetails = async () => {
    try {
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
      setCartItemsWithDetails(itemsWithDetails);
      
      // 计算正确的总金额
      const total = itemsWithDetails.reduce((sum, item) => {
        if (item.product) {
          return sum + (item.product.price * item.quantity) / 100;
        }
        return sum;
      }, 0);
      setCalculatedTotal(total);
    } catch (error) {
      console.error('Failed to fetch product details:', error);
    }
  };

  const fetchAddresses = async () => {
    setLoading(true);
    try {
      const response = await getAddress({});
      setAddresses(response.data.addresses);
      if (response.data.addresses.length > 0) {
        setSelectedAddressUuid(response.data.addresses[0].uuid);
      }
    } catch (error) {
      console.error('Failed to fetch addresses:', error);
      setError('获取地址失败');
    } finally {
      setLoading(false);
    }
  };

  const handleAddAddress = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    
    if (!streetAddress || !city || !state || !country || !zipCode) {
      setError('请填写所有地址字段');
      return;
    }
    
    try {
      const response = await createAddress({
        street_address: streetAddress,
        city,
        state,
        country,
        zip_code: parseInt(zipCode)
      });
      
      setAddresses([...addresses, response.data.address]);
      setSelectedAddressUuid(response.data.address.uuid);
      setShowAddressForm(false);
      
      // 清空表单
      setStreetAddress('');
      setCity('');
      setState('');
      setCountry('');
      setZipCode('');
    } catch (error) {
      console.error('Failed to add address:', error);
      setError('添加地址失败');
    }
  };

  const handlePlaceOrder = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    
    if (!firstName || !lastName || !email || !selectedAddressUuid) {
      setError('请填写所有必填字段');
      return;
    }
    
    try {
      const orderItems: OrderItem[] = items.map(item => ({
        product_uuid: item.product_uuid,
        quantity: item.quantity
      }));
      
      const response = await checkout({
        first_name: firstName,
        last_name: lastName,
        email,
        address_uuid: selectedAddressUuid,
        items: orderItems
      });
      
      setOrderUuid(response.data.order_uuid);
      setOrderPlaced(true);
    } catch (error) {
      console.error('Failed to place order:', error);
      setError('下单失败');
    }
  };

  const handlePayment = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    
    if (!creditCardNumber || !creditCardCvv || !creditCardExpMonth || !creditCardExpYear) {
      setError('请填写所有支付信息');
      return;
    }
    
    try {
      const creditCard: CreditCard = {
        credit_card_number: creditCardNumber,
        credit_card_cvv: parseInt(creditCardCvv),
        credit_card_exp_month: parseInt(creditCardExpMonth),
        credit_card_exp_year: parseInt(creditCardExpYear)
      };
      
      const response = await charge({
        order_uuid: orderUuid,
        credit_card: creditCard
      });
      
      // 设置支付成功状态和交易流水号
      setTransactionUuid(response.data.transaction_uuid);
      setPaymentSuccess(true);
      
      // 这里不立即清空购物车，而是在用户点击按钮离开支付成功页面时清空
    } catch (error) {
      console.error('Payment failed:', error);
      setError('支付失败(Mock随机失败，不是BUG)');
    }
  };

  if (loading) {
    return (
      <div className="checkout-page">
        <div className="loading">
          <div className="loading-spinner"></div>
          <span>正在加载结账信息...</span>
        </div>
      </div>
    );
  }

  if (paymentSuccess) {
    return (
      <div className="checkout-page">
        <div className="checkout-success">
          <CheckCircleOutlineIcon sx={{ fontSize: 60, color: '#52c41a', marginBottom: '1rem' }} />
          <h2>支付成功！</h2>
          <p>您的订单已完成，感谢您的购买。</p>
          <p className="transaction-id">交易流水号: {transactionUuid}</p>
          <button onClick={() => {
            // 直接导航到订单页面，不清空购物车
            // 将购物车清空操作存储在localStorage中，在目标页面加载时执行
            localStorage.setItem('clear_cart_after_payment', 'true');
            navigate('/orders');
          }}>
            查看我的订单
          </button>
          <button onClick={() => {
            // 直接导航到商品页面，不清空购物车
            // 将购物车清空操作存储在localStorage中，在目标页面加载时执行
            localStorage.setItem('clear_cart_after_payment', 'true');
            navigate('/products');
          }} style={{ marginLeft: '1rem', background: 'linear-gradient(45deg, #36cfc9, #5cdbd3)' }}>
            继续购物
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="checkout-page">
      <h2>订单结账</h2>
      {error && <div className="error-message">{error}</div>}
      
      {!orderPlaced ? (
        <div className="checkout-form">
          <h3>填写订单信息</h3>
          <form onSubmit={handlePlaceOrder}>
            <div className="personal-info">
              <h4><PersonIcon sx={{ verticalAlign: 'middle', marginRight: '0.5rem' }} />个人信息</h4>
              <div className="form-row">
                <div className="form-group">
                  <label htmlFor="firstName">名</label>
                  <input
                    type="text"
                    id="firstName"
                    value={firstName}
                    onChange={(e) => setFirstName(e.target.value)}
                    required
                    placeholder="请输入您的名"
                  />
                </div>
                <div className="form-group">
                  <label htmlFor="lastName">姓</label>
                  <input
                    type="text"
                    id="lastName"
                    value={lastName}
                    onChange={(e) => setLastName(e.target.value)}
                    required
                    placeholder="请输入您的姓"
                  />
                </div>
              </div>
              <div className="form-group">
                <label htmlFor="email"><EmailIcon sx={{ verticalAlign: 'middle', marginRight: '0.5rem', fontSize: '1rem' }} />邮箱</label>
                <input
                  type="email"
                  id="email"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  required
                  placeholder="请输入您的邮箱地址"
                />
              </div>
            </div>
            
            <div className="address-selection">
              <h4><HomeIcon sx={{ verticalAlign: 'middle', marginRight: '0.5rem' }} />送货地址</h4>
              {addresses.length > 0 ? (
                <div className="saved-addresses">
                  <select
                    value={selectedAddressUuid}
                    onChange={(e) => setSelectedAddressUuid(e.target.value)}
                  >
                    {addresses.map((address) => (
                      <option key={address.uuid} value={address.uuid}>
                        {address.street_address}, {address.city}, {address.state}, {address.country}, {address.zip_code}
                      </option>
                    ))}
                  </select>
                </div>
              ) : (
                <p>没有保存的地址，请添加新地址</p>
              )}
              
              <button
                type="button"
                onClick={() => setShowAddressForm(!showAddressForm)}
                className="add-address-btn"
              >
                {showAddressForm ? '取消添加' : '添加新地址'}
              </button>
              
              {showAddressForm && (
                <div className="new-address-form">
                  <h4>新地址信息</h4>
                  <div className="form-group">
                    <label htmlFor="streetAddress">街道地址</label>
                    <input
                      type="text"
                      id="streetAddress"
                      value={streetAddress}
                      onChange={(e) => setStreetAddress(e.target.value)}
                      placeholder="请输入街道地址"
                    />
                  </div>
                  <div className="form-row">
                    <div className="form-group">
                      <label htmlFor="city">城市</label>
                      <input
                        type="text"
                        id="city"
                        value={city}
                        onChange={(e) => setCity(e.target.value)}
                        placeholder="请输入城市"
                      />
                    </div>
                    <div className="form-group">
                      <label htmlFor="state">省/州</label>
                      <input
                        type="text"
                        id="state"
                        value={state}
                        onChange={(e) => setState(e.target.value)}
                        placeholder="请输入省/州"
                      />
                    </div>
                  </div>
                  <div className="form-row">
                    <div className="form-group">
                      <label htmlFor="country">国家</label>
                      <input
                        type="text"
                        id="country"
                        value={country}
                        onChange={(e) => setCountry(e.target.value)}
                        placeholder="请输入国家"
                      />
                    </div>
                    <div className="form-group">
                      <label htmlFor="zipCode">邮编</label>
                      <input
                        type="text"
                        id="zipCode"
                        value={zipCode}
                        onChange={(e) => setZipCode(e.target.value)}
                        placeholder="请输入邮编"
                      />
                    </div>
                  </div>
                  <button type="button" onClick={handleAddAddress}>
                    <AddCircleOutlineIcon sx={{ marginRight: '0.5rem' }} />
                    保存地址
                  </button>
                </div>
              )}
            </div>
            
            <div className="order-summary">
              <h4><ShoppingCartIcon sx={{ verticalAlign: 'middle', marginRight: '0.5rem' }} />订单摘要</h4>
              <div className="cart-items-preview">
                {cartItemsWithDetails.map((item, index) => (
                  <div key={index} className="cart-item-preview">
                    <span className="item-name">{item.product?.name}</span>
                    <span className="item-quantity">x {item.quantity}</span>
                    <span className="item-price">¥{item.product ? (item.product.price * item.quantity / 100).toFixed(2) : '0.00'}</span>
                  </div>
                ))}
              </div>
              <div className="summary-total">
                <p>商品数量: <strong>{items.reduce((sum, item) => sum + item.quantity, 0)}</strong></p>
                <p className="total-price">总计: <LocalOfferIcon sx={{ verticalAlign: 'middle', marginRight: '0.5rem', color: '#ff4a6b' }} />¥{calculatedTotal.toFixed(2)}</p>
              </div>
            </div>
            
            <button type="submit" className="place-order-btn">
              生成订单
            </button>
          </form>
        </div>
      ) : (
        <div className="payment-form">
          <h3><CreditCardIcon sx={{ verticalAlign: 'middle', marginRight: '0.5rem' }} />支付信息</h3>
          <form onSubmit={handlePayment}>
            <div className="credit-card-info">
              <h4>信用卡详情</h4>
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
              <div className="form-row">
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
                <div className="form-group">
                  <label htmlFor="creditCardExpMonth">到期月份</label>
                  <input
                    type="text"
                    id="creditCardExpMonth"
                    value={creditCardExpMonth}
                    onChange={(e) => setCreditCardExpMonth(e.target.value)}
                    required
                    placeholder="MM"
                    maxLength={2}
                  />
                </div>
                <div className="form-group">
                  <label htmlFor="creditCardExpYear">到期年份</label>
                  <input
                    type="text"
                    id="creditCardExpYear"
                    value={creditCardExpYear}
                    onChange={(e) => setCreditCardExpYear(e.target.value)}
                    required
                    placeholder="YYYY"
                    maxLength={4}
                  />
                </div>
              </div>
            </div>
            
            <div className="order-total">
              <h4>订单总计</h4>
              <p className="total-amount">¥{calculatedTotal.toFixed(2)}</p>
            </div>
            
            <button type="submit" className="pay-btn">
              {loading ? '处理中...' : '确认支付'}
            </button>
          </form>
        </div>
      )}
    </div>
  );
};

export default CheckoutPage; 