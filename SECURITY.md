# Security Policy

AnchorState is a security-focused open-source project designed to help engineering teams detect runtime security drift in cloud-native environments.

Security is a core principle of the project. We appreciate security researchers and contributors who responsibly report vulnerabilities.

---

# Reporting a Vulnerability

Please **do not publicly disclose security vulnerabilities** through:

- GitHub Issues
- Pull Requests
- Discussions
- Public forums

Instead, report vulnerabilities privately through:

- GitHub Security Advisories
- Email: biti_gchar@simplelogin.com

Please include:

- Description of the vulnerability
- Affected component or module
- AnchorState version
- Kubernetes version and environment details (if applicable)
- Steps to reproduce
- Proof-of-concept or relevant logs (if available)
- Potential security impact

Providing detailed information helps us investigate and respond quickly.

---

# Response Process

After receiving a vulnerability report, we will:

1. Acknowledge receipt of the report.
2. Validate and investigate the issue.
3. Determine severity and impact.
4. Develop and test a fix where appropriate.
5. Publish security updates and advisories when possible.

Response times may vary depending on project availability and vulnerability complexity.

---

# Supported Versions

Security fixes are provided for actively maintained releases.

Currently supported:

| Version | Supported |
|---------|-----------|
| Latest release | Yes |
| Previous releases | Best effort |
| Development versions | Best effort |

---

# Responsible Disclosure

We encourage responsible security research.

Please allow reasonable time for investigation and remediation before publicly disclosing vulnerabilities.

We appreciate researchers who help improve AnchorState's security.

---

# Security Scope

This policy applies to vulnerabilities affecting:

- AnchorState source code
- CloudGuard components
- CloudGuard deployment configurations
- CloudGuard security mechanisms

Issues unrelated to CloudGuard itself, such as vulnerabilities in Kubernetes, cloud providers, or third-party dependencies, should be reported to the appropriate maintainers.
