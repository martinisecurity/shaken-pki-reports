# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Primo Dialler LLC 249K

Tested At: 28 Nov 22 20:39 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 10 day(s)\
Subject: CN=SHAKEN Primo Dialler LLC 249K, OU=Service Provider, O=Primo Dialler LLC, ST=Wyoming, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Primo_0809

[View certificate details](https://understandingwebpki.com/?cert=MIIDkDCCAzagAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkVOIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTEwODE4MTgxNVoXDTIyMTIwODE4MTgxNVowfjELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB1d5b21pbmcxGjAYBgNVBAoMEVByaW1vIERpYWxsZXIgTExDMRkwFwYDVQQLDBBTZXJ2aWNlIFByb3ZpZGVyMSYwJAYDVQQDDB1TSEFLRU4gUHJpbW8gRGlhbGxlciBMTEMgMjQ5SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABFg6xHfZ62p9eFM8GOqgnFF0%2FABqcNt5W1BS2%2BAcWpSKYYtpHghoFeDjZm7wyekVORtp83RzXIvGAU357%2FtK8iajggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYEMjQ5SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFK41QtiLRoRGP0Mw%2BKibFhdIuQrWMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BnDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAP2pPbXtEKhR5ByEYMX6BYfI2j6ylFrPa2dbKAZIRjaQAiBEdClvLyRopXAwYsJWuTSKbNASCMyPaIfc6iAoF6Be1w%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 249K' |


Generated: 28 Nov 22 20:41 UTC