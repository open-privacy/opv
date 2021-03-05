---
title: "Source Code"
description: "Go Source Code Development."
lead: "Go Source Code Development."
images: []
menu:
  docs:
    parent: "development"
weight: 3
toc: true
---

Development related commands can be found in the [Makefile](https://github.com/open-privacy/opv/blob/main/Makefile) file.

## Compile from the source

Make sure you have `go` and `make` installed.

```sh
# Prepare dependencies and compile opv

make deps
make vendor
make run
```

One can open the local swagger UI to test the APIs:

- Default DataPlane Swagger URL: [http://127.0.0.1:28000/swagger/index.html](http://127.0.0.1:28000/swagger/index.html)
- Default ControlPlane Swagger URL: [http://127.0.0.1:27999/swagger/index.html](http://127.0.0.1:27999/swagger/index.html)

## Generate swagger 2.0

When you make any change to APIs (e.g. `pkg/dataplane`, `pkg/controlplane`, `pkg/apimodel`), you may want to regenerate
the swagger 2.0 API doc. This is done by [swag](https://github.com/swaggo/swag).

```sh
# Regenerate swagger 2.0 docs

make swag
```

Notice that dataplane and controlplane's entrypoints are separated, and they are defined in `cmd/dataplane` and `cmd/controlplane`.

## Generate ent models

If you want to migrate or change the schema of the data models, please change the [schema files](https://github.com/open-privacy/opv/tree/main/pkg/ent/schema)
and then run the following command to re-generate the ent related data models.

```sh
# Regenerate ent data models

make ent
```

Note that [ent](https://entgo.io/docs/migrate/#auto-migration) will run the auto migration for schema changes.

## Tests

Unit tests:

```sh
make test
```

Functional tests (i.e. the integration tests):

```sh
make local_functional_test
```
