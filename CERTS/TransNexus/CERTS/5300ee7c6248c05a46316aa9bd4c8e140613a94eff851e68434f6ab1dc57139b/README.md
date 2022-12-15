# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 0226

Tested At: 15 Dec 22 18:11 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -57 day(s)\
Subject: CN=SHAKEN 0226, OU=SHAKEN, O=Lumos Networks, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/f8a68e44-8fb0-4f28-b533-c4df27ed8e1b/c50c7d0604149700d478cf3b266f4273.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8DCCApagAwIBAgIQSv0Agbr2%2FhH5Cws6BUWXjjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTEyMDIyNDdaFw0yMjEwMTgyMDIyNDZaME0xCzAJBgNVBAYTAlVTMRcwFQYDVQQKEw5MdW1vcyBOZXR3b3JrczEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMDIyNjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHka4Bc1XUaCitndudB%2FLq1wJhckfphBjlXiy0fUpx3wiuMTtbtmCN1H6FmPwOOdbQHtWNgpgF5Ml8IyuDRpixqjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUQuh%2BUAJevoIrKYjaMzCK1oDUErkwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDIyNjAKBggqhkjOPQQDAgNIADBFAiEAqaiJ5TtsRPFjtSs5MLUKB3cUKiFp%2FmGGgO4dd3Wd51kCICW2IYp%2Bc30h5bWc0ZonrvPtYC1EZRku3qPDwOMfJqYr)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 15 Dec 22 18:35 UTC