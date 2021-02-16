.PHONY: deps ent

deps:
	curl -sf https://gobinaries.com/myitcv/gobin | sh
	gobin entgo.io/ent/cmd/ent@v0.6.0
	gobin github.com/swaggo/swag/cmd/swag@v1.7.0

ent:
	ent generate --feature privacy,entql,schema/snapshot ./ent/schema
