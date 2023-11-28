# STIR/SHAKEN CA Ecosystem Compliance

## Certificate COMMTRUNKS LLC 622K

Tested At: 28 Nov 23 10:27 UTC\
Initial Validity Period: 367 day(s)\
Remaining Validity Period: 201 day(s)\
Subject: CN=COMMTRUNKS LLC 622K\
Issuer: CN=Telonium STI-CA Intermediate CA, O=Telonium STI-CA\
Link: https://certificates.peeringhub.io/622K/622K.crt

[View certificate details](https://understandingwebpki.com/?cert=MIICnTCCAkKgAwIBAgIRALSSfAKiBBXlbo%2BVen2ZTqMwCgYIKoZIzj0EAwIwRDEYMBYGA1UEChMPVGVsb25pdW0gU1RJLUNBMSgwJgYDVQQDEx9UZWxvbml1bSBTVEktQ0EgSW50ZXJtZWRpYXRlIENBMB4XDTIzMDYxNTE4MzMyM1oXDTI0MDYxNTE4MzQyM1owHjEcMBoGA1UEAxMTQ09NTVRSVU5LUyBMTEMgNjIySzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABN%2Bff%2FQCzK%2FKMn8KXfKNvvjjwi7WcvdJAJyKQmsDfQNR9FEkVL4zSqhQClK%2BXOfE0VVYU1AqCiAryJSbXr%2FFE7ujggE5MIIBNTAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH%2FBAIwADAdBgNVHQ4EFgQUutDWDDAoLsBe6IJxBkcICUHaufcwHwYDVR0jBBgwFoAUqiS7%2FxR1QHkth2%2FoDUF3yrvNiLAwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAUBggrBgEFBQcBGgQIMAYTBDYyMkswXQYMKwYBBAGCpGTGKEABBE0wSwIBAQQZVGVsb25pdW0gU1RJLUNBIFNQQyBUb2tlbgQrWFZDVmJzWDZpa3dKbGl2QnZ5WWhleEphYjIzYlRxZnIyYXRsZDNiVS12RTAKBggqhkjOPQQDAgNJADBGAiEAl%2BzT76HkgOQcT8ahGkvol%2Bl0NPU%2FfR3F%2BNlGZt5VXTkCIQCySXd%2FhmDrHPENKycku%2BPfE7BBrDRhwQ2KOHBgKf84%2Bg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_c_iso](../../ISSUES/e_atis_subject_c_iso/README.md) | error | ATIS1000080 | Subject MUST be present and MUST contain exactly one value for Country (C=). |
| [e_atis_ext_not_specified](../../ISSUES/e_atis_ext_not_specified/README.md) | error | ATIS1000080 | Certificate contains extensions that are not specified: 1.3.6.1.4.1.37476.9000.64.1 |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'COMMTRUNKS LLC 622K' does not contain 'SHAKEN' |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Cannot get SPC value from the TNAuthList extension, TNAuthorizationList shall have only one TN Entry |
| [e_atis_tn_auth_list](../../ISSUES/e_atis_tn_auth_list/README.md) | error | ATIS1000080 | TNAuthorizationList shall have only one TN Entry |
| [e_atis_subject_dn](../../ISSUES/e_atis_subject_dn/README.md) | error | ATIS1000080 | Subject DN does not contain a Country (C=) attribute |
| [e_atis_tn_auth_list_spc_format](../../ISSUES/e_atis_tn_auth_list_spc_format/README.md) | error | ATIS1000080 | TNAuthorizationList shall have only one TN Entry |
| [e_atis_subject_o_required](../../ISSUES/e_atis_subject_o_required/README.md) | error | ATIS1000080 | The DN does not contain exactly one Organization (O=) attribute. |


Generated: 28 Nov 23 10:53 UTC