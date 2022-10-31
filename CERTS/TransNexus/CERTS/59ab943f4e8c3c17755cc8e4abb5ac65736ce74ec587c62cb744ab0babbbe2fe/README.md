# STIR/SHAKEN CA Ecosystem Compliance

## Certificate TransNexus, Inc. SHAKEN Root CA1

Tested At: 31 Oct 22 16:43 UTC\
Initial Validity Period: 7305 day(s)\
Remaining Validity Period: 6868 day(s)\
Subject: CN=TransNexus\\, Inc. SHAKEN Root CA1, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Root CA1, OU=SHAKEN, O=TransNexus\\, Inc., C=US

View: [Click to view](https://understandingwebpki.com/?cert=MIICBzCCAa6gAwIBAgIQSJU%2F1evcbb3BlGgsmoxKtjAKBggqhkjOPQQDAjBkMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSkwJwYDVQQDEyBUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBSb290IENBMTAeFw0yMTA4MjAwMDAwMDBaFw00MTA4MTkyMzU5NTlaMGQxCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBUcmFuc05leHVzLCBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xKTAnBgNVBAMTIFRyYW5zTmV4dXMsIEluYy4gU0hBS0VOIFJvb3QgQ0ExMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEJWqZunl3%2BTOvG5v1B7qA8%2Ft%2FLLhATxSqZzt8EOVX67trV%2FU9UocZc8U6MVm3yRg0MSEo8D%2F3Br4o7nJTHzkcLaNCMEAwDwYDVR0TAQH%2FBAUwAwEB%2FzAOBgNVHQ8BAf8EBAMCAAYwHQYDVR0OBBYEFJqMShmfZMReM7YpjBbKp7gQSy5bMAoGCCqGSM49BAMCA0cAMEQCICr3%2BhNmSquUHkNspgT5tCi%2FfvlMgHMxV0Btv7TDXuuvAiAq6bUMx7jsydNYJEjd6IWBfP36w0n9cRwjxUGSh%2BQewA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [n_pki_ca_key_usage](../../ISSUES/n_pki_ca_key_usage/README.md) | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 1. Raw Bytes: [3 2 0 6], Raw Binary: [00000011 00000010 00000000 00000110] |
| [e_cp1_3_ca_key_usage_crl_sign](../../ISSUES/e_cp1_3_ca_key_usage_crl_sign/README.md) | error | United States SHAKEN CP | The model for managing and communicating the status of revoked certificates is in the form of a distributed Certificate Revocation List (CRL) that is maintained by the STI-PA as described in ATIS-1000080 |
| [w_pki_ca_subject_rdn_unknown](../../ISSUES/w_pki_ca_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


### Not Effective

- e_sti_ca_serial_number
- e_sti_ca_subject_cn
- e_sti_root_certificate_policies
- e_sti_root_extension_unknown

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31/10/2022 at 16:43:22