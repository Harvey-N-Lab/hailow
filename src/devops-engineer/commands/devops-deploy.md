# Command: Deploy

## Purpose

Execute a deployment to a specified environment with appropriate validation and monitoring.

## Usage

```
/deploy <environment> [options]
```

## Parameters

- `<environment>` - Target environment: `dev`, `staging`, `production`
- `[options]` - Additional deployment options

## What This Command Does

1. **Pre-Deployment Checks**
   - Verify all tests passed
   - Check deployment approval (for production)
   - Validate configuration
   - Check for active incidents

2. **Backup Phase**
   - Create database backup (if applicable)
   - Tag current version
   - Record current state

3. **Deployment Phase**
   - Deploy new version using configured strategy
   - Monitor rollout progress
   - Perform health checks
   - Run smoke tests

4. **Validation Phase**
   - Verify service health
   - Check error rates
   - Monitor performance metrics
   - Validate business metrics

5. **Post-Deployment**
   - Update deployment tracking
   - Notify team of completion
   - Monitor for issues
   - Document results

## Examples

```bash
# Deploy to development
/deploy dev

# Deploy to staging
/deploy staging

# Deploy to production (requires approval)
/deploy production

# Deploy with specific version
/deploy production --version v1.2.3

# Deploy with rollback on error
/deploy production --auto-rollback

# Dry-run deployment
/deploy production --dry-run
```

## Options

- `--version <tag>` - Deploy specific version
- `--dry-run` - Show what would happen
- `--auto-rollback` - Automatically rollback on errors
- `--skip-tests` - Skip pre-deployment tests (not recommended)
- `--no-backup` - Skip backup phase (not recommended)

## Environment-Specific Behavior

### Development
- Fast deployment
- Minimal validation
- No approval required
- Immediate rollout

### Staging
- Standard deployment
- Full validation
- No approval required
- Gradual rollout available

### Production
- Safe deployment
- Comprehensive validation
- Approval required
- Canary or blue-green deployment
- Extended monitoring

## Rollback

If deployment fails or issues detected:

```bash
/rollback <environment>
```

This reverts to the previous version immediately.

## Integration

This command integrates with:
- CI/CD pipelines
- Version control
- Monitoring systems
- Incident management
- Notification channels

## Safety Features

- Pre-deployment validation
- Health check verification
- Automatic rollback on failure
- Rate limiting for production
- Approval gates
- Audit logging

## Monitoring Dashboard

After deployment, view:
- Deployment progress
- Health check status
- Error rate trends
- Performance metrics
- Rollout percentage (for canary)

## Troubleshooting

**Deployment stuck:**
- Check pod status: `kubectl get pods`
- View logs: `kubectl logs <pod>`
- Check events: `kubectl get events`

**Health checks failing:**
- Verify service configuration
- Check dependencies availability
- Review application logs

**Performance degradation:**
- Compare metrics to baseline
- Check resource utilization
- Review recent changes

## Best Practices

- Always deploy to dev first
- Test in staging before production
- Deploy during low-traffic windows
- Have on-call engineer available
- Monitor closely for first 30 minutes
- Document any issues encountered

## Related Commands

- `/rollback` - Revert to previous version
- `/health-check` - Check system health
- `/validate` - Run validation tests
