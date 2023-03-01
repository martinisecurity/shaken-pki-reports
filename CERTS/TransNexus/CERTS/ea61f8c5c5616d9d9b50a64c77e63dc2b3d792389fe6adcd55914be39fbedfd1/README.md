# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 747J

Tested At: 01 Mar 23 18:13 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -134 day(s)\
Subject: CN=SHAKEN 747J, OU=SHAKEN, O=Magic Apple, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/9e31f6fe-cfd3-49cc-b9fc-22963012a8d7/98b19b690a0b2a441d77d8772aeb821a.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7jCCApOgAwIBAgIQdNGUf5PQIPi8NgMpse%2FcrTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTEwNDU0MjNaFw0yMjEwMTgwNDU0MjJaMEoxCzAJBgNVBAYTAlVTMRQwEgYDVQQKEwtNYWdpYyBBcHBsZTEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNzQ3SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHSm4glalDZbFPMopOQtV9zlhLGd%2Ff1EgouEfLLtwvn50sThvl1F3BPpbviRAQdMg76K7xeDOJ7hmBbWew%2BDk6%2BjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUWQTnoytTuW18j0OQGjAtAnzjf0UwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENzQ3SjAKBggqhkjOPQQDAgNJADBGAiEAh9KGbuMRi6QwE4hjNaxuF3dMgKAQSU%2FHkNcWQxx0%2FowCIQCy7%2F2pFbZ%2FKEUs6GeF9B8ZhbhU7TNFaL2aJrL0BR05qg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 01 Mar 23 18:22 UTC