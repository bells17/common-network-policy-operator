
# Image URL to use all building/pushing image targets
VERSION = 1.0.2
IMG ?= bells17/common-network-policy-controller:${VERSION}

all: test manager

# Run tests
test: generate fmt vet manifests
	go test ./pkg/... ./cmd/... -coverprofile cover.out

# Build manager binary
manager: generate fmt vet
	go build -o bin/manager github.com/bells17/common-network-policy-operator/cmd/manager

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet
	go run ./cmd/manager/main.go

# Only run controller
run-only:
	go run ./cmd/manager/main.go

# Install CRDs into a cluster
install: manifests
	kubectl apply -f config/crds

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests
	kubectl apply -f config/crds
	kubectl apply -k config/default

delete: manifests
	kubectl delete -f config/crds
	kubectl delete -k config/default

# Generate manifests e.g. CRD, RBAC etc.
manifests:
	go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go all
	rm -fr config/default/rbac
	rm -fr config/default/crds
	mv config/rbac config/default/rbac
	mv config/crds config/default/crds

# Run go fmt against code
fmt:
	go fmt ./pkg/... ./cmd/...

# Run go vet against code
vet:
	go vet ./pkg/... ./cmd/...

# Generate code
generate:
	go generate ./pkg/... ./cmd/...

# Generate client code
genclient:
	./vendor/k8s.io/code-generator/generate-groups.sh all \
		github.com/bells17/common-network-policy-operator/pkg/client \
		github.com/bells17/common-network-policy-operator/pkg/apis \
		commonnetworkpolicies:v1alpha1

# Build the docker image
docker-build: test
	docker build . -t ${IMG}
	@echo "updating kustomize image patch file for manager resource"
	sed -i'' -e 's@image: .*@image: '"${IMG}"'@' ./config/default/manager_image_patch.yaml
	rm ./config/default/manager_image_patch.yaml-e

# Push the docker image
docker-push:
	docker push ${IMG}

gendeploy: manifests
	kubectl kustomize config/default/ > config/deploy.yaml
