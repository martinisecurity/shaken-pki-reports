# STIR/SHAKEN CA Ecosystem Compliance

## Certificate NetNumber SHAKEN Root CA 1

Tested At: 08 Feb 23 19:45 UTC\
Initial Validity Period: 9125 day(s)\
Remaining Validity Period: 8627 day(s)\
Subject: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root CA 1\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root CA 1

[View certificate details](https://understandingwebpki.com/?cert=MIIDTzCCArCgAwIBAgIISf7Z%2FTlgXrcwCgYIKoZIzj0EAwQwgYExIzAhBgNVBAMMGk5ldE51bWJlciBTSEFLRU4gUm9vdCBDQSAxMQswCQYDVQQGEwJVUzEWMBQGA1UECgwNTmV0TnVtYmVyIEluYzELMAkGA1UECwwCVVMxFzAVBgNVBAgMDk1hc3NhY2h1c2V0dGVzMQ8wDQYDVQQHDAZMb3dlbGwwHhcNMjEwOTI3MTk0NTU1WhcNNDYwOTIxMTk0NTU1WjCBgTEjMCEGA1UEAwwaTmV0TnVtYmVyIFNIQUtFTiBSb290IENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDCBmzAQBgcqhkjOPQIBBgUrgQQAIwOBhgAEAdTfxsBujQ9g7Wq3w65HUq4rfsSIB1ZvmFdq37XYx6EhbhJHZqYv9JXuzgdpSb7LA4TkpvoDxRteduiSmFUfLXC8AWPudzD4upqXx1nLqI6w6QG2slldNJZhyxO%2B6PULGC8TX40zKvCFG3jE2XuxSYXVTlLi2WdsIej7T915JjwGThuso4HMMIHJMBIGA1UdEwEB%2FwQIMAYBAf8CAQMwDgYDVR0PAQH%2FBAQDAgIEMB8GA1UdIwQYMBaAFH%2FSxmcvsw47rbQTPz8vahAgvo8UMB0GA1UdDgQWBBR%2F0sZnL7MOO620Ez8%2FL2oQIL6PFDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwGgYDVR0gAQH%2FBBAwDjAMBgpghkgBhv8JAQEBMAoGCCqGSM49BAMEA4GMADCBiAJCAUaYZrJYA0P%2FCeus%2FlAFSp3h4zszZ8LGgVscwecmmmxYfT0eP1UH8dJLk67oNvMtQRJBbeYpNlu306rC8rF%2F%2B9NOAkIBNqlzppquQhm%2FxpT7ghMCnsqzDZ5tGM5OUsgsIIzfm489%2Br3B4bLZg9XyQkaRL4fYdjxaZbVmZyqurCINCDzQKDQ%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ca_signature_algorithm](../../ISSUES/e_atis_ca_signature_algorithm/README.md) | error | ATIS1000080 | STI certificates shall contain a Signature Algorithm field with the value 'ecdsa-with-SHA256' |
| [w_shaken_ca_subject_rdn_unknown](../../ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_ca_subject_public_key](../../ISSUES/e_atis_ca_subject_public_key/README.md) | error | ATIS1000080 | STI certificates shall contain a Subject Public Key Info field specifying a Public Key Algorithm of "id-ecPublicKey" and containing a 256-bit public key |

### Not Effective

- e_atis_ca_serial_number
- e_atis_ca_subject_cn
- e_atis_root_certificate_policies
- e_atis_root_extension_unknown

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 08 Feb 23 19:45 UTC