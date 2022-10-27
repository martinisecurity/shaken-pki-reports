# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate 36dc4ae1d521b8a5aedd10498e6ce757581b197f
Tested At: 2022-10-27 21:42:13 +0000 UTC\
Initial Validity Period: 7305 day(s)\
Remaining Validity Period: 6872 day(s)\
Subject: CN=TransNexus\\, Inc. SHAKEN Root CA1, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Root CA1, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: 

View: [Click to view](https://understandingwebpki.com/?cert=MIICBzCCAa6gAwIBAgIQSJU%2F1evcbb3BlGgsmoxKtjAKBggqhkjOPQQDAjBkMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSkwJwYDVQQDEyBUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBSb290IENBMTAeFw0yMTA4MjAwMDAwMDBaFw00MTA4MTkyMzU5NTlaMGQxCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBUcmFuc05leHVzLCBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xKTAnBgNVBAMTIFRyYW5zTmV4dXMsIEluYy4gU0hBS0VOIFJvb3QgQ0ExMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEJWqZunl3%2BTOvG5v1B7qA8%2Ft%2FLLhATxSqZzt8EOVX67trV%2FU9UocZc8U6MVm3yRg0MSEo8D%2F3Br4o7nJTHzkcLaNCMEAwDwYDVR0TAQH%2FBAUwAwEB%2FzAOBgNVHQ8BAf8EBAMCAAYwHQYDVR0OBBYEFJqMShmfZMReM7YpjBbKp7gQSy5bMAoGCCqGSM49BAMCA0cAMEQCICr3%2BhNmSquUHkNspgT5tCi%2FfvlMgHMxV0Btv7TDXuuvAiAq6bUMx7jsydNYJEjd6IWBfP36w0n9cRwjxUGSh%2BQewA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 1. Raw Bytes: [3 2 0 6], Raw Binary: [00000011 00000010 00000000 00000110] |
| n_pki_ca_key_usage | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |

### Not Effective

- w_cp1_3_ca_subject_rdn_unknown
- e_sti_ca_serial_number
- e_sti_root_extension_unknown
- e_cp1_3_ca_key_usage_crl_sign
- e_sti_root_certificate_policies
- e_sti_ca_subject_cn

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 27/10/2022 at 21:42:52