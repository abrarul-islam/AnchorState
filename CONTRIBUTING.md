# Contributing to AnchorState

Thank you for your interest in contributing to AnchorState.

AnchorState is an open-source, founder-led project focused on runtime trust verification for Kubernetes and cloud-native infrastructure. The project is in its early stages, and contributions are appreciated as it continues to evolve.

## Ways to Contribute

Contributions are welcome in many forms, including:

* Reporting bugs
* Improving documentation
* Suggesting new ideas
* Improving tests
* Enhancing examples
* Fixing small issues
* Improving the developer experience

Not every contribution needs to involve writing production code.

## Before You Start

For larger changes, please open an issue first to discuss the proposal before beginning implementation. This helps avoid duplicate work and ensures the change aligns with the project's long-term direction.

Small bug fixes, documentation improvements, and minor enhancements generally do not require prior discussion.

## Development Workflow

1. Fork the repository.
2. Create a feature branch.
3. Make focused, well-scoped changes.
4. Add or update tests where appropriate.
5. Update documentation if behavior changes.
6. Submit a pull request.

## Local Development Setup
To build and test AnchorState locally:
1. Ensure you have Go 1.22+ installed.
2. Clone your fork: `git clone https://github.com/abrarul-islam/AnchorState.git`
3. Download dependencies: `go mod download`
4. Run the test suite: `go test ./...`
*(Note: Full Kubernetes local-cluster testing requirements using Kind/Minikube will be added as the controller architecture matures).*

## Developer Certificate of Origin (DCO)
To legally accept your contributions, CloudGuard requires all commits to be signed off, adhering to the Developer Certificate of Origin (DCO). 

This is a simple statement that you wrote the contribution or have the right to contribute it under the open-source license.

To sign off on a commit, simply use the `-s` flag when committing:
`git commit -s -m "feat(hasher): add sha256 module"`

## Code Quality & Formatting
Before submitting a pull request, ensure your code matches standard Go formatting:
- Run `go fmt ./...` on your changes.
- Ensure all existing and new tests pass via `go test ./...`.

## Coding Principles

When contributing, please aim to:

* Write clear and readable code.
* Prefer simplicity over unnecessary complexity.
* Keep changes focused on a single problem.
* Follow the existing project structure and style.
* Include tests where practical.
* Update documentation when functionality changes.

## Commit Messages

AnchorState follows the Conventional Commits specification.

Examples:

```text
feat(controller): add secret watcher
fix(hash): ensure deterministic key ordering
docs(readme): clarify trust model
test(controller): add secret drift detection tests
```

## Pull Requests

A good pull request should:

* Solve one clearly defined problem.
* Include a concise explanation of the change.
* Pass all existing tests.
* Include new tests when appropriate.
* Update documentation if required.

Pull requests may receive review comments before being merged.

## Project Direction

AnchorState is currently maintained as a founder-led open-source project.

Community contributions are encouraged, while the project's architecture, long-term vision, and core security model are intentionally guided by the primary maintainer to ensure consistency as the platform evolves.

## Reporting Security Issues

Please do **not** publicly disclose security vulnerabilities.

If you discover a security issue in AnchorState itself, follow the responsible disclosure process described in **[SECURITY.md](https://github.com/abrarul-islam/AnchorState/blob/main/SECURITY.md)**.

## Thank You

Every contribution—whether it's code, documentation, testing, or feedback—helps improve AnchorState and is genuinely appreciated.

