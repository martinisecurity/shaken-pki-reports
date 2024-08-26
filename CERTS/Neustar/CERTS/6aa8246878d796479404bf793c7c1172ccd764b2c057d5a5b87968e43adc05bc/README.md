# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 305K

Tested At: 26 Aug 24 18:12 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -311 day(s)\
Subject: CN=SHAKEN 305K, O=2Talk LLC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/189.253

[View certificate details](https://x509.io/?cert=MIIC%2FDCCAqOgAwIBAgIULzrDaPTE2KMzbw46nMezfZs%2BKVswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMB4XDTIyMTAyMDE4NDUzNFoXDTIzMTAyMDE4NDUzNFowNzELMAkGA1UEBhMCVVMxEjAQBgNVBAoMCTJUYWxrIExMQzEUMBIGA1UEAwwLU0hBS0VOIDMwNUswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASmam9aOO7DrBt7FHQUHfQJeoNmSH2QXCIpuIslFFxepaGXPAzMNgL05J6i56Lwgy%2F1b1AF7sely7XtghA2Y1glo4IBPDCCATgwFgYIKwYBBQUHARoECjAIoAYWBDMwNUswDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSCThX%2F%2Fp9qZ1HkyJfHbXTD%2FGVFeDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMB0GA1UdDgQWBBRfxT5eqFgilo%2B2ZqxTn8jynJwSnzAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgfWpDmPiEPMlzrMd%2BzvOeNpsUMu%2FH%2FuS2IaZQ%2FwtVXrkCIB61onTPp58cG46S8lIxRRi7%2FRCmOv2OluSdo7dbNQcO)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 26 Aug 24 18:49 UTC