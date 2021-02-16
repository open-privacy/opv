.PHONY: deps ent vendor test build swag run

deps:
	curl -sf https://gobinaries.com/myitcv/gobin | sh
	gobin entgo.io/ent/cmd/ent@v0.6.0
	gobin github.com/swaggo/swag/cmd/swag@v1.7.0

ent:
	ent generate --feature privacy,entql,schema/snapshot ./pkg/ent/schema

vendor:
	go mod tidy
	go mod vendor

test:
	go test -race -covermode=atomic -coverprofile=coverage.txt ./pkg/...

build:
	go build -o build/data_plane ./cmd/data_plane
	go build -o build/control_plane ./cmd/control_plane
	go build -o build/proxy_plane ./cmd/proxy_plane

swag:
	swag init -d ./cmd/data_plane -o ./cmd/data_plane/docs
	swag init -d ./cmd/control_plane -o ./cmd/control_plane/docs

run: build
	$(MAKE) -j _run_data_plane

_run_data_plane:
	./build/data_plane
