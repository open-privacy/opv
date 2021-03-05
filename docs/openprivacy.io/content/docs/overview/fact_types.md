---
title: "9. PII Fact Types"
description: "OPV PII Fact Types"
lead: ""
images: []
menu: 
  docs:
    parent: "overview"
weight: 9
toc: true
---

## Built-in Fact Types

OPV supports a long list of built-in fact types to determine . For example `email`, `ssn`, `ssnstrict`, `phonenumber`, and etc.
The full list of the current fact types can be found at

```sh
curl --request GET \
  --url http://127.0.0.1:28000/api/v1/fact_types \
  --header 'x-opv-grant-token: v1:example.com:yourtoken'
```

Example of the built-in fact type slugs:

```json
[
  "IMEI",
  "ISO3166Alpha2",
  "ISO3166Alpha3",
  "ISO4217",
  "address",
  "alpha",
  "alphanum",
  "ascii",
  "base64",
  "creditcard",
  "datauri",
  "dialstring",
  "dns",
  "email",
  "float",
  "fullwidth",
  "halfwidth",
  "hexadecimal",
  "hexcolor",
  "host",
  "int",
  "ip",
  "ipv4",
  "ipv6",
  "isbn10",
  "isbn13",
  "json",
  "latitude",
  "longitude",
  "lowercase",
  "mac",
  "multibyte",
  "notnull",
  "null",
  "numeric",
  "phonenumber",
  "photourl",
  "port",
  "printableascii",
  "requri",
  "requrl",
  "rfc3339",
  "rfc3339WithoutZone",
  "rgbcolor",
  "semver",
  "ssn",
  "ssnstrict",
  "uppercase",
  "url",
  "utfdigit",
  "utfletter",
  "utfletternum",
  "utfnumeric",
  "uuid",
  "uuidv3",
  "uuidv4",
  "uuidv5",
  "variablewidth",
]
```
