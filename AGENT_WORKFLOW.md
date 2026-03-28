# Agent Workflow and Collaboration Protocol

## Purpose

This document defines the recommended workflow for AI agents collaborating on software engineering tasks. It establishes a structured approach that ensures quality, thoroughness, and maintainability across all engineering domains.

## Core Workflow

All engineering tasks should follow this five-phase workflow:

```
┌──────────────┐
│  Researcher  │  Phase 1: Investigation and Discovery
└──────┬───────┘
       │
       ▼
┌──────────────┐
│   Planner    │  Phase 2: Strategy and Task Breakdown
└──────┬───────┘
       │
       ▼
┌──────────────┐
│  Architect   │  Phase 3: Design and Technical Decisions
└──────┬───────┘
       │
       ▼
┌──────────────┐
│ Implementer  │  Phase 4: Execution and Building
└──────┬───────┘
       │
       ▼
┌──────────────┐
│   Reviewer   │  Phase 5: Validation and Quality Assurance
└──────────────┘
```

## Phase 1: Researcher

**Role:** Investigate, gather context, and understand the problem space.

**Responsibilities:**
- Analyze the current codebase and infrastructure
- Identify existing patterns, conventions, and technologies
- Research requirements and constraints
- Document findings and insights
- Surface potential risks and dependencies
- Gather relevant documentation and examples

**Deliverables:**
- Research findings document
- Context summary
- Requirement analysis
- Risk assessment
- Relevant code/infrastructure snapshots

**Key Questions:**
- What exists currently?
- What are the constraints?
- What are the risks?
- What patterns should we follow?
- What dependencies exist?

## Phase 2: Planner

**Role:** Create a strategic plan and break down work into manageable tasks.

**Responsibilities:**
- Synthesize research findings into actionable insights
- Break down the problem into discrete tasks
- Sequence tasks in logical order
- Identify dependencies between tasks
- Estimate complexity and effort
- Define success criteria for each task
- Establish milestones

**Deliverables:**
- Detailed task breakdown
- Dependency graph
- Execution sequence
- Success criteria
- Risk mitigation plan
- Timeline estimates (if applicable)

**Key Questions:**
- What needs to be done?
- In what order should work proceed?
- What are the dependencies?
- How will we measure success?
- What could go wrong?

## Phase 3: Architect

**Role:** Design the technical solution and make architectural decisions.

**Responsibilities:**
- Design system architecture and components
- Make technology and framework choices
- Define interfaces, APIs, and contracts
- Plan data models and schemas
- Consider scalability, performance, and maintainability
- Document design decisions and tradeoffs
- Create technical specifications
- Ensure alignment with existing patterns

**Deliverables:**
- Architecture diagram(s)
- Technical specifications
- API/interface definitions
- Data models and schemas
- Design decision records
- Integration plan
- Technology choices with justifications

**Key Questions:**
- How should this be structured?
- What technologies/patterns fit best?
- What are the tradeoffs?
- How does this integrate with existing systems?
- What are the non-functional requirements?

## Phase 4: Implementer

**Role:** Execute the plan and build the solution.

**Responsibilities:**
- Write production-quality code
- Implement designs according to specifications
- Follow coding standards and best practices
- Write tests (unit, integration, e2e as appropriate)
- Handle error cases and edge conditions
- Document code and APIs
- Commit changes with clear messages
- Ensure security best practices

**Deliverables:**
- Working code implementation
- Tests (unit, integration, e2e)
- Documentation (inline and external)
- Configuration files
- Migration scripts (if applicable)
- Deployment artifacts (if applicable)

**Key Questions:**
- Does this match the design?
- Is this code maintainable?
- Are edge cases handled?
- Is this properly tested?
- Is this secure?

## Phase 5: Reviewer

**Role:** Validate quality, completeness, and correctness.

**Responsibilities:**
- Review code for correctness and quality
- Verify requirements are met
- Check adherence to standards and best practices
- Validate tests are comprehensive
- Review error handling and edge cases
- Check documentation completeness
- Verify security considerations
- Identify technical debt
- Suggest improvements
- Approve or request changes

**Deliverables:**
- Review feedback and findings
- List of issues (critical, major, minor)
- Improvement suggestions
- Approval or request for revisions
- Documentation of technical debt
- Final quality report

**Key Questions:**
- Does this meet the requirements?
- Is the code quality acceptable?
- Are tests sufficient?
- Are there security issues?
- Is documentation adequate?
- What could be improved?

## Collaboration Rules

### 1. Explicit Handoffs

Each phase must explicitly hand off to the next:
- Summarize what was accomplished
- Provide artifacts and deliverables
- Highlight important findings or decisions
- Call out risks or concerns
- State readiness for next phase

### 2. Iteration is Allowed

The workflow is not strictly linear:
- Implementer may return to Architect for clarification
- Reviewer may send back to Implementer for fixes
- New research may be needed at any phase
- Planner may revise plan based on discoveries

### 3. Documentation is Mandatory

Every phase must document:
- What was done
- Why it was done
- Key decisions and tradeoffs
- Risks and concerns
- Next steps

### 4. Context Preservation

Agents must preserve and pass context:
- Link to previous phase outputs
- Reference relevant files and systems
- Maintain a trail of decisions
- Keep context accessible

### 5. Domain Expertise

Each agent should:
- Apply domain-specific knowledge (DevOps, Backend, ML, etc.)
- Follow domain-specific rules and patterns
- Use domain-appropriate tools and practices
- Consider domain-specific concerns

## Workflow Variations

### Small Changes

For small, low-risk changes, phases can be condensed:
- Research + Planning can be combined
- Architecture may be implicit
- Review can be lightweight

### Emergency Fixes

For critical production issues:
- Researcher: Rapid diagnosis
- Planner: Quick mitigation strategy
- Architect: Minimal safe design
- Implementer: Fix with tests
- Reviewer: Fast-track review with follow-up items

### Exploratory Work

For spikes or experiments:
- Researcher: Define experiment goals
- Planner: Scope exploration
- Architect: Light design
- Implementer: Prototype
- Reviewer: Evaluate learnings and next steps

## Quality Gates

Each phase has quality gates that must be satisfied:

**Researcher Exit Criteria:**
- [ ] Problem is well understood
- [ ] Context is documented
- [ ] Constraints are identified
- [ ] Risks are surfaced

**Planner Exit Criteria:**
- [ ] Tasks are clearly defined
- [ ] Dependencies are identified
- [ ] Success criteria are established
- [ ] Plan is reviewable and actionable

**Architect Exit Criteria:**
- [ ] Design is complete and documented
- [ ] Tradeoffs are justified
- [ ] Integration points are defined
- [ ] Design is implementable

**Implementer Exit Criteria:**
- [ ] Code is complete and tested
- [ ] Standards are followed
- [ ] Documentation is written
- [ ] Implementation matches design

**Reviewer Exit Criteria:**
- [ ] Quality bar is met
- [ ] Requirements are satisfied
- [ ] Issues are documented
- [ ] Decision to approve or revise is clear

## Anti-Patterns

### ❌ Skip Research
Don't start implementing without understanding context.

### ❌ No Planning
Don't jump from research directly to coding without a plan.

### ❌ Architecture by Implementation
Don't figure out the design while coding. Design first.

### ❌ Undocumented Changes
Don't implement without documenting what and why.

### ❌ No Review
Don't skip review, even for small changes. Quality matters.

### ❌ Superficial Review
Don't rubber-stamp reviews. Actually validate quality.

### ❌ Lost Context
Don't discard context from previous phases. Build on it.

## Success Indicators

A well-executed workflow shows:
- ✅ Clear progression through phases
- ✅ Documented decisions and tradeoffs
- ✅ Tests accompanying implementation
- ✅ Meaningful review feedback
- ✅ Quality improving over iterations
- ✅ Knowledge retained and shared

## Domain-Specific Considerations

Each engineering domain (DevOps, Backend, ML, Data, JS/TS) has specific concerns:

- **DevOps**: Infrastructure, deployment, observability, security, disaster recovery
- **Backend**: API design, data models, business logic, performance, scalability
- **JS/TS**: UI/UX, component design, state management, browser compatibility
- **Data**: Pipeline design, data quality, schema evolution, performance, governance
- **ML**: Model design, training, evaluation, serving, monitoring, reproducibility

Agents should adapt the workflow to emphasize domain-relevant concerns at each phase.

## Continuous Improvement

The workflow itself should evolve:
- Retrospect on what worked and what didn't
- Adapt processes to project needs
- Document lessons learned
- Update guidelines based on experience
- Share improvements across domains

## Getting Started

1. Read this workflow document
2. Understand your domain's specific rules and patterns
3. Start with Researcher phase for new work
4. Follow the workflow through each phase
5. Document as you go
6. Review and iterate
7. Learn and improve

## Summary

This workflow ensures:
- **Thoughtfulness** through research
- **Planning** through deliberate task breakdown
- **Quality** through design-first approach
- **Correctness** through careful implementation
- **Excellence** through thorough review

Follow this workflow to deliver high-quality engineering work across all domains.
