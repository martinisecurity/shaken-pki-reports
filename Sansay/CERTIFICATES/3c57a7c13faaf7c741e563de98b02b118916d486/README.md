# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 3c57a7c13faaf7c741e563de98b02b118916d486
Tested At: 2022-10-28 18:22:27 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 362 day(s)\
Subject: CN=SHAKEN Arbeit 816J, OU=TAC, O=Arbeit, ST=New York, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Arbeit_816J

View: [Click to view](https://understandingwebpki.com/?cert=MIICvzCCAmagAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkUtwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNDIwMjMwNloXDTIzMTAyNDIwMjMwNlowXDELMAkGA1UEBhMCVVMxETAPBgNVBAgMCE5ldyBZb3JrMQ8wDQYDVQQKDAZBcmJlaXQxDDAKBgNVBAsMA1RBQzEbMBkGA1UEAwwSU0hBS0VOIEFyYmVpdCA4MTZKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEyAi4mR9yB6bXrTeLnDA75UTvOJcWnm6KduqJ4ghDZAu%2FNUqTIvDhOS1AhGc6Yty0QgIl5PNQi4LTHHEh622%2F3KOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEODE2SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFDj7hlp21bpCUJdnFF8Xkr8FmgVRMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiBZP6a25IC3Sed7e9iEG8p0J3Cc1hlOY8wJQWIzCBO86QIgYi235uDcfryDZ5WTCMZ6cRrpDOomD%2FpwzhutNpXdSlI%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 816J' |
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 28/10/2022 at 18:22:55