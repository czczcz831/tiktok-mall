import React from 'react';
import '../styles/HomePage.css';

const HomePage: React.FC = () => {
  // 技术栈列表
  const techStack = [
    'React', 'TypeScript', 'Material-UI', 'Go', 'Kitex', 'Hertz','CasBin','Sentinel',
    'RocketMQ', 'MySQL', 'Redis', 'Docker','ELK','Prometheus', 'Consul', 'Kubernetes',
  ];

  return (
    <div className="home-page">
      <div className="container">
        <div className="home-content">
          <h1 className="title">CZCZCZ</h1>
          <h2 className="subtitle">欢迎来到我的TikTok商城演示项目</h2>
          
          <p className="description">
            本项目仅作为技术演示使用，无任何商业目的。
            这是一个基于微服务架构的电商系统Demo，集成了用户认证、商品管理、购物车、订单处理等功能。
            采用前后端分离设计，前端使用React+TypeScript构建，后端采用Go语言微服务架构。
          </p>
          
          <a 
            href="https://github.com/czczcz831/tiktok-mall" 
            target="_blank" 
            rel="noopener noreferrer"
            className="github-link"
          >
            GitHub Repository
          </a>
          
          <div className="tech-stack">
            <h3>技术栈</h3>
            <div className="tech-list">
              {techStack.map((tech, index) => (
                <div key={index} className="tech-item">
                  {tech}
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default HomePage;