# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 769J

Tested At: 02 Dec 22 07:26 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -50 day(s)\
Subject: CN=SHAKEN 769J, OU=SHAKEN, O=Affiliated Technology Solutions LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/471a46b3-3de4-4aa2-964b-2ff03ddbc8fe/d56bd1655cdcf43a507902221b01f753.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDBTCCAqugAwIBAgIQbZ0isgb%2B6mmMw7hPPML2BzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDUyMDE5NDlaFw0yMjEwMTIyMDE5NDhaMGIxCzAJBgNVBAYTAlVTMSwwKgYDVQQKEyNBZmZpbGlhdGVkIFRlY2hub2xvZ3kgU29sdXRpb25zIExMQzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNzY5SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNoezDsN7doNPkBRbS7Olbp8ozjMqmbjQlQ9Q8iqJMhAUqdkHrG47R0Y%2BOi%2BWfcqvRlyCM0W8nPumBOzkTXJXwqjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQU%2FYUuKo6FZahKpl2HZTXSDcwUkVIwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENzY5SjAKBggqhkjOPQQDAgNIADBFAiBvEP1ry2dEPZE%2FIKyepEgMcTXEkr0aNqKcMXuxNS9MkAIhAP9vwNITu2bz51zO26xtkVgrkeKzgcq%2B1Yb8bEvovjRl)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 02 Dec 22 07:30 UTC