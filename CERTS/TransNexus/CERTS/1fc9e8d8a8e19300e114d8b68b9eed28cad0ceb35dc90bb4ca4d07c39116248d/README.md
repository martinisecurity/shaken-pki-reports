# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 749J

Tested At: 05 Jan 23 20:55 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -63 day(s)\
Subject: CN=SHAKEN 749J, OU=SHAKEN, O=iTalkGlobal, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/d7bfa171-ff4f-483d-849b-3f987c919d43/d231dd34f8310b359216e546f8265010.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7TCCApOgAwIBAgIQVNLtyqf4fIKQSODuBK6dnzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjcxOTA3NDJaFw0yMjExMDMxOTA3NDFaMEoxCzAJBgNVBAYTAlVTMRQwEgYDVQQKEwtpVGFsa0dsb2JhbDEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNzQ5SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABA8rfxCTaWnDvGAOxob1eh87KH5u6QBxGzCCO07emzPHE0zAFGnhpmIEtU%2FKLF0yDlgRAzMbyUHi5w1M9Bz6sgmjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUNsfBY7q3i9LnVcgrYftdU7%2FU54owHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENzQ5SjAKBggqhkjOPQQDAgNIADBFAiEAqjVf7Vy%2F8k8ynKnhRthwu4lq1tH5W7gGmqHLQ%2BQYrWYCIBDzO6ccfJWG8tv1MyryM2WjlnABSgJVsGpTTa310Tze)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 05 Jan 23 21:05 UTC