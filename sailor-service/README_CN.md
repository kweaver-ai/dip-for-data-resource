# AF Sailor 服务

## 语言选择

[English](README.md) | [中文](./README_CN.md)

## 简介
AF Sailor Service 是 AnyFabric 生态系统中的智能认知助手微服务，提供先进的 AI 驱动功能，用于数据理解、知识发现和智能推荐。它利用大型语言模型、知识图谱和认知搜索等前沿技术，使用户能够更直观、更高效地与数据交互。

## 功能特性

### 数据理解
基于 AI 的数据分析和理解能力，帮助用户理解复杂的数据结构和关系。

### 知识网络
用于可视化和导航复杂知识网络的图分析和探索工具。

### 智能助手
支持多轮对话、问答和任务自动化的会话 AI 界面。

### 认知搜索
跨数据目录、资源和表单视图的高级搜索功能，具有语义理解能力。

### 智能推荐
针对主题模型、表格、流程和代码标准的 AI 驱动推荐。

### 知识构建
用于构建和管理知识图谱和数据模型的工具。

## 架构设计

### 系统架构
**核心组件：**
- **API 层：** 为外部和内部服务提供 RESTful 接口
- **服务层：** 各种功能的业务逻辑实现
- **领域层：** 核心业务模型和规则
- **适配器层：** 与外部系统和数据源的集成
- **基础设施层：** 数据库访问和资源管理

## 快速开始

### 前提条件
- Go 1.18+
- MySQL/MariaDB
- Elasticsearch/OpenSearch
- Kubernetes（用于生产部署）

### 安装部署
```bash
# 克隆仓库
git clone https://github.com/your-org/af-sailor-service.git
cd af-sailor-service

# 安装依赖
go mod download

# 构建服务
go build -o sailor-service ./cmd/main.go

# 运行服务
./sailor-service server --config=./cmd/server/config/config.yaml
```

### API 文档
服务提供以下 API 端点：

| 路径 | 方法 | 描述 |
|------|------|------|
| /api/af-sailor-service/v1/comprehension | GET | 基于 AI 的数据理解 |
| /api/af-sailor-service/v1/assistant/qa | GET | 会话式 AI 助手 |
| /api/af-sailor-service/v1/cognitive/datacatalog/search | POST | 数据目录的认知搜索 |
| /api/af-sailor-service/v1/recommend/subject_model | POST | 推荐主题模型 |

## 构建与测试
```bash
# 运行测试
go test ./...

# 构建 Docker 镜像
docker build -t af-sailor-service:latest .

# 在 Docker 中运行
docker run -p 8080:8080 --name sailor-service af-sailor-service:latest
```

## 贡献指南
我们欢迎社区贡献！请遵循以下指南：
1. Fork 仓库
2. 创建功能分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

## 许可证
Apache License 2.0