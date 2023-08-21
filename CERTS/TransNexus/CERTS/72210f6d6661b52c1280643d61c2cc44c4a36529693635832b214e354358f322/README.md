# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 056K

Tested At: 21 Aug 23 20:05 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -301 day(s)\
Subject: CN=SHAKEN 056K, OU=SHAKEN, O=Logista Solutions, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/d673a99f-ad2d-4256-8d65-e20ff91adba4/8140550d516aca64f042df75c37894ed.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8jCCApmgAwIBAgIQab3wmvHEqE8nJHHHPJTrlTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTcxODM1MjZaFw0yMjEwMjQxODM1MjVaMFAxCzAJBgNVBAYTAlVTMRowGAYDVQQKExFMb2dpc3RhIFNvbHV0aW9uczEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMDU2SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNLF0Q%2FXpafRl%2Bs7VBR7P0MyrxN7IEgHoVAgy5tUbVXDmpt8EsKzbL7KE6EGltke6EIHvdeThWDuOhUnyz%2F4Ve6jggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUSBGTAHOkwPXFllzArjuuUeYoFKMwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDU2SzAKBggqhkjOPQQDAgNHADBEAiBiHzgKXRThScvXnFHQFxKUPpKf891jFzZpFIp%2B79YcIgIgErQVfeJSQK3%2FDWpflOdvp8i5o3mdp5t6Ox8r9cSgzuo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 21 Aug 23 20:18 UTC