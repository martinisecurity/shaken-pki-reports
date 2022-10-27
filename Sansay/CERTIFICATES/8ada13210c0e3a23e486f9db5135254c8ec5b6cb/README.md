# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 8ada13210c0e3a23e486f9db5135254c8ec5b6cb
Tested At: 2022-10-27 21:42:34 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 30 day(s)\
Subject: CN=SHAKEN Airespring 996H, OU=Airespring NOC, O=Airespring, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Airespring_996H

View: [Click to view](https://understandingwebpki.com/?cert=MIIC1TCCAnugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU4MwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNzAwNDAzMFoXDTIyMTEyNjAwNDAzMFowcTELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEzARBgNVBAoMCkFpcmVzcHJpbmcxFzAVBgNVBAsMDkFpcmVzcHJpbmcgTk9DMR8wHQYDVQQDDBZTSEFLRU4gQWlyZXNwcmluZyA5OTZIMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEO%2FDtGCxkRKOrL%2FCO4UTvqQyw1p4MqYMIqk92awEPU6Zj%2BIbwKfhahHQHw%2BYo3o%2BOQupLUMCFOnweTLm%2B2rEXzKOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEOTk2SDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFGXP4f0G2020cYTsBTAPzTqgd8AUMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEA2iQ6z2Agw%2BzAywXpUZ%2Fsm%2FfDQPa5oG755gDjvklWKDkCIAj6xZ1%2FGB9q2qQw%2BA6jhfRuRq42x5r0VEY6H0aTrlnK)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 996H' |
| w_cp1_3_subject_rdn_unknown | warn | United States SHAKEN CP | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 27/10/2022 at 21:42:52