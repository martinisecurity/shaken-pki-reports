# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate df7d871ff60d213820a96308346eda870d6e8ed2
Tested At: 2022-10-26 22:31:18 +0000 UTC\
Initial Validity Period: 3652 day(s)\
Remaining Validity Period: 3220 day(s)\
Subject: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Root CA1, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/06ebd24a-1f0a-46d5-8a2f-a7ae49be8ed6/08d04d1bfffa8ddf52073599e552f0bc.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC8TCCApigAwIBAgIQaeMkSbXPfDSfEkC20T6VZTAKBggqhkjOPQQDAjBkMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSkwJwYDVQQDEyBUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBSb290IENBMTAeFw0yMTA4MjAwMDAwMDBaFw0zMTA4MTkyMzU5NTlaMGcxCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBUcmFuc05leHVzLCBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xLDAqBgNVBAMTI1RyYW5zTmV4dXMsIEluYy4gU0hBS0VOIElzc3VpbmcgQ0EzMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEedxAVLKoKQD8g8QPsb9EqRyITRIDArijIRVn1QSXV3Oh7H5HsWihLlTqgbnVM7zF%2FnXicWvV%2FkkgvIKOfmCpW6OCAScwggEjMA8GA1UdEwEB%2FwQFMAMBAf8wDgYDVR0PAQH%2FBAQDAgAGMB0GA1UdDgQWBBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAfBgNVHSMEGDAWgBSajEoZn2TEXjO2KYwWyqe4EEsuWzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMAoGCCqGSM49BAMCA0cAMEQCIGgZROhV4BF%2FKGMwnKGbSUJ0VMdMavpq1jSifXhtc7B3AiA6ODY5dkKtrUbywLLH%2BZJX1UnDad6FZwwQVQpUD0oZHA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| n_pki_ca_key_usage | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 1. Raw Bytes: [3 2 0 6], Raw Binary: [00000011 00000010 00000000 00000110] |

### Not Effective

- e_sti_basic_constraints
- e_sti_ca_crl_distribution
- e_sti_ca_authority_key_identifier
- n_sti_ca_certificate_policy_critical
- e_sti_ca_serial_number
- e_sti_ca_key_usage
- e_sti_ca_extension_unknown
- e_sti_ca_signature_algorithm
- e_sti_ca_subject_cn
- e_sti_ca_version
- w_cp1_3_ca_subject_rdn_unknown
- e_sti_ca_subject
- e_sti_ca_certificate_policies
- e_sti_ca_subject_public_key
- e_sti_ca_subject_key_identifier
- e_cp1_3_ca_key_usage_crl_sign
- e_sti_ca_issuer_dn

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 26/10/2022 at 22:31:35