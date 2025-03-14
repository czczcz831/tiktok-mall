## 一、项目介绍

### 项目概述
<div align=center>
    
![](https://oss.czczcz.xyz/blog/202503060038102.png)

</div>

<p align="center">TikTok Mall 是一个基于微服务架构的电商平台，提供用户注册登录、商品浏览、购物车管理、订单结算和支付等完整电商功能，采用现代化的技术栈和架构设计，具有高可用性、可扩展性和安全性。
</p>

项目飞书文档：https://gagjcxhxrb.feishu.cn/docx/IU6BdqBERoPYSCxycb1cBqDQnLc

### 项目服务地址

演示地址: https://mall.czczcz.xyz

API 地址: https://tiktok-mall-api.czczcz.xyz

API 文档:

https://apifox.com/apidoc/shared-0b153f12-af64-4717-9b6c-96c5f504f7d8

### 项目仓库地址
[GitHub - czczcz831/tiktok-mall: tiktok-mall](https://github.com/czczcz831/tiktok-mall)

### 3.1 技术选型
#### 3.1.1 场景分析

Demo演示集群配置：三台E52666V3节点,每台32G内存,120G NFS存储

Demo演示集群瓶颈： 单点NFS云存储,'阿里云5MB,家庭线路内网穿透是网络瓶颈
![image.png](https://oss.czczcz.xyz/blog/202503062142578.png)

本项目设计为支持中等规模的电商平台，系统需要处理商品、用户、订单等数据，预计需要以下资源：

- 存储空间：初期约50GB，包括用户数据、商品信息、订单记录等
    
- 服务器配置：8核16G内存的服务器5台，用于部署微服务集群
    
- 网络带宽：100Mbps，满足正常访问需求

#### 3.1.2 技术栈
##### 3.1.2.1 后端技术栈：

- 编程语言：Golang
    
- 微服务框架：
    
    - Kitex：字节跳动开源的高性能RPC框架，用于服务间通信
        
    - Hertz：字节跳动开源的HTTP框架，用于API网关
        
    - Eino : 字节跳动开源的大模型框架，实现Ark大模型接入和工具调用
        
- 数据存储：
    
    - MySQL：关系型数据库，存储用户、商品、订单等核心数据
        
    - Redis：缓存层，用于存储会话信息、库存缓存等
        
- 消息队列：
    
    - RocketMQ：分布式事务处理和延时消息发送，处理异步消息，如订单状态变更通知、库存变更等
        
- 服务发现与配置：
    
    - Consul：服务注册、发现和KV配置管理
        
- 安全与权限：
    
    - Casbin：细粒度的访问控制框架，实现RBAC权限模型
        
    - JWT+EeDSA ：JWT非对称加密，只有Auth服务有资格(Private Key)签发Token,就算其他微服务被攻破也只有Public key,无法签发token
        
- 系统稳定性：
    
    - Sentinel：流量控制、熔断降级、系统负载保护
        
    - Kubernetes-lstio: 分布式限流
        
- 监控与日志：
    
    - Prometheus + Grafana：系统监控和可视化
        
    - ELK Stack(ElasticSearch + Logstash + Kibana)：日志收集、存储和分析
        
- 分布式UUID生产 : Snowflake
    
- 容器编排：
    
    - Kubernetes + Helm：容器化部署和编排
        
    - Docker-Compose: 单机部署容器化
        
- CI/CD：
    
    - Github Actions: 实现代码上传自动编译推送镜像
        
- 负载均衡：
    
    - Ingress 控制器（如 ingress-nginx）: 处理外部流量并将其路由到后端 Service
        
    - Kube-proxy：在 Service 层面通过IPVS实现负载均衡的具体逻辑,将流量分发到后端 Pod
        
- 前端部署：
    
    - 腾讯云OSS: 通过静态站点桶的方式部署前端
        
    - CDN 和安全防护: 通过Cloudflare 作为安全代理，Cloudflare CDN全球加速访问
        

  

##### 3.1.2.2 前端技术栈(接口未完全对接）：

- React 19：用户界面构建
    
- TypeScript：类型安全的JavaScript超集
    
- Material UI 6：UI组件库
    
- React Router 7：前端路由管理
    
- Axios：HTTP客户端
    
- React Toastify：通知提示组件：

### 3.2 架构设计

**核心理念**： 所有的微服务必须是无状态的，所有的配置都在Consul上动态配置，微服务环境变量只用注入Consul地址，ACL token,和Config Key。
#### 3.2.1 请求的路程

浏览器->Cloudflare->阿里云Frp->Ingress-Nginx->Kube-Proxy->Hertz
红色部分为链路延迟瓶颈
![image.png](https://oss.czczcz.xyz/blog/202503062146019.png)

#### 3.2.2 微服务架构设计及链路

##### Api(Hertz Gateway)

Api服务负责作为外部请求的入口，是整个架构的核心，负责权限校验，限流，RPC调用各个微服务。
![image.png](https://oss.czczcz.xyz/blog/202503062147510.png)

##### Auth(Token签发续期服务)

负责JWT Token签发和RefreshToken续期，采用非对称加密EdDSA算法，Auth服务拥有私钥，API网关通过公钥对请求携带的Token进行校验，这样一方面减少了RPC调用，网关可以快速鉴权，一方面即使网关配置泄露也无法签发Token
![image.png](https://oss.czczcz.xyz/blog/202503062147573.png)

##### Cart(购物车服务)

负责管理用户的购物车，可以添加产品到购物车，清空购物车，获取用户购物车。

RocketMQ消费者： 监听订单事务队列，收到创建订单事件自动清空购物车
![image.png](https://oss.czczcz.xyz/blog/202503062147157.png)


##### Checkout(结算服务)

负责结算用户的商品，计算价格，调用Order订单服务创建订单，还负责用户地址管理(增删改查)
![image.png](https://oss.czczcz.xyz/blog/202503062147806.png)


##### Eino(AI服务)

负责调用大模型AI解决用户问题，利用下单工具，产品信息工具，用户地址和用户订单等工具通过大语言模型帮助用户操作，提升用户体验，外部可以对接Ark,OpenAI,Deepseek等大语言模型
![image.png](https://oss.czczcz.xyz/blog/202503062148297.png)

##### Order(订单服务)

负责管理用户订单数据，创建订单进行调用Product商品库存服务进行缓存预扣

RocketMQ消费者：一个消费者负责监听订单延时队列，超时未支付取消订单，另一个消费者负责监听支付成功事件，支付成功后标记订单为已支付

RocketMQ生产者： 生产者负责在订单事务队列生产创建订单事件
![image.png](https://oss.czczcz.xyz/blog/202503062148063.png)

##### Payment(支付服务）

负责用户支付订单，生成账单，接收用户信用卡数据调用第三方支付服务支付

RocketMQ消费者： 监听支付延时队列，超时未支付则取消账单

RocketMQ生产者： 支付成功后生产支付成功事件到支付事务队列
![image.png](https://oss.czczcz.xyz/blog/202503062148761.png)

##### Product(商品和库存服务)

负责管理商品和库存数据，负责商品的增删改查和保障库存一致性，缓存预扣，商品库存增减时更新预扣缓存

RocketMQ消费者：负责监听支付成功事件，在数据库层面扣减库存，并更新缓存
![image.png](https://oss.czczcz.xyz/blog/202503062148048.png)

##### User(用户服务)

负责用户的登录和注册，获取用户信息等，采用MD5加密存储密码,Redis记录用户登录状态
![image.png](https://oss.czczcz.xyz/blog/202503062148804.png)
#### 3.2.3 表设计+ID生成

采用Snowflake算法生成分布式唯一UUID,UUID有唯一索引加快查询速度

之所以还要有普通自增id作为主键是因为在并发插入下由于机器ID不一致虽然时间是自增的，但依然会造成页分裂，所以采用两个ID结合一定程度用空间换取了查询插入效率
![image.png](https://oss.czczcz.xyz/blog/202503062149962.png)
![image.png](https://oss.czczcz.xyz/blog/202503062150065.png)


### 3.3 项目代码和核心功能实现
#### 3.3.1 项目结构

```bash
/tiktok-mall
├── app/                 # 微服务实现
│   ├── api/             # API网关服务
│   ├── auth/            # 认证服务
│   ├── cart/            # 购物车服务
│   ├── checkout/        # 结算服务
│   ├── eino/            # 智能助手服务
│   ├── order/           # 订单服务
│   ├── payment/         # 支付服务
│   ├── product/         # 商品服务
│   └── user/            # 用户服务
├── client/              # 各服务的客户端SDK
├── common/              # 公共代码和工具
├── deploy-env/          # 部署环境配置
│   ├── consul/          # Consul配置
│   ├── docker-elk/      # ELK Stack配置
│   ├── prome-grafana/   # Prometheus和Grafana配置
│   └── rocketmq/        # RocketMQ配置
├── docker/              # Docker构建文件
├── frontend/            # 前端React应用
├── helm-chart/          # Kubernetes Helm部署配置
├── idl/                 # 接口定义文件
└── scripts/             # 部署和管理脚本
```

  

  

#### 3.3.2 Consul服务注册和服务发现

相关代码: conf/conf.go, main.go, client/*

每个服务的conf/conf.go下都是这个服务的配置结构体，这个结构体根据环境的配置Key去Consul里面获取YAML配置，序列化到结构体中，单例模式供全局使用。

在main函数中注册 ：

```Go

        r, err := consul.NewConsulRegisterWithConfig(conf.GetConsulCfg())
        if err != nil {
                klog.Fatalf("new consul register failed: %v", err)
        }
        opts = append(opts, server.WithRegistry(r)) 
```

项目根目录client包则是各个微服务的Client,与app下的代码完全解耦，任何服务想用某个微服务的client只需要引入这个包就可以,十分优雅方便

e.g: Auth Client

```go
package auth

var (
        // todo edit custom config
        defaultClient     RPCClient
        defaultDstService = "auth"
        defaultClientOpts = []client.Option{
                client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
                client.WithTransportProtocol(transport.TTHeader),
        }
        once sync.Once
)

func init() {
        //Consul client resolve
        consulHost := os.Getenv("CONSUL_HOST")
        consulPort := os.Getenv("CONSUL_PORT")
        r, err := consul.NewConsulResolver(net.JoinHostPort(consulHost, consulPort))

        if err != nil {
                klog.Fatalf("new consul resolver failed: %v", err)
        }

        //Sentinel middleware
        bf := func(ctx context.Context, req, resp interface{}, blockErr error) error {
                return errors.New("circuit break! ")
        }
        defaultClientOpts = append(defaultClientOpts, client.WithMiddleware(sentinel.SentinelClientMiddleware(
                sentinel.WithBlockFallback(bf),
        )))

        defaultClientOpts = append(defaultClientOpts, client.WithResolver(r))

        DefaultClient()
}
```

调用只需要引入就可以了

```Go
import (
        auth "github.com/czczcz831/tiktok-mall/client/auth/kitex_gen/auth"
        authAgent "github.com/czczcz831/tiktok-mall/client/auth/rpc/auth"
)

...
deliveryTokenResp, err := authAgent.DeliverTokenByRPC(h.Context, &auth.DeliverTokenReq{
                UserUuid: uuid,
        })

        if err != nil {
                return nil, &packer.MyError{
                        Code: packer.AUTH_DELIBER_TOKEN_ERROR,
                        Err:  err,
                }
        }
...
```
![image.png](https://oss.czczcz.xyz/blog/202503062151818.png)
![image.png](https://oss.czczcz.xyz/blog/202503062151051.png)
![image.png](https://oss.czczcz.xyz/blog/202503062151319.png)


#### 3.3.3 用户鉴权、权限控制、黑名单、登录状态管理

相关代码: app/auth , app/api/dal/casbin, app/api/router/api/middleware.go 等

鉴权采用JWT+EdDSA非对称签名加密，payload为用户uuid,只有Auth服务拥有私钥签发资格，Api网关只有公钥验证资格，Token签发只有24小时有效期，RefreshToken有一星期有效期，Token过期后前端可以拿着RefreshToken继续续期。

  

权限控制采用Casbin RBAC权限控制，通过regexMatch实现正则表达式权限匹配，分为顾客，卖家，管理员三个角色，用户初始注册为顾客角色，采用Hertz Casbin Middleware快速集成。

在middleware.go下就可以配置哪些接口需要鉴权，哪些不需要。

```go
[request_definition]
r = sub, obj

[policy_definition]
p = sub, obj

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && regexMatch(r.obj, p.obj)
```

  

黑名单则是有一个单独的角色Banned，如果用户拥有这个角色则说明被封禁。

  

登录状态管理是登录时会在Redis里设置这个token键，TTL为token过期时间，如果用户登出的话则删除这个键,如果没有这个键则登录失败。

e.g :
![image.png](https://oss.czczcz.xyz/blog/202503062151449.png)
![image.png](https://oss.czczcz.xyz/blog/202503062151875.png)
![image.png](https://oss.czczcz.xyz/blog/202503062151070.png)

#### 3.3.4 RocketMQ分布式事务

在本项目的很多异步场景中，如何保证本地事务和RocketMQ消息的提交是这个项目维护一致性的关键问题

比如在订单服务中，如果订单在数据库创建了但是没有成功发送到消息队列，那么购物车服务就不能清空，如果发送到消息队列了但是数据库事务失败，购物车又被误清除了。

同样在支付服务中，如果支付成功但是事件没有发送到队列中，库存服务就无法扣减库存，订单服务也无法标记完成支付，造成超卖和重复支付的结果。

所以在本项目中引入了RocketMQ分布式事务保障项目的数据一致性。

![image.png](https://oss.czczcz.xyz/blog/202503062152396.png)

生产者回查e.g:

![image.png](https://oss.czczcz.xyz/blog/202503062152898.png)
![image.png](https://oss.czczcz.xyz/blog/202503062152703.png)

#### 3.3.5 订单，支付 定时取消

RocketMQ除了可以用来做分布式事务外，还可以利用它的延时队列实现订单，支付的超时取消

生产者eg:

![image.png](https://oss.czczcz.xyz/blog/202503062152252.png)


消费者eg:

![image.png](https://oss.czczcz.xyz/blog/202503062152435.png)


#### 3.3.6 库存缓存预扣+补偿

库存防止超买超卖问题是电商场景的核心需求，本项目基于Redis + Lua 实现了缓存预扣和补偿机制防止超买超卖。

订单创建前，订单服务会去调用库存服务进行缓存预扣，预扣成功后才会创建订单，失败则提示库存不足。库存数量存储在Redis中，如果键不存在会有服务去拿分布式锁构建缓存(防止秒杀时大量请求绕过缓存打到数据库)。

只有用户支付成功后才会异步通知库存服务去数据库里真正扣减库存。

库存增减时也会去补偿Redis的库存，保障数据一致。

![image.png](https://oss.czczcz.xyz/blog/202503062152922.png)


![image.png](https://oss.czczcz.xyz/blog/202503062153732.png)


#### 3.3.7 AI大模型

采用Cloudwego Eino框架，将服务划分为两部分，Agent和tools，分别编写逻辑，使用React Agent调用Ark豆包模型进行函数工具调用

![image.png](https://oss.czczcz.xyz/blog/202503062153022.png)


![image.png](https://oss.czczcz.xyz/blog/202503062153542.png)


#### 3.3.8 K8s Istio分布式限流和Sentinel本地限流,服务熔断

基于K8s lstio实现分布式限流

通过Sentinel实现可配置(sentinel/rules.go)限流，还有基于错误率的服务熔断,与Hertz快速集成

服务熔断可实现全闭，半开，全开三种状态，当熔断发生时，会切到全闭，过一段时间切到半开，放一点请求出去，成功的话又回到全开。

![image.png](https://oss.czczcz.xyz/blog/202503062153392.png)
![image.png](https://oss.czczcz.xyz/blog/202503062153847.png)

![image.png](https://oss.czczcz.xyz/blog/202503062153503.png)


#### 3.3.9 可观测性

本项目采用Prometheus+Grafana+ELK实现可观测性

Grafana可以查看各服务的实例数，Hertz Kitex P99延迟,AVG延迟，吞吐量

![image.png](https://oss.czczcz.xyz/blog/202503062154301.png)


ELK日志采集

![image.png](https://oss.czczcz.xyz/blog/202503062154105.png)


![image.png](https://oss.czczcz.xyz/blog/202503062154061.png)



#### 3.3.10 容器化,CI/CD，K8s部署

本项目配置了Github Action，当代码推到main分支后可以点击自动编译推送到阿里云的镜像仓库，极大地方便了部署和开发，由于没钱承担流量费用，就不发布Demo镜像地址了。

项目docker文件夹下可以单独编译每个微服务为镜像，每个镜像约50MB，可以在目录下直接用Docker-compose直接全部编译拉起，由于中间件过多，不建议在单机上跑全部服务，所以Docker-compose里只有微服务，没有其他中间件，其他所需中间件和配置在deploy-env下可选配。

本项目配置了Helm Chart，在values.yaml里填上你的Consul参数和镜像地址则可以在你的k8s集群上自动部署，此外还配置了Prometheus注解，如果你的项目有Prometheus-Operator的话可自动识别抓取Metrics

![image.png](https://oss.czczcz.xyz/blog/202503062154327.png)


![image.png](https://oss.czczcz.xyz/blog/202503062154163.png)

![image.png](https://oss.czczcz.xyz/blog/202503062154104.png)

