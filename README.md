# Golib Sample Project

A sample project for Golib based on Clean Architecture, include the following layers:

- [Business Rules Layer: Core](./src/core)
- [Adapter Layer: Adapter](./src/adapter)
- [Framework Layer: Internal API (using Basic Auth)](./src/internal)
- [Framework Layer: Public API module (using Jwt Auth)](./src/public)
- [Framework Layer: Worker module](./src/worker)

#### And provide samples for

- [Database migration](./src/migration)
- [Internal API integration test](./src/internal/testing)
- [Public API integration test](./src/public/testing)
- [Worker integration test](./src/worker/testing)
- [Kubernetes deployment](./k8s)
- [CICD template (Gitlab CI)](./.gitlab-ci.yml)
