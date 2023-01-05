# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 0226

Tested At: 05 Jan 23 20:55 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -106 day(s)\
Subject: CN=SHAKEN 0226, OU=SHAKEN, O=Lumos Networks, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/f8a68e44-8fb0-4f28-b533-c4df27ed8e1b/934d9581bbfc5644a50387ca280ebfc0.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8TCCApagAwIBAgIQfuzXAlnTE0P9Cz1Wmux7fTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTQyMDE5NTJaFw0yMjA5MjEyMDE5NTFaME0xCzAJBgNVBAYTAlVTMRcwFQYDVQQKEw5MdW1vcyBOZXR3b3JrczEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMDIyNjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABBTxPhucUjDtnMwlXCQaNzSl6%2BkOaQeYdAqKS%2BprO4BHOpMnIkpyJNJKGMR4cbhCdqmW%2FoJQjq0z1TjXxJqXLJqjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUd4O%2B3nGe7R%2Bwnd9cTK1d8ZEqobEwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDIyNjAKBggqhkjOPQQDAgNJADBGAiEAkh8zZUObbSMhc6oVgMcdBlkpBCLMkS4flW5jmWHnCkQCIQCJZbQluKz1Ddh%2BxYUJ%2Bt5wj1lEf5AUa5Id%2FlLHgZt6xw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 05 Jan 23 21:05 UTC