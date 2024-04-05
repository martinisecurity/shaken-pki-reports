# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Mitel Cloud Services, Inc. 670J

Tested At: 05 Apr 24 18:52 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 205 day(s)\
Subject: CN=SHAKEN Mitel Cloud Services\\, Inc. 670J, OU=ComNet, O=Mitel Cloud Services\\, Inc., ST=Arizona, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Mitel_Cloud_Services_Inc_670J_1

[View certificate details](https://x509.io/?cert=MIIDmjCCAz%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhcUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTAyNzIyMjYyN1oXDTI0MTAyNjIyMjYyN1owgYYxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdBcml6b25hMSMwIQYDVQQKDBpNaXRlbCBDbG91ZCBTZXJ2aWNlcywgSW5jLjEPMA0GA1UECwwGQ29tTmV0MS8wLQYDVQQDDCZTSEFLRU4gTWl0ZWwgQ2xvdWQgU2VydmljZXMsIEluYy4gNjcwSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABA0q7tFRXvcFnnAjZdoOwcKszXHRzLj4IaGVP7WfRjQp57mRQKv%2BH0Ej8OXJazvc6uksjt8cC5MqnSBgO%2Ftyck%2BjggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYENjcwSjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFDQXqBPmwAVEFWsIBckuZWbf%2Fuk8MIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BnDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAP6l2uK%2FFabgivZYyCGJsLTGzFnuBRQPeeGNwf9aUgHYAiEAkDAVAWfRGQg1P8Dls4%2Fw9aVG0ZJaxzSqvqGzq1a6Cvw%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 670J', but common name is 'SHAKEN Mitel Cloud Services, Inc. 670J' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 05 Apr 24 19:04 UTC