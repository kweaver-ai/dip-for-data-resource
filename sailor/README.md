# AF Sailor Cognitive Assistant Service

## Language Selection
[English](./README.md) | [中文](README_CN.md)

## Introduction

AF Sailor is a powerful cognitive assistant microservice built on FastAPI, designed to provide a comprehensive suite of AI-powered capabilities for data intelligence platforms. It integrates multiple advanced AI features to enable intelligent interactions with data systems.

## Core Features

### 1. Cognitive Assistant
- Intelligent question-answering system
- Context-aware conversational capabilities
- Multi-round dialogue support

### 2. Cognitive Search
- Advanced asset search functionality
- Semantic understanding of search queries
- High-precision search results ranking

### 3. Text-to-SQL (T2S)
- Natural language to SQL conversion
- Support for complex query generation
- Schema-aware SQL generation

### 4. Data Understanding & Comprehension
- Automated data categorization
- Intelligent data asset recommendation
- Data relationship analysis

### 5. Chat & Chart Generation
- Interactive chat functionality
- Natural language to chart conversion
- Visual data representation

### 6. Prompt Management
- Centralized prompt repository
- Prompt lifecycle management
- Prompt version control

## Technology Stack

- **Framework**: FastAPI
- **Language**: Python
- **Deployment**: Docker & Kubernetes
- **API Documentation**: Swagger UI

## Getting Started

### Prerequisites
- Python 3.8+
- Docker (optional)
- Kubernetes (optional)

### Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   pip install -r requirements.txt
   ```
3. Configure environment variables:
   ```bash
   cp .env.example .env
   # Edit .env file with your configuration
   ```
4. Start the service:
   ```bash
   python main.py
   ```

### API Access

- **Swagger UI**: http://localhost:9797/docs
- **ReDoc**: http://localhost:9797/redoc

## Build and Test

### Build Docker Image

```bash
docker build -t sailor .
```

### Run Tests

```bash
python -m pytest tests/
```

## Deployment

### Kubernetes Deployment

```bash
helm install sailor helm/sailor
```

## Contribute

We welcome contributions! Please feel free to submit issues and pull requests.

## License

TODO: Add license information