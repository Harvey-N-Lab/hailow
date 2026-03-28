# Collaboration Protocol

## Purpose

This document establishes how AI agents and human engineers collaborate effectively on software projects. It defines communication patterns, handoff procedures, and quality standards for collaborative work.

## Core Principles

### 1. Transparent Communication

**Principle:** All agents must communicate clearly about what they're doing and why.

**Practices:**
- State your current phase explicitly (researcher, planner, architect, implementer, reviewer)
- Summarize what you've learned or accomplished
- Call out uncertainties and assumptions
- Ask for clarification when needed
- Document decisions and reasoning

### 2. Context Preservation

**Principle:** Context must be maintained and passed between phases.

**Practices:**
- Link to relevant files and previous work
- Summarize key findings from previous phases
- Maintain a decision log
- Reference tickets/issues consistently
- Keep documentation updated

### 3. Iterative Refinement

**Principle:** Work improves through iteration and feedback.

**Practices:**
- Accept that first drafts need improvement
- Provide constructive feedback
- Be open to revising your work
- Don't take criticism personally
- Focus on outcome quality

### 4. Respect Expertise

**Principle:** Each agent and human has domain expertise to contribute.

**Practices:**
- Acknowledge good ideas
- Ask questions to learn
- Share knowledge generously
- Defer to domain experts
- Contribute your unique perspective

## Agent-to-Agent Handoffs

### Researcher → Planner Handoff

**Researcher Provides:**
- Research findings document
- Current state analysis
- Identified constraints and requirements
- Risk assessment
- Relevant code/config examples

**Handoff Message Template:**
```markdown
## Research Complete

### Summary
[What was investigated and key findings]

### Current State
[What exists now]

### Requirements
[What needs to be done]

### Constraints
[Technical, business, or resource constraints]

### Risks
[Potential issues to watch for]

### Recommendations
[Suggested approaches based on research]

Ready for planning phase.
```

**Planner Acknowledges:**
- Confirms receipt of research findings
- Clarifies any ambiguities
- Identifies missing information (if any)

### Planner → Architect Handoff

**Planner Provides:**
- Detailed task breakdown
- Task dependencies
- Success criteria
- Timeline/milestones
- Risk mitigation strategies

**Handoff Message Template:**
```markdown
## Plan Complete

### Task Breakdown
[Ordered list of tasks with dependencies]

### Success Criteria
[How we measure completion for each task]

### Dependencies
[What depends on what]

### Risks and Mitigation
[Potential issues and how to address them]

### Estimates
[Complexity/effort estimates if applicable]

Ready for architecture phase.
```

**Architect Acknowledges:**
- Confirms plan is clear and actionable
- Identifies any design questions
- Confirms approach alignment

### Architect → Implementer Handoff

**Architect Provides:**
- Architecture diagrams
- Technical specifications
- API/interface definitions
- Data models
- Design decisions with justifications
- Integration points

**Handoff Message Template:**
```markdown
## Architecture Complete

### Design Overview
[High-level architecture summary]

### Components
[Key components and responsibilities]

### Interfaces/APIs
[Definitions and contracts]

### Data Models
[Schemas and relationships]

### Technology Choices
[Selected technologies and why]

### Integration Points
[How this fits with existing systems]

### Non-Functional Considerations
[Performance, security, scalability notes]

Ready for implementation phase.
```

**Implementer Acknowledges:**
- Confirms design is clear and complete
- Asks clarification questions
- Identifies any ambiguities

### Implementer → Reviewer Handoff

**Implementer Provides:**
- Working code
- Tests (passing)
- Documentation
- Migration/deployment notes (if applicable)
- List of tradeoffs made

**Handoff Message Template:**
```markdown
## Implementation Complete

### What Was Built
[Summary of implementation]

### Key Files
[List of created/modified files]

### Tests
[Test coverage and approach]

### Documentation
[What was documented and where]

### Deployment Notes
[How to deploy or run this]

### Known Limitations
[Any technical debt or shortcuts taken]

### Tradeoffs
[Decisions made and alternatives considered]

Ready for review phase.
```

**Reviewer Acknowledges:**
- Confirms submission is complete
- Begins review process
- May ask for clarifications during review

### Reviewer Feedback Loop

**Reviewer Provides:**
- Detailed feedback on code quality
- List of issues (blocking, non-blocking)
- Suggestions for improvement
- Approval or change request

**Feedback Message Template:**
```markdown
## Review Complete

### Overall Assessment
[Summary of code quality]

### Blocking Issues
[Must be fixed before approval]
- Issue 1
- Issue 2

### Non-Blocking Issues
[Should be addressed but not blockers]
- Issue 1
- Issue 2

### Strengths
[What was done well]

### Recommendations
[Suggestions for future improvements]

### Decision
[APPROVED | CHANGES REQUESTED]
```

**If Changes Requested:**
- Implementer addresses feedback
- Provides updated code
- Documents what was changed
- Requests re-review

**If Approved:**
- Work is ready for deployment
- Update documentation
- Close related tickets

## Human-Agent Collaboration

### When Agents Should Ask Humans

**Critical Decisions:**
- Business logic choices
- Product requirements clarification
- Budget/resource allocation
- Major architectural changes
- Security/privacy decisions

**Ambiguities:**
- Conflicting requirements
- Unclear specifications
- Multiple valid approaches with significant tradeoffs
- Domain expertise beyond agent knowledge

**Blockers:**
- Missing access or permissions
- Unavailable resources
- External dependencies
- Policy or compliance questions

### How to Ask

**Good Question Format:**
```markdown
## Question for Human Review

### Context
[What we're working on]

### Question
[Clear, specific question]

### Options Considered
[Possible approaches with pros/cons]

### Recommendation
[Agent's suggested approach with reasoning]

### Impact
[What's blocked waiting for answer]
```

**Bad Question Format:**
```markdown
What should I do? [Too vague]
```

### When Humans Should Intervene

**Course Correction:**
- Agent is solving wrong problem
- Approach is suboptimal
- Missing critical context
- Violating standards or policies

**Knowledge Sharing:**
- Teaching agent about domain patterns
- Explaining business context
- Sharing institutional knowledge
- Providing examples

**Quality Gates:**
- Final approval for deployments
- Security review
- Compliance verification
- Stakeholder sign-off

## Conflict Resolution

### When Agents Disagree

**Process:**
1. State disagreement explicitly
2. Each agent explains reasoning
3. Evaluate against project goals
4. Seek objective criteria (standards, benchmarks, data)
5. Escalate to human if no consensus

**Example:**
```markdown
## Technical Disagreement

### Topic
[What we disagree about]

### Position A
[First agent's view and reasoning]

### Position B
[Second agent's view and reasoning]

### Evaluation Criteria
[How should we decide?]

### Proposed Resolution
[Suggested way forward or escalation to human]
```

### When Feedback Seems Wrong

**If you receive feedback you disagree with:**
1. Assume good intent
2. Ask for clarification
3. Explain your reasoning
4. Provide evidence if available
5. Be open to being wrong
6. Escalate if needed

**Don't:**
- Dismiss feedback immediately
- Get defensive
- Ignore the feedback
- Argue without reasoning

## Communication Standards

### Be Explicit

**Bad:** "This might have issues."
**Good:** "The authentication logic doesn't handle token expiration, which could leave users logged in indefinitely."

### Be Specific

**Bad:** "The code needs improvement."
**Good:** "The `processData()` function is 200 lines long. Consider extracting smaller functions for each processing step."

### Be Actionable

**Bad:** "This isn't good."
**Good:** "Replace the nested loops with a hash table lookup to improve performance from O(n²) to O(n)."

### Be Respectful

**Bad:** "This code is terrible."
**Good:** "This implementation works, but could be more maintainable if we applied the strategy pattern here."

### Document Decisions

Every significant decision should be documented:

```markdown
## Decision: [Title]

### Context
[What situation prompted this decision]

### Options Considered
1. Option A: [pros and cons]
2. Option B: [pros and cons]
3. Option C: [pros and cons]

### Decision
[What we chose]

### Reasoning
[Why we chose it]

### Tradeoffs
[What we're accepting/sacrificing]

### Revisit Conditions
[What would make us reconsider this decision]
```

## Work Distribution

### Parallelization

When multiple agents can work simultaneously:

**Scenario 1: Multiple Independent Components**
- Architect designs all components
- Divide components among implementers
- Each implementer handles their components
- Separate reviews for each

**Scenario 2: Research Multiple Options**
- Each researcher investigates one option
- Compile findings
- Planner synthesizes into unified plan

**Scenario 3: Large Codebase Changes**
- Architect designs overall structure
- Divide by module/package
- Coordinate interfaces between modules
- Integration testing after individual completion

### Prerequisites and Dependencies

**Rule:** Don't start work that depends on unfinished work.

**Exceptions:**
- Can proceed with mocks/stubs
- Can work on independent parts
- Can prepare adjacent work

**Communication:**
- Clearly state what you're blocked on
- Notify when blocking work is complete
- Update task status regularly

## Quality Standards

### Definition of Done

Work is complete when:
- [ ] Functionality is implemented and working
- [ ] Tests are written and passing
- [ ] Code is reviewed and approved
- [ ] Documentation is updated
- [ ] No security vulnerabilities
- [ ] Performance is acceptable
- [ ] Adheres to coding standards
- [ ] Integrated with existing codebase
- [ ] Deployment/rollout plan exists

### Review Checklist

Reviewers must verify:
- [ ] **Correctness:** Does it work as specified?
- [ ] **Tests:** Are tests comprehensive and passing?
- [ ] **Security:** Are there security issues?
- [ ] **Performance:** Are there performance problems?
- [ ] **Maintainability:** Is code readable and well-structured?
- [ ] **Documentation:** Is it adequately documented?
- [ ] **Standards:** Does it follow our conventions?
- [ ] **Edge Cases:** Are edge cases handled?

## Anti-Patterns to Avoid

### 1. ❌ Working in Isolation

**Problem:** Agent completes work without checking in or sharing progress.

**Solution:** Regular status updates, share work-in-progress, ask for early feedback.

### 2. ❌ Assumption Without Verification

**Problem:** Agent assumes requirements or design without confirming.

**Solution:** Ask clarifying questions, document assumptions explicitly, verify with human/previous agent.

### 3. ❌ Rubber-Stamp Approvals

**Problem:** Reviewer approves without actually reviewing.

**Solution:** Spend adequate time on review, provide meaningful feedback, ask questions if unclear.

### 4. ❌ Design by Implementation

**Problem:** Making up architecture while coding.

**Solution:** Architecture phase before implementation, document design first.

### 5. ❌ Silent Disagreement

**Problem:** Agent disagrees but doesn't voice concerns.

**Solution:** Speak up respectfully, explain concerns, seek resolution.

### 6. ❌ Lost Context

**Problem:** Each phase starts fresh without building on previous work.

**Solution:** Explicit handoffs, reference previous phases, maintain context thread.

### 7. ❌ Scope Creep

**Problem:** Expanding work beyond planned scope without discussion.

**Solution:** Stick to plan, discuss scope changes before implementing, track as separate work.

## Emergency Protocols

### Critical Production Issue

**Immediate Response:**
1. Researcher: Quick diagnosis and impact assessment
2. Planner: Immediate mitigation vs. proper fix
3. Architect: Fast-path design for fix
4. Implementer: Implement with tests
5. Reviewer: Expedited review focused on correctness and safety

**Post-Incident:**
- Blameless post-mortem
- Root cause analysis
- Preventive measures
- Documentation update

### Uncertainty or Blocked

**If Stuck:**
1. Clearly state the problem
2. Explain what's blocking you
3. List what you've tried
4. Ask specific questions
5. Request human intervention if needed

**Example:**
```markdown
## Blocked: Need Clarification

### Task
Implement user authentication

### Blocker
Unclear whether to use OAuth2 or JWT

### Analysis
- OAuth2: Better for third-party integration, more complex
- JWT: Simpler, sufficient for our use case

### Question
Do we need third-party authentication (Google, GitHub)?
Or is email/password sufficient?

### Impact
Blocking implementation of auth system.
Approximately 4 hours of work.

### Recommendation
If no third-party auth needed: Use JWT
If third-party auth needed: Use OAuth2

Awaiting human decision.
```

## Success Metrics

Effective collaboration shows:
- ✅ Clear communication at all phases
- ✅ Context preserved across handoffs
- ✅ Decisions documented with reasoning
- ✅ Issues raised and resolved quickly
- ✅ Quality improving through feedback
- ✅ Work completed without rework
- ✅ Knowledge shared across team

## Summary

Effective collaboration requires:
1. **Clear communication** - Say what you mean
2. **Context preservation** - Build on previous work
3. **Explicit handoffs** - Pass work cleanly between phases
4. **Constructive feedback** - Help each other improve
5. **Documented decisions** - Record the "why"
6. **Respect** - Value everyone's contributions
7. **Iteration** - Embrace continuous improvement

Follow this protocol to ensure smooth, productive collaboration across all project phases and team members.
