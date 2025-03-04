import React, { useEffect, useState } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate, useLocation, useNavigate } from 'react-router-dom';
import { AuthProvider, useAuth } from './contexts/AuthContext';
import { CartProvider } from './contexts/CartContext';
import Navbar from './components/Navbar';
import AIAssistant from './components/AIAssistant';
import LoginPage from './pages/LoginPage';
import RegisterPage from './pages/RegisterPage';
import ProductListPage from './pages/ProductListPage';
import ProductDetailPage from './pages/ProductDetailPage';
import CartPage from './pages/CartPage';
import CheckoutPage from './pages/CheckoutPage';
import OrderListPage from './pages/OrderListPage';
import AdminPage from './pages/AdminPage';
import HomePage from './pages/HomePage';
import { getUserInfo } from './api/userApi';
import { UserRole } from './types/api';
import './index.css';
import './styles/Navbar.css';
import './styles/ProductList.css';
import './styles/ProductDetail.css';
import './styles/AIAssistant.css';
import './styles/CartPage.css';

// 需要登录的路由保护组件
const PrivateRoute: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const { isAuthenticated } = useAuth();
  const location = useLocation();
  
  if (!isAuthenticated) {
    return <Navigate to="/login" state={{ from: location }} replace />;
  }
  
  return <>{children}</>;
};

// 管理员路由保护组件
const AdminRoute: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const { isAuthenticated } = useAuth();
  const location = useLocation();
  const navigate = useNavigate();
  const [loading, setLoading] = useState(true);
  const [isAdmin, setIsAdmin] = useState(false);
  
  useEffect(() => {
    const checkRole = async () => {
      if (!isAuthenticated) {
        navigate('/login', { state: { from: location } });
        return;
      }
      
      try {
        const response = await getUserInfo({});
        if (response.data.user) {
          const role = response.data.user.role;
          setIsAdmin(role !== UserRole.Customer);
        }
      } catch (error) {
        console.error('获取用户角色失败:', error);
      } finally {
        setLoading(false);
      }
    };
    
    checkRole();
  }, [isAuthenticated, location, navigate]);
  
  if (loading) {
    return <div>Loading...</div>;
  }
  
  if (!isAdmin) {
    return <Navigate to="/" replace />;
  }
  
  return <>{children}</>;
};

const App: React.FC = () => {
  return (
    <Router>
      <AuthProvider>
        <CartProvider>
          <div className="app">
            <Navbar />
            <Routes>
              <Route path="/" element={<HomePage />} />
              <Route path="/products" element={<ProductListPage />} />
              <Route path="/login" element={<LoginPage />} />
              <Route path="/register" element={<RegisterPage />} />
              <Route path="/products" element={<ProductListPage />} />
              <Route path="/product/:uuid" element={<ProductDetailPage />} />
              <Route path="/cart" element={
                <PrivateRoute>
                  <CartPage />
                </PrivateRoute>
              } />
              <Route path="/checkout" element={
                <PrivateRoute>
                  <CheckoutPage />
                </PrivateRoute>
              } />
              <Route path="/orders" element={
                <PrivateRoute>
                  <OrderListPage />
                </PrivateRoute>
              } />
              <Route path="/admin" element={
                <AdminRoute>
                  <AdminPage />
                </AdminRoute>
              } />
            </Routes>
            <AIAssistant />
          </div>
        </CartProvider>
      </AuthProvider>
    </Router>
  );
};

export default App;
