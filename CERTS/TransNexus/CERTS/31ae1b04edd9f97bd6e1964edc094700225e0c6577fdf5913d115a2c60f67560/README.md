# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 0226

Tested At: 04 Nov 22 01:09 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 5 day(s)\
Subject: CN=SHAKEN 0226, OU=SHAKEN, O=Lumos Networks, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/f8a68e44-8fb0-4f28-b533-c4df27ed8e1b/3308d999b65eab97b5f345c45f50851f.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8DCCApagAwIBAgIQde02Fyun6Yq%2BMi6XE2EJtjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDEyMDI1NDhaFw0yMjExMDgyMDI1NDdaME0xCzAJBgNVBAYTAlVTMRcwFQYDVQQKEw5MdW1vcyBOZXR3b3JrczEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMDIyNjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKVlLBZnfaEty5LtRLmjvYNz2ufoqPRzhV5NRV4BEFpB9Ic0WbfWdXVrpFxQ1AgHpnpYgL%2F0yeMi1i8iBMACWnejggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUGIbe%2F2ui%2Fx%2Fz2SO%2B7OoblIHO84EwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDIyNjAKBggqhkjOPQQDAgNIADBFAiEAvU%2FndN%2B4k3Q62iMhbX4%2BeS2FX12zV7bNYvUb1VBKxToCIEtjz%2BW3XOsauqpFq8WIvvVrcDNnkm2%2B1h%2BioInhXDcM)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 04 Nov 22 01:11 UTC