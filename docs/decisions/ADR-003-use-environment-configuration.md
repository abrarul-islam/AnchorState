# ADR-003: Use Environment-Based Configuration (12-Factor App)

- Status: Accepted
- Date: 2026-07-28

## Context

AnchorState is intended to run in multiple environments throughout its lifecycle, including:

- Local development
- Continuous Integration (CI)
- Kubernetes clusters
- Cloud environments
- Future production deployments

These environments require different configuration values such as:

- Log level
- Kubernetes namespace
- Controller behavior
- Feature flags
- API endpoints
- Authentication credentials
- Timeouts and retry limits

Hardcoding configuration values or maintaining multiple configuration files increases operational complexity, creates deployment inconsistencies, and risks exposing sensitive information.

The project also aims to follow modern cloud-native engineering practices and align with the principles of the Twelve-Factor App methodology.

## Decision

AnchorState will use environment variables as the primary mechanism for runtime configuration.

Application configuration will be centralized within an internal configuration package responsible for:

- Reading environment variables
- Applying default values
- Validating required configuration
- Exposing a typed configuration structure to the rest of the application

Individual packages must not read environment variables directly.

Instead, configuration will be loaded once during application startup and passed explicitly to components that require it.

## Consequences

### Positive

- Aligns with cloud-native and Kubernetes deployment practices
- Compatible with Docker, Kubernetes ConfigMaps, and Secrets
- Separates configuration from application code
- Simplifies deployments across development, testing, and production
- Enables reproducible builds without environment-specific binaries
- Reduces configuration duplication
- Improves testability through explicit dependency injection
- Centralizes validation and error handling

### Negative

- Requires configuration validation during startup
- Slightly increases initialization complexity
- Missing environment variables must be handled explicitly

## Alternatives Considered

### Option A — Configuration Files

Examples include:

- YAML
- JSON
- TOML

Rejected because:

- Introduces additional parsing logic
- Creates multiple sources of configuration
- Less portable for containerized deployments
- More difficult to manage securely in Kubernetes

Configuration files may still be supported in the future for local development, but they will not be the primary configuration mechanism.

### Option B — Package-Level Global Variables

Rejected because:

- Creates hidden dependencies
- Makes testing more difficult
- Couples packages to process-wide state
- Reduces architectural clarity

### Option C — Environment Variables (Selected)

Chosen because it:

- Follows Twelve-Factor App principles
- Is the standard approach for cloud-native applications
- Integrates naturally with Kubernetes, Docker, and CI/CD systems
- Keeps configuration external to application code
- Encourages explicit dependency injection and clean package boundaries

## Implementation Notes

A dedicated `internal/config` package will:

- Define a typed `Config` structure
- Load configuration from environment variables
- Apply sensible defaults where appropriate
- Validate required values during startup
- Return a fully initialized configuration object

The application's `main` package will load configuration once and pass it explicitly to all components that require it.

## References

- The Twelve-Factor App — https://12factor.net/config
- Go Standard Library (`os`)
- Kubernetes ConfigMaps
- Kubernetes Secrets
