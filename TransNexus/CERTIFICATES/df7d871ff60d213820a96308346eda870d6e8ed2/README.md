# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate df7d871ff60d213820a96308346eda870d6e8ed2
Tested At: 2022-10-28 18:54:55 +0000 UTC\
Initial Validity Period: 3652 day(s)\
Remaining Validity Period: 3218 day(s)\
Subject: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Root CA1, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/06ebd24a-1f0a-46d5-8a2f-a7ae49be8ed6/08d04d1bfffa8ddf52073599e552f0bc.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC8TCCApigAwIBAgIQaeMkSbXPfDSfEkC20T6VZTAKBggqhkjOPQQDAjBkMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSkwJwYDVQQDEyBUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBSb290IENBMTAeFw0yMTA4MjAwMDAwMDBaFw0zMTA4MTkyMzU5NTlaMGcxCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBUcmFuc05leHVzLCBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xLDAqBgNVBAMTI1RyYW5zTmV4dXMsIEluYy4gU0hBS0VOIElzc3VpbmcgQ0EzMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEedxAVLKoKQD8g8QPsb9EqRyITRIDArijIRVn1QSXV3Oh7H5HsWihLlTqgbnVM7zF%2FnXicWvV%2FkkgvIKOfmCpW6OCAScwggEjMA8GA1UdEwEB%2FwQFMAMBAf8wDgYDVR0PAQH%2FBAQDAgAGMB0GA1UdDgQWBBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAfBgNVHSMEGDAWgBSajEoZn2TEXjO2KYwWyqe4EEsuWzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMAoGCCqGSM49BAMCA0cAMEQCIGgZROhV4BF%2FKGMwnKGbSUJ0VMdMavpq1jSifXhtc7B3AiA6ODY5dkKtrUbywLLH%2BZJX1UnDad6FZwwQVQpUD0oZHA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 1. Raw Bytes: [3 2 0 6], Raw Binary: [00000011 00000010 00000000 00000110] |
| e_cp1_3_ca_key_usage_crl_sign | error | United States SHAKEN CP | The model for managing and communicating the status of revoked certificates is in the form of a distributed Certificate Revocation List (CRL) that is maintained by the STI-PA as described in ATIS-1000080 |
| w_pki_ca_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| n_pki_ca_key_usage | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |

### Not Effective

- e_sti_ca_subject_cn
- e_sti_ca_serial_number
- n_sti_ca_certificate_policy_critical
- e_sti_ca_crl_distribution
- e_sti_ca_certificate_policies
- e_sti_ca_extension_unknown

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 28/10/2022 at 18:55:01