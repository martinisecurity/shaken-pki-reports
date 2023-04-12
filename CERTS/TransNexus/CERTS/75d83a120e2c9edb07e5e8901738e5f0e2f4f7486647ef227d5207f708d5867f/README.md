# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 366G

Tested At: 12 Apr 23 01:03 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -145 day(s)\
Subject: CN=SHAKEN 366G, OU=SHAKEN, O=USA Digital, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/e9497931-96b3-4ff8-8306-1b1273847a4d/1299df638a754b6f8cb6b0830f5e9bdb.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7DCCApOgAwIBAgIQeHgm%2F%2FZFi752D%2FhcV366yjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMTAyMDI2MTVaFw0yMjExMTcyMDI2MTRaMEoxCzAJBgNVBAYTAlVTMRQwEgYDVQQKEwtVU0EgRGlnaXRhbDEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMzY2RzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJ0ZsJHXwdgKW9sUNwKmcB4jC3PlfGsb0Zp%2Bm5VoOfytPqFmvlmpZBq0jGFyKoyRam69qlb85HQo%2BXkM2vMnXrejggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUYMfSRGCAzuBglX0CwSlWCoe5iLQwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMzY2RzAKBggqhkjOPQQDAgNHADBEAiAQQPcb4AYUYw5FeQcbcwbD%2F2%2BanMxaLVdVZptE1BwbvAIgJ5GZQw07b85lK3Vf%2FVF5pniBKWJqU45a9AwzMBw83N0%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 12 Apr 23 01:46 UTC