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