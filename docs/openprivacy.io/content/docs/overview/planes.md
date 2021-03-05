---
title: "3. Planes"
description: "OPV Planes"
images: []
menu: 
  docs:
    parent: "overview"
weight: 3
toc: true
---

Planes are separated entrypoints and ports to access different segments of OPV.
Currently OPV is planning to support the following planes:

- Control Plane
- Data Plane
- Proxy Plane (TODO)

## Control Plane

Control plane (`default port: 27999`) is a group of controllers that handle the admin related logic.
In your network, you should do your best to isolate the access to the control plane.

{{< alert icon="👉" text="Please don't expose control plane to the public network." >}}

Control plane securely manages the following data models:

- `Grant`
- `Grant Token`
- `Grant Permissions`

## Data Plane

Data plane (`default port: 28000`) is a group of controllers that handle the crud logic related to
PII information.

Data plane securely manages the following data models:

- `Scope`
- `Fact`
- `Fact Types`

## Proxy Plane (TODO)

Proxy plane will handle multiple protocols' (e.g. HTTP, SMTP, and etc.) requests and responses rewrite on-the-fly with PII tokenization rules.
