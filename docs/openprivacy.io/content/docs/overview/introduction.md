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

### Running Local OPV

Start from source code:

```sh
git clone https://github.com/open-privacy/opv
cd opv
make vendor
make run
```

Start from docker image. You may need to start control plane and data plane separatedly.

```sh
docker run -it -v /tmp/opv_data:/data -p 28000:28000 openprivacyio/opv dataplane
docker run -it -v /tmp/opv_data:/data -p 27999:27999 openprivacyio/opv controlplane
```

### APIs

For more information, please take a look at

- [Data Plane API →]({{< ref "dataplane_api" >}})
- [Control Plane API →]({{< ref "controlplane_api" >}})

Now you can test the APIs with `curl`.

```sh
# Create a new grant token from the control plane http://127.0.0.1:27999

curl -X POST 'http://127.0.0.1:27999/api/v1/grants' \
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
# Store a new fact by calling the data plane http://127.0.0.1:28000
# Please replace the token with the token you just got above

curl -X POST 'http://127.0.0.1:28000/api/v1/facts' \
-H 'Content-Type: application/json' \
-H 'X-OPV-GRANT-TOKEN: v1:test.com:your_new_token' \
--data-raw '{
        "fact_type_slug": "ssn",
        "value": "123-45-6789"
}'
```

### Swagger UI

One can open the local swagger UI to test the APIs:

- Default DataPlane Swagger URL: [http://127.0.0.1:28000/swagger/index.html](http://127.0.0.1:28000/swagger/index.html)
- Default ControlPlane Swagger URL: [http://127.0.0.1:27999/swagger/index.html](http://127.0.0.1:27999/swagger/index.html)

## Performance

We are expecting to see `P99 < 10ms` latency when sending GET requests to `/api/v1/facts/:id`,
which is the most heavily used endpoint to retrieve PIIs from a tokinized `fact`.

```sh
#!/bin/sh

# Make sure you have vegeta downlaoded. https://github.com/tsenart/vegeta
# Replace the X-Opv-Grant-Token and fact_id for the benchmark script

echo $'GET http://127.0.0.1:28000/api/v1/facts/fact_1LqMuvudjA1xdtqbjd0l \nX-Opv-Grant-Token: v1:example.com:gCPMdjk1650km2IA3sgZ' \
    | vegeta attack -duration=10s | vegeta report

# Example of the result
# $ sh vegeta.sh

# Requests      [total, rate, throughput]         500, 50.11, 50.11
# Duration      [total, attack, wait]             9.978s, 9.977s, 1.17ms
# Latencies     [min, mean, 50, 90, 95, 99, max]  753.995µs, 1.679ms, 1.415ms, 2.564ms, 3.061ms, 6.926ms, 12.096ms
# Bytes In      [total, mean]                     62000, 124.00
# Bytes Out     [total, mean]                     0, 0.00
# Success       [ratio]                           100.00%
# Status Codes  [code:count]                      200:500
# Error Set:
```

<img src="/images/BenchmarkFlameGraph.png" class="img-fluid" alt="benchmark_flamegraph.png">

One can also launch the benchmark tests with prometheus and grafana to closely monitor the performance.

```sh
cd ./benchmark
docker-compse up --build
```

<img src="/images/Prometheus.png" class="img-fluid" alt="prometheus.png">
