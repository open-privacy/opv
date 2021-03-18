.PHONY: ent vendor swag test build gen run

ent:
	docker-compose run ent

swag:
	docker-compose run swag

vendor:
	go mod tidy
	go mod vendor

test:
	go test -race -covermode=atomic -coverprofile=coverage.txt ./pkg/...

build:
	go build -o build/dataplane    ./cmd/dataplane
	go build -o build/controlplane ./cmd/controlplane
	go build -o build/proxyplane   ./cmd/proxyplane

gen: ent swag

local_functional_test:
	go clean -testcache
	go test ./functional_test/...

run:
	go build -o build/dataplane    ./cmd/dataplane
	go build -o build/controlplane ./cmd/controlplane
	OPV_DB_CONNECTION_STR="_opv.sqlite?cache=shared&_fk=1" $(MAKE) -j _run_controlplane _run_dataplane

run_proxyplane:
	go build -o build/proxyplane ./cmd/proxyplane
	./build/proxyplane

_run_dataplane:
	./build/dataplane

_run_controlplane:
	./build/controlplane