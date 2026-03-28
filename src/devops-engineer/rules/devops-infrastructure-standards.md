# DevOps Infrastructure Standards

## Purpose

This document defines standards for infrastructure implementation, ensuring consistency, reliability, and maintainability across all infrastructure code.

## Infrastructure as Code Standards

### Terraform Standards

**File Organization:**
```
terraform/
├── modules/              # Reusable modules
│   ├── vpc/
│   ├── eks/
│   └── rds/
├── environments/         # Environment-specific
│   ├── dev/
│   ├── staging/
│   └── prod/
├── main.tf
├── variables.tf
├── outputs.tf
├── versions.tf
└── terraform.tfvars.example
```

**Naming Conventions:**
- Resources: `<environment>-<application>-<resource>` (e.g., `prod-api-server`)
- Variables: lowercase with underscores (e.g., `instance_type`)
- Modules: descriptive names (e.g., `vpc`, `database`, `load_balancer`)

**Required Practices:**
- Pin provider versions
- Use remote state (S3, Terraform Cloud)
- Enable state locking
- Tag all resources
- Use variables for environment differences
- Document modules with README
- Use data sources over hardcoded values

### CloudFormation Standards

**Stack Naming:** `<environment>-<application>-<component>`

**Required Practices:**
- Use parameters for reusability
- Output stack resources
- Use nested stacks for complex infrastructure
- Version control templates
- Use change sets for review

### Ansible Standards

**Playbook Structure:**
```
ansible/
├── inventories/
│   ├── dev/
│   ├── staging/
│   └── prod/
├── roles/
├── playbooks/
└── ansible.cfg
```

**Required Practices:**
- Use roles for reusability
- Variables in group_vars and host_vars
- Use vault for secrets
- Idempotent tasks
- Tag tasks for selective execution

## Resource Naming Standards

### AWS Resources

**EC2 Instances:** `<env>-<app>-<purpose>-<number>`
- Example: `prod-api-server-01`

**S3 Buckets:** `<org>-<env>-<app>-<purpose>`
- Example: `mycompany-prod-api-logs`

**IAM Roles:** `<app>-<env>-<service>-role`
- Example: `api-prod-ec2-role`

**Security Groups:** `<env>-<app>-<purpose>-sg`
- Example: `prod-api-web-sg`

**Load Balancers:** `<env>-<app>-<type>-lb`
- Example: `prod-api-app-lb`

### Kubernetes Resources

**Namespaces:** `<environment>` or `<team>-<env>`
- Examples: `production`, `data-team-dev`

**Deployments:** `<app-name>`
- Example: `user-service`

**Services:** `<app-name>-svc`
- Example: `user-service-svc`

**ConfigMaps:** `<app-name>-config`
- Example: `user-service-config`

**Secrets:** `<app-name>-secrets`
- Example: `user-service-secrets`

## Tagging Standards

### Required Tags (All Resources)

```hcl
tags = {
  Environment = "production"
  Application = "api-server"
  ManagedBy   = "terraform"
  Owner       = "platform-team"
  CostCenter  = "engineering"
  Project     = "api-v2"
}
```

### Tag Naming Convention

- PascalCase for tag keys
- Consistent values across resources
- Avoid special characters
- Keep values short but descriptive

## Network Architecture Standards

### VPC Design

**CIDR Blocks:**
- VPC: `/16` (e.g., `10.0.0.0/16`)
- Public subnets: `/24` per AZ
- Private subnets: `/20` per AZ
- Database subnets: `/24` per AZ

**Subnet Strategy:**
- Minimum 2 AZs (3 AZs for production)
- Public subnets for load balancers, NAT gateways
- Private subnets for application servers
- Isolated subnets for databases

**Example:**
```
VPC: 10.0.0.0/16
├── Public Subnet 1:  10.0.0.0/24   (AZ-a)
├── Public Subnet 2:  10.0.1.0/24   (AZ-b)
├── Private Subnet 1: 10.0.16.0/20  (AZ-a)
├── Private Subnet 2: 10.0.32.0/20  (AZ-b)
├── Data Subnet 1:    10.0.48.0/24  (AZ-a)
└── Data Subnet 2:    10.0.49.0/24  (AZ-b)
```

### Security Group Rules

**Principle:** Least privilege, explicit allow

**Best Practices:**
- No 0.0.0.0/0 ingress except for load balancers (ports 80/443)
- Reference security groups instead of CIDR blocks
- Document each rule with description
- Regular audit and cleanup

**Example:**
```hcl
resource "aws_security_group_rule" "app_from_lb" {
  type                     = "ingress"
  from_port                = 8080
  to_port                  = 8080
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.lb.id
  security_group_id        = aws_security_group.app.id
  description              = "Allow HTTP from load balancer"
}
```

## High Availability Standards

### Production Requirements

**Multi-AZ Deployment:**
- Minimum 2 AZs for all production services
- 3 AZs preferred for critical services
- Load balancers across all AZs
- Databases with multi-AZ failover

**Auto Scaling:**
- Minimum 2 instances at all times
- Target utilization 70% (scale before saturation)
- Cool-down periods to prevent thrashing
- Health checks for scale-in protection

**Load Balancing:**
- Application Load Balancer for HTTP/HTTPS
- Network Load Balancer for TCP/UDP
- Health checks with appropriate thresholds
- Connection draining enabled

## Disaster Recovery Standards

### Backup Requirements

**RTO/RPO Targets:**
- Critical systems: RTO < 1 hour, RPO < 5 minutes
- Important systems: RTO < 4 hours, RPO < 1 hour
- Non-critical: RTO < 24 hours, RPO < 24 hours

**Backup Strategy:**
- Automated backups daily minimum
- Retention: 7 days (production), 30 days (compliance)
- Cross-region replication for critical data
- Backup restoration tested quarterly

**Required Backups:**
- Database snapshots
- EBS volumes
- S3 versioning enabled
- Configuration backups
- Application state

### Disaster Recovery Testing

**Frequency:**
- DR drills quarterly
- Failover testing semi-annually
- Backup restoration monthly

**Documentation:**
- DR runbooks updated
- Contact lists current
- RTO/RPO measured and reported

## Security Standards

### IAM Best Practices

**Policies:**
- Least privilege principle
- Use managed policies where possible
- Custom policies version controlled
- Regular access reviews

**Roles:**
- Services use roles, not users
- Assume role for cross-account access
- External ID for third-party access
- Role names follow convention

**Users:**
- MFA enforced
- Password policy enforced
- Access keys rotated
- Unused IAM credentials removed

### Encryption Standards

**At Rest:**
- All databases encrypted (RDS, DynamoDB)
- All storage encrypted (S3, EBS, EFS)
- Use KMS customer-managed keys for sensitive data
- Key rotation enabled

**In Transit:**
- TLS 1.2 minimum
- TLS 1.3 preferred
- Valid certificates (no self-signed in production)
- HTTPS enforced (redirect HTTP to HTTPS)

### Secrets Management

**Requirements:**
- No secrets in code or config files
- Use secrets manager (AWS Secrets Manager, HashiCorp Vault)
- Automatic rotation where possible
- Access logged and audited

**Approved Methods:**
- AWS Secrets Manager
- HashiCorp Vault
- Kubernetes Secrets (with encryption at rest)
- Environment variables (from secrets manager)

**Prohibited:**
- Hardcoded secrets
- Secrets in git
- Secrets in container images
- Plaintext secrets in config files

## Cost Optimization Standards

### Resource Right-Sizing

**Regular Reviews:**
- Monthly compute utilization review
- Quarterly storage audit
- Eliminate orphaned resources
- Downsize over-provisioned resources

**Compute Optimization:**
- Use reserved instances for steady workloads
- Use spot instances for fault-tolerant workloads
- Enable auto-scaling
- Stop dev/test instances outside business hours

**Storage Optimization:**
- S3 lifecycle policies for infrequent access
- Delete old snapshots
- Use appropriate storage classes
- Compress and deduplicate where appropriate

### Cost Monitoring

**Required:**
- Cost allocation tags on all resources
- Budget alerts configured
- Monthly cost review
- Anomaly detection enabled

## Compliance Standards

### Audit Logging

**CloudTrail:**
- Enabled in all regions
- Log file validation enabled
- Logs stored in secure S3 bucket
- Retention minimum 1 year

**VPC Flow Logs:**
- Enabled for all VPCs
- Capture accepted and rejected traffic
- Logs analyzed for security threats

**Application Logs:**
- Centralized log aggregation
- Retention per compliance requirements
- Access controlled
- Tamper-proof storage

### Compliance Frameworks

**SOC 2:**
- Access controls documented
- Change management process
- Incident response plan
- Regular security audits

**HIPAA (if applicable):**
- PHI encrypted at rest and in transit
- Access logging enabled
- Business Associate Agreements in place
- Annual risk assessments

**GDPR (if applicable):**
- Data residency requirements met
- Right to deletion implemented
- Data processing agreements
- Privacy impact assessments

## Documentation Standards

### Required Documentation

**Infrastructure:**
- Architecture diagrams (current state)
- Network topology
- Data flow diagrams
- Disaster recovery procedures

**Operations:**
- Runbooks for common operations
- Troubleshooting guides
- Incident response procedures
- On-call escalation paths

**Security:**
- Security policies
- Access control matrix
- Incident response plan
- Compliance documentation

## Change Management Standards

### Infrastructure Changes

**Process:**
1. Change request submitted
2. Peer review required
3. Tested in non-production
4. Approved by team lead
5. Deployed during change window
6. Post-deployment validation
7. Documentation updated

**Change Windows:**
- Production: Outside business hours
- Staging: Anytime
- Development: Anytime

**Emergency Changes:**
- Critical security patches: Immediate
- Production outage: Immediate with communication
- Follow-up documentation required

## Monitoring Standards

### Required Metrics

**Infrastructure:**
- CPU utilization (all compute)
- Memory utilization
- Disk utilization
- Network throughput
- Load balancer request count
- Auto-scaling metrics

**Application:**
- Request rate
- Error rate
- Response time (p50, p95, p99)
- Queue depth
- Database connections
- Cache hit rate

**Business:**
- User signups
- Transactions completed
- Revenue metrics
- API usage

### Alerting Standards

**Alert Levels:**
- **Critical:** Page on-call immediately
- **High:** Notify team channel
- **Medium:** Log and review daily
- **Low:** Log and review weekly

**Alert Quality:**
- Actionable (clear what to do)
- Meaningful (not noise)
- Properly routed
- Documented response

## Summary Checklist

Before deploying infrastructure:

- [ ] Infrastructure as Code written and reviewed
- [ ] Resources properly named and tagged
- [ ] Multi-AZ for production
- [ ] Auto-scaling configured
- [ ] Monitoring and alerting configured
- [ ] Backups automated and tested
- [ ] Security groups follow least privilege
- [ ] Encryption enabled (at rest and transit)
- [ ] Secrets managed properly
- [ ] Documentation updated
- [ ] Cost tags applied
- [ ] Compliance requirements met
- [ ] Change management followed
- [ ] Tested in non-production
- [ ] Rollback procedure documented
