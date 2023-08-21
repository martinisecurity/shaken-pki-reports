# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Medtel Communications 994J

Tested At: 21 Aug 23 20:14 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 284 day(s)\
Subject: CN=SHAKEN Medtel Communications 994J, OU=Medtel Communications, O=Medtel Communications, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Medtel_994J

[View certificate details](https://understandingwebpki.com/?cert=MIIC8TCCApagAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkcu8wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDYwMTE5MzMyMloXDTI0MDUzMTE5MzMyMlowgYsxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdGbG9yaWRhMR4wHAYDVQQKDBVNZWR0ZWwgQ29tbXVuaWNhdGlvbnMxHjAcBgNVBAsMFU1lZHRlbCBDb21tdW5pY2F0aW9uczEqMCgGA1UEAwwhU0hBS0VOIE1lZHRlbCBDb21tdW5pY2F0aW9ucyA5OTRKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAElm6J72YV8ZxxBnNp7QlEtl3zd3aebc7OaFXTCCcicGIATnRyMDuOBaqokxhTPe2Ogs2do65ULNI%2BBjPAZGR3%2FqOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEOTk0SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFDFsYyOs9FH1WYU5EPLLV2JmEwTdMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEAx9BKJNyivVKlsXNcjlAOKAZVds1Hda%2B6TLfmrOQbyfYCIQDXzlzUNy3sVAACIL2tlvpbYfhKAXfQ33g%2FWJiObzVnfw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 994J' |


Generated: 21 Aug 23 20:18 UTC