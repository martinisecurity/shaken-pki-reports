# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN  XCast Labs 689J

Tested At: 17 Nov 22 19:11 UTC\
Initial Validity Period: 31 day(s)\
Remaining Validity Period: 24 day(s)\
Subject: CN=SHAKEN  XCast Labs 689J, OU=XCast Labs, O=XCast Labs, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.xcastlabs.net/1670740200/xclsshaken.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC0TCCAnigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkVNgwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTExMDA2MzAwMFoXDTIyMTIxMTA2MzAwMFowbjELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEzARBgNVBAoMClhDYXN0IExhYnMxEzARBgNVBAsMClhDYXN0IExhYnMxIDAeBgNVBAMMF1NIQUtFTiAgWENhc3QgTGFicyA2ODlKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAECSaRM8oTiHWzZxnWX07VK9ClMrfygu659qCz3%2FNdIZetgYKmv2%2Fi%2B%2F3bRH4jjvo5INmhtTlMMWWnkQ0T%2Bn9R4qOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENjg5SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFELtX4ZXV9APsoUdJG5BLBEV4aPuMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiBzPm%2BC0tw5SuW7RGeFoeT7DSvTrmCnuWpmN8KOZY3U0gIgC8DdN0OaOfSRfZS5E6ghc9RBUEViazn2Pyo9NuELXFo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 689J' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 17 Nov 22 19:20 UTC