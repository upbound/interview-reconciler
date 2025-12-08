export CGO_ENABLED ?= 0

.PHONY: help
help: ## Display this help message
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9.-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: generate
generate: controller-gen ## Generate code
	$(CONTROLLER_GEN) object paths="./..."
	$(CONTROLLER_GEN) crd:generateEmbeddedObjectMeta=true paths="./..." output:crd:artifacts:config=config/crd

.PHONY: lint
lint: golangci-lint ## Run golangci-lint linter
	$(GOLANGCI_LINT) run

.PHONY: build
build: ## Build controller binary
	go build -o bin/controller ./cmd/controller

##@ Local Development

.PHONY: cluster.up
cluster.up: kind ## Create KinD cluster
	$(KIND) create cluster --config=kind.yaml --name=interview-reconciler || true
	kubectl cluster-info --context kind-interview-reconciler

.PHONY: cluster.down
cluster.down: kind ## Delete KinD cluster
	$(KIND) delete cluster --name=interview-reconciler

.PHONY: install
install: ## Install CRDs into the cluster

.PHONY: run
run: install ## Run controller locally against the cluster
	go run ./cmd/controller

##@ Tools

.PHONY: install-tools
install-tools: controller-gen golangci-lint kind ## Install required tools
	@echo "All tools installed successfully!"

TOOLSDIR ?= $(shell pwd)/.tools
$(TOOLSDIR):
	mkdir -p $(TOOLSDIR)

CONTROLLER_GEN ?= $(TOOLSDIR)/controller-gen
GOLANGCI_LINT ?= $(TOOLSDIR)/golangci-lint
KIND ?= $(TOOLSDIR)/kind

CONTROLLER_GEN_VERSION ?= v0.16.5
GOLANGCI_LINT_VERSION ?= v1.62.2
KIND_VERSION ?= v0.25.0

$(CONTROLLER_GEN): $(TOOLSDIR)
	@echo "Installing controller-gen $(CONTROLLER_GEN_VERSION)..."
	@GOBIN=$(TOOLSDIR) go install sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_GEN_VERSION)

$(GOLANGCI_LINT): $(TOOLSDIR)
	@echo "Installing golangci-lint $(GOLANGCI_LINT_VERSION)..."
	@GOBIN=$(TOOLSDIR) go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

$(KIND): $(TOOLSDIR)
	@echo "Installing kind $(KIND_VERSION)..."
	@GOBIN=$(TOOLSDIR) go install sigs.k8s.io/kind@$(KIND_VERSION)

.PHONY: controller-gen
controller-gen: $(CONTROLLER_GEN) ## Download controller-gen locally if necessary

.PHONY: golangci-lint
golangci-lint: $(GOLANGCI_LINT) ## Download golangci-lint locally if necessary

.PHONY: kind
kind: $(KIND) ## Download kind locally if necessary
