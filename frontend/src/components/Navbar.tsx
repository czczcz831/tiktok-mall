import React, { useEffect, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { useCart } from '../contexts/CartContext';
import { getUserInfo } from '../api/userApi';
import { UserRole } from '../types/api';

const Navbar: React.FC = () => {
  const { isAuthenticated, logout } = useAuth();
  const { items } = useCart();
  const navigate = useNavigate();
  const [userRole, setUserRole] = useState<UserRole | null>(null);

  useEffect(() => {
    if (isAuthenticated) {
      fetchUserInfo();
    } else {
      setUserRole(null);
    }
  }, [isAuthenticated]);

  const fetchUserInfo = async () => {
    try {
      const response = await getUserInfo({});
      if (response.data.user) {
        setUserRole(response.data.user.role);
      }
    } catch (error) {
      console.error('获取用户信息失败:', error);
    }
  };

  const handleLogout = async () => {
    try {
      await logout();
      navigate('/login');
    } catch (error) {
      console.error('Logout failed:', error);
    }
  };

  const isAdmin = userRole !== UserRole.Customer;

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
          {isAuthenticated && isAdmin && (
            <Link to="/admin" className="navbar-item">
              管理后台
            </Link>
          )}
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