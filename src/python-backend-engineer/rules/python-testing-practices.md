# Python Testing Practices

## Test Types
- Unit tests: Test individual functions/methods
- Integration tests: Test component interactions
- End-to-end tests: Test complete workflows
- API tests: Test HTTP endpoints

## Testing Standards
- Use pytest framework
- Fixtures for test data
- Parametrize tests for multiple scenarios
- Mock external dependencies
- Test both success and failure cases
- Test edge cases and boundary conditions

## Coverage
- Minimum 80% code coverage
- 100% coverage for critical paths
- Use coverage.py for measurement

## Test Organization
tests/
  unit/
  integration/
  e2e/
  conftest.py (shared fixtures)
