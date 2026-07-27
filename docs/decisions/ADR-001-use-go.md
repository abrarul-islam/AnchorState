# ADR-001: Use Go as the Primary Implementation Language

* **Status:** Accepted
* **Date:** 2026-07-27
* **Decision Makers:** CloudGuard Maintainers

---

## Context

CloudGuard is designed as a cloud-native security platform focused on runtime trust verification across Kubernetes and cloud infrastructure.

The initial capability of CloudGuard requires building components that:

* Interact directly with the Kubernetes API.
* Run continuously as background services.
* Process resource events.
* Implement reconciliation loops.
* Operate efficiently in production environments.
* Integrate with cloud-native tooling.

The implementation language must support:

* Reliable systems programming.
* Cloud-native ecosystem integration.
* Strong concurrency support.
* Maintainability for long-term open-source development.

---

## Decision

CloudGuard will use **Go (Golang)** as its primary implementation language.

Go will be used for:

* Kubernetes controllers and operators.
* Core CloudGuard services.
* Cloud-native integrations.
* Security tooling components.
* Command-line interfaces.

---

## Reasons

### 1. Kubernetes Ecosystem Alignment

Kubernetes itself is written in Go, and the majority of the Kubernetes ecosystem is built around Go.

Using Go provides direct access to mature tooling including:

* Kubernetes client libraries.
* Controller-runtime.
* Operator SDK.
* Cloud-native libraries.

This reduces implementation complexity and aligns CloudGuard with existing Kubernetes engineering practices.

---

### 2. Systems Programming Capabilities

CloudGuard requires components that run continuously and interact with infrastructure.

Go provides:

* Efficient compiled binaries.
* Low runtime overhead.
* Strong concurrency primitives.
* Simple deployment through static binaries and containers.

These characteristics are suitable for security infrastructure.

---

### 3. Maintainability

CloudGuard is intended to become an open-source project.

Go prioritizes:

* Readability.
* Simplicity.
* Consistent project structure.
* Ease of onboarding for contributors.

The language's opinionated tooling encourages consistent engineering practices.

---

### 4. Career and Technical Alignment

CloudGuard is intended to demonstrate expertise for roles including:

* Cloud Security Engineer.
* DevSecOps Engineer.
* Platform Security Engineer.
* Cloud Infrastructure Engineer.

Go is widely used in:

* Kubernetes engineering.
* Cloud infrastructure.
* Security tooling.
* Developer platforms.

Using Go allows CloudGuard to demonstrate relevant industry skills.

---

## Alternatives Considered

### Python

**Advantages:**

* Faster prototyping.
* Large security ecosystem.
* Excellent scripting capabilities.

**Rejected because:**

Python is better suited for automation, analysis, and tooling scripts. The core CloudGuard runtime requires a systems-oriented implementation.

Python may still be used for:

* Testing.
* Attack simulations.
* Automation scripts.
* Research tooling.

---

### Rust

**Advantages:**

* Strong memory safety guarantees.
* Excellent performance.
* Growing security ecosystem.

**Rejected because:**

Rust introduces additional complexity and has a smaller Kubernetes ecosystem compared to Go.

Rust may be reconsidered for future performance-critical components.

---

### Java

**Advantages:**

* Mature ecosystem.
* Enterprise adoption.

**Rejected because:**

Java is less aligned with Kubernetes-native infrastructure development compared to Go.

---

## Consequences

### Positive Consequences

* Strong Kubernetes ecosystem compatibility.
* Easier integration with cloud-native tooling.
* Efficient runtime performance.
* Strong alignment with target engineering roles.
* Easier deployment as containerized services.

---

### Negative Consequences

* Requires learning Go deeply.
* Smaller general-purpose ecosystem compared to Python.
* Some rapid prototypes may take longer initially.

---

## Implementation Guidelines

CloudGuard Go projects should follow:

* Standard Go project layout.
* `gofmt` formatting.
* Idiomatic Go practices.
* Automated testing.
* Clear package responsibilities.

Additional language decisions should be documented through future ADRs.

---

## Related Decisions

Future ADRs may cover:

* Kubernetes controller architecture.
* Repository structure.
* Storage choices.
* Event processing design.
* Security model.
* Deployment architecture.

