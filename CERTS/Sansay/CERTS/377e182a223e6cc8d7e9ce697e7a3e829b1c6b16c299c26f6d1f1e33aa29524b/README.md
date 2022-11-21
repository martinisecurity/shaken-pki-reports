# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Arbeit 816J

Tested At: 21 Nov 22 23:25 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 337 day(s)\
Subject: CN=SHAKEN Arbeit 816J, OU=TAC, O=Arbeit, ST=New York, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Arbeit_816J

[View certificate details](https://understandingwebpki.com/?cert=MIICvzCCAmagAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkUtwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNDIwMjMwNloXDTIzMTAyNDIwMjMwNlowXDELMAkGA1UEBhMCVVMxETAPBgNVBAgMCE5ldyBZb3JrMQ8wDQYDVQQKDAZBcmJlaXQxDDAKBgNVBAsMA1RBQzEbMBkGA1UEAwwSU0hBS0VOIEFyYmVpdCA4MTZKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEyAi4mR9yB6bXrTeLnDA75UTvOJcWnm6KduqJ4ghDZAu%2FNUqTIvDhOS1AhGc6Yty0QgIl5PNQi4LTHHEh622%2F3KOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEODE2SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFDj7hlp21bpCUJdnFF8Xkr8FmgVRMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiBZP6a25IC3Sed7e9iEG8p0J3Cc1hlOY8wJQWIzCBO86QIgYi235uDcfryDZ5WTCMZ6cRrpDOomD%2FpwzhutNpXdSlI%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 816J' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 21 Nov 22 23:27 UTC