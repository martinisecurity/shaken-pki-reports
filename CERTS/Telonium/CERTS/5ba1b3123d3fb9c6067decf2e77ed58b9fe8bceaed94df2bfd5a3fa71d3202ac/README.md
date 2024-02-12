# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 620K

Tested At: 12 Feb 24 18:56 UTC\
Initial Validity Period: 367 day(s)\
Remaining Validity Period: 139 day(s)\
Subject: serialNumber=479da13f-6233-4ecd-a4d7-e766c3ab6f28, CN=SHAKEN 620K, O=Miami Telecom, C=US\
Issuer: CN=Telonium STI-CA Intermediate CA, O=Telonium STI-CA\
Link: https://certs.telonium.net/23/b3ddbf00.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIC6DCCAo6gAwIBAgIRALKPvntr8%2F8CP2PtrnLzEPAwCgYIKoZIzj0EAwIwRDEYMBYGA1UEChMPVGVsb25pdW0gU1RJLUNBMSgwJgYDVQQDEx9UZWxvbml1bSBTVEktQ0EgSW50ZXJtZWRpYXRlIENBMB4XDTIzMDYyOTIwMjQxMFoXDTI0MDYyOTIwMjUxMFowajELMAkGA1UEBhMCVVMxFjAUBgNVBAoTDU1pYW1pIFRlbGVjb20xFDASBgNVBAMTC1NIQUtFTiA2MjBLMS0wKwYDVQQFEyQ0NzlkYTEzZi02MjMzLTRlY2QtYTRkNy1lNzY2YzNhYjZmMjgwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASRKMWyfkdulWIplZ0mzTHx3y9LEJzX2jiSZgA4uJRs7l%2B1miWmz4BDxN0ETlRG%2BPazE12SNOpHSdpbfHD5zQaeo4IBOTCCATUwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFHD8V1uV%2FEReB6Nc%2FMnP%2BhAsFx%2FRMB8GA1UdIwQYMBaAFKoku%2F8UdUB5LYdv6A1Bd8q7zYiwMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFAYIKwYBBQUHARoECDAGEwQ2MjBLMF0GDCsGAQQBgqRkxihAAQRNMEsCAQEEGVRlbG9uaXVtIFNUSS1DQSBTUEMgVG9rZW4EK1hWQ1Zic1g2aWt3SmxpdkJ2eVloZXhKYWIyM2JUcWZyMmF0bGQzYlUtdkUwCgYIKoZIzj0EAwIDSAAwRQIgWtRkRx8qwRhfZ3aLdEeAKxjMmB%2F8XrHpUnyMMy4BhJQCIQD%2F8SfbQKfBcktMMfQOGsjQyrDFpZNjOUps2QaDuXneRg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Cannot get SPC value from the TNAuthList extension, TNAuthorizationList shall have only one TN Entry |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_ext_not_specified](../../ISSUES/e_atis_ext_not_specified/README.md) | error | ATIS1000080 | Certificate contains extensions that are not specified: 1.3.6.1.4.1.37476.9000.64.1 |
| [e_atis_tn_auth_list](../../ISSUES/e_atis_tn_auth_list/README.md) | error | ATIS1000080 | TNAuthorizationList shall have only one TN Entry |
| [e_atis_tn_auth_list_spc_format](../../ISSUES/e_atis_tn_auth_list_spc_format/README.md) | error | ATIS1000080 | TNAuthorizationList shall have only one TN Entry |


Generated: 12 Feb 24 19:26 UTC