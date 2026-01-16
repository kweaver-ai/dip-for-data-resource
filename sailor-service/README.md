# AF Sailor Service

## Language Selection

[English](./README.md) | [中文](README_CN.md)

## Introduction
AF Sailor Service is an intelligent cognitive assistant microservice within the AnyFabric ecosystem, providing advanced AI-powered capabilities for data understanding, knowledge discovery, and intelligent recommendations. It leverages cutting-edge technologies like large language models, knowledge graphs, and cognitive search to enable users to interact with data more intuitively and efficiently.

## Features

### Data Comprehension
AI-powered data analysis and comprehension capabilities to help users understand complex data structures and relationships.

### Knowledge Network
Graph analysis and exploration tools for visualizing and navigating complex knowledge networks.

### Intelligent Assistant
Conversational AI interface supporting multi-turn chat, question answering, and task automation.

### Cognitive Search
Advanced search capabilities across data catalogs, resources, and form views with semantic understanding.

### Smart Recommendations
AI-driven recommendations for subject models, tables, flows, and code standards.

### Knowledge Building
Tools for building and managing knowledge graphs and data models.

## Architecture

### System Architecture
**Core Components:**
- **API Layer:** RESTful interfaces for external and internal services
- **Service Layer:** Business logic implementation for various capabilities
- **Domain Layer:** Core business models and rules
- **Adapter Layer:** Integration with external systems and data sources
- **Infrastructure Layer:** Database access and resource management

## Getting Started

### Prerequisites
- Go 1.18+
- MySQL/MariaDB
- Elasticsearch/OpenSearch
- Kubernetes (for production deployment)

### Installation
```bash
# Clone the repository
git clone https://github.com/your-org/af-sailor-service.git
cd af-sailor-service

# Install dependencies
go mod download

# Build the service
go build -o sailor-service ./cmd/main.go

# Run the service
./sailor-service server --config=./cmd/server/config/config.yaml
```

### API Documentation
The service provides the following API endpoints:

| Path | Method | Description |
|------|--------|-------------|
| /api/af-sailor-service/v1/comprehension | GET | AI-powered data comprehension |
| /api/af-sailor-service/v1/assistant/qa | GET | Conversational AI assistant |
| /api/af-sailor-service/v1/cognitive/datacatalog/search | POST | Cognitive search for data catalogs |
| /api/af-sailor-service/v1/recommend/subject_model | POST | Recommend subject models |

## Build and Test
```bash
# Run tests
go test ./...

# Build Docker image
docker build -t af-sailor-service:latest .

# Run in Docker
docker run -p 8080:8080 --name sailor-service af-sailor-service:latest
```

## Contribution
We welcome contributions from the community! Please follow these guidelines:
1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License
Apache License 2.0