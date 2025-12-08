# Reconciler Technical Exercise

## Overview

This project provides a complete scaffolding for a Kubernetes controller with
local development tooling.
The controller manages `PostgresConnection` custom resources.
The reconciliation logic is intentionally left empty as a starting point for
implementation.

## Prerequisites

Before you begin, ensure you have the following tools installed:

- **Go 1.21+**: [Download and install Go](https://golang.org/dl/)
- **Docker**: [Install Docker](https://docs.docker.com/get-docker/)
- **kubectl**: [Install kubectl](https://kubernetes.io/docs/tasks/tools/)
- **Tilt**: [Install Tilt](https://docs.tilt.dev/install.html)

## Quick Start

### 1. Install Tools

First, install the development tools:

```bash
make install-tools
```

This will install:
- `controller-gen`: For generating CRDs and DeepCopy methods
- `golangci-lint`: For linting Go code
- `ko`: For building container images
- `kind`: For creating local Kubernetes clusters
- Verification that `tilt` is installed

### 2. Create Local Cluster

Create a KinD cluster for local development:

```bash
make cluster.up
```

This creates a KinD cluster named `interview-reconciler`.

### 3. Start Development Environment

Start Tilt to build and deploy the controller:

```bash
make tilt.up
```

Tilt will:
- Build the controller using `ko` (fast Go container builds)
- Install the CRDs
- Deploy the RBAC configuration
- Deploy the controller
- Stream logs from the controller

You can view the Tilt UI by opening http://localhost:10350 in your browser.

## Testing the Controller

### Apply Sample Resources

Create the credentials secret:

```bash
kubectl apply -f config/samples/secret.yaml
```

Create a PostgresConnection resource:

```bash
kubectl apply -f config/samples/postgresconnection.yaml
```

### Verify Controller Behavior

Check the controller logs:

```bash
kubectl logs -l app=interview-reconciler-controller -f
```

View the PostgresConnection resource:

```bash
kubectl get postgresconnection my-app-db -o yaml
```

## Development Workflow

1. Start an environment with `make cluster.up tilt.up`
2. Edit code and save your changes
3. Tilt automatically rebuilds and redeploys in a few seconds
4. Test your changes with kubectl commands

## Project Structure

```
api/                  # API type definitions
cmd/controller/       # Controller main entrypoint
config/
├── crd/              # Generated CRD manifests
├── controller/       # Controller deployment
├── rbac/             # RBAC configuration
└── samples/          # Example resources
internal/controller/  # Reconciler implementation
kind.yaml             # KinD cluster configuration
Makefile              # Build targets
Tiltfile              # Tilt configuration
```

## Makefile Targets

### Tools

- `make install-tools` - Install required development tools

### Development

- `make generate` - Generate code and CRDs
- `make lint` - Run golangci-lint linter
- `make build` - Build the controller binary locally

### Local Development

- `make cluster.up` - Create KinD cluster
- `make cluster.down` - Delete KinD cluster
- `make tilt.up` - Start Tilt development environment
- `make tilt.down` - Stop Tilt

## API Reference

### PostgresConnection

The `PostgresConnection` CRD represents a PostgreSQL database connection configuration.

**Example:**

```yaml
apiVersion: example.com/v1alpha1
kind: PostgresConnection
metadata:
  name: my-app-db
spec:
  host: postgres.default.svc.cluster.local
  port: 5432
  database: myapp
  credentials:
    secretRef:
      name: db-creds
```

**Spec Fields:**

- `host` (string, required): PostgreSQL server hostname
- `port` (int32, required): PostgreSQL server port (1-65535)
- `database` (string, required): Database name
- `credentials.secretRef.name` (string, required): Name of Secret containing `username` and `password` keys

## Cleanup

Stop Tilt:

```bash
make tilt.down
```

Delete the KinD cluster:

```bash
make cluster.down
```
