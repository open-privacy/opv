---
title: "FAQ"
description: "Answers to frequently asked questions."
lead: "Answers to frequently asked questions."
images: []
menu:
  docs:
    parent: "development"
weight: 5
toc: true
---

## How to migrate database?

By default, OPV uses [ent](https://entgo.io/) for auto migration. If the `OPV_DB_DRIVER` and `OPV_DB_CONNECTION_STR` is set corrently
with the right database permission, OPV will run the database migration on startup (including both dataplane and controlplane).
