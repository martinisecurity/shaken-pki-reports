# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 719k

Tested At: 12 Feb 24 16:55 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 311 day(s)\
Subject: CN=SHAKEN 719k, O=925 Telecom LLC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://ssc.getsipnav.com/certs/765ab8defea02c4a4de45048909cf0c21946154f

[View certificate details](https://understandingwebpki.com/?cert=MIIDAzCCAqmgAwIBAgIUPP9AL8AJVZ01jrzhZXJyhatHzzQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMB4XDTIzMTIyMDE0NDAyMloXDTI0MTIxOTE0NDAyMlowPTELMAkGA1UEBhMCVVMxGDAWBgNVBAoMDzkyNSBUZWxlY29tIExMQzEUMBIGA1UEAwwLU0hBS0VOIDcxOWswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARpwvgjPMv0dSWaVQNBrOeeX4YzLtuZf9kD7ufoG0OEkV15HWGXUism8po1EGbxCmMtUStQwiYWXpIZMb5K9Irlo4IBPDCCATgwFgYIKwYBBQUHARoECjAIoAYWBDcxOWswDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSCThX%2F%2Fp9qZ1HkyJfHbXTD%2FGVFeDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMB0GA1UdDgQWBBTm2LpOKnTMsmapBI40u91qebxtKzAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAMDMzwmxxU6ihCueekod2h%2FP%2F%2B10y8o5fVGGouU9nptVAiAfzkzQ37l9ZEzG9VV2EjXXsEV6K9oywK6ZFwBiJHEh3g%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_tn_auth_list_spc_format](../../ISSUES/e_atis_tn_auth_list_spc_format/README.md) | error | ATIS1000080 | the SPC value '719k' contains characters other than uppercase letters and numbers |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 12 Feb 24 17:02 UTC