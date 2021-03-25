---
title: "Data Plane API"
description: "Data Plane API"
images: []
menu:
  docs:
    parent: "development"
weight: 2
toc: false
---

Sample Instructions:

- Click "Authorize" and add a sandbox testing grant token `v1:sandbox.example.com:Iy8TJZcuhicocCklFdwA`.
- Click "POST /facts" and "Try it out".
- Test APIs like create fact with a payload like the following

```json
{
  "fact_type_slug": "ssn",
  "scope_custom_id": "customer_123",
  "value": "123-45-6789"
}
```

<iframe src="https://playground.openprivacy.io/swagger/index.html"
        class="mb-10"
        frameBorder="0" title="OPV Data Plane API" width="750" height="900">
</iframe>
