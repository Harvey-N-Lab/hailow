# Command: Health Check

## Purpose

Perform comprehensive health checks across infrastructure and applications.

## Usage

```
/health-check [environment] [options]
```

## Parameters

- `[environment]` - Target environment: `dev`, `staging`, `production` (default: all)
- `[options]` - Additional check options

## What This Command Does

1. **Infrastructure Health**
   - Check all compute instances
   - Verify load balancers healthy
   - Check database connectivity
   - Verify cache availability
   - Check storage systems

2. **Application Health**
   - HTTP health endpoint checks
   - Dependency availability
   - Background job processing
   - Queue depths

3. **Monitoring Health**
   - Metrics collection active
   - Logs flowing
   - Alerts functioning
   - Dashboards accessible

4. **Security Health**
   - Certificates valid
   - Secrets accessible
   - Security groups configured
   - IAM permissions working

## Examples

```bash
# Check all environments
/health-check

# Check production only
/health-check production

# Detailed health check
/health-check production --detailed

# Check specific component
/health-check production --component api
```

## Health Status

**Healthy** ✅
- All checks passing
- No alerts firing
- Metrics within normal range

**Degraded** ⚠️
- Some non-critical checks failing
- Minor alerts firing
- Performance slightly below baseline

**Unhealthy** ❌
- Critical checks failing
- Major alerts firing
- Service unavailable or severely degraded

## Output Example

```
Health Check Results - Production
=================================

Infrastructure:
  ✅ Compute: 10/10 instances healthy
  ✅ Load Balancer: Healthy, 250 req/sec
  ✅ Database: Healthy, 45ms average query time
  ✅ Cache: Healthy, 95% hit rate
  ✅ Storage: Healthy, 45% utilization

Applications:
  ✅ API: Healthy, 150ms p95 latency
  ✅ Worker: Healthy, processing 50 jobs/min
  ⚠️  Background Task: Degraded, queue backed up

Monitoring:
  ✅ Metrics: Collecting normally
  ✅ Logs: Flowing normally
  ✅ Alerts: 2 warnings, 0 critical

Security:
  ✅ Certificates: Valid, expire in 45 days
  ✅ Secrets: Accessible
  ✅ Security: No issues detected

Overall Status: HEALTHY ✅
```

## Related Commands

- `/deploy` - Deploy new version
- `/rollback` - Revert deployment
