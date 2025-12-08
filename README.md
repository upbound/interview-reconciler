# Reconciler Technical Exercise

## Overview

This project provides a complete scaffolding for a Kubernetes controller with
local development tooling.
The reconciliation logic is intentionally left empty as a starting point for
implementation.

## Prerequisites

Before you begin, ensure you have the following tools installed:

- **Go 1.21+**: [Download and install Go](https://golang.org/dl/)
- **kubectl**: [Install kubectl](https://kubernetes.io/docs/tasks/tools/)
- **Docker**: [Install Docker](https://docs.docker.com/get-docker/)

## Quick Start

### 1. Install Tools

First, install the development tools:

```bash
make install-tools
```

This will install:
- `controller-gen`: For generating CRDs and DeepCopy methods
- `golangci-lint`: For linting Go code
- `kind`: For creating local Kubernetes clusters

### 2. Create Local Cluster

Create a KinD cluster for local development:

```bash
make cluster.up
```

This creates a KinD cluster named `interview-reconciler`.

### 3. Run the Controller

Run the controller locally (this also installs CRDs):

```bash
make run
```

The controller will connect to your KinD cluster using your kubeconfig.

## Development Workflow

1. Start the environment with `make cluster.up`
2. Run the controller with `make run`
3. Edit code, then stop the controller (Ctrl+C) and re-run `make run`
4. Test your changes with kubectl commands

## Project Structure

```
cmd/controller/       # Controller main entrypoint
internal/controller/  # Reconciler implementation
kind.yaml             # KinD cluster configuration
Makefile              # Build targets
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
- `make install` - Install CRDs into cluster
- `make run` - Run controller locally (also installs CRDs)

## Cleanup

Delete the KinD cluster:

```bash
make cluster.down
```
