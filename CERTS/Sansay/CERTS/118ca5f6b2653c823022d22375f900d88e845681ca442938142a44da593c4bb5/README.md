# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Ytel Inc. 703J

Tested At: 18 Aug 25 20:38 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -170 day(s)\
Subject: CN=SHAKEN Ytel Inc. 703J, OU=Connect, O=Ytel Inc., ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/703J/429C7C70711E3820F0B8E1DEAE6FF326226493BE.pem

[View certificate details](https://x509.io/?cert=MIICyzCCAnKgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkk74wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDIyOTIzNTIxOFoXDTI1MDIyODIzNTIxOFowaDELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAoMCVl0ZWwgSW5jLjEQMA4GA1UECwwHQ29ubmVjdDEeMBwGA1UEAwwVU0hBS0VOIFl0ZWwgSW5jLiA3MDNKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEpyjhdlT3rpIHVGPi6cZ52IOv2y7dhv3ls7n%2FPO6TNuhp2DFuGw%2FSYYptOMaYbyw8Y8rudlGpq3cHa7rJ1%2FENcKOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENzAzSjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFDPhXure7G2QobiIsc4JnXnK%2FH%2FuMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiA2WLp%2Bvih3WJqmfvMjNFVbB62%2FTKixWHAIdaYuGu5%2BJgIgSPEPw%2BvfRSPn6HJ9IL2klsm8V2rEG4I6RWGYG%2FRwQQA%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 703J', but common name is 'SHAKEN Ytel Inc. 703J' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 18 Aug 25 21:13 UTC