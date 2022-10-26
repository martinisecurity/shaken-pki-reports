# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate 2d540c7179aad19290a2098ccac8053645527154
Tested At: 2022-10-26 20:21:19 +0000 UTC\
Initial Validity Period: 3652 day(s)\
Remaining Validity Period: 3220 day(s)\
Subject: CN=TransNexus\\, Inc. SHAKEN Issuing CA1, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Root CA1, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.transnexus.com/663J/0b51e9f5-0e35-4d02-acef-5c91c3a6b903.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC8jCCApigAwIBAgIQYSYdBnmeSS4hKUnugtJ52jAKBggqhkjOPQQDAjBkMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSkwJwYDVQQDEyBUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBSb290IENBMTAeFw0yMTA4MjAwMDAwMDBaFw0zMTA4MTkyMzU5NTlaMGcxCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBUcmFuc05leHVzLCBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xLDAqBgNVBAMTI1RyYW5zTmV4dXMsIEluYy4gU0hBS0VOIElzc3VpbmcgQ0ExMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEwU4ngjABUA071A%2FXOkLFaIEqSCXPlIUeXkoijzyvlyxsRXWYQbtbLXJ9UJ%2BXKJ3KWQFAWO4%2BbkNeiN7BfJt%2Bf6OCAScwggEjMA8GA1UdEwEB%2FwQFMAMBAf8wDgYDVR0PAQH%2FBAQDAgAGMB0GA1UdDgQWBBSUhjk%2F5PWSoJ%2F%2F3Cd1GppG8HnhYjAfBgNVHSMEGDAWgBSajEoZn2TEXjO2KYwWyqe4EEsuWzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMAoGCCqGSM49BAMCA0gAMEUCIEVTrs0QN6oaaz28fi59EsaWwtZf8r%2BxtO8jcc6exkE3AiEA4%2FO%2BExWV%2BYvctFSw6WXhHjSaa%2FNiGRFyqo2m67Qy4aY%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| n_pki_ca_key_usage | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 1. Raw Bytes: [3 2 0 6], Raw Binary: [00000011 00000010 00000000 00000110] |

### Not Effective

- e_sti_ca_signature_algorithm
- w_cp1_3_ca_subject_rdn_unknown
- e_cp1_3_ca_key_usage_crl_sign
- e_sti_ca_certificate_policies
- e_sti_ca_version
- e_sti_ca_serial_number
- e_sti_ca_crl_distribution
- e_sti_ca_issuer_dn
- n_sti_ca_certificate_policy_critical
- e_sti_basic_constraints
- e_sti_ca_key_usage
- e_sti_ca_subject_cn
- e_sti_ca_authority_key_identifier
- e_sti_ca_subject
- e_sti_ca_subject_public_key
- e_sti_ca_subject_key_identifier
- e_sti_ca_extension_unknown

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 26/10/2022 at 20:22:11