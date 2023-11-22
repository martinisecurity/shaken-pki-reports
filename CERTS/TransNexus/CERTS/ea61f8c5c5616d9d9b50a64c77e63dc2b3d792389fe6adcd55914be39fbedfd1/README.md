# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 747J

Tested At: 22 Nov 23 03:20 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -399 day(s)\
Subject: CN=SHAKEN 747J, OU=SHAKEN, O=Magic Apple, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/9e31f6fe-cfd3-49cc-b9fc-22963012a8d7/98b19b690a0b2a441d77d8772aeb821a.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7jCCApOgAwIBAgIQdNGUf5PQIPi8NgMpse%2FcrTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTEwNDU0MjNaFw0yMjEwMTgwNDU0MjJaMEoxCzAJBgNVBAYTAlVTMRQwEgYDVQQKEwtNYWdpYyBBcHBsZTEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNzQ3SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHSm4glalDZbFPMopOQtV9zlhLGd%2Ff1EgouEfLLtwvn50sThvl1F3BPpbviRAQdMg76K7xeDOJ7hmBbWew%2BDk6%2BjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUWQTnoytTuW18j0OQGjAtAnzjf0UwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENzQ3SjAKBggqhkjOPQQDAgNJADBGAiEAh9KGbuMRi6QwE4hjNaxuF3dMgKAQSU%2FHkNcWQxx0%2FowCIQCy7%2F2pFbZ%2FKEUs6GeF9B8ZhbhU7TNFaL2aJrL0BR05qg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 22 Nov 23 03:38 UTC