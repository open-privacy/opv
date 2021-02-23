.PHONY: deps ent vendor swag test build gen run

deps:
	curl -sf https://gobinaries.com/myitcv/gobin | sh
	gobin entgo.io/ent/cmd/ent@v0.6.0
	gobin github.com/swaggo/swag/cmd/swag@v1.7.0

ent:
	ent generate --feature privacy,entql,schema/snapshot ./pkg/ent/schema

vendor:
	go mod tidy
	go mod vendor

swag:
	swag init --parseDependency -d ./cmd/dataplane    -o ./cmd/dataplane/docs
	swag init --parseDependency -d ./cmd/controlplane -o ./cmd/controlplane/docs

test:
	go test -race -covermode=atomic -coverprofile=coverage.txt ./pkg/...

build:
	go build -o build/dataplane    ./cmd/dataplane
	go build -o build/controlplane ./cmd/controlplane
	go build -o build/proxyplane   ./cmd/proxyplane

gen: ent swag

run: build
	$(MAKE) -j _run_controlplane _run_dataplane

_run_dataplane:
	./build/dataplane

_run_controlplane:
	./build/controlplane
