import React, { useState, useEffect, useRef } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { LoginReq } from '../types/api';
import '../styles/Auth.css';
// Material UI 图标导入
import VisibilityIcon from '@mui/icons-material/Visibility';
import VisibilityOffIcon from '@mui/icons-material/VisibilityOff';
import EmailIcon from '@mui/icons-material/Email';
import LockIcon from '@mui/icons-material/Lock';
import LoginIcon from '@mui/icons-material/Login';

const LoginPage: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [showPassword, setShowPassword] = useState(false);
  const { login, isLoading } = useAuth();
  const navigate = useNavigate();
  const cardRef = useRef<HTMLDivElement>(null);

  // 3D 卡片倾斜效果
  const handleMouseMove = (e: React.MouseEvent<HTMLDivElement>) => {
    if (!cardRef.current) return;
    
    const card = cardRef.current;
    const rect = card.getBoundingClientRect();
    const x = e.clientX - rect.left;
    const y = e.clientY - rect.top;
    
    const centerX = rect.width / 2;
    const centerY = rect.height / 2;
    
    const rotateX = (y - centerY) / 20;
    const rotateY = (centerX - x) / 20;
    
    card.style.transform = `perspective(1000px) rotateX(${rotateX}deg) rotateY(${rotateY}deg)`;
  };
  
  const handleMouseLeave = () => {
    if (!cardRef.current) return;
    cardRef.current.style.transform = 'perspective(1000px) rotateX(0) rotateY(0)';
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');

    if (!email || !password) {
      setError('请填写所有字段');
      return;
    }

    try {
      const loginData: LoginReq = {
        email,
        password,
      };
      await login(loginData);
      navigate('/');
    } catch (err) {
      setError('登录失败，请检查您的凭据');
      console.error(err);
    }
  };

  const togglePasswordVisibility = () => {
    setShowPassword(!showPassword);
  };

  return (
    <div className="login-page">
      <div 
        ref={cardRef}
        className="auth-card tilt-effect"
        onMouseMove={handleMouseMove}
        onMouseLeave={handleMouseLeave}
      >
        <div className="auth-card-inner tilt-content">
          <div className="auth-header">
            <h2>欢迎回来</h2>
            <p>请登录您的账户以继续购物之旅</p>
          </div>
          
          {error && <div className="error-message">{error}</div>}
          
          <form onSubmit={handleSubmit} className="auth-form">
            <div className="form-group">
              <label htmlFor="email">
                <EmailIcon sx={{ fontSize: '0.9rem', verticalAlign: 'middle', marginRight: '0.3rem' }} />
                邮箱地址
              </label>
              <input
                type="email"
                id="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                placeholder="请输入您的邮箱"
                required
              />
            </div>
            
            <div className="form-group">
              <label htmlFor="password">
                <LockIcon sx={{ fontSize: '0.9rem', verticalAlign: 'middle', marginRight: '0.3rem' }} />
                密码
              </label>
              <input
                type={showPassword ? "text" : "password"}
                id="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder="请输入您的密码"
                required
              />
              <div className="password-toggle" onClick={togglePasswordVisibility}>
                {showPassword ? <VisibilityOffIcon /> : <VisibilityIcon />}
              </div>
            </div>
            
            <button type="submit" className="submit-button" disabled={isLoading}>
              {isLoading ? (
                <>
                  <span className="spinner"></span>
                  登录中...
                </>
              ) : (
                <>
                  <LoginIcon sx={{ marginRight: '0.5rem' }} />
                  登录
                </>
              )}
            </button>
          </form>
          
          <div className="auth-link">
            <p>
              还没有账号？ <a onClick={() => navigate('/register')}>立即注册</a>
            </p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default LoginPage; 