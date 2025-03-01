import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useCart } from '../contexts/CartContext';
import { useAuth } from '../contexts/AuthContext';
import { getAddress, createAddress, checkout } from '../api/checkoutApi';
import { charge } from '../api/paymentApi';
import { getProduct } from '../api/productApi';
import { Address, OrderItem, CreditCard, Product } from '../types/api';

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
      
      await charge({
        order_uuid: orderUuid,
        credit_card: creditCard
      });
      
      setPaymentSuccess(true);
      await clearCart(); // 清空购物车
    } catch (error) {
      console.error('Payment failed:', error);
      setError('支付失败');
    }
  };

  if (loading) {
    return <div className="loading">加载中...</div>;
  }

  if (paymentSuccess) {
    return (
      <div className="checkout-success">
        <h2>支付成功！</h2>
        <p>您的订单已完成。</p>
        <button onClick={() => navigate('/')}>继续购物</button>
      </div>
    );
  }

  return (
    <div className="checkout-page">
      <h2>结账</h2>
      {error && <div className="error-message">{error}</div>}
      
      {!orderPlaced ? (
        <div className="checkout-form">
          <h3>订单信息</h3>
          <form onSubmit={handlePlaceOrder}>
            <div className="personal-info">
              <h4>个人信息</h4>
              <div className="form-group">
                <label htmlFor="firstName">名</label>
                <input
                  type="text"
                  id="firstName"
                  value={firstName}
                  onChange={(e) => setFirstName(e.target.value)}
                  required
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
                />
              </div>
              <div className="form-group">
                <label htmlFor="email">邮箱</label>
                <input
                  type="email"
                  id="email"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  required
                />
              </div>
            </div>
            
            <div className="address-selection">
              <h4>送货地址</h4>
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
                <p>没有保存的地址</p>
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
                  <h4>新地址</h4>
                  <div className="form-group">
                    <label htmlFor="streetAddress">街道地址</label>
                    <input
                      type="text"
                      id="streetAddress"
                      value={streetAddress}
                      onChange={(e) => setStreetAddress(e.target.value)}
                    />
                  </div>
                  <div className="form-group">
                    <label htmlFor="city">城市</label>
                    <input
                      type="text"
                      id="city"
                      value={city}
                      onChange={(e) => setCity(e.target.value)}
                    />
                  </div>
                  <div className="form-group">
                    <label htmlFor="state">省/州</label>
                    <input
                      type="text"
                      id="state"
                      value={state}
                      onChange={(e) => setState(e.target.value)}
                    />
                  </div>
                  <div className="form-group">
                    <label htmlFor="country">国家</label>
                    <input
                      type="text"
                      id="country"
                      value={country}
                      onChange={(e) => setCountry(e.target.value)}
                    />
                  </div>
                  <div className="form-group">
                    <label htmlFor="zipCode">邮编</label>
                    <input
                      type="text"
                      id="zipCode"
                      value={zipCode}
                      onChange={(e) => setZipCode(e.target.value)}
                    />
                  </div>
                  <button type="button" onClick={handleAddAddress}>
                    保存地址
                  </button>
                </div>
              )}
            </div>
            
            <div className="order-summary">
              <h4>订单摘要</h4>
              <p>商品数量: {items.length}</p>
              <p>总计: ¥{calculatedTotal.toFixed(2)}</p>
            </div>
            
            <button type="submit" className="place-order-btn">
              下单
            </button>
          </form>
        </div>
      ) : (
        <div className="payment-form">
          <h3>支付信息</h3>
          <form onSubmit={handlePayment}>
            <div className="credit-card-info">
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
              <div className="form-group">
                <label htmlFor="creditCardExpMonth">到期月份</label>
                <input
                  type="text"
                  id="creditCardExpMonth"
                  value={creditCardExpMonth}
                  onChange={(e) => setCreditCardExpMonth(e.target.value)}
                  required
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
                />
              </div>
            </div>
            
            <div className="order-total">
              <h4>订单总计</h4>
              <p className="total-amount">¥{calculatedTotal.toFixed(2)}</p>
            </div>
            
            <button type="submit" className="pay-btn">
              支付
            </button>
          </form>
        </div>
      )}
    </div>
  );
};

export default CheckoutPage; 