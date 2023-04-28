# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 366G

Tested At: 28 Apr 23 02:05 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -158 day(s)\
Subject: CN=SHAKEN 366G, OU=SHAKEN, O=USA Digital, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/e9497931-96b3-4ff8-8306-1b1273847a4d/afe87d15ad0f75a87d520c582d107a86.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7TCCApOgAwIBAgIQe%2Fg7hQSAGpf%2BK1ldFsHAGDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMTMyMDI2MjBaFw0yMjExMjAyMDI2MTlaMEoxCzAJBgNVBAYTAlVTMRQwEgYDVQQKEwtVU0EgRGlnaXRhbDEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMzY2RzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABEbRFYyKbd05a9hrhxp3E0iP2Kj4eLFmJggrwSIFr%2BKC9UqCzS2lvL7k2fsgxY5PEgK80ovQhyR50ZKyYssOQR2jggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUC4NwHBHVDZmLMKg%2FGKrTAp5UtVcwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMzY2RzAKBggqhkjOPQQDAgNIADBFAiEAyvtwhH6jwkArfxVbayx%2BE0jtDk0%2BV5cmAKiliQF8gnYCIAnIq%2Fx0dtmdTMdPLsb78BsXMHiaZ3eFkqtxitWWoxJL)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 28 Apr 23 02:17 UTC