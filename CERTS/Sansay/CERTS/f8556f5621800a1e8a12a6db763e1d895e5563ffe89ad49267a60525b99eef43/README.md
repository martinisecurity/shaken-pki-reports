# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Zray Technologies Corporation 862J

Tested At: 15 Nov 23 18:03 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -241 day(s)\
Subject: emailAddress=jhansen@ztelco.com, CN=SHAKEN Zray Technologies Corporation 862J, OU=Ztelco, O=Zray Technologies Corporation, ST=California, C=US, emailAddress=jhansen@ztelco.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/862J/order/18_862J_58

[View certificate details](https://understandingwebpki.com/?cert=MIIDFzCCAr2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkYakwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDIxNjE5MDAxN1oXDTIzMDMxODE5MDAxN1owgbIxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMSYwJAYDVQQKDB1acmF5IFRlY2hub2xvZ2llcyBDb3Jwb3JhdGlvbjEPMA0GA1UECwwGWnRlbGNvMTIwMAYDVQQDDClTSEFLRU4gWnJheSBUZWNobm9sb2dpZXMgQ29ycG9yYXRpb24gODYySjEhMB8GCSqGSIb3DQEJARYSamhhbnNlbkB6dGVsY28uY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEGwlDbaswwZmUeqGiVZFghXZ6vA0esmcuWAkunny4Y5JeOWy4o7zGeMh8F0bmzHnM6pDRMwcziGauzxH7wuyYq6OB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEODYySjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFC9G9LSW3%2BPB%2Bx6Yo9I3NU3ubKC5MB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiBChxxZsSoBiJSkqrK8J%2F%2FGWh%2BxyGqlvpgGf%2FdMJN%2BH%2BAIhANwg22Bk5QT%2BKZH6ERe94C4vZ9oDBHRBAMaw99ERAMGG)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 862J' |


Generated: 15 Nov 23 18:10 UTC