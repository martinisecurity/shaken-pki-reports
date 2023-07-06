# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 620K

Tested At: 06 Jul 23 13:57 UTC\
Initial Validity Period: 367 day(s)\
Remaining Validity Period: 360 day(s)\
Subject: serialNumber=479da13f-6233-4ecd-a4d7-e766c3ab6f28, CN=SHAKEN 620K, O=Miami Telecom, C=US\
Issuer: CN=Telonium STI-CA Intermediate CA, O=Telonium STI-CA\
Link: https://certs.telonium.net/23/b3ddbf00.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIC6DCCAo6gAwIBAgIRALKPvntr8%2F8CP2PtrnLzEPAwCgYIKoZIzj0EAwIwRDEYMBYGA1UEChMPVGVsb25pdW0gU1RJLUNBMSgwJgYDVQQDEx9UZWxvbml1bSBTVEktQ0EgSW50ZXJtZWRpYXRlIENBMB4XDTIzMDYyOTIwMjQxMFoXDTI0MDYyOTIwMjUxMFowajELMAkGA1UEBhMCVVMxFjAUBgNVBAoTDU1pYW1pIFRlbGVjb20xFDASBgNVBAMTC1NIQUtFTiA2MjBLMS0wKwYDVQQFEyQ0NzlkYTEzZi02MjMzLTRlY2QtYTRkNy1lNzY2YzNhYjZmMjgwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASRKMWyfkdulWIplZ0mzTHx3y9LEJzX2jiSZgA4uJRs7l%2B1miWmz4BDxN0ETlRG%2BPazE12SNOpHSdpbfHD5zQaeo4IBOTCCATUwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFHD8V1uV%2FEReB6Nc%2FMnP%2BhAsFx%2FRMB8GA1UdIwQYMBaAFKoku%2F8UdUB5LYdv6A1Bd8q7zYiwMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFAYIKwYBBQUHARoECDAGEwQ2MjBLMF0GDCsGAQQBgqRkxihAAQRNMEsCAQEEGVRlbG9uaXVtIFNUSS1DQSBTUEMgVG9rZW4EK1hWQ1Zic1g2aWt3SmxpdkJ2eVloZXhKYWIyM2JUcWZyMmF0bGQzYlUtdkUwCgYIKoZIzj0EAwIDSAAwRQIgWtRkRx8qwRhfZ3aLdEeAKxjMmB%2F8XrHpUnyMMy4BhJQCIQD%2F8SfbQKfBcktMMfQOGsjQyrDFpZNjOUps2QaDuXneRg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_tn_auth_list](../../ISSUES/e_atis_tn_auth_list/README.md) | error | ATIS1000080 | TNAuthorizationList shall have only one TN Entry |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | TNAuthorizationList shall have only one TN Entry |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Cannot get SPC value from the TNAuthList extension, TNAuthorizationList shall have only one TN Entry |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 06 Jul 23 14:08 UTC