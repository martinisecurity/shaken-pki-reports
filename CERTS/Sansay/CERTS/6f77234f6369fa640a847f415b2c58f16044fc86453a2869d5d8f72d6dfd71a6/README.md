# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Medtel Communications 994J

Tested At: 22 Aug 24 15:53 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 279 day(s)\
Subject: CN=SHAKEN Medtel Communications 994J, OU=Medtel Communications, O=Medtel Communications, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/994J/429C7C70711E3820F0B8E1DEAE6FF32622649DE8.pem

[View certificate details](https://x509.io/?cert=MIIC7zCCApagAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJknegwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDUyODEyMjMzMVoXDTI1MDUyODEyMjMzMVowgYsxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdGbG9yaWRhMR4wHAYDVQQKDBVNZWR0ZWwgQ29tbXVuaWNhdGlvbnMxHjAcBgNVBAsMFU1lZHRlbCBDb21tdW5pY2F0aW9uczEqMCgGA1UEAwwhU0hBS0VOIE1lZHRlbCBDb21tdW5pY2F0aW9ucyA5OTRKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAElm6J72YV8ZxxBnNp7QlEtl3zd3aebc7OaFXTCCcicGIATnRyMDuOBaqokxhTPe2Ogs2do65ULNI%2BBjPAZGR3%2FqOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEOTk0SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFDFsYyOs9FH1WYU5EPLLV2JmEwTdMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiBgPfBoe0AQyAuGdMN177LjWOaBWOcKNlVlIOsNDE2BnwIgMVwqPJPcjiObRXwRSKtuVM39zoYpnmYPPtyqqrl1GdE%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 994J', but common name is 'SHAKEN Medtel Communications 994J' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 22 Aug 24 16:06 UTC