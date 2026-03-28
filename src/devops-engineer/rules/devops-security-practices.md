# DevOps Security Practices

## Purpose

Security is everyone's responsibility. This document defines security practices that must be followed when building, deploying, and operating infrastructure.

## Core Security Principles

### 1. Defense in Depth
Multiple layers of security controls to protect systems.

### 2. Least Privilege
Grant minimum permissions necessary for tasks.

### 3. Zero Trust
Never trust, always verify - authenticate and authorize every request.

### 4. Fail Securely
When failures occur, default to secure state.

### 5. Security by Default
Secure configurations should be the default, not opt-in.

## Secrets Management

### Never Store Secrets In

- ❌ Git repositories
- ❌ Container images
- ❌ Configuration files (plaintext)
- ❌ Environment variables (hardcoded)
- ❌ CI/CD logs
- ❌ Application logs
- ❌ Documentation
- ❌ Code comments

### Approved Secrets Storage

- ✅ AWS Secrets Manager
- ✅ HashiCorp Vault
- ✅ Azure Key Vault
- ✅ GCP Secret Manager
- ✅ Kubernetes Secrets (with encryption at rest)

### Secrets Rotation

**Requirements:**
- Rotate secrets every 90 days minimum
- Automated rotation preferred
- Test rotation process regularly
- Document rotation procedures
- Monitor for rotation failures

## IAM and Access Control

### AWS IAM Best Practices

**For Users:**
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ec2:DescribeInstances",
        "ec2:DescribeImages"
      ],
      "Resource": "*"
    }
  ]
}
```

**Avoid:**
```json
{
  "Effect": "Allow",
  "Action": "*",
  "Resource": "*"
}
```

**Requirements:**
- MFA enforced for all human users
- Use roles for services, not users
- Regular access reviews (quarterly)
- Inactive credentials removed after 90 days
- Use IAM policy conditions to restrict access

### Kubernetes RBAC

**Least Privilege Example:**
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: production
  name: pod-reader
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "watch", "list"]
```

**Avoid Cluster-Admin:**
- Don't use cluster-admin except for cluster operations
- Create specific roles for specific needs
- Use namespaced roles, not cluster roles when possible
- Regularly audit RBAC permissions

## Network Security

### Security Group Rules

**Good Practice:**
```hcl
resource "aws_security_group_rule" "app_from_lb" {
  type                     = "ingress"
  from_port                = 8080
  to_port                  = 8080
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.lb.id
  security_group_id        = aws_security_group.app.id
  description              = "App from load balancer"
}
```

**Bad Practice:**
```hcl
resource "aws_security_group_rule" "open_to_world" {
  type        = "ingress"
  from_port   = 0
  to_port     = 65535
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]  # ❌ Never do this
}
```

### Network Segmentation

**Best Practices:**
- Public subnets for load balancers only
- Private subnets for application servers
- Isolated subnets for databases
- No direct internet access to databases
- NAT Gateway for outbound from private subnets
- VPC Flow Logs enabled

### Kubernetes Network Policies

**Restrict Traffic:**
```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: api-network-policy
  namespace: production
spec:
  podSelector:
    matchLabels:
      app: api
  policyTypes:
  - Ingress
  - Egress
  ingress:
  - from:
    - podSelector:
        matchLabels:
          app: frontend
    ports:
    - protocol: TCP
      port: 8080
  egress:
  - to:
    - podSelector:
        matchLabels:
          app: database
    ports:
    - protocol: TCP
      port: 5432
```

## Encryption

### At Rest

**Required Encryption:**
- All database storage (RDS, DynamoDB, MongoDB)
- All object storage (S3 buckets)
- All block storage (EBS volumes, EFS)
- Secrets storage
- Backup storage

**KMS Key Management:**
- Use customer-managed keys for sensitive data
- Enable automatic key rotation
- Audit key usage with CloudTrail
- Principle of least privilege for key access

### In Transit

**Requirements:**
- TLS 1.2 minimum (TLS 1.3 preferred)
- Valid certificates (no self-signed in production)
- HTTPS enforced (HTTP redirects to HTTPS)
- HSTS headers enabled
- Certificate expiration monitoring

**Load Balancer TLS:**
```hcl
resource "aws_lb_listener" "https" {
  load_balancer_arn = aws_lb.main.arn
  port              = "443"
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-TLS-1-2-2017-01"
  certificate_arn   = aws_acm_certificate.cert.arn

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.main.arn
  }
}
```

## Container Security

### Image Security

**Base Images:**
- Use official images from trusted registries
- Use minimal base images (alpine, distroless)
- Scan images for vulnerabilities (Trivy, Snyk, Clair)
- Pin base image versions
- Regularly update base images

**Dockerfile Best Practices:**
```dockerfile
# Good
FROM node:18-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production
COPY . .
RUN npm run build

FROM node:18-alpine
RUN addgroup -g 1001 -S nodejs && adduser -S nodejs -u 1001
WORKDIR /app
COPY --from=builder --chown=nodejs:nodejs /app/dist ./dist
COPY --from=builder --chown=nodejs:nodejs /app/node_modules ./node_modules
USER nodejs
EXPOSE 3000
CMD ["node", "dist/index.js"]
```

**Avoid:**
```dockerfile
# Bad
FROM ubuntu:latest  # ❌ Don't use 'latest' or full OS
WORKDIR /app
COPY . .
RUN apt-get update && apt-get install -y nodejs npm  # ❌ Unnecessary packages
USER root  # ❌ Never run as root
CMD ["node", "index.js"]
```

### Runtime Security

**Pod Security:**
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: secure-pod
spec:
  securityContext:
    runAsNonRoot: true
    runAsUser: 1000
    fsGroup: 1000
    seccompProfile:
      type: RuntimeDefault
  containers:
  - name: app
    image: myapp:1.0.0
    securityContext:
      allowPrivilegeEscalation: false
      capabilities:
        drop:
        - ALL
      readOnlyRootFilesystem: true
    volumeMounts:
    - name: tmp
      mountPath: /tmp
  volumes:
  - name: tmp
    emptyDir: {}
```

## CI/CD Security

### Pipeline Security

**Requirements:**
- Secrets stored in secret manager, not pipeline config
- Least privilege for CI/CD service accounts
- Audit logs enabled for all pipeline actions
- Require approval for production deployments
- Scan code for vulnerabilities before deployment
- Sign and verify artifacts

**GitHub Actions Security:**
```yaml
name: Secure Pipeline

on:
  push:
    branches: [main]

# Restrict permissions
permissions:
  contents: read
  id-token: write

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      # Use OIDC instead of long-lived credentials
      - name: Configure AWS Credentials
        uses: aws-actions/configure-aws-credentials@v2
        with:
          role-to-assume: arn:aws:iam::123456789012:role/GitHubActionsRole
          aws-region: us-east-1
      
      # Security scanning
      - name: Run security scan
        run: |
          trivy fs --security-checks vuln,config .
          
      # Build and push
      - name: Build and push
        run: |
          docker build -t myapp:${{ github.sha }} .
          docker push myapp:${{ github.sha }}
```

### Artifact Security

**Requirements:**
- Sign container images
- Verify signatures before deployment
- Scan artifacts for vulnerabilities
- Store artifacts in secure registries
- Implement artifact retention policies

## Vulnerability Management

### Scanning Requirements

**Infrastructure:**
- Scan IaC before apply (tfsec, checkov)
- Scan for misconfigurations
- Scan for compliance violations

**Containers:**
- Scan images on build
- Scan images in registry (continuous)
- Fail builds on critical vulnerabilities
- Track vulnerability remediation

**Dependencies:**
- Scan application dependencies
- Monitor for new vulnerabilities
- Automated dependency updates
- Test updates before deployment

### Patch Management

**Timeline:**
- Critical vulnerabilities: 24 hours
- High vulnerabilities: 7 days
- Medium vulnerabilities: 30 days
- Low vulnerabilities: 90 days

**Process:**
1. Vulnerability identified
2. Impact assessment
3. Patch/mitigation identified
4. Test in non-production
5. Deploy to production
6. Verify remediation

## Incident Response

### Security Incident Process

1. **Detection:** Automated alerts or manual discovery
2. **Containment:** Isolate affected systems
3. **Investigation:** Determine scope and impact
4. **Eradication:** Remove threat
5. **Recovery:** Restore systems
6. **Post-Mortem:** Learn and improve

### Required Capabilities

- Incident response plan documented
- On-call rotation for security incidents
- Communication channels established
- Forensic logging enabled
- Backup and recovery tested
- Runbooks for common scenarios

## Compliance and Auditing

### Audit Logging

**What to Log:**
- Authentication attempts (success and failure)
- Authorization decisions
- Resource access
- Configuration changes
- Security group changes
- IAM policy changes
- Secrets access

**Log Storage:**
- Centralized and secure
- Tamper-proof (write-once)
- Encrypted
- Retained per compliance requirements
- Monitored for security events

### Cloud Audit Trails

**AWS CloudTrail:**
```hcl
resource "aws_cloudtrail" "main" {
  name                          = "main-trail"
  s3_bucket_name                = aws_s3_bucket.cloudtrail.id
  include_global_service_events = true
  is_multi_region_trail         = true
  enable_log_file_validation    = true
  
  event_selector {
    read_write_type           = "All"
    include_management_events = true
  }
}
```

## Security Testing

### Types of Testing

**SAST (Static Application Security Testing):**
- CodeQL, SonarQube, Semgrep
- Run in CI/CD pipeline
- Block on high severity issues

**DAST (Dynamic Application Security Testing):**
- OWASP ZAP, Burp Suite
- Test running applications
- Run in staging environment

**Dependency Scanning:**
- Dependabot, Snyk, WhiteSource
- Scan for vulnerable dependencies
- Automated updates when safe

**Infrastructure Scanning:**
- tfsec, checkov, Terratest
- Scan IaC before apply
- Security policy as code

## Security Checklist

Before deploying:

- [ ] No secrets in code or containers
- [ ] Secrets stored in secrets manager
- [ ] Encryption at rest enabled
- [ ] Encryption in transit enforced
- [ ] IAM follows least privilege
- [ ] Security groups restrictive
- [ ] Network segmentation applied
- [ ] Containers run as non-root
- [ ] Images scanned for vulnerabilities
- [ ] Dependencies scanned
- [ ] IaC scanned for security issues
- [ ] Audit logging enabled
- [ ] Monitoring and alerting configured
- [ ] Incident response plan exists
- [ ] Compliance requirements met
- [ ] Security review completed

Remember: Security is not a one-time task but a continuous process.
