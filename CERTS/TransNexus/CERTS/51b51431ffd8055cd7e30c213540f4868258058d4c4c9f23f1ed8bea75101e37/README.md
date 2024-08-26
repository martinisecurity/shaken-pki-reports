# STIR/SHAKEN CA Ecosystem Compliance

## Certificate TransNexus, Inc. SHAKEN Issuing CA3

Tested At: 26 Aug 24 18:48 UTC\
Initial Validity Period: 3652 day(s)\
Remaining Validity Period: 2550 day(s)\
Subject: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Root CA1, OU=SHAKEN, O=TransNexus\\, Inc., C=US

[View certificate details](https://x509.io/?cert=MIIC8TCCApigAwIBAgIQaeMkSbXPfDSfEkC20T6VZTAKBggqhkjOPQQDAjBkMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSkwJwYDVQQDEyBUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBSb290IENBMTAeFw0yMTA4MjAwMDAwMDBaFw0zMTA4MTkyMzU5NTlaMGcxCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBUcmFuc05leHVzLCBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xLDAqBgNVBAMTI1RyYW5zTmV4dXMsIEluYy4gU0hBS0VOIElzc3VpbmcgQ0EzMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEedxAVLKoKQD8g8QPsb9EqRyITRIDArijIRVn1QSXV3Oh7H5HsWihLlTqgbnVM7zF%2FnXicWvV%2FkkgvIKOfmCpW6OCAScwggEjMA8GA1UdEwEB%2FwQFMAMBAf8wDgYDVR0PAQH%2FBAQDAgAGMB0GA1UdDgQWBBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAfBgNVHSMEGDAWgBSajEoZn2TEXjO2KYwWyqe4EEsuWzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMAoGCCqGSM49BAMCA0cAMEQCIGgZROhV4BF%2FKGMwnKGbSUJ0VMdMavpq1jSifXhtc7B3AiA6ODY5dkKtrUbywLLH%2BZJX1UnDad6FZwwQVQpUD0oZHA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 1. Raw Bytes: [3 2 0 6], Raw Binary: [00000011 00000010 00000000 00000110] |

### Not Effective

- e_atis_ext_not_specified_ca
- e_atis_serial_number_size_ca
- e_atis_subject_c_iso_ca
- e_atis_subject_key_identifier_size_ca
- e_atis_subject_o_required_ca

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 26 Aug 24 18:49 UTC