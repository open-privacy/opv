---
title: "6. Hash"
description: "OPV Hash"
images: []
menu: 
  docs:
    parent: "overview"
weight: 6
toc: true
---

## Supported Hashing Algorithms

OPV's hash function is designed to generate consistent hash so that we can enable optional equal-match search
without storing the plaintext of PII information. Currently supported hash algorithms are

- Scrypt
- SHA3 (keccak256)
