# Agent Configuration Domains

This directory contains metadata about available domains.

## Available Domains

### 1. DevOps Engineer
Infrastructure, CI/CD, deployment, and operational configurations.

### 2. Python Backend Engineer
Python backend development, API design, and service architecture.

### 3. JavaScript/TypeScript Software Engineer
JS/TS development for frontend and backend applications.

### 4. Data Engineer
Data pipelines, ETL, data quality, and data infrastructure.

### 5. Machine Learning Engineer
ML model development, training, deployment, and monitoring.

## Domain Structure

Each domain contains:
- **agents/**: Five agent prompts (researcher, planner, architect, implementer, reviewer)
- **rules/**: Domain-specific coding standards and practices
- **skills/**: Skill definitions and recommendations
- **commands/**: Workflow commands and shortcuts
- **contexts/**: Placeholder for user-specific context files

## Using Domains

Install domains using the `hailow` CLI:

```bash
# Install single domain
hailow install devops-engineer

# Install multiple domains
hailow install devops-engineer python-backend-engineer

# Install all domains
hailow install --all
```

## Domain Metadata

See `domains.yaml` for complete domain metadata including versions, descriptions, and aliases.
