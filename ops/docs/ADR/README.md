# ADR

### How use

```shell
export ADR_TEMPLATE=${PWD}/docs/ADR/template/template.md
adr new Implement as Unix shell scripts
```

### Docs

- [How to install](https://github.com/npryce/adr-tools/blob/master/INSTALL.md)

### Update

For update ADR we use GIT, so we can get date each updated and we use git message
for information team about cases for updated ADR.

### Architecture Decision Log (ADL):

  - **CI/CD**
    - [ADR-0001](./decisions/0001-ci-cd.md) - Use CI/CD
    - [ADR-0016](./decisions/0016-lifecycle-deploy.md) - Lifecycle deploy
    - [ADR-0028](./decisions/0028-application-lifecycle-orchestration.md) - Implementing Stage-based Promotion Mechanisms with Kargo for CD in Kubernetes
  - **Kubernetes**
    - [ADR-0022](./decisions/0022-kubernetes.md) - Kubernetes
    - [ADR-0027](./decisions/0027-local-kubernetes-development-tools.md) - Local Kubernetes Development Tools
  - **Security**
    - [ADR-0029](./decisions/0029-ids-and-ips.md) - Intrusion Detection System (IDS) and Intrusion Prevention System (IPS) for Kubernetes (k8s)
  - **DataBase**
    - [ADR-0002](./decisions/0002-cache-system.md) - Cache System
    - [ADR-0005](./decisions/0005-postgres-optimization.md) - Postgres optimization
    - [ADR-0020](./decisions/0020-postgresql-operator.md) - Research Summary and Decision for PostgreSQL Operators: Choosing Crunchy Postgres Operator
