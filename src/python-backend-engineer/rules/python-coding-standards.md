# Python Backend Coding Standards

## Code Style
- Follow PEP 8
- Use Black for formatting  
- Use isort for imports
- Max line length: 100

## Type Hints
- Use type hints for all functions
- Use mypy for type checking
- Required for public APIs

## Testing
- pytest framework
- Minimum 80% coverage
- Test naming: test_<function>_<scenario>

## Documentation
- Docstrings for all public functions (Google style)
- API documentation with OpenAPI/Swagger
- README for setup and usage

## Security
- Validate all inputs
- Use parameterized queries (prevent SQL injection)
- Hash passwords (bcrypt/argon2)
- Sanitize outputs (prevent XSS)
- Use environment variables for secrets

## Performance
- Use async/await for I/O
- Cache expensive operations
- Optimize database queries (select_related, prefetch_related)
- Use database indexes appropriately
