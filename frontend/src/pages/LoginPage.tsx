import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { LoginReq } from '../types/api';

const LoginPage: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const { login, isLoading } = useAuth();
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');

    if (!email || !password ) {
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

  return (
    <div className="login-page">
      <div className="login-container">
        <h2>登录</h2>
        {error && <div className="error-message">{error}</div>}
        <form onSubmit={handleSubmit}>
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
          <div className="form-group">
            <label htmlFor="password">密码</label>
            <input
              type="password"
              id="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>
          <button type="submit" disabled={isLoading}>
            {isLoading ? '登录中...' : '登录'}
          </button>
        </form>
        <div className="register-link">
          <p>
            还没有账号？ <a onClick={() => navigate('/register')}>注册</a>
          </p>
        </div>
      </div>
    </div>
  );
};

export default LoginPage; 