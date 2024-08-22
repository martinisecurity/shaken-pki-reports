# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN WWT INC dba VoIP Networks 053K

Tested At: 22 Aug 24 15:30 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 237 day(s)\
Subject: CN=SHAKEN WWT INC dba VoIP Networks 053K, OU=Tier3, O=WWT INC dba VoIP Networks, ST=New Jersey, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/053K/429C7C70711E3820F0B8E1DEAE6FF32622649907.pem

[View certificate details](https://x509.io/?cert=MIIC7DCCApGgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkmQcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDQxNjAxNDQxMVoXDTI1MDQxNjAxNDQxMVowgYYxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApOZXcgSmVyc2V5MSIwIAYDVQQKDBlXV1QgSU5DIGRiYSBWb0lQIE5ldHdvcmtzMQ4wDAYDVQQLDAVUaWVyMzEuMCwGA1UEAwwlU0hBS0VOIFdXVCBJTkMgZGJhIFZvSVAgTmV0d29ya3MgMDUzSzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHqQUmpKdeH4pDq4Gor6kKgFEkIlU%2BmdBJ2SSMDjvO0yskQPfW%2BNrC60OaqOcGSwM4rnqQuDa%2FoGPJlN2NoOTWijgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDA1M0swFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBSA4DeNDRAWiEJtd14RPDLqdfsvTjAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAIwr62cLGX%2B6ETiljAhfODTlqww5K%2Fd47stTnX1hypvlAiEAkg3r6YSLutCVuqaX2XQ85U0V%2BgyG6uiSu7Rvxmqp23o%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 053K', but common name is 'SHAKEN WWT INC dba VoIP Networks 053K' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 22 Aug 24 15:44 UTC