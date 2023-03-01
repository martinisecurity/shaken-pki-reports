# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 291K

Tested At: 01 Mar 23 18:13 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -100 day(s)\
Subject: CN=SHAKEN 291K, OU=SHAKEN, O=Hypercore Networks\\, Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/b66a7496-cdd4-46e3-b219-cb0ecdc80d22/683b5c97b341c10e960f09e32c13d17b.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BDCCAp%2BgAwIBAgIQWUdGxxUNucbUYBb1J8%2FrfDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMTQxNjM1MDNaFw0yMjExMjExNjM1MDJaMFYxCzAJBgNVBAYTAlVTMSAwHgYDVQQKExdIeXBlcmNvcmUgTmV0d29ya3MsIEluYzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMjkxSzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABIIIP27Y7dqCwrMXEijz87A4q9yT0UaaXDLvJsOf9a9YBFK2ln5wvNMpcYQFFkwZkyPUTZ7W0BBjmbV38X4BFKKjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUfQ7%2F%2Bopxw%2B6aPOHGlxgAZs0O0NkwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMjkxSzAKBggqhkjOPQQDAgNHADBEAiBchA5VN78h8BWIcJys8y6AaxEOxAwjDlZKW33TgazKQQIgPmQEcwPtqRLR29rSzvLw%2B5or393ry8klTFCALFkhnZs%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 01 Mar 23 18:22 UTC