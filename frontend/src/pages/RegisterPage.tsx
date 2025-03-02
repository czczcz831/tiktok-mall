import React, { useState, useRef, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { RegisterReq } from '../types/api';
import '../styles/Auth.css';
// Material UI 图标导入
import VisibilityIcon from '@mui/icons-material/Visibility';
import VisibilityOffIcon from '@mui/icons-material/VisibilityOff';
import EmailIcon from '@mui/icons-material/Email';
import LockIcon from '@mui/icons-material/Lock';
import PersonAddIcon from '@mui/icons-material/PersonAdd';
import SecurityIcon from '@mui/icons-material/Security';

const RegisterPage: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [error, setError] = useState('');
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const [passwordStrength, setPasswordStrength] = useState<'weak' | 'medium' | 'strong' | 'very-strong' | ''>('');
  const { register, isLoading } = useAuth();
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

  // 计算密码强度
  useEffect(() => {
    if (!password) {
      setPasswordStrength('');
      return;
    }

    const hasLetters = /[a-zA-Z]/.test(password);
    const hasNumbers = /\d/.test(password);
    const hasSpecialChars = /[!@#$%^&*(),.?":{}|<>]/.test(password);
    
    if (password.length < 6) {
      setPasswordStrength('weak');
    } else if (password.length < 8) {
      setPasswordStrength('medium');
    } else if (hasLetters && hasNumbers && hasSpecialChars) {
      setPasswordStrength('very-strong');
    } else if ((hasLetters && hasNumbers) || (hasLetters && hasSpecialChars) || (hasNumbers && hasSpecialChars)) {
      setPasswordStrength('strong');
    } else {
      setPasswordStrength('medium');
    }
  }, [password]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');

    if (!email || !password || !confirmPassword) {
      setError('请填写所有字段');
      return;
    }

    if (password !== confirmPassword) {
      setError('两次输入的密码不一致');
      return;
    }

    if (passwordStrength === 'weak') {
      setError('密码强度太弱，请设置更强的密码');
      return;
    }

    try {
      const registerData: RegisterReq = {
        email,
        password,
        confirm_password: confirmPassword
      };
      await register(registerData);
      navigate('/login');
    } catch (err) {
      setError('注册失败，请稍后再试');
      console.error(err);
    }
  };

  const togglePasswordVisibility = () => {
    setShowPassword(!showPassword);
  };

  const toggleConfirmPasswordVisibility = () => {
    setShowConfirmPassword(!showConfirmPassword);
  };

  const getPasswordStrengthText = () => {
    switch (passwordStrength) {
      case 'weak':
        return '弱 - 请尝试包含字母、数字和特殊字符';
      case 'medium':
        return '中等 - 尝试加入特殊字符增强安全性';
      case 'strong':
        return '强 - 非常好的密码';
      case 'very-strong':
        return '非常强 - 完美的密码';
      default:
        return '';
    }
  };

  return (
    <div className="register-page">
      <div 
        ref={cardRef}
        className="auth-card tilt-effect"
        onMouseMove={handleMouseMove}
        onMouseLeave={handleMouseLeave}
      >
        <div className="auth-card-inner tilt-content">
          <div className="auth-header">
            <h2>创建账户</h2>
            <p>注册一个新账户，开始您的购物体验</p>
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
                placeholder="设置您的密码"
                required
              />
              <div className="password-toggle" onClick={togglePasswordVisibility}>
                {showPassword ? <VisibilityOffIcon /> : <VisibilityIcon />}
              </div>
              
              {password && (
                <>
                  <div className="password-strength">
                    <div className={`password-strength-bar ${passwordStrength}`}></div>
                  </div>
                  <div className="password-strength-text">
                    <SecurityIcon sx={{ fontSize: '0.8rem', verticalAlign: 'middle', marginRight: '0.3rem' }} />
                    {getPasswordStrengthText()}
                  </div>
                </>
              )}
            </div>
            
            <div className="form-group">
              <label htmlFor="confirmPassword">
                <LockIcon sx={{ fontSize: '0.9rem', verticalAlign: 'middle', marginRight: '0.3rem' }} />
                确认密码
              </label>
              <input
                type={showConfirmPassword ? "text" : "password"}
                id="confirmPassword"
                value={confirmPassword}
                onChange={(e) => setConfirmPassword(e.target.value)}
                placeholder="再次输入密码"
                required
              />
              <div className="password-toggle" onClick={toggleConfirmPasswordVisibility}>
                {showConfirmPassword ? <VisibilityOffIcon /> : <VisibilityIcon />}
              </div>
            </div>
            
            <button type="submit" className="submit-button" disabled={isLoading}>
              {isLoading ? (
                <>
                  <span className="spinner"></span>
                  注册中...
                </>
              ) : (
                <>
                  <PersonAddIcon sx={{ marginRight: '0.5rem' }} />
                  注册
                </>
              )}
            </button>
          </form>
          
          <div className="auth-link">
            <p>
              已有账号？ <a onClick={() => navigate('/login')}>立即登录</a>
            </p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default RegisterPage; 