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

OPV uses [casbin](https://github.com/casbin/casbin) for authorization. We leverage the RBAC with multi domain model
to have fine-grained access control of the `Grant Token`. See RBAC model definition.

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
m = g(r.sub, p.sub, r.dom) && keyMatch(r.dom, p.dom) && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act)
```
