import React, { useState, useRef, useEffect } from 'react';
import { aiChat } from '../api/aiApi';
import '../styles/AIAssistant.css';

interface Message {
  role: 'user' | 'assistant';
  content: string;
}

const AIAssistant: React.FC = () => {
  const [isOpen, setIsOpen] = useState<boolean>(false);
  const [inputMessage, setInputMessage] = useState<string>('');
  const [chatHistory, setChatHistory] = useState<Message[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [showEntrance, setShowEntrance] = useState<boolean>(true);
  const chatContainerRef = useRef<HTMLDivElement>(null);
  const inputRef = useRef<HTMLInputElement>(null);
  
  // 控制聊天窗口的开关
  const toggleChat = () => {
    setIsOpen(!isOpen);
    setShowEntrance(false);
    
    // 聊天窗口打开时，自动聚焦输入框
    if (!isOpen) {
      setTimeout(() => {
        inputRef.current?.focus();
      }, 300);
    }
  };
  
  // 发送消息
  const sendMessage = async () => {
    if (inputMessage.trim() === '') return;
    
    // 添加用户消息到聊天历史
    const userMessage: Message = { role: 'user', content: inputMessage };
    setChatHistory(prev => [...prev, userMessage]);
    setInputMessage('');
    setIsLoading(true);
    
    try {
      // 拼接历史聊天记录和当前消息
      let concatenatedContent = '';
      
      // 先添加历史聊天记录
      chatHistory.forEach(msg => {
        const rolePrefix = msg.role === 'user' ? '用户: ' : 'AI助手: ';
        concatenatedContent += `${rolePrefix}${msg.content}\n`;
      });
      
      // 添加当前消息
      concatenatedContent += `用户: ${inputMessage}\n`;
      
      // 调用AI聊天API
      const response = await aiChat({
        content: concatenatedContent
      });
      
      // 添加AI响应到聊天历史
      if (response.data && response.data.reply) {
        const assistantMessage: Message = {
          role: 'assistant',
          content: response.data.reply
        };
        setChatHistory(prev => [...prev, assistantMessage]);
      }
    } catch (error) {
      console.error('AI聊天请求失败:', error);
      // 添加错误消息
      const errorMessage: Message = {
        role: 'assistant',
        content: '抱歉，我遇到了一些问题，请稍后再试。'
      };
      setChatHistory(prev => [...prev, errorMessage]);
    } finally {
      setIsLoading(false);
    }
  };
  
  // 重置对话
  const resetChat = () => {
    setChatHistory([]);
  };
  
  // 处理输入框按键事件
  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') {
      sendMessage();
    }
  };
  
  // 聊天历史更新后滚动到底部
  useEffect(() => {
    if (chatContainerRef.current) {
      chatContainerRef.current.scrollTop = chatContainerRef.current.scrollHeight;
    }
  }, [chatHistory]);

  // 组件加载时的入场动画
  useEffect(() => {
    const timer = setTimeout(() => {
      setShowEntrance(false);
    }, 2000);
    
    return () => clearTimeout(timer);
  }, []);
  
  return (
    <div className="ai-assistant">
      {/* 助手图标按钮 */}
      <button
        className={`ai-assistant-icon ${showEntrance ? 'ai-assistant-entrance' : ''}`}
        onClick={toggleChat}
        aria-label="AI助手"
      >
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24">
          <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10s10-4.48 10-10S17.52 2 12 2zm0 3c1.66 0 3 1.34 3 3s-1.34 3-3 3s-3-1.34-3-3s1.34-3 3-3zm0 14.2c-2.5 0-4.71-1.28-6-3.22.03-1.99 4-3.08 6-3.08c1.99 0 5.97 1.09 6 3.08-1.29 1.94-3.5 3.22-6 3.22z" />
        </svg>
      </button>
      
      {/* 聊天窗口 */}
      {isOpen && (
        <div className="ai-chat-container">
          <div className="ai-chat-header">
            <h3>智能AI助手</h3>
            <div>
              <button
                className="ai-chat-reset"
                onClick={resetChat}
                aria-label="重置对话"
              >
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="18" height="18">
                  <path fill="currentColor" d="M17.65 6.35C16.2 4.9 14.21 4 12 4c-4.42 0-7.99 3.58-7.99 8s3.57 8 7.99 8c3.73 0 6.84-2.55 7.73-6h-2.08c-.82 2.33-3.04 4-5.65 4-3.31 0-6-2.69-6-6s2.69-6 6-6c1.66 0 3.14.69 4.22 1.78L13 11h7V4l-2.35 2.35z" />
                </svg>
              </button>
              <button
                className="ai-chat-close"
                onClick={toggleChat}
                aria-label="关闭"
              >
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="18" height="18">
                  <path fill="currentColor" d="M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z" />
                </svg>
              </button>
            </div>
          </div>
          
          {/* 聊天记录 */}
          <div className="ai-chat-messages" ref={chatContainerRef}>
            {chatHistory.length === 0 && (
              <div className="ai-chat-welcome">
                <div className="ai-welcome-icon">
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="32" height="32">
                    <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10s10-4.48 10-10S17.52 2 12 2zm0 3c1.66 0 3 1.34 3 3s-1.34 3-3 3s-3-1.34-3-3s1.34-3 3-3zm0 14.2c-2.5 0-4.71-1.28-6-3.22.03-1.99 4-3.08 6-3.08c1.99 0 5.97 1.09 6 3.08-1.29 1.94-3.5 3.22-6 3.22z" />
                  </svg>
                </div>
                <h4>您好！我是您的智能助手</h4>
                <p>有任何问题请随时向我咨询，我将竭诚为您服务！</p>
              </div>
            )}
            
            {chatHistory.map((message, index) => (
              <div
                key={index}
                className={`ai-chat-message ${
                  message.role === 'user' ? 'user-message' : 'assistant-message'
                }`}
              >
                {message.role === 'assistant' && (
                  <div className="message-avatar">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="20" height="20">
                      <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10s10-4.48 10-10S17.52 2 12 2zm0 3c1.66 0 3 1.34 3 3s-1.34 3-3 3s-3-1.34-3-3s1.34-3 3-3zm0 14.2c-2.5 0-4.71-1.28-6-3.22.03-1.99 4-3.08 6-3.08c1.99 0 5.97 1.09 6 3.08-1.29 1.94-3.5 3.22-6 3.22z" />
                    </svg>
                  </div>
                )}
                <div className="message-content">{message.content}</div>
              </div>
            ))}
            
            {/* 加载动画 */}
            {isLoading && (
              <div className="ai-chat-message assistant-message loading-message">
                <div className="message-avatar">
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="16" height="16">
                    <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10s10-4.48 10-10S17.52 2 12 2zm0 3c1.66 0 3 1.34 3 3s-1.34 3-3 3s-3-1.34-3-3s1.34-3 3-3zm0 14.2c-2.5 0-4.71-1.28-6-3.22.03-1.99 4-3.08 6-3.08c1.99 0 5.97 1.09 6 3.08-1.29 1.94-3.5 3.22-6 3.22z" />
                  </svg>
                </div>
                <div className="message-content loading">
                  <div className="loading-dots">
                    <span className="dot"></span>
                    <span className="dot"></span>
                    <span className="dot"></span>
                  </div>
                </div>
              </div>
            )}
          </div>
          
          {/* 输入框 */}
          <div className="ai-chat-input">
            <input
              ref={inputRef}
              type="text"
              value={inputMessage}
              onChange={(e) => setInputMessage(e.target.value)}
              onKeyDown={handleKeyDown}
              placeholder="请输入您的问题..."
              disabled={isLoading}
            />
            <button
              onClick={sendMessage}
              disabled={inputMessage.trim() === '' || isLoading}
              aria-label="发送"
              className={inputMessage.trim() !== '' ? 'active-send' : ''}
            >
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="20" height="20">
                <path fill="currentColor" d="M2.01 21L23 12 2.01 3 2 10l15 2-15 2z" />
              </svg>
            </button>
          </div>
        </div>
      )}
    </div>
  );
};

export default AIAssistant; 