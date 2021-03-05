---
title: "5. Encryption"
description: "OPV Encryption"
lead: ""
images: []
menu: 
  docs:
    parent: "overview"
weight: 5
toc: true
---

## Supported Encryption Engine

OPV's encryption is designed to work with multiple encryption engines. Currently supported encryption engines are

- NaCl
  - Secretbox
- Hashicorp Vault (TODO)
  - Transit Secret

### Built-in secretbox encryption engine

[Secretbox](https://pkg.go.dev/golang.org/x/crypto/nacl/secretbox) uses XSalsa20 and Poly1305 to encrypt and authenticate messages with secret-key cryptography.

For example, to configure the built-in secretbox encryption engine, you can set a list of secret keys to use.
The first key is always the secret key for encryption, and all the keys will be used for decryption, which allows key rotation.

```sh
OPV_ENCRYPTOR_SECRETBOX_KEYS="GKVB40Yk4JUQno9CVIKxH2uK343GHMVE,Y3XkLGAVQYnRNt1kXO6xyHhoY2pYeieh"
```

Make sure you use a secure random string generator with 32 bytes.

### Hashicorp Vault Transit Secret Engine

TODO

### Encryption Engine Configuration

For more details, see [Env Configuration →]({{< ref "env" >}}).
