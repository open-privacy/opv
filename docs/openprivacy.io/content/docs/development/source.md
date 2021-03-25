---
title: "Working with Source Code"
description: "Go Source Code Development."
lead: "Go Source Code Development."
images: []
menu:
  docs:
    parent: "development"
weight: 5
toc: true
---

Development related commands can be found in the [Makefile](https://github.com/open-privacy/opv/blob/main/Makefile) file.

## Compile from the source

Make sure you have `go`, `make`, `docker`, and `docker-compose` installed.

```sh
# Prepare dependencies
make vendor

# Compile and run the data plane and control plane
make run
```

One can open the local swagger UI to test the APIs:

- Default DataPlane Swagger URL: [http://127.0.0.1:28000/swagger/index.html](http://127.0.0.1:28000/swagger/index.html)
- Default ControlPlane Swagger URL: [http://127.0.0.1:27999/swagger/index.html](http://127.0.0.1:27999/swagger/index.html)

```sh
# Compile and run the proxy plane
# It requires at least a "GET/POST /api/v1/facts" grant token to work

OPV_PROXY_PLANE_DEFAULT_DP_GRANT_TOKEN=v1:sandbox.example.com:Iy8TJZcuhicocCklFdwA make run_proxyplane
```

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

## Generate dbdoc with tbls

This is to generate the dbdoc for [Database Visualization →]({{< ref "database" >}}). We leverage a tool
called [https://github.com/k1LoW/tbls](https://github.com/k1LoW/tbls).

```sh
make run

# Once make run is running, open a new terminal to generate the
# database tables visualization from the default _opv.sqlite db locally.
make tbls
```

## Tests

Unit tests:

```sh
make test
```

Functional tests (i.e. the integration tests):

```sh
# Note the functional tests will try to start 3 planes if they don't open ports locally

make local_functional_test
```

## CI/CD

All the CI are run on github. Please refer to [https://github.com/open-privacy/opv/blob/main/.github/workflows/ci.yml](https://github.com/open-privacy/opv/blob/main/.github/workflows/ci.yml).

Currently the playground is deployed and triggered by [https://github.com/open-privacy/opv/blob/main/.github/workflows/deploy_opv_playground.yml](https://github.com/open-privacy/opv/blob/main/.github/workflows/deploy_opv_playground.yml).