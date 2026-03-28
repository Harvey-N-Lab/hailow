# DevOps Deployment Guidelines

## Purpose

Safe, reliable deployments are critical to system uptime. This document defines deployment strategies, procedures, and guardrails.

## Deployment Strategies

### Rolling Deployment

**When to Use:** Standard deployments with minimal downtime tolerance

**How it Works:**
1. Update instances one at a time or in small batches
2. Wait for health checks to pass
3. Continue to next batch
4. Rollback if issues detected

**Pros:**
- Simple to implement
- No additional infrastructure
- Gradual rollout

**Cons:**
- Multiple versions running simultaneously
- Longer deployment time
- Harder to rollback

**Kubernetes Example:**
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
spec:
  replicas: 10
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 2
      maxUnavailable: 1
  template:
    spec:
      containers:
      - name: app
        image: myapp:v2
```

### Blue-Green Deployment

**When to Use:** Zero-downtime requirement, need quick rollback

**How it Works:**
1. Deploy new version to identical "green" environment
2. Test green environment thoroughly
3. Switch traffic from "blue" to "green"
4. Keep blue environment for quick rollback

**Pros:**
- Zero downtime
- Instant rollback
- Test in production-like environment

**Cons:**
- Double infrastructure cost
- Database migration complexity
- Requires load balancer switching

**AWS Example:**
```hcl
# Two target groups
resource "aws_lb_target_group" "blue" {
  name     = "blue-tg"
  port     = 8080
  protocol = "HTTP"
  vpc_id   = aws_vpc.main.id
}

resource "aws_lb_target_group" "green" {
  name     = "green-tg"
  port     = 8080
  protocol = "HTTP"
  vpc_id   = aws_vpc.main.id
}

# Switch traffic by updating listener
resource "aws_lb_listener_rule" "main" {
  listener_arn = aws_lb_listener.main.arn
  
  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.green.arn  # Switch to green
  }
  
  condition {
    path_pattern {
      values = ["/*"]
    }
  }
}
```

### Canary Deployment

**When to Use:** High-risk changes, want to limit blast radius

**How it Works:**
1. Deploy new version to small percentage of traffic (e.g., 5%)
2. Monitor metrics closely
3. Gradually increase percentage if healthy
4. Rollback if issues detected

**Pros:**
- Limits blast radius
- Real production traffic testing
- Data-driven promotion

**Cons:**
- Complex to implement
- Longer deployment time
- Requires sophisticated monitoring

**Flagger Example (Istio/Linkerd):**
```yaml
apiVersion: flagger.app/v1beta1
kind: Canary
metadata:
  name: myapp
spec:
  targetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: myapp
  service:
    port: 8080
  analysis:
    interval: 1m
    threshold: 5
    maxWeight: 50
    stepWeight: 10
    metrics:
    - name: request-success-rate
      thresholdRange:
        min: 99
    - name: request-duration
      thresholdRange:
        max: 500
```

### Feature Flags

**When to Use:** Deploy code but control feature activation

**How it Works:**
1. Deploy code with features disabled
2. Enable features for specific users/percentage
3. Monitor impact
4. Gradually roll out or instantly disable

**Pros:**
- Decouple deployment from release
- A/B testing capability
- Instant kill-switch

**Cons:**
- Code complexity
- Technical debt (old flags)
- Testing complexity

## Pre-Deployment Checklist

Before deploying to production:

### Code Quality
- [ ] All tests passing (unit, integration, e2e)
- [ ] Code review approved
- [ ] No known critical bugs
- [ ] Performance tested
- [ ] Security scan passed

### Infrastructure Ready
- [ ] Resources provisioned and healthy
- [ ] Configuration updated
- [ ] Database migrations tested
- [ ] Dependencies available
- [ ] SSL certificates valid

### Monitoring Ready
- [ ] Dashboards created
- [ ] Alerts configured
- [ ] Logging working
- [ ] Tracing enabled
- [ ] Baseline metrics established

### Rollback Ready
- [ ] Rollback procedure documented
- [ ] Previous version available
- [ ] Database rollback plan (if needed)
- [ ] Rollback tested in staging

### Communication
- [ ] Deployment scheduled
- [ ] Stakeholders notified
- [ ] On-call engineer available
- [ ] Incident channels ready

## Deployment Procedures

### Standard Deployment Procedure

**1. Pre-Deployment (T-30 min)**
```bash
# Verify staging
kubectl get pods -n staging
curl https://staging.example.com/health

# Backup database (if applicable)
pg_dump mydb > backup_$(date +%Y%m%d_%H%M%S).sql

# Tag release
git tag v1.2.3
git push origin v1.2.3

# Notify team
echo "Deploying v1.2.3 to production"
```

**2. Deployment (T=0)**
```bash
# Deploy new version
kubectl set image deployment/myapp myapp=myapp:v1.2.3

# Watch rollout
kubectl rollout status deployment/myapp

# Verify pods healthy
kubectl get pods -l app=myapp
```

**3. Validation (T+5 min)**
```bash
# Health check
curl https://api.example.com/health

# Smoke tests
./scripts/smoke-tests.sh production

# Check metrics
# - Request rate normal
# - Error rate < 0.1%
# - Response time < baseline + 10%
# - No alerts firing
```

**4. Post-Deployment (T+15 min)**
```bash
# Monitor for issues
# Watch for 15-30 minutes

# If all good:
echo "Deployment successful" | notify-team

# If issues:
kubectl rollout undo deployment/myapp
echo "Rolled back due to: <reason>" | notify-team
```

### Emergency Rollback Procedure

**Immediate Actions:**
```bash
# Kubernetes
kubectl rollout undo deployment/myapp

# Verify previous version running
kubectl rollout status deployment/myapp
kubectl get pods -l app=myapp

# Check health
curl https://api.example.com/health

# Verify metrics returning to normal
```

**Communicate:**
- Notify team immediately
- Update status page if customer-facing
- Document incident for post-mortem

### Database Migration Deployment

**Strategy:** Backward-compatible migrations

**Process:**
1. **Deploy migration that's compatible with old code**
   - Add new column (nullable)
   - Don't remove old column yet
   
2. **Deploy application code that writes to both**
   - Use new column
   - Keep old column for safety
   
3. **Backfill data** (if needed)
   - Populate new column with data
   
4. **Deploy code that only uses new column**
   - Stop writing to old column
   
5. **Remove old column** (separate deployment)
   - After verification period

**Example Migration:**
```sql
-- Step 1: Add new column
ALTER TABLE users ADD COLUMN email_verified BOOLEAN DEFAULT FALSE;

-- Step 2: Backfill (in batches)
UPDATE users 
SET email_verified = (verified_at IS NOT NULL)
WHERE id BETWEEN 1 AND 10000;

-- Step 3: Make NOT NULL (after all rows updated)
ALTER TABLE users ALTER COLUMN email_verified SET NOT NULL;

-- Step 4: Remove old column (later, after app deployed)
-- ALTER TABLE users DROP COLUMN verified_at;
```

## Deployment Windows

### Production
- **Preferred:** Monday-Thursday, 10am-4pm
- **Avoid:** Fridays, weekends, holidays
- **Emergency:** Anytime (with incident process)

### Staging
- **Anytime:** No restrictions

### Development
- **Anytime:** No restrictions

### Exception
High-priority security patches deployed immediately regardless of window.

## Deployment Approval Requirements

### Production
- Code review approved
- All tests passing
- Staging validation complete
- Team lead approval (for significant changes)
- Security review (for infrastructure/security changes)

### Staging
- Code review approved
- All tests passing

### Development
- No approval required

## Monitoring During Deployment

### Key Metrics to Watch

**Immediate (0-5 minutes):**
- Deployment success/failure
- Pod/instance health
- Error count/rate
- Response time (p50, p95, p99)

**Short-term (5-30 minutes):**
- Request rate
- Error rate
- Response time trends
- Resource utilization
- Database performance

**Medium-term (30 minutes - 2 hours):**
- Business metrics
- User-reported issues
- Background job success rate
- Downstream service impact

### Alert Configuration

**Critical Alerts:**
- Error rate > 1%
- Response time > 2x baseline
- Deployment failure
- Health checks failing

**Warning Alerts:**
- Error rate > 0.5%
- Response time > 1.5x baseline
- Resource utilization > 80%

## Rollback Decision Matrix

| Condition | Action | Timeline |
|-----------|--------|----------|
| Error rate > 1% | Rollback immediately | < 5 min |
| Response time > 2x | Rollback immediately | < 5 min |
| Critical bug discovered | Rollback immediately | < 5 min |
| Minor bug, workaround exists | Hot-fix or next release | 1-4 hours |
| Performance degradation < 50% | Investigate, may rollback | 15-30 min |
| No issues detected | Proceed | Monitor 30-60 min |

## Deployment Documentation

### Required Documentation

**Deployment Plan:**
- What is being deployed
- Why it's being deployed
- When it's being deployed
- Who is responsible
- Rollback procedure
- Risk assessment

**Runbook:**
- Step-by-step deployment commands
- Validation checks
- Troubleshooting steps
- Rollback commands
- Points of contact

**Post-Deployment Report:**
- What was deployed
- Deployment duration
- Issues encountered
- Metrics comparison (before/after)
- Lessons learned

## Common Deployment Issues

### Issue: Pods CrashLooping

**Diagnosis:**
```bash
kubectl logs <pod-name>
kubectl describe pod <pod-name>
```

**Common Causes:**
- Configuration error
- Missing environment variable
- Database connection failure
- Insufficient resources

**Resolution:**
- Fix configuration
- Rollback if quick fix not available

### Issue: Deployment Stuck

**Diagnosis:**
```bash
kubectl rollout status deployment/myapp
kubectl get pods -l app=myapp
kubectl describe deployment myapp
```

**Common Causes:**
- Image pull failure
- Resource quota exceeded
- Pod security policy rejection

**Resolution:**
- Check image availability
- Increase quotas if needed
- Fix security policy issues

### Issue: Performance Degradation

**Diagnosis:**
- Check response time metrics
- Check resource utilization
- Check database slow queries
- Compare to baseline

**Resolution:**
- Rollback if severe
- Optimize if gradual
- Scale up if resource constrained

## Deployment Metrics

### Track These Metrics

**Deployment Frequency:**
- How often we deploy
- Goal: Daily or multiple times per day

**Lead Time:**
- Time from commit to production
- Goal: < 1 hour for small changes

**Change Failure Rate:**
- Percentage of deployments causing issues
- Goal: < 15%

**Mean Time to Recovery (MTTR):**
- Time to recover from failed deployment
- Goal: < 1 hour

## Continuous Deployment

### Requirements for CD

- Comprehensive automated tests
- Robust monitoring and alerting
- Automated rollback capability
- Feature flags for risk mitigation
- High team confidence in process

### CD Pipeline Example

```yaml
on:
  push:
    branches: [main]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Run tests
        run: npm test
      
      - name: Build and push image
        run: |
          docker build -t myapp:${{ github.sha }} .
          docker push myapp:${{ github.sha }}
      
      - name: Deploy to production
        run: |
          kubectl set image deployment/myapp myapp=myapp:${{ github.sha }}
          kubectl rollout status deployment/myapp
      
      - name: Run smoke tests
        run: ./scripts/smoke-tests.sh
      
      - name: Monitor for 5 minutes
        run: ./scripts/monitor-deployment.sh
```

## Summary

Successful deployments require:
1. **Preparation:** Testing, review, validation
2. **Strategy:** Choose appropriate deployment method
3. **Monitoring:** Watch metrics closely
4. **Rollback:** Be ready to rollback quickly
5. **Communication:** Keep stakeholders informed
6. **Documentation:** Record what happened

Remember: **If in doubt, rollback.** It's better to rollback and investigate than to leave a broken deployment in production.
