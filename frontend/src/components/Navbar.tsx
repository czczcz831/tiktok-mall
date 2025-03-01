import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { useCart } from '../contexts/CartContext';

const Navbar: React.FC = () => {
  const { isAuthenticated, logout } = useAuth();
  const { items } = useCart();
  const navigate = useNavigate();

  const handleLogout = async () => {
    try {
      await logout();
      navigate('/login');
    } catch (error) {
      console.error('Logout failed:', error);
    }
  };

  return (
    <nav className="navbar">
      <div className="navbar-brand">
        <Link to="/">TikTok商城</Link>
      </div>
      
      <div className="navbar-menu">
        <div className="navbar-start">
          <Link to="/" className="navbar-item">
            首页
          </Link>
          <Link to="/products" className="navbar-item">
            商品
          </Link>
        </div>
        
        <div className="navbar-end">
          {isAuthenticated ? (
            <>
              <Link to="/cart" className="navbar-item cart-link">
                购物车
                {items && items.length > 0 && <span className="cart-badge">{items.length}</span>}
              </Link>
              <Link to="/orders" className="navbar-item">
                我的订单
              </Link>
              <button onClick={handleLogout} className="navbar-item logout-btn">
                登出
              </button>
            </>
          ) : (
            <>
              <Link to="/login" className="navbar-item">
                登录
              </Link>
              <Link to="/register" className="navbar-item">
                注册
              </Link>
            </>
          )}
        </div>
      </div>
    </nav>
  );
};

export default Navbar; 