# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 551G

Tested At: 17 Dec 22 16:58 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -34 day(s)\
Subject: CN=SHAKEN 551G, OU=SHAKEN, O=Brightlink Communications LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/551G/1f8e3805-a178-47e6-8e1e-9b1d3d948e2b.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDADCCAqWgAwIBAgIQe%2FBCSDYYNqeL5F8vnB0QHTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDUxNzIwNTBaFw0yMjExMTIxNzIwNDlaMFwxCzAJBgNVBAYTAlVTMSYwJAYDVQQKEx1CcmlnaHRsaW5rIENvbW11bmljYXRpb25zIExMQzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNTUxRzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABFa6uxSFb3gQ27%2BomVHxz%2BSqdG9g54YBLnVDb44S3gndUGDgdvoexY%2By%2FJTl7CK%2F8wmK9k0Nzx%2FFB3EW%2FOHU%2B5ijggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUjEjxBO%2BZhr4XkrXUv241z3XQu9wwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENTUxRzAKBggqhkjOPQQDAgNJADBGAiEAu2XcdxHmJywGvlz%2FdV562iEny%2Fz6ymAkUYuryqJs7e4CIQDjrrWU58JqOcze0qhxDWjGKcgsd4wDNaD8u3oNHaVJ5A%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 17 Dec 22 17:07 UTC