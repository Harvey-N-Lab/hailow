# Command: Rollback

## Purpose

Quickly revert to the previous stable version to restore service.

## Usage

```
/rollback <environment> [options]
```

## Parameters

- `<environment>` - Target environment: `dev`, `staging`, `production`
- `[options]` - Additional rollback options

## What This Command Does

1. **Verification**
   - Confirm rollback intent
   - Identify previous version
   - Check rollback eligibility

2. **Rollback Execution**
   - Revert to previous version
   - Update configuration if needed
   - Monitor rollback progress

3. **Validation**
   - Verify previous version running
   - Check health and metrics
   - Confirm service restored

4. **Communication**
   - Notify team of rollback
   - Update incident if applicable
   - Log rollback reason

## Examples

```bash
# Rollback production immediately
/rollback production

# Rollback to specific version
/rollback production --to-version v1.2.0

# Dry-run rollback
/rollback production --dry-run
```

## Options

- `--to-version <tag>` - Rollback to specific version
- `--dry-run` - Show what would happen
- `--reason <text>` - Document reason for rollback

## When to Rollback

Immediate rollback if:
- Error rate > 1%
- Response time > 2x baseline  
- Critical functionality broken
- Security vulnerability discovered
- Data loss risk

Consider rollback if:
- Performance degradation > 50%
- Significant user complaints
- Unexpected behavior
- Failed validation tests

## Rollback Decision Matrix

| Issue Severity | Action | Timeline |
|----------------|--------|----------|
| Critical | Rollback immediately | < 2 min |
| High | Rollback if no quick fix | < 10 min |
| Medium | Investigate, may rollback | < 30 min |
| Low | Plan fix for next release | N/A |

## Post-Rollback Actions

1. **Verify Stability**
   - Monitor metrics for 15-30 minutes
   - Confirm error rates normal
   - Check user reports

2. **Root Cause Analysis**
   - Document what went wrong
   - Identify why it wasn't caught earlier
   - Plan preventive measures

3. **Fix and Redeploy**
   - Fix issue in code
   - Test thoroughly
   - Deploy fix

## Related Commands

- `/deploy` - Deploy new version
- `/health-check` - Check system health
