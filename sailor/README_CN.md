# AF Sailor 认知助手服务

## 语言选择

[English](README.md) | [中文](./README_CN.md)

## 介绍

AF Sailor 是一个基于 FastAPI 构建的强大认知助手微服务，旨在为数据智能平台提供一套全面的 AI 驱动功能。它集成了多种先进的 AI 特性，实现了与数据系统的智能交互。

## 核心功能

### 1. 认知助手
- 智能问答系统
- 上下文感知对话能力
- 多轮对话支持

### 2. 认知搜索
- 高级资产搜索功能
- 搜索查询的语义理解
- 高精度搜索结果排序

### 3. 文本转 SQL (T2S)
- 自然语言到 SQL 转换
- 支持复杂查询生成
-  schema 感知的 SQL 生成

### 4. 数据理解与 comprehension
- 自动数据分类
- 智能数据资产推荐
- 数据关系分析

### 5. 聊天与图表生成
- 交互式聊天功能
- 自然语言到图表转换
- 可视化数据表示

### 6. 提示词管理
- 集中式提示词仓库
- 提示词生命周期管理
- 提示词版本控制

## 技术栈

- **框架**: FastAPI
- **语言**: Python
- **部署**: Docker & Kubernetes
- **API 文档**: Swagger UI

## 快速开始

### 前置条件
- Python 3.8+
- Docker (可选)
- Kubernetes (可选)

### 安装

1. 克隆仓库
2. 安装依赖:
   ```bash
   pip install -r requirements.txt
   ```
3. 配置环境变量:
   ```bash
   cp .env.example .env
   # 编辑 .env 文件配置您的参数
   ```
4. 启动服务:
   ```bash
   python main.py
   ```

### API 访问

- **Swagger UI**: http://localhost:9797/docs
- **ReDoc**: http://localhost:9797/redoc

## 构建与测试

### 构建 Docker 镜像

```bash
docker build -t sailor .
```

### 运行测试

```bash
python -m pytest tests/
```

## 部署

### Kubernetes 部署

```bash
helm install sailor helm/sailor
```

## 贡献

我们欢迎贡献！请随时提交问题和拉取请求。

## 许可证

TODO: 添加许可证信息