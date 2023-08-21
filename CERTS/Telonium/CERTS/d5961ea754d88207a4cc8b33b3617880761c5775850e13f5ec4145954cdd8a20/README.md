# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 627K

Tested At: 21 Aug 23 20:07 UTC\
Initial Validity Period: 367 day(s)\
Remaining Validity Period: 315 day(s)\
Subject: serialNumber=21030943-cc85-46da-b291-cc73aaad2540, CN=SHAKEN 627K, O=Come 2 LLC, C=US\
Issuer: CN=Telonium STI-CA Intermediate CA, O=Telonium STI-CA\
Link: https://certs.telonium.net/23/7dfe33b6.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIC5DCCAougAwIBAgIRANNs25W%2FtclqJc2dPtVrPKswCgYIKoZIzj0EAwIwRDEYMBYGA1UEChMPVGVsb25pdW0gU1RJLUNBMSgwJgYDVQQDEx9UZWxvbml1bSBTVEktQ0EgSW50ZXJtZWRpYXRlIENBMB4XDTIzMDYzMDIxNTY0MFoXDTI0MDYzMDIxNTc0MFowZzELMAkGA1UEBhMCVVMxEzARBgNVBAoTCkNvbWUgMiBMTEMxFDASBgNVBAMTC1NIQUtFTiA2MjdLMS0wKwYDVQQFEyQyMTAzMDk0My1jYzg1LTQ2ZGEtYjI5MS1jYzczYWFhZDI1NDAwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARF%2FLGFA4NGmwx0GORQk9lmHCU1P0irZuGMtob9Cdh%2Fe%2BA93mOQIRTqC5pLpei5cQt2bS%2BKGPBM0kWK3bpjD6WGo4IBOTCCATUwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFFg3JwCs5X2c6g7U6%2FeaT9DGcmauMB8GA1UdIwQYMBaAFKoku%2F8UdUB5LYdv6A1Bd8q7zYiwMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFAYIKwYBBQUHARoECDAGEwQ2MjdLMF0GDCsGAQQBgqRkxihAAQRNMEsCAQEEGVRlbG9uaXVtIFNUSS1DQSBTUEMgVG9rZW4EK1hWQ1Zic1g2aWt3SmxpdkJ2eVloZXhKYWIyM2JUcWZyMmF0bGQzYlUtdkUwCgYIKoZIzj0EAwIDRwAwRAIgV%2BSxNbumjNMlQkPmOlmd9G4Mk3UM%2BoU75x4Sgj%2BnaQUCIDsHFYlhMvQlv795yxWaLa3TIro%2FdfHNSgGbGPn6BN%2Bf)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Cannot get SPC value from the TNAuthList extension, TNAuthorizationList shall have only one TN Entry |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | TNAuthorizationList shall have only one TN Entry |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_tn_auth_list](../../ISSUES/e_atis_tn_auth_list/README.md) | error | ATIS1000080 | TNAuthorizationList shall have only one TN Entry |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 21 Aug 23 20:18 UTC