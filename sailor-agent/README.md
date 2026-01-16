# Sailor Agent Microservice

## Language Selection
[English](./README.md) | [中文](README_CN.md)

## Introduction

Sailor Agent is an intelligent AI microservice built on FastAPI, designed to provide advanced conversational capabilities, prompt management, and Text-to-SQL functionality. It serves as a core component for AI-powered applications, enabling seamless integration of large language models (LLMs) into various business scenarios.

## Core Features

### Conversational AI
- Multi-turn dialogue management
- Context-aware responses
- Support for multiple LLM models
- Customizable prompt templates

### Prompt Management
- Centralized prompt storage and retrieval
- Version control for prompts
- Dynamic prompt generation
- Prompt optimization tools

### Text-to-SQL
- Natural language to SQL conversion
- Database schema understanding
- Query validation and optimization
- Support for multiple database types

### Agent Management
- Create and manage AI agents
- Configure agent capabilities
- Monitor agent performance
- Scale agent deployment

## Architecture

Sailor Agent follows a modular architecture design, consisting of the following key components:

1. **API Layer**: FastAPI-based endpoints for client interactions
2. **Core Services**: Business logic implementation for various features
3. **LLM Integration**: Connectors to different large language models
4. **Prompt Engine**: Prompt management and optimization system
5. **Data Access Layer**: Database interactions and ORM mappings
6. **Session Management**: Conversation context handling
7. **Logging and Monitoring**: Comprehensive logging and performance tracking

## Technology Stack

- **Framework**: FastAPI
- **Language**: Python 3.8+
- **LLM Support**: Multiple models (e.g., Qwen, etc.)
- **Database**: MySQL, Redis
- **Deployment**: Docker, Kubernetes
- **API Documentation**: OpenAPI/Swagger

## Getting Started

### Prerequisites
- Python 3.8 or higher
- Docker (optional, for containerized deployment)
- Kubernetes (optional, for orchestration)

### Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   pip install -r requirements.txt
   ```
3. Set up environment variables (refer to `.env.example`)
4. Start the service:
   ```bash
   python main.py
   ```

The service will be available at `http://localhost:9595`

### API Documentation

Access the Swagger UI at `http://localhost:9595/docs` for interactive API documentation.

## Build and Test

### Build Docker Image
```bash
docker build -t sailor-agent .
```

### Run Tests
```bash
python -m pytest
```

## Deployment

### Docker
```bash
docker run -d -p 9595:9595 --env-file .env sailor-agent
```

### Kubernetes
```bash
helm install sailor-agent ./helm/sailor-agent
```

## API Endpoints

### Agent Management
- `GET /api/v1/agents` - List all agents
- `POST /api/v1/agents` - Create a new agent
- `GET /api/v1/agents/{agent_id}` - Get agent details
- `PUT /api/v1/agents/{agent_id}` - Update agent configuration
- `DELETE /api/v1/agents/{agent_id}` - Delete an agent

### Conversational API
- `POST /api/v1/chat` - Send a message to the agent
- `GET /api/v1/chat/{session_id}` - Get chat history
- `DELETE /api/v1/chat/{session_id}` - Clear chat history

### Text-to-SQL
- `POST /api/v1/text2sql` - Convert natural language to SQL

## Configuration

The service is configured through environment variables. Refer to `.env.example` for the full list of configuration options.

## Contribute

We welcome contributions from the community! Please refer to our contribution guidelines for more information.

## License

[Insert License Information Here]

## Contact

For questions or support, please contact our team at [insert contact information here].
