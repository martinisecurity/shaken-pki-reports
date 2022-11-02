# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Xchange Telecom LLC 325B

Tested At: 02 Nov 22 17:24 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 343 day(s)\
Subject: CN=SHAKEN Xchange Telecom LLC 325B, OU=Bulk Solutions STI-AS, O=Xchange Telecom LLC, ST=New York, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/325B_20211101.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7TCCApOgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkT0IwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MjQyMFoXDTIzMTAxMTE3MjQyMFowgYgxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazEcMBoGA1UECgwTWGNoYW5nZSBUZWxlY29tIExMQzEeMBwGA1UECwwVQnVsayBTb2x1dGlvbnMgU1RJLUFTMSgwJgYDVQQDDB9TSEFLRU4gWGNoYW5nZSBUZWxlY29tIExMQyAzMjVCMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE0ufYApoQnxM2ww8zQ%2Be%2BEyvMooxbmK9ISQFoyTcGN6WoPxmvIXWBGBu1zv3aKBksXVrKOdAsE5j3ONV%2BWqorXqOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMzI1QjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFMQWXI5nFolUBJpN5j%2BbQWR3TIohMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAuQ1Rrf0y%2B4Iy0ZaiMk729eZ01%2Bizbyosd2b3wfC%2B9SsCIAoD%2F1GhbggaNRQ%2FW7xTTPejbObSDq02A5x0lKd6BEvM)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 325B' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 02 Nov 22 17:25 UTC