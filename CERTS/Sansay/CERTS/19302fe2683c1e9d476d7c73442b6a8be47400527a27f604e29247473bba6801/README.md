# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Arbeit 816J

Tested At: 28 Nov 23 16:06 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 1 day(s)\
Subject: CN=SHAKEN Arbeit 816J, OU=TAC, O=Arbeit, ST=New York, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Arbeit_816J

[View certificate details](https://understandingwebpki.com/?cert=MIICwDCCAmagAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhdEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTAzMDE0MDMwM1oXDTIzMTEyOTE0MDMwM1owXDELMAkGA1UEBhMCVVMxETAPBgNVBAgMCE5ldyBZb3JrMQ8wDQYDVQQKDAZBcmJlaXQxDDAKBgNVBAsMA1RBQzEbMBkGA1UEAwwSU0hBS0VOIEFyYmVpdCA4MTZKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEyAi4mR9yB6bXrTeLnDA75UTvOJcWnm6KduqJ4ghDZAu%2FNUqTIvDhOS1AhGc6Yty0QgIl5PNQi4LTHHEh622%2F3KOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEODE2SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFDj7hlp21bpCUJdnFF8Xkr8FmgVRMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAyhldlyLOgxVZw6GzYU5KrZaceu59I8y1%2BJqNbd4duGUCIGQRgOS2INnRe9yx99B4FJNWXlMzCVD%2Bs%2BWYX2zOqPGs)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 816J', but common name is 'SHAKEN Arbeit 816J' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 28 Nov 23 16:15 UTC