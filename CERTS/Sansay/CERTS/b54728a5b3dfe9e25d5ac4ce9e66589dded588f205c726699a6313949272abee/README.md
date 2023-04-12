# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Ace Innovative Networks, Inc. 040K

Tested At: 12 Apr 23 21:49 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -5 day(s)\
Subject: emailAddress=elevitt@1pcom.com, CN=SHAKEN Ace Innovative Networks\\, Inc. 040K, OU=NOC, O=Ace Innovative Networks\\, Inc., ST=New York, C=US, emailAddress=elevitt@1pcom.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/040K/order/274_040K_84

[View certificate details](https://understandingwebpki.com/?cert=MIIDETCCAregAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkZAYwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDMwODIxMTcxOFoXDTIzMDQwNzIxMTcxOFowgawxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazEmMCQGA1UECgwdQWNlIElubm92YXRpdmUgTmV0d29ya3MsIEluYy4xDDAKBgNVBAsMA05PQzEyMDAGA1UEAwwpU0hBS0VOIEFjZSBJbm5vdmF0aXZlIE5ldHdvcmtzLCBJbmMuIDA0MEsxIDAeBgkqhkiG9w0BCQEWEWVsZXZpdHRAMXBjb20uY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEOqyYRuppi4NWGh0Dk0LShTDAr4sxeglOkve2zQwZPQJfpBYQOwwROCoWfqq%2BJEldZwIV9zMtR2PhzXqB%2B22djKOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMDQwSzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFI4Gh%2FAp%2BjtuoawM59l0d1Wq6oMoMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAjBzhInjZx%2BD%2FRRMEp8cHJhceWXkbhavmMKNMi6Bua90CIGaeHcKB1RdDCkgnX9zS7OXvUsQqmnSMdrstcAgi9cX2)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 040K' |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 12 Apr 23 22:02 UTC