# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 7e46cda05ca9cc6162b35e3bca5473f3b116f894
Tested At: 2022-10-26 20:31:42 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 350 day(s)\
Subject: CN=SHAKEN Xchange Telecom LLC 325B, OU=Bulk Solutions STI-AS, O=Xchange Telecom LLC, ST=New York, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/325B_20211101.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC7TCCApOgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkT0IwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MjQyMFoXDTIzMTAxMTE3MjQyMFowgYgxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazEcMBoGA1UECgwTWGNoYW5nZSBUZWxlY29tIExMQzEeMBwGA1UECwwVQnVsayBTb2x1dGlvbnMgU1RJLUFTMSgwJgYDVQQDDB9TSEFLRU4gWGNoYW5nZSBUZWxlY29tIExMQyAzMjVCMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE0ufYApoQnxM2ww8zQ%2Be%2BEyvMooxbmK9ISQFoyTcGN6WoPxmvIXWBGBu1zv3aKBksXVrKOdAsE5j3ONV%2BWqorXqOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMzI1QjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFMQWXI5nFolUBJpN5j%2BbQWR3TIohMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAuQ1Rrf0y%2B4Iy0ZaiMk729eZ01%2Bizbyosd2b3wfC%2B9SsCIAoD%2F1GhbggaNRQ%2FW7xTTPejbObSDq02A5x0lKd6BEvM)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 325B' |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 26/10/2022 at 20:32:17