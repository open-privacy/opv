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
Currently OPV supports the following planes:

- Control Plane
- Data Plane
- Proxy Plane

## Control Plane

Control plane (`default port: 27999`) is a group of controllers that handle the admin related logic.
In your network, you should do your best to isolate the access to the control plane.

{{< alert icon="🔥️" text="Please don't expose control plane to the public network." >}}

Control plane securely manages the following data models:

- `Grant`
- `API Audit`

To run the control plane:

```sh
docker run -it \
  -v /tmp/opv_data:/data \
  -p 27999:27999 \
  openprivacyio/opv \
  controlplane
```

Control Plane Swagger

- [Control Plane API →]({{< ref "controlplane_api" >}})
- Local Swagger UI: [http://127.0.0.1:27999/swagger/index.html](http://127.0.0.1:27999/swagger/index.html)

For example, once the control plane is ready, you can create a grant token for the data plane and the proxy plane.

```sh
curl -X POST 'http://127.0.0.1:27999/api/v1/grants' \
--header 'Content-Type: application/json' \
--data-raw '{
        "allowed_http_methods": ["*"],
        "paths": ["*"],
        "domain": "sandbox.example.com"
}'
```

## Data Plane

Data plane (`default port: 28000`) is a group of controllers that handle the crud logic related to
PII information.

Data plane securely manages the following data models:

- `Scope`
- `Fact`
- `Fact Types`

To run the data plane:

```sh
docker run -it \
  -v /tmp/opv_data:/data \
  -p 28000:28000 \
  openprivacyio/opv \
  dataplane
```

Data Plane Swagger

- [Data Plane API →]({{< ref "dataplane_api" >}})
- Local Swagger UI: [http://127.0.0.1:28000/swagger/index.html](http://127.0.0.1:28000/swagger/index.html)

For example, to tokenize your PII is equivalent to creating a `fact` with the sensitive PII value encrypted.
`fact_type_slug` is the slug representing the type of the fact and its validation rules. Currently, the full list of
built-in fact types can be found in [Fact Types →]({{< ref "fact_types" >}}).

```sh
# Tokenize PII into fact

curl -X POST 'http://127.0.0.1:28000/api/v1/facts' \
-H 'Content-Type: application/json' \
-H 'X-OPV-GRANT-TOKEN: v1:sandbox.example.com:Iy8TJZcuhicocCklFdwA' \
--data-raw '{
        "fact_type_slug": "ssn",
        "value": "123-45-6789"
}'
```

```sh
# Detokenize PII from fact

curl -X GET 'http://127.0.0.1:28000/api/v1/facts/fact_rawLfXBSJ0DQXUbtNPl4' \
-H 'X-OPV-GRANT-TOKEN: v1:sandbox.example.com:Iy8TJZcuhicocCklFdwA'
```

## Proxy Plane

Proxy plane modifies the payload on-the-fly according to the routing rules specified.
You can think of proxy plane as a man-in-the-middle that does the transformation of
sensitive PII.

The rules are flexible that the action can be applied to specific parts of the HTTP
requests and responses. For example, one can define a `opv.body.Modifier` rule like
the following JSON config (for more information [Proxy Plane Config →]({{< ref "proxyplane_config" >}})):

```json
"github.com/roney492/opv": {
  "opv.body.Modifier": {
    "scope": [
      "request"
    ],
    "items": [
      {
        "json_pointer_path": "/first_name",
        "fact_type_slug": "ascii",
        "action": "tokenize"
      }
    ]
  }
}
```

It will automatically convert the the payload that sends to this proxy route.

```json
// "action": "tokenize"

{
  "first_name": "John"
}

====> proxy as

{
  "first_name": "fact_rawLfXBSJ0DQXUbtNPl4"
}
```

```json
// "action": "detokenize"

{
  "first_name": "fact_rawLfXBSJ0DQXUbtNPl4"
}

====> proxy as

{
  "first_name": "John"
}
```

Proxy plane currently supports the following protocols and content types.

```sh
HTTP
Content-Type: application/json
Content-Type: text/html                  // coming soon: v1 roadmap
Content-Type: text/xml, application/xml  // coming soon: v1 roadmap

SMTP                                     // coming soon: v1 roadmap
```

For more information about the proxy plane, please see

- [Proxy Plane Config →]({{< ref "proxyplane_config" >}})
- [Proxy Plane Playground →]({{< ref "playground" >}})
