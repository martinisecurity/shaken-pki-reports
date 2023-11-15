# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Carrier One Inc. 705J

Tested At: 15 Nov 23 16:34 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 357 day(s)\
Subject: CN=SHAKEN Carrier One Inc. 705J, OU=Voice NOC, O=Carrier One Inc., ST=Wyoming, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Carrier_One_Inc._705J

[View certificate details](https://understandingwebpki.com/?cert=MIIDiDCCAy2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhfIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTEwNzE1NTUxMVoXDTI0MTEwNjE1NTUxMVowdTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB1d5b21pbmcxGTAXBgNVBAoMEENhcnJpZXIgT25lIEluYy4xEjAQBgNVBAsMCVZvaWNlIE5PQzElMCMGA1UEAwwcU0hBS0VOIENhcnJpZXIgT25lIEluYy4gNzA1SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABB424KM1Fpxx%2FW8QXNzarsuHIfoz1dq4yQZYg2%2FVKqjOnjgP8Ki2ySz0EnsRrod4mu82CD4vsUjdyvxwpjRzFN6jggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYENzA1SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFH6eW3DcQiBGPpZUenPGyDP4EYEiMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BnDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAL0EPOhuYpS%2FpzcrGOXJshA7Cc6dZb9L9tJBW9%2B%2FmbmUAiEAsUyfWXqHdSYD5ZDsEy6r48hQ2ELph8JY9iWloJlvKMI%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 705J' |


Generated: 15 Nov 23 17:17 UTC