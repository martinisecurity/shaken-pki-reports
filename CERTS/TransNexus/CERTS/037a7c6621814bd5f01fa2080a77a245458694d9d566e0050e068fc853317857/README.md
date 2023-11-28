# STIR/SHAKEN CA Ecosystem Compliance

## Certificate TransNexus, Inc. SHAKEN Issuing CA1

Tested At: 28 Nov 23 10:51 UTC\
Initial Validity Period: 3652 day(s)\
Remaining Validity Period: 2822 day(s)\
Subject: CN=TransNexus\\, Inc. SHAKEN Issuing CA1, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Root CA1, OU=SHAKEN, O=TransNexus\\, Inc., C=US

[View certificate details](https://understandingwebpki.com/?cert=MIIC8jCCApigAwIBAgIQYSYdBnmeSS4hKUnugtJ52jAKBggqhkjOPQQDAjBkMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSkwJwYDVQQDEyBUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBSb290IENBMTAeFw0yMTA4MjAwMDAwMDBaFw0zMTA4MTkyMzU5NTlaMGcxCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBUcmFuc05leHVzLCBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xLDAqBgNVBAMTI1RyYW5zTmV4dXMsIEluYy4gU0hBS0VOIElzc3VpbmcgQ0ExMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEwU4ngjABUA071A%2FXOkLFaIEqSCXPlIUeXkoijzyvlyxsRXWYQbtbLXJ9UJ%2BXKJ3KWQFAWO4%2BbkNeiN7BfJt%2Bf6OCAScwggEjMA8GA1UdEwEB%2FwQFMAMBAf8wDgYDVR0PAQH%2FBAQDAgAGMB0GA1UdDgQWBBSUhjk%2F5PWSoJ%2F%2F3Cd1GppG8HnhYjAfBgNVHSMEGDAWgBSajEoZn2TEXjO2KYwWyqe4EEsuWzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMAoGCCqGSM49BAMCA0gAMEUCIEVTrs0QN6oaaz28fi59EsaWwtZf8r%2BxtO8jcc6exkE3AiEA4%2FO%2BExWV%2BYvctFSw6WXhHjSaa%2FNiGRFyqo2m67Qy4aY%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 1. Raw Bytes: [3 2 0 6], Raw Binary: [00000011 00000010 00000000 00000110] |
| [e_shaken_certificate_policies_id_ca](../../ISSUES/e_shaken_certificate_policies_id_ca/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |

### Not Effective

- e_atis_ext_not_specified_ca
- e_atis_serial_number_size_ca
- e_atis_subject_c_iso_ca
- e_atis_subject_key_identifier_size_ca
- e_atis_subject_o_required_ca

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 10:53 UTC