# ADR-002: Use Go Standard Library `log/slog` for Structured Logging

- **Status:** Accepted
- **Date:** 2026-07-28
- **Decision Makers:** AnchorState Project

---

## Context

AnchorState is a long-running security service that continuously observes runtime state, detects security drift, and produces evidence of unexpected changes.

Logging is a foundational capability of the system.

It supports:

- operational troubleshooting
- security investigations
- debugging
- observability
- auditability

Because logging is integrated throughout the application, the choice of logging framework is an architectural decision that affects maintainability, dependencies, performance, and long-term evolution.

Several mature logging libraries exist within the Go ecosystem, including:

- `zap`
- `zerolog`
- `logrus`

Since Go 1.21, the standard library also provides `log/slog`, offering structured logging without requiring third-party dependencies.

---

## Decision

AnchorState will use Go's standard library `log/slog` package as its primary logging framework.

Structured logging will be implemented using JSON output by default.

The logging subsystem will expose a single initialization function responsible for:

- configuring log levels
- creating the logger
- registering it as the application's default logger

All application components should emit structured log records using `slog` rather than unstructured console output.

---

## Rationale

### 1. Minimize Dependencies

AnchorState aims to keep its core runtime as lightweight and maintainable as possible.

Using the standard library avoids introducing additional external dependencies for a foundational capability.

---

### 2. Long-Term Stability

The Go standard library maintains strong backward compatibility guarantees.

Using `log/slog` reduces maintenance burden associated with third-party logging libraries and minimizes future migration risk.

---

### 3. Structured Logging by Default

AnchorState is intended to operate as an infrastructure service.

Structured JSON logs integrate naturally with modern observability platforms such as:

- OpenTelemetry
- Elasticsearch
- Loki
- Splunk
- Cloud-native logging backends

---

### 4. Consistent Logging Across the Codebase

A single logging implementation ensures consistent log structure, formatting, severity levels, and metadata throughout the project.

This improves debugging, automation, and future integration with monitoring systems.

---

### 5. Simplicity

The standard library provides all functionality currently required by AnchorState.

Additional complexity should only be introduced when justified by concrete engineering requirements.

---

## Alternatives Considered

### Option A — `zap`

Pros:

- Extremely high performance
- Widely adopted
- Mature ecosystem

Cons:

- Additional dependency
- More complex API
- Performance advantages are unlikely to matter at AnchorState's current scale

---

### Option B — `zerolog`

Pros:

- Very low allocation
- Excellent performance
- Small footprint

Cons:

- Additional dependency
- Less aligned with the standard library ecosystem

---

### Option C — `logrus`

Pros:

- Familiar API
- Large ecosystem

Cons:

- Older design
- Slower than newer alternatives
- No longer the preferred choice for new Go projects

---

### Option D — `log/slog` (Chosen)

Pros:

- Standard library
- Zero external dependencies
- Structured logging
- Stable API
- Supported by the Go team
- Sufficient performance for current project requirements

Cons:

- Smaller ecosystem than long-established third-party libraries
- Fewer advanced features than specialized logging frameworks

---

## Consequences

### Positive

- Smaller dependency graph
- Easier maintenance
- Standardized logging throughout the project
- Simpler onboarding for contributors
- Strong compatibility with future observability tooling

### Negative

- Future migration may be required if project requirements significantly exceed the capabilities of `log/slog`
- Some advanced performance optimizations available in specialized logging libraries are not currently utilized

---

## Future Considerations

The logging abstraction should remain isolated so that the underlying implementation can be replaced if future operational requirements justify migration.

Any future migration should preserve the application's structured logging interface and minimize changes to calling code.

---

## Decision Summary

AnchorState adopts Go's standard library `log/slog` as its default structured logging framework to prioritize simplicity, maintainability, standardization, and long-term stability while avoiding unnecessary third-party dependencies.
