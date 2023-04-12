# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN CIMA Telecom, Inc 313K

Tested At: 12 Apr 23 21:51 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 21 day(s)\
Subject: CN=SHAKEN CIMA Telecom\\, Inc 313K, OU=NOC, O=CIMA Telecom\\, Inc, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/CIMA_Telecom_Inc_313K

[View certificate details](https://understandingwebpki.com/?cert=MIIC1TCCAnugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkaLwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDQwMzE0NTUwOVoXDTIzMDUwMzE0NTUwOVowcTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0Zsb3JpZGExGjAYBgNVBAoMEUNJTUEgVGVsZWNvbSwgSW5jMQwwCgYDVQQLDANOT0MxJjAkBgNVBAMMHVNIQUtFTiBDSU1BIFRlbGVjb20sIEluYyAzMTNLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEWZxYiUIGA%2BQCeT097eUc88lbP9RiuLcm4VLtHrIQk5%2B%2Ful91lqbaBIoCPigzfPkiiHThsByt287tLmZDVCxhHqOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMzEzSzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFG2v9Azvt9PJnHi3QS1utwBKsBviMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiBElXKd3JeVYsFNEApuKh6tPNKFEOWzFdF5szha2aZruwIhAIV8ROwumm9diumMjSxhsDd3ukU1EKamId9lBw4h5Z5o)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 313K' |


Generated: 12 Apr 23 22:02 UTC