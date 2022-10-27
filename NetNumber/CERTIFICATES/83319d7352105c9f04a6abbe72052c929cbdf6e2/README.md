# STIR/SHAKEN CA Ecosystem Compliance
## NetNumber

### Certificate 83319d7352105c9f04a6abbe72052c929cbdf6e2
Tested At: 2022-10-27 22:31:35 +0000 UTC\
Initial Validity Period: 9125 day(s)\
Remaining Validity Period: 8730 day(s)\
Subject: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root CA 1\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root CA 1

Link: 

View: [Click to view](https://understandingwebpki.com/?cert=MIIDTzCCArCgAwIBAgIISf7Z%2FTlgXrcwCgYIKoZIzj0EAwQwgYExIzAhBgNVBAMMGk5ldE51bWJlciBTSEFLRU4gUm9vdCBDQSAxMQswCQYDVQQGEwJVUzEWMBQGA1UECgwNTmV0TnVtYmVyIEluYzELMAkGA1UECwwCVVMxFzAVBgNVBAgMDk1hc3NhY2h1c2V0dGVzMQ8wDQYDVQQHDAZMb3dlbGwwHhcNMjEwOTI3MTk0NTU1WhcNNDYwOTIxMTk0NTU1WjCBgTEjMCEGA1UEAwwaTmV0TnVtYmVyIFNIQUtFTiBSb290IENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDCBmzAQBgcqhkjOPQIBBgUrgQQAIwOBhgAEAdTfxsBujQ9g7Wq3w65HUq4rfsSIB1ZvmFdq37XYx6EhbhJHZqYv9JXuzgdpSb7LA4TkpvoDxRteduiSmFUfLXC8AWPudzD4upqXx1nLqI6w6QG2slldNJZhyxO%2B6PULGC8TX40zKvCFG3jE2XuxSYXVTlLi2WdsIej7T915JjwGThuso4HMMIHJMBIGA1UdEwEB%2FwQIMAYBAf8CAQMwDgYDVR0PAQH%2FBAQDAgIEMB8GA1UdIwQYMBaAFH%2FSxmcvsw47rbQTPz8vahAgvo8UMB0GA1UdDgQWBBR%2F0sZnL7MOO620Ez8%2FL2oQIL6PFDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwGgYDVR0gAQH%2FBBAwDjAMBgpghkgBhv8JAQEBMAoGCCqGSM49BAMEA4GMADCBiAJCAUaYZrJYA0P%2FCeus%2FlAFSp3h4zszZ8LGgVscwecmmmxYfT0eP1UH8dJLk67oNvMtQRJBbeYpNlu306rC8rF%2F%2B9NOAkIBNqlzppquQhm%2FxpT7ghMCnsqzDZ5tGM5OUsgsIIzfm489%2Br3B4bLZg9XyQkaRL4fYdjxaZbVmZyqurCINCDzQKDQ%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_ca_subject_public_key | error | ATIS-1000080 | STI certificates shall contain a Subject Public Key Info field specifying a Public Key Algorithm of "id-ecPublicKey" and containing a 256-bit public key |
| w_cp1_3_ca_subject_rdn_unknown | warn | United States SHAKEN CP | Only CN, C, and O can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_ca_signature_algorithm | error | ATIS-1000080 | STI certificates shall contain a Signature Algorithm field with the value 'ecdsa-with-SHA256' |

### Not Effective

- e_sti_root_certificate_policies
- e_sti_ca_serial_number
- e_sti_root_extension_unknown
- e_sti_ca_subject_cn

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 27/10/2022 at 22:33:03