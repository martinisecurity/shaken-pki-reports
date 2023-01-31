# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 0226

Tested At: 31 Jan 23 21:39 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -81 day(s)\
Subject: CN=SHAKEN 0226, OU=SHAKEN, O=Lumos Networks, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/f8a68e44-8fb0-4f28-b533-c4df27ed8e1b/20c3d38176f6dea405fc3505a86f2d39.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8DCCApagAwIBAgIQR12szJxD21EpO0J4lZbrODAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDQyMDI1NTNaFw0yMjExMTEyMDI1NTJaME0xCzAJBgNVBAYTAlVTMRcwFQYDVQQKEw5MdW1vcyBOZXR3b3JrczEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMDIyNjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJJRX8bfyQy479PmZ3Y3hhifzmG9MdAJ%2Fn5Le6GEM%2F81saSjyJCFCrqYEZ2VD7m4mIYK6lHrjY2lanVHQJG7%2F%2FejggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUST15EnVAxn1g%2BcTsn56yLZB3JkAwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDIyNjAKBggqhkjOPQQDAgNIADBFAiA4Lj9KykAZsHQZyCxNnUSEsEaf0Tu8aelRN4ld9GQZawIhANbLf4%2BCIyF8k4uYPXzIGWArsPxKXBk89dFqTDNzLg2r)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 31 Jan 23 21:50 UTC