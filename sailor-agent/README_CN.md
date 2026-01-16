# Sailor Agent 微服务

## 语言选择

[English](README.md) | [中文](./README_CN.md)
## 简介

Sailor Agent 是一个基于 FastAPI 构建的智能 AI 微服务，旨在提供高级对话功能、提示词管理和文本转 SQL 功能。它作为 AI 驱动应用的核心组件，支持将大型语言模型 (LLM) 无缝集成到各种业务场景中。

## 核心功能

### 对话式 AI
- 多轮对话管理
- 上下文感知响应
- 支持多种 LLM 模型
- 可定制的提示词模板

### 提示词管理
- 集中式提示词存储和检索
- 提示词版本控制
- 动态提示词生成
- 提示词优化工具

### 文本转 SQL
- 自然语言到 SQL 转换
- 数据库模式理解
- 查询验证和优化
- 支持多种数据库类型

### 智能体管理
- 创建和管理 AI 智能体
- 配置智能体功能
- 监控智能体性能
- 扩展智能体部署

## 架构

Sailor Agent 采用模块化架构设计，包含以下关键组件：

1. **API 层**：基于 FastAPI 的端点，用于客户端交互
2. **核心服务**：各种功能的业务逻辑实现
3. **LLM 集成**：与不同大型语言模型的连接器
4. **提示词引擎**：提示词管理和优化系统
5. **数据访问层**：数据库交互和 ORM 映射
6. **会话管理**：对话上下文处理
7. **日志和监控**：全面的日志记录和性能跟踪

## 技术栈

- **框架**：FastAPI
- **语言**：Python 3.8+
- **LLM 支持**：多种模型（如 Qwen 等）
- **数据库**：MySQL、Redis
- **部署**：Docker、Kubernetes
- **API 文档**：OpenAPI/Swagger

## 快速开始

### 先决条件
- Python 3.8 或更高版本
- Docker（可选，用于容器化部署）
- Kubernetes（可选，用于编排）

### 安装

1. 克隆仓库
2. 安装依赖：
   ```bash
   pip install -r requirements.txt
   ```
3. 设置环境变量（参考 `.env.example`）
4. 启动服务：
   ```bash
   python main.py
   ```

服务将在 `http://localhost:9595` 可用

### API 文档

访问 `http://localhost:9595/docs` 查看 Swagger UI 交互式 API 文档。

## 构建和测试

### 构建 Docker 镜像
```bash
docker build -t sailor-agent .
```

### 运行测试
```bash
python -m pytest
```

## 部署

### Docker
```bash
docker run -d -p 9595:9595 --env-file .env sailor-agent
```

### Kubernetes
```bash
helm install sailor-agent ./helm/sailor-agent
```

## API 端点

### 智能体管理
- `GET /api/v1/agents` - 列出所有智能体
- `POST /api/v1/agents` - 创建新智能体
- `GET /api/v1/agents/{agent_id}` - 获取智能体详情
- `PUT /api/v1/agents/{agent_id}` - 更新智能体配置
- `DELETE /api/v1/agents/{agent_id}` - 删除智能体

### 对话 API
- `POST /api/v1/chat` - 向智能体发送消息
- `GET /api/v1/chat/{session_id}` - 获取聊天历史
- `DELETE /api/v1/chat/{session_id}` - 清除聊天历史

### 文本转 SQL
- `POST /api/v1/text2sql` - 将自然语言转换为 SQL

## 配置

服务通过环境变量进行配置。请参考 `.env.example` 获取完整的配置选项列表。

## 贡献

我们欢迎社区贡献！有关更多信息，请参考我们的贡献指南。

## 许可证

[在此插入许可证信息]

## 联系我们

如有问题或需要支持，请通过 [在此插入联系信息] 联系我们的团队。
