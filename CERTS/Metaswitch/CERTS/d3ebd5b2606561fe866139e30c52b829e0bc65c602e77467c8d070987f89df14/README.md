# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Plant Telephone Company SHAKEN 0379

Tested At: 15 Nov 23 15:51 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 901 day(s)\
Subject: CN=Plant Telephone Company SHAKEN 0379, O=Plant Telephone Company, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/81f3a3da944c5778c3db724e7f2a7a58accbd30b

[View certificate details](https://understandingwebpki.com/?cert=MIICZDCCAgqgAwIBAgIQevgDTWjTHT9tA1me3TscyDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDUwNTE1MzkzNVoXDTI2MDUwNDE1MzkzNVowXTELMAkGA1UEBhMCVVMxIDAeBgNVBAoMF1BsYW50IFRlbGVwaG9uZSBDb21wYW55MSwwKgYDVQQDDCNQbGFudCBUZWxlcGhvbmUgQ29tcGFueSBTSEFLRU4gMDM3OTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPwZ4VEuXfIs8Sd5c2%2B7arqxDp0sJpPpaxygUvjED3%2FSm0OzlF%2F8ph3vlj85Crwo7Yg%2BmjznV7dEBA13n0%2FGXoejgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDAzNzkwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUQMe0%2FkE0A2NVV8NKHsAGWUrANBgwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSAAwRQIhAOfrvWDfjWsu0boZpMYsoWVHcvfT3Bra4c7HS23zo6jAAiAPxXMNTH88DvT5IziY0WMqcVchheaFQgtk788AF51tEQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |


Generated: 15 Nov 23 16:51 UTC