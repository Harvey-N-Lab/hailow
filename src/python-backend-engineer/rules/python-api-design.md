# Python API Design Standards

## REST API Design
- Use proper HTTP methods (GET, POST, PUT, DELETE, PATCH)
- Use plural resource names (/users, not /user)
- Use nested resources for relationships
- Version APIs (/api/v1/)
- Return appropriate status codes
- Include pagination for lists
- Use consistent error response format

## Request/Response
- Accept and return JSON
- Validate request data with Pydantic/Marshmallow
- Include rate limiting headers
- Support filtering, sorting, pagination

## Authentication
- Use JWT or OAuth2
- Include authentication in headers
- Implement rate limiting
- Log authentication attempts

## Error Handling
- Consistent error response format
- Include error codes and messages
- Log errors with context
- Don't expose internal details
