# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN California Telecom 319K

Tested At: 05 Apr 24 18:52 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 267 day(s)\
Subject: CN=SHAKEN California Telecom 319K, OU=CaliforniaTelecom, O=California Telecom, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/California_Telecom_319K

[View certificate details](https://x509.io/?cert=MIIC6TCCAo%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJki9UwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTIyODIxNDkxMFoXDTI0MTIyNzIxNDkxMFowgYQxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJDYWxpZm9ybmlhIFRlbGVjb20xGjAYBgNVBAsMEUNhbGlmb3JuaWFUZWxlY29tMScwJQYDVQQDDB5TSEFLRU4gQ2FsaWZvcm5pYSBUZWxlY29tIDMxOUswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASARI%2FByFwiJV9tLoOKjSNGDY97dFOsdg6%2FNSu1Ve0AI6Q7yvIajg5pQwKhWqlqGDbEc3BT7cPXBoCWkglSvSv%2Fo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQzMTlLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUh%2F6stqcI6yhZbhr4pJsRGsNYUT0wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIBHKwx%2BrMLrPGl%2Fqf9VT%2BS4CE%2BmD%2FaDVOWf%2BIpyPED9PAiEA4FGf%2BuOlsgBjQkeeT7%2B8RzGsIBnsveWn3rkjyaJ5d4A%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 319K', but common name is 'SHAKEN California Telecom 319K' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 05 Apr 24 19:04 UTC