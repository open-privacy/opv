---
title: "1. Introduction"
description: "Open Privacy Vault - Secure, Performant, Open Source PII as a Service."
images: []
menu:
  docs:
    parent: "overview"
weight: 1
toc: true
---

Open Privacy Vault - Secure, Performant, Open Source PII as a Service.

## Quick Start

### Running OPV locally

Start from source code:

```sh
git clone https://github.com/open-privacy/opv
cd opv
make vendor
make run
```

Start from docker image (TODO):

```sh
docker run -it -p 27999-28001:27999-28001 open-privacy/opv
```

### Test Local APIs

Now you can test the APIs with `curl`.

```sh
# Create a new grant token from the control plane http://localhost:27999

curl -X POST 'http://localhost:27999/api/v1/grants' \
--header 'Content-Type: application/json' \
--data-raw '{
        "allowed_http_methods": ["*"],
        "domain": "test.com"
}'


# The response will give you a grant token for data plane access
# You can pass the token via HTTP header X-OPV-GRANT-TOKEN
{
  "token": "v1:test.com:6yBQzIcZUaypri8iysut",
  "domain": "test.com",
  "allowed_http_methods": ["*"]
}
```

```sh
# Store a new fact by calling the data plane http://localhost:28000
# Please replace the token with the token you just got above

curl -X POST 'http://localhost:28000/api/v1/facts' \
-H 'Content-Type: application/json' \
-H 'X-OPV-GRANT-TOKEN: v1:test.com:your_new_token' \
--data-raw '{
        "fact_type_slug": "ssn",
        "value": "123-45-6789"
}'
```

### Open Local Swagger UI

One can open the local swagger UI to test the APIs:

- Default DataPlane Swagger URL: [http://127.0.0.1:28000/swagger/index.html](http://127.0.0.1:28000/swagger/index.html)
- Default ControlPlane Swagger URL: [http://127.0.0.1:27999/swagger/index.html](http://127.0.0.1:27999/swagger/index.html)
