# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 1563

Tested At: 21 Nov 23 18:49 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 108 day(s)\
Subject: CN=SHAKEN 1563, O=Hooper Telephone Company, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/208.287

[View certificate details](https://understandingwebpki.com/?cert=MIIDDDCCArKgAwIBAgIUNGfcyaC5hbY%2FCxloFAaMgqEdffIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMB4XDTIzMDMwODIyMDMzN1oXDTI0MDMwNzIyMDMzN1owRjELMAkGA1UEBhMCVVMxITAfBgNVBAoMGEhvb3BlciBUZWxlcGhvbmUgQ29tcGFueTEUMBIGA1UEAwwLU0hBS0VOIDE1NjMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQbD6NHxdle5NRG%2BBmFdIkzCGJMgiGuorW%2FTg0AgrBomhAurTiqwKMG5ib57DewYeNCY%2Bn5oGDi7IBvaKRgkZJYo4IBPDCCATgwFgYIKwYBBQUHARoECjAIoAYWBDE1NjMwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSCThX%2F%2Fp9qZ1HkyJfHbXTD%2FGVFeDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMB0GA1UdDgQWBBQyuMlI8QqDqFBH6cfkM5iX5E575zAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAMwU%2FPj1dpQXATO5o9YvdiFDIfU%2FM41A2NRsTm%2BnrWhbAiAgQBRTbfm2AIqfqX4qkUV%2FxUZH9NVTfhhnMK1gd2cW4A%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 19:18 UTC