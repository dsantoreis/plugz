# Security Policy

## Supported Versions

Latest `main` branch only.

## Reporting a Vulnerability

Please open a private security advisory in GitHub or email maintainers with:
- impact summary
- reproduction steps
- suggested remediation

## Baseline Controls

- Bearer auth required for API routes
- Rate limiting per source IP
- Structured JSON logs for audits
- Containerized deployment with non-root runtime recommended
