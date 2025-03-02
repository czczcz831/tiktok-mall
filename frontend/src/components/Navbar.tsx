import React, { useEffect, useState } from 'react';
import { Link, useNavigate, useLocation } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { useCart } from '../contexts/CartContext';
import { getUserInfo } from '../api/userApi';
import { UserRole } from '../types/api';
import '../styles/Navbar.css';
// Material UI 图标导入
import ShoppingCartIcon from '@mui/icons-material/ShoppingCart';
import HomeIcon from '@mui/icons-material/Home';
import StorefrontIcon from '@mui/icons-material/Storefront';
import DashboardIcon from '@mui/icons-material/Dashboard';
import ExitToAppIcon from '@mui/icons-material/ExitToApp';
import ReceiptIcon from '@mui/icons-material/Receipt';
import PersonIcon from '@mui/icons-material/Person';
import PersonAddIcon from '@mui/icons-material/PersonAdd';

const Navbar: React.FC = () => {
  const { isAuthenticated, logout } = useAuth();
  const { items } = useCart();
  const navigate = useNavigate();
  const location = useLocation();
  const [userRole, setUserRole] = useState<UserRole | null>(null);
  const [scrolled, setScrolled] = useState<boolean>(false);

  useEffect(() => {
    if (isAuthenticated) {
      fetchUserInfo();
    } else {
      setUserRole(null);
    }
  }, [isAuthenticated]);

  useEffect(() => {
    const handleScroll = () => {
      if (window.scrollY > 50) {
        setScrolled(true);
      } else {
        setScrolled(false);
      }
    };

    window.addEventListener('scroll', handleScroll);
    
    return () => {
      window.removeEventListener('scroll', handleScroll);
    };
  }, []);

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
  const isActive = (path: string) => location.pathname === path;

  return (
    <nav className={`navbar ${scrolled ? 'scrolled' : ''}`}>
      <div className="navbar-brand">
        <Link to="/">TikTok商城</Link>
      </div>
      
      <div className="navbar-menu">
        <div className="navbar-start">
          <Link to="/" className={`navbar-item ${isActive('/') ? 'active' : ''}`}>
            <HomeIcon className="navbar-icon" />
            <span>首页</span>
          </Link>
          <Link to="/products" className={`navbar-item ${isActive('/products') ? 'active' : ''}`}>
            <StorefrontIcon className="navbar-icon" />
            <span>商品</span>
          </Link>
          {isAuthenticated && isAdmin && (
            <Link to="/admin" className={`navbar-item ${isActive('/admin') ? 'active' : ''}`}>
              <DashboardIcon className="navbar-icon" />
              <span>管理后台</span>
            </Link>
          )}
        </div>
        
        <div className="navbar-end">
          {isAuthenticated ? (
            <>
              <Link to="/cart" className={`navbar-item cart-link ${isActive('/cart') ? 'active' : ''}`}>
                <ShoppingCartIcon className="navbar-icon cart-icon" />
                <span>购物车</span>
                {items && items.length > 0 && <span className="cart-badge">{items.length}</span>}
              </Link>
              <Link to="/orders" className={`navbar-item ${isActive('/orders') ? 'active' : ''}`}>
                <ReceiptIcon className="navbar-icon" />
                <span>我的订单</span>
              </Link>
              <button onClick={handleLogout} className="navbar-item logout-btn">
                <ExitToAppIcon className="navbar-icon" />
                <span>登出</span>
              </button>
            </>
          ) : (
            <>
              <Link to="/login" className={`navbar-item ${isActive('/login') ? 'active' : ''}`}>
                <PersonIcon className="navbar-icon" />
                <span>登录</span>
              </Link>
              <Link to="/register" className={`navbar-item ${isActive('/register') ? 'active' : ''}`}>
                <PersonAddIcon className="navbar-icon" />
                <span>注册</span>
              </Link>
            </>
          )}
        </div>
      </div>
    </nav>
  );
};

export default Navbar; 