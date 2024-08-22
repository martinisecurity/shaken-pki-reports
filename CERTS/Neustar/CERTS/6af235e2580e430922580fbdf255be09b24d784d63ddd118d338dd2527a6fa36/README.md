# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 558a

Tested At: 22 Aug 24 15:30 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 331 day(s)\
Subject: CN=SHAKEN 558a, O=D&P Communications, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11569.10207.pem

[View certificate details](https://x509.io/?cert=MIIDBjCCAqygAwIBAgIUOZ7LtNE7qqcXaTvTjphSqcJ%2FY8AwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMB4XDTI0MDcxODIxMjAzOFoXDTI1MDcxODIxMjAzOFowQDELMAkGA1UEBhMCVVMxGzAZBgNVBAoMEkQmUCBDb21tdW5pY2F0aW9uczEUMBIGA1UEAwwLU0hBS0VOIDU1OGEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT2P3oSSrF88UhPDq2Xn0D7KI8NDHjJO%2FjPmcfPWlDNTQlVPHE1OieB%2BOhZg6dHPOEsFTKBvuMrMK5GO3yCYhQIo4IBPDCCATgwFgYIKwYBBQUHARoECjAIoAYWBDU1OGEwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSCThX%2F%2Fp9qZ1HkyJfHbXTD%2FGVFeDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMB0GA1UdDgQWBBTr0Fmu1HCo0g%2BL0CIDxSliy4KXUDAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgLeAoMFDh8SThDtnXWrB1yX33pATxy%2FhLftktrhB3VEICIQDAn4yG%2BVtHFXaslnaZIoNuU%2FuoKZzXjBCUkj0T2WG3cg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_tn_auth_list_spc_format](../../ISSUES/e_atis_tn_auth_list_spc_format/README.md) | error | ATIS1000080 | the SPC value '558a' contains characters other than uppercase letters and numbers |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 22 Aug 24 15:44 UTC