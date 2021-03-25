---
title: "7. Auth"
description: "OPV Authn and Authz"
lead: ""
images: []
menu: 
  docs:
    parent: "overview"
weight: 7
toc: true
---

## Authentication

OPV uses Key-Auth (HTTP Header `X-OPV-GRANT-TOKEN`) for authentication.
Notice that by default any request to the data plane (except `/api/v1/healthz` checks) should be authenticated.

## Authorization

### Grant Token

A grant token is required to (1) access to the data plane and (2) deploy the proxy plane.
A grant token can only be created from the control plane. For example, you can create a grant token
with fine-grained control of what to access.

#### Omni Grant Token

This creates an omni grant token to the `sandbox.example.com` domain that has access to
all `allowed_http_methods` and all `paths`.

```sh

curl -X POST 'http://127.0.0.1:27999/api/v1/grants' \
--header 'Content-Type: application/json' \
--data-raw '{
        "allowed_http_methods": ["*"],
        "paths": ["*"],
        "domain": "sandbox.example.com"
}'
```

#### Normal Backend Application Grant Token

This creates a normal backend application grant token to the `sandbox.example.com` domain that has
access to `POST/GET /api/v1/facts or /api/v1/facts/*`.

```sh

curl -X POST 'http://127.0.0.1:27999/api/v1/grants' \
--header 'Content-Type: application/json' \
--data-raw '{
        "allowed_http_methods": ["POST", "GET"],
        "paths": ["/api/v1/facts", "/api/v1/facts/*"],
        "domain": "sandbox.example.com"
}'
```

#### Single Fact Access Grant Token

This creates a single fact access grant token to the `sandbox.example.com` domain that has
access to `GET /api/v1/facts/fact_1LqMuvudjA1xdtqbjd0l`.

```sh

curl -X POST 'http://127.0.0.1:27999/api/v1/grants' \
--header 'Content-Type: application/json' \
--data-raw '{
        "allowed_http_methods": ["GET"],
        "paths": ["/api/v1/facts/fact_1LqMuvudjA1xdtqbjd0l"],
        "domain": "sandbox.example.com"
}'
```

#### Public JS Only Grant Token

This creates a public JS only grant token to the `sandbox.example.com` domain that has
access to `POST /js/v1/facts`.

```sh

curl -X POST 'http://127.0.0.1:27999/api/v1/grants' \
--header 'Content-Type: application/json' \
--data-raw '{
        "allowed_http_methods": ["POST"],
        "paths": ["/js/v1/facts"],
        "domain": "sandbox.example.com"
}'
```

We note that the difference between `/js/v1/facts` and `/api/v1/facts` is that
`/js/v1/facts` will automatically omit the `scope_custom_id` field in the fact creation
payload. This is by design to allow duplicated facts created under the anonymous `scope`
within the domain.

### Casbin Implementation

OPV uses [casbin](https://github.com/casbin/casbin) for the implementation of the grant
token and its policy enforcement. We leverage the **RBAC with multi domains** model
to have fine-grained access control of the `Grant Token`. See RBAC model definition below.

```text
# RBAC with domain pattern model
# https://github.com/casbin/casbin/blob/master/examples/rbac_with_domain_pattern_model.conf

[request_definition]
r = sub, dom, obj, act

[policy_definition]
p = sub, dom, obj, act, eft

[role_definition]
g = _, _, _

[policy_effect]
e = some(where (p.eft == allow)) && !some(where (p.eft == deny))

[matchers]
m = g(r.sub, p.sub, r.dom) && keyMatch2(r.dom, p.dom) && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act)
```
