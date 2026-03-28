# Universal Engineering Practices

## Purpose

This document defines engineering practices that apply across all domains: DevOps, Backend, Frontend, Data, and Machine Learning. These are foundational principles that every engineer should follow regardless of their specialization.

## Code Quality Standards

### 1. Readability First

**Principle:** Code is read far more often than it is written.

**Practices:**
- Use clear, descriptive names for variables, functions, and classes
- Keep functions small and focused (single responsibility)
- Avoid clever code that requires deep understanding
- Add comments for "why" not "what" (the code shows what)
- Use consistent formatting and style

**Example - Bad:**
```python
def p(x): return x*x  # What does 'p' mean?
```

**Example - Good:**
```python
def calculate_area(side_length):
    """Calculate the area of a square."""
    return side_length * side_length
```

### 2. Fail Fast and Clearly

**Principle:** Errors should be detected as early as possible with clear messages.

**Practices:**
- Validate inputs at boundaries
- Use type hints/annotations where available
- Throw exceptions for invalid states
- Provide actionable error messages
- Log errors with context

**Example:**
```python
def divide(a: float, b: float) -> float:
    if b == 0:
        raise ValueError(f"Cannot divide {a} by zero. Divisor must be non-zero.")
    return a / b
```

### 3. Don't Repeat Yourself (DRY)

**Principle:** Avoid duplication of logic and data.

**Practices:**
- Extract repeated code into functions
- Use constants for magic numbers
- Create shared libraries for common logic
- But: Don't abstract prematurely (wait for 3+ occurrences)

### 4. Separation of Concerns

**Principle:** Different responsibilities should be in different modules.

**Practices:**
- Business logic separate from infrastructure
- Data access separate from business logic
- Presentation separate from logic
- Configuration separate from code

## Testing Standards

### 1. Tests Are Mandatory

**Principle:** All production code must have tests.

**Requirements:**
- Unit tests for business logic
- Integration tests for system interactions
- End-to-end tests for critical paths
- Test coverage target: 80% minimum

### 2. Tests Should Be Fast

**Principle:** Slow tests don't get run.

**Practices:**
- Unit tests should run in milliseconds
- Use mocks/stubs for external dependencies
- Parallelize tests when possible
- Reserve long-running tests for CI/CD only

### 3. Tests Should Be Independent

**Principle:** Tests should not depend on each other.

**Practices:**
- Each test sets up its own data
- Tests can run in any order
- Tests clean up after themselves
- No shared mutable state between tests

### 4. Tests Should Be Readable

**Principle:** Tests are documentation of how code should work.

**Practices:**
- Use descriptive test names
- Follow Arrange-Act-Assert pattern
- One assertion per test (generally)
- Clear failure messages

## Security Standards

### 1. Never Trust User Input

**Principle:** All external input is potentially malicious.

**Practices:**
- Validate all inputs
- Sanitize data before use
- Use parameterized queries (prevent SQL injection)
- Escape output (prevent XSS)
- Validate file uploads

### 2. Secrets Management

**Principle:** Secrets must never be in code or version control.

**Practices:**
- Use environment variables for secrets
- Use secret management services (AWS Secrets Manager, Vault, etc.)
- Rotate secrets regularly
- Never log secrets
- Use .gitignore for sensitive files

### 3. Principle of Least Privilege

**Principle:** Grant minimum necessary permissions.

**Practices:**
- Use role-based access control
- Limit service account permissions
- Regularly audit permissions
- Remove unused accounts and permissions

### 4. Security by Default

**Principle:** Secure configurations should be the default.

**Practices:**
- HTTPS by default
- Authentication required by default
- Encryption at rest and in transit
- Secure headers configured
- Regular security updates

## Performance Standards

### 1. Profile Before Optimizing

**Principle:** Don't optimize without data.

**Practices:**
- Measure actual performance
- Identify real bottlenecks
- Optimize the slow parts, not guesses
- Re-measure after optimization
- Document performance requirements

### 2. Design for Scale

**Principle:** Consider future growth in design.

**Practices:**
- Avoid N+1 queries
- Use pagination for large datasets
- Consider caching strategies
- Plan for horizontal scaling
- Monitor resource usage

### 3. Fail Gracefully

**Principle:** System should degrade gracefully under load.

**Practices:**
- Implement timeouts
- Use circuit breakers
- Provide fallback mechanisms
- Rate limiting where appropriate
- Queue work instead of rejecting

## Documentation Standards

### 1. Document the "Why"

**Principle:** Code shows how, documentation explains why.

**What to Document:**
- Architecture decisions and tradeoffs
- Non-obvious business rules
- Setup and installation steps
- API contracts and examples
- Security considerations

### 2. Keep Documentation Close to Code

**Principle:** Documentation should be easy to find and update.

**Practices:**
- README in every directory
- Inline documentation for complex logic
- API documentation from code (Swagger, JSDoc, docstrings)
- Architecture Decision Records (ADRs) in repo

### 3. Documentation Must Be Maintained

**Principle:** Outdated documentation is worse than no documentation.

**Practices:**
- Update docs with code changes
- Review docs in code reviews
- Mark deprecated features clearly
- Remove outdated documentation

## Version Control Standards

### 1. Commit Messages Matter

**Principle:** Good commit messages enable future understanding.

**Format:**
```
<type>: <short summary> (max 50 chars)

<detailed description if needed>

<reference to issue/ticket if applicable>
```

**Types:** feat, fix, docs, refactor, test, chore, perf

**Example:**
```
feat: add user authentication with JWT

Implements JWT-based authentication for API endpoints.
Includes token refresh mechanism and role-based access.

Closes #123
```

### 2. Branch Strategy

**Principle:** Structured branching enables collaboration.

**Practices:**
- `main` branch is always deployable
- Feature branches from `main`
- Name branches descriptively: `feature/user-auth`, `fix/login-bug`
- Delete branches after merge
- Pull request required for merging

### 3. Small, Focused Commits

**Principle:** Small commits are easier to review and revert.

**Practices:**
- One logical change per commit
- Commit working code (tests pass)
- Separate refactoring from feature changes
- Avoid mixing unrelated changes

## Code Review Standards

### 1. All Code Must Be Reviewed

**Principle:** Every change requires a second pair of eyes.

**Requirements:**
- At least one approval before merge
- Reviewer understands the change
- Tests pass in CI
- No obvious issues

### 2. Reviews Should Be Timely

**Principle:** Fast feedback enables fast iteration.

**Targets:**
- First response within 4 hours
- Complete review within 24 hours
- Small PRs reviewed first

### 3. Reviews Should Be Constructive

**Principle:** Reviews improve code and share knowledge.

**Practices:**
- Ask questions, don't demand changes
- Explain reasoning for suggestions
- Praise good solutions
- Distinguish blocking vs. non-blocking feedback
- Focus on correctness, maintainability, and readability

### 4. Reviewers Check For

- [ ] Correctness: Does it work?
- [ ] Tests: Are they comprehensive?
- [ ] Readability: Is it clear?
- [ ] Performance: Any obvious issues?
- [ ] Security: Any vulnerabilities?
- [ ] Documentation: Is it adequate?

## Error Handling Standards

### 1. Handle Errors, Don't Ignore Them

**Principle:** Silent failures are dangerous.

**Practices:**
- Never use empty catch blocks
- Log errors with context
- Return errors explicitly (Go, Rust) or throw exceptions
- Provide recovery mechanisms where appropriate

### 2. Error Messages Should Be Actionable

**Principle:** Users should know what to do next.

**Bad:** `Error: invalid input`
**Good:** `Error: Email address 'user@example' is invalid. Please provide a valid email in format 'user@domain.com'`

### 3. Fail Appropriately

**Principle:** Match failure response to severity.

**Practices:**
- Fatal errors: Crash fast with clear message
- Expected errors: Handle gracefully and inform user
- Transient errors: Retry with backoff
- Log all errors for debugging

## Monitoring and Observability

### 1. Log Meaningfully

**Principle:** Logs should help debug issues.

**Practices:**
- Log at appropriate levels (DEBUG, INFO, WARN, ERROR)
- Include context (user ID, request ID, etc.)
- Structured logging (JSON) for parsing
- Don't log sensitive data
- Log both successes and failures

### 2. Monitor What Matters

**Principle:** Track metrics that indicate health.

**Key Metrics:**
- Request rate and latency
- Error rate
- Resource utilization (CPU, memory, disk)
- Business metrics (signups, transactions, etc.)

### 3. Set Up Alerts

**Principle:** Know about problems before users do.

**Practices:**
- Alert on symptoms, not causes
- Set meaningful thresholds
- Avoid alert fatigue
- Make alerts actionable
- Document response procedures

## Dependency Management

### 1. Keep Dependencies Up to Date

**Principle:** Outdated dependencies have security vulnerabilities.

**Practices:**
- Regular dependency updates (monthly)
- Automated vulnerability scanning
- Pin dependency versions
- Test after updates

### 2. Minimize Dependencies

**Principle:** Every dependency is a potential liability.

**Practices:**
- Evaluate necessity before adding
- Consider maintenance status
- Prefer well-maintained libraries
- Remove unused dependencies

### 3. Understand Your Dependencies

**Principle:** You're responsible for your dependencies' behavior.

**Practices:**
- Read documentation
- Check license compatibility
- Review security advisories
- Understand what it does

## Backward Compatibility

### 1. Don't Break Existing Users

**Principle:** Breaking changes require major version bump.

**Practices:**
- Deprecate before removing
- Provide migration guides
- Support old versions during transition
- Communicate changes clearly

### 2. Versioning Strategy

**Principle:** Semantic versioning communicates compatibility.

**Format:** MAJOR.MINOR.PATCH
- MAJOR: Breaking changes
- MINOR: New features (backward compatible)
- PATCH: Bug fixes (backward compatible)

## Continuous Improvement

### 1. Retrospect Regularly

**Principle:** Learn from both successes and failures.

**Practices:**
- Post-mortems for incidents (blameless)
- Regular team retrospectives
- Document lessons learned
- Share knowledge across teams

### 2. Refactor Continuously

**Principle:** Code quality erodes without maintenance.

**Practices:**
- Boy Scout Rule: Leave code better than you found it
- Regular refactoring sessions
- Address technical debt
- Don't let perfect be enemy of good

### 3. Stay Current

**Principle:** Technology evolves, so should we.

**Practices:**
- Read documentation and blogs
- Experiment with new tools
- Share learnings with team
- Attend conferences/meetups
- Contribute to open source

## Summary Checklist

Before merging any code, verify:

- [ ] Code is readable and well-structured
- [ ] Tests are comprehensive and passing
- [ ] Security best practices followed
- [ ] Performance is acceptable
- [ ] Documentation is updated
- [ ] Error handling is robust
- [ ] Logging is appropriate
- [ ] Code is reviewed and approved
- [ ] Secrets are not in code
- [ ] Dependencies are justified and secure

These practices ensure consistent, high-quality engineering across all domains and teams.
