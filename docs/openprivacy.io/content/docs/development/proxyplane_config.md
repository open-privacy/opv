---
title: "Proxy Plane Config"
description: "Proxy Plane Config"
images: []
menu:
  docs:
    parent: "development"
weight: 4
toc: true
---

The configuration that the Proxy Plane needs to start is a single configuration
file (e.g. [opv-proxyplane-http.example.json](https://github.com/roney492/opv/blob/53eb70c1ce9aaaa897863982efb468df487ce7c0/cmd/proxyplane/opv-proxyplane-http.example.json#L105)).

OPV currently enables the proxy plane via [KrakenD](https://www.krakend.io/docs/configuration/overview/),
thus the configuration file needs to be compatible with the KrakenD config format. One
can learn about the structure of the JSON configuration file at [Understanding the KrakenD configuration file](https://www.krakend.io/docs/configuration/structure/).

## Example Configuration

## Root Level

At the root level, one can enable many configuration like `debug`, `cache_ttl`, and `extra_config` for `CORS` support.

```json
{
  "version": 2,
  "name": "OPV Proxy Plane",
  "debug": false,
  "cache_ttl": 3600,
  "timeout": "3s",
  "extra_config": {
    "github_com/krakendio/krakend-cors": {
      "allow_origins": [
        "http*"
      ],
      "allow_headers": [
        "Origin",
        "Authorization",
        "Content-Type",
        "Accept"
      ],
      "expose_headers": [
        "Content-Type",
        "Content-Length"
      ],
      "allow_credentials": true
    }
  }
}
```

## Endpoints Level

`"endpoints"` defines a set of routes that the proxy plane knows how to react to.

Notes

- `headers_to_pass` is usually required if you want to proxy headers more than the default minimal headers.
- `output_encoding: no-op` (endpoints level) and `encoding: no-op` (backends level) are usually required if you want to respond back with non-2XX status code from the backends. Please refer to [Proxying directly to the backends with no-op](https://www.krakend.io/docs/endpoints/no-op/#using-no-op-to-proxy-requests).

```json
  "endpoints": [
    {
      "endpoint": "/tokenize",
      "headers_to_pass": [
        "*"
      ],
      "output_encoding": "no-op",
      "method": "POST",
      "backend": [
        {
            "encoding": "no-op",
            "host": [ "localhost:8080" ],
            "url_pattern": "/__debug/login"
        }
      ]
    }
  ]
```

## Backends Level

`"backend: [...]"` within the `endpoints` level defines a set of upstream servers that the proxy plane knows how to proxy to.

- `host` defines the upstream host.
- `url_pattern` defines the upstream path. One can also leverage KrakenD's advance url pattern matching to build templates for url paths. See [Parameter forwarding](https://www.krakend.io/docs/endpoints/parameter-forwarding/#mandatory-query-string-parameters).
- `extra_config > github.com/roney492/opv` defines a set of OPV proxy modifiers. We follow the standard of [https://github.com/google/martian](https://github.com/google/martian), which means the full list of modifiers can be found here:
  - [Built-in martian modifiers](https://github.com/google/martian/wiki/Modifier-Reference)
    - `log.Logger`
    - `cookie.Modifier`
    - `header.Modifier`
    - `header.Blacklist`
    - `querystring.Modifier`
    - `status.Modifier`
    - `url.Modifier`
    - `body.Modifier`
    - `fifo.Group`
    - `priority.Group`
    - `header.Filter`
  - OPV specific modifiers
    - `opv.body.Modifier`
      - `opv_dataplane_grant_token_from_env`
        - This is optional. If it's not set, the grant token will be using the global default, which is defined by the environment variable `OPV_PROXY_PLANE_DEFAULT_DP_GRANT_TOKEN`. If it's set, the grant token that's used will be derived from the environment variable's value. For example, setting `opv_dataplane_grant_token_from_env=SOME_GRANT_TOKEN_SECRET` is equivalent of using `token := os.Getenv("SOME_GRANT_TOKEN_SECRET")` as the actual grant token.
      - `opv_dataplane_base_url`
        - This is optional. If it's not set, the dataplane base URL will be using the global default, which is defined by the environment variable `OPV_PROXY_PLANE_DEFAULT_DP_BASE_URL`.
      - `scope`
        - It's an array. The possible values for the array item are `request` and `response`, which indicates which part of the request -> response can the `opv.body.Modifier` apply to.
      - `items`
        - `json_pointer_path`: a standard [JSON Pointer](https://tools.ietf.org/html/rfc6901) path indicate which field of the JSON payload should be applied for the `opv.body.Modifier`. Currently this only supports `Content-type: application/json`.
        - `fact_type_slug`: a fact type for `tokenize` action. The built-in list can be found on [PII Fact Types →]({{< ref "fact_types" >}}).
        - `action`: currenly only supports `tokenize` and `detokenize`.

```json
    {
      "endpoint": "/tokenize",
      "method": "POST",
      "backend": [
        {
          "host": [
            "https://httpbin.org"
          ],
          "url_pattern": "/post",
          "extra_config": {
            "github.com/roney492/opv": {
              "opv.body.Modifier": {
                "opv_dataplane_grant_token_from_env": "SOME_GRANT_TOKEN_SECRET",
                "opv_dataplane_base_url": "http://127.0.0.1:28000",
                "scope": [
                  "request"
                ],
                "items": [
                  {
                    "json_pointer_path": "/user/ssn",
                    "fact_type_slug": "ssn",
                    "action": "tokenize"
                  }
                ]
              }
            }
          }
        }
      ]
    }
```
