# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 5512

Tested At: 06 Jan 23 02:53 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -73 day(s)\
Subject: CN=SHAKEN 5512, OU=SHAKEN, O=Andrew Ward Consulting LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/19aed76a-a067-4715-8a05-1993cc9d939e/26060a26b4428eed5762dc72a3dec527.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BzCCAqKgAwIBAgIQWt7G3lXNKEwV%2BPqvmf9LTDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTcyMDE4MDZaFw0yMjEwMjQyMDE4MDVaMFkxCzAJBgNVBAYTAlVTMSMwIQYDVQQKExpBbmRyZXcgV2FyZCBDb25zdWx0aW5nIExMQzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNTUxMjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNH%2FsgwgUo%2BXiumUK62y6%2FqYzC8Oq6ygcdNcQ2JOFrjDFyRYroO%2FGkiaCg3wJrtBIY%2FyRnl5PZSM60m4iYlYT2ejggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUiclQSPEhy9Mj6yYG8118Xoj7lw4wHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENTUxMjAKBggqhkjOPQQDAgNHADBEAiAVEzefWXtqv9khZjwYIqYNb2Zdx65Kk27jEPIApThTVQIgB%2BV58Tumhee81YcRBTVA2z24K9SjzzIkhXm%2BBns0500%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 06 Jan 23 03:03 UTC