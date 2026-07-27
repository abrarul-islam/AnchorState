# Engineering & Development Standards

## 1. Commit Message Convention
We follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

- `feat:` A new feature for the application or API
- `fix:` A bug fix
- `docs:` Documentation changes only
- `refactor:` Code restructuring without changing functionality or adding features
- `test:` Adding or updating unit/integration tests
- `chore:` Maintenance tasks, dependency updates, or build script tweaks

**Example:** `feat(hasher): add SHA-256 state comparison algorithm`

---

## 2. Branch Naming Strategy
- `main` — Primary production-ready branch (protected).
- `feat/<short-description>` — For new capabilities (e.g., `feat/yaml-parser`).
- `fix/<short-description>` — For bug or security fixes (e.g., `fix/nil-pointer-leak`).
- `docs/<short-description>` — For documentation updates.

---

## 3. Versioning Policy
This project adheres to [Semantic Versioning (v2.0.0)](https://semver.org/):
- **Pre-1.0.0 (`v0.x.x`):** Rapid prototyping phase. Breaking changes may occur between minor versions.
- **Post-1.0.0 (`vMAJOR.MINOR.PATCH`):**
  - **MAJOR:** Breaking API changes.
  - **MINOR:** Backward-compatible features.
  - **PATCH:** Backward-compatible bug/security fixes.
