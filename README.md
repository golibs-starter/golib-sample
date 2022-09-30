# Golib Sample Project

A sample project for Golib based on Clean Architecture, include the following layers:

- [Business Rules Layer: Core](./src/core)
- [Adapter Layer: Adapter](./src/adapter)
- [Framework Layer: Internal API Service (using Basic Auth)](./src/internal)
- [Framework Layer: Public API Service (using Jwt Auth)](./src/public)
- [Framework Layer: Worker Service](./src/worker)

#### And provide samples tests for

- [Internal API integration test](./src/internal/testing) (Functional style)
- [Public API integration test](./src/public/testing) (TestSuite style)
- [Worker integration test](./src/worker/testing) (Functional style)
- [Database migration](./src/migration)
- [Kubernetes deployment](./k8s)
- [CICD template (Gitlab CI)](./.gitlab-ci.yml)

> Note: Integration tests for **Worker Service** with Kafka Consumer will not work correctly with TestSuite.
> Because TestSuite may start multiple instances of consumer group.
> **So Functional style is suggested for this case.**
