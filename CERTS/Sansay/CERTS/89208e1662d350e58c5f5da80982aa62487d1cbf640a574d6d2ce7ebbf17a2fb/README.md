# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Carrier One Inc. 705J

Tested At: 28 Nov 23 20:11 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 344 day(s)\
Subject: CN=SHAKEN Carrier One Inc. 705J, OU=Voice NOC, O=Carrier One Inc., ST=Wyoming, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Carrier_One_Inc._705J

[View certificate details](https://understandingwebpki.com/?cert=MIIDiDCCAy2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhfIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTEwNzE1NTUxMVoXDTI0MTEwNjE1NTUxMVowdTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB1d5b21pbmcxGTAXBgNVBAoMEENhcnJpZXIgT25lIEluYy4xEjAQBgNVBAsMCVZvaWNlIE5PQzElMCMGA1UEAwwcU0hBS0VOIENhcnJpZXIgT25lIEluYy4gNzA1SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABB424KM1Fpxx%2FW8QXNzarsuHIfoz1dq4yQZYg2%2FVKqjOnjgP8Ki2ySz0EnsRrod4mu82CD4vsUjdyvxwpjRzFN6jggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYENzA1SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFH6eW3DcQiBGPpZUenPGyDP4EYEiMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BnDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAL0EPOhuYpS%2FpzcrGOXJshA7Cc6dZb9L9tJBW9%2B%2FmbmUAiEAsUyfWXqHdSYD5ZDsEy6r48hQ2ELph8JY9iWloJlvKMI%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 705J', but common name is 'SHAKEN Carrier One Inc. 705J' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 28 Nov 23 20:21 UTC