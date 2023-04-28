# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 749J

Tested At: 28 Apr 23 02:05 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -157 day(s)\
Subject: CN=SHAKEN 749J, OU=SHAKEN, O=iTalkGlobal, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/d7bfa171-ff4f-483d-849b-3f987c919d43/1df6267275cf78943a25d7cfe8e44f62.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7jCCApOgAwIBAgIQSe60Ppti7dVQ84AAM5kvWzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMTQxOTA5NDNaFw0yMjExMjExOTA5NDJaMEoxCzAJBgNVBAYTAlVTMRQwEgYDVQQKEwtpVGFsa0dsb2JhbDEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNzQ5SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABOYbDj6mGprWJ%2BTn6D02CVjIoPw1noTXqAgpXp4q5nfPRBI77cqf5IYtvsc9LEK%2FKau8wZ56UX2sSw22wQ%2BXBvSjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUhfWDpp4KyFPZUiKkQpyKAIFZruswHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENzQ5SjAKBggqhkjOPQQDAgNJADBGAiEA%2FlUU0uWekgPBdlIa%2BagCnkeZhAC%2FZhHf2qR3myWE%2BfkCIQCj8WIiZNSjpmm0hSeCDZFsT%2FT0h%2FSgxYlk8BSSTmJKsw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 28 Apr 23 02:17 UTC