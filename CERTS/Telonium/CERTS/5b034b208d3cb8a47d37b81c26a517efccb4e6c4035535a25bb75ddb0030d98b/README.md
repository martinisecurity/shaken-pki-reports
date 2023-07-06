# STIR/SHAKEN CA Ecosystem Compliance

## Certificate COMMTRUNKS LLC 622K

Tested At: 06 Jul 23 13:57 UTC\
Initial Validity Period: 367 day(s)\
Remaining Validity Period: 346 day(s)\
Subject: CN=COMMTRUNKS LLC 622K\
Issuer: CN=Telonium STI-CA Intermediate CA, O=Telonium STI-CA\
Link: https://certificates.peeringhub.io/622K/622K.crt

[View certificate details](https://understandingwebpki.com/?cert=MIICnTCCAkKgAwIBAgIRALSSfAKiBBXlbo%2BVen2ZTqMwCgYIKoZIzj0EAwIwRDEYMBYGA1UEChMPVGVsb25pdW0gU1RJLUNBMSgwJgYDVQQDEx9UZWxvbml1bSBTVEktQ0EgSW50ZXJtZWRpYXRlIENBMB4XDTIzMDYxNTE4MzMyM1oXDTI0MDYxNTE4MzQyM1owHjEcMBoGA1UEAxMTQ09NTVRSVU5LUyBMTEMgNjIySzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABN%2Bff%2FQCzK%2FKMn8KXfKNvvjjwi7WcvdJAJyKQmsDfQNR9FEkVL4zSqhQClK%2BXOfE0VVYU1AqCiAryJSbXr%2FFE7ujggE5MIIBNTAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH%2FBAIwADAdBgNVHQ4EFgQUutDWDDAoLsBe6IJxBkcICUHaufcwHwYDVR0jBBgwFoAUqiS7%2FxR1QHkth2%2FoDUF3yrvNiLAwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAUBggrBgEFBQcBGgQIMAYTBDYyMkswXQYMKwYBBAGCpGTGKEABBE0wSwIBAQQZVGVsb25pdW0gU1RJLUNBIFNQQyBUb2tlbgQrWFZDVmJzWDZpa3dKbGl2QnZ5WWhleEphYjIzYlRxZnIyYXRsZDNiVS12RTAKBggqhkjOPQQDAgNJADBGAiEAl%2BzT76HkgOQcT8ahGkvol%2Bl0NPU%2FfR3F%2BNlGZt5VXTkCIQCySXd%2FhmDrHPENKycku%2BPfE7BBrDRhwQ2KOHBgKf84%2Bg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_subject](../../ISSUES/e_atis_subject/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Cannot get SPC value from the TNAuthList extension, TNAuthorizationList shall have only one TN Entry |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_tn_auth_list](../../ISSUES/e_atis_tn_auth_list/README.md) | error | ATIS1000080 | TNAuthorizationList shall have only one TN Entry |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | TNAuthorizationList shall have only one TN Entry |


Generated: 06 Jul 23 14:08 UTC