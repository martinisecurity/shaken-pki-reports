# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN  XCast Labs 689J

Tested At: 04 Oct 24 16:15 UTC\
Initial Validity Period: 31 day(s)\
Remaining Validity Period: -228 day(s)\
Subject: CN=SHAKEN  XCast Labs 689J, OU=XCast Labs, O=XCast Labs, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.xcastlabs.net/1708324200/xclsshaken.pem

[View certificate details](https://x509.io/?cert=MIIC0jCCAnigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkjkcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDExOTA2MzAwMFoXDTI0MDIxOTA2MzAwMFowbjELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEzARBgNVBAoMClhDYXN0IExhYnMxEzARBgNVBAsMClhDYXN0IExhYnMxIDAeBgNVBAMMF1NIQUtFTiAgWENhc3QgTGFicyA2ODlKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAECSaRM8oTiHWzZxnWX07VK9ClMrfygu659qCz3%2FNdIZetgYKmv2%2Fi%2B%2F3bRH4jjvo5INmhtTlMMWWnkQ0T%2Bn9R4qOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENjg5SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFELtX4ZXV9APsoUdJG5BLBEV4aPuMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAqRu61cZv%2FrZS%2BkQRLw7pkwp%2BfaN0Svv%2BpdU146ZTo%2F4CIFHwfFgdBsYasBf9iSRZX77Y9gFJkG2DcSRw0SZ9dhZh)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 689J', but common name is 'SHAKEN  XCast Labs 689J' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 04 Oct 24 16:29 UTC